// Copyright 2018 ETH Zurich
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resvmgr

import (
	"net"
	"sync"

	"github.com/scionproto/scion/go/lib/addr"
	"github.com/scionproto/scion/go/lib/common"
	"github.com/scionproto/scion/go/lib/ctrl/sibra_mgmt"
	"github.com/scionproto/scion/go/lib/hpkt"
	"github.com/scionproto/scion/go/lib/infra"
	"github.com/scionproto/scion/go/lib/log"
	"github.com/scionproto/scion/go/lib/sibra"
	"github.com/scionproto/scion/go/lib/sibra/sbextn"
	"github.com/scionproto/scion/go/lib/sibra/sbreq"
	"github.com/scionproto/scion/go/lib/snet"
	"github.com/scionproto/scion/go/lib/spkt"
)

var _ infra.Handler = (*ephemRepHandler)(nil)
var _ repMaster = (*ephemRepHandler)(nil)

type ephemRepHandler struct {
	sync.Mutex
	log.Logger
	// notify is a mapping from notifyKey to a notify channel
	// for a listener.
	notify map[string]chan notifyEvent
}

// Register allows listeners to register for certain notify keys.
func (h *ephemRepHandler) Register(key *notifyKey, c chan notifyEvent) error {
	h.Lock()
	defer h.Unlock()
	if _, ok := h.notify[key.String()]; ok {
		return common.NewBasicError("Unable to add multiple listeners", nil,
			"id", key.Id, "idx", key.Idx, "rt", key.ReqType)
	}
	h.notify[key.String()] = c
	return nil
}

// Deregister removes the listener from the mapping for the given key.
func (h *ephemRepHandler) Deregister(key *notifyKey) {
	h.Lock()
	defer h.Unlock()
	delete(h.notify, key.String())
}

func (h *ephemRepHandler) Handle(r *infra.Request) {
	saddr := r.Peer.(*snet.Addr)
	pld := r.Message.(*sibra_mgmt.EphemRep)
	h.Debug("Received reply", "addr", saddr)
	if err := h.handle(saddr, pld); err != nil {
		h.logDropRep(saddr, pld, err)
	}
}

func (h *ephemRepHandler) handle(saddr *snet.Addr, pld *sibra_mgmt.EphemRep) error {
	pkt, err := pld.ScnPkt()
	if err != nil {
		return err
	}
	event, base, err := parseRep(pkt)
	if err != nil {
		return err
	}
	log.Debug("Received event", "addr", saddr, "extn", event.extn, "pld", event.pld)
	key, err := h.getNotifyKey(base, event.pld)
	if err != nil {
		return err
	}
	h.Lock()
	defer h.Unlock()
	if c, ok := h.notify[key.String()]; ok {
		select {
		case c <- event:
		default:
			return common.NewBasicError("Drop reply due to full buffer", nil)
		}
		return nil
	}
	return common.NewBasicError("No listener registered", nil, "key", key)
}

func (h *ephemRepHandler) getNotifyKey(base *sbextn.Base, pld *sbreq.Pld) (notifyKey, error) {
	var id sibra.ID
	var idx sibra.Index
	switch e := pld.Data.(type) {
	case *sbreq.EphemReq:
		id = e.ID
		idx = e.Block.Info.Index
	case *sbreq.EphemFailed:
		id = e.ID
		idx = e.Info.Index
	case *sbreq.EphemClean:
		id = e.ID
		idx = e.Info.Index
	default:
		return notifyKey{}, common.NewBasicError("Request type not supported", nil, "req", e)
	}
	if id == nil {
		if base.Steady {
			return notifyKey{}, common.NewBasicError("No ID for request on steady reservation", nil)
		}
		id = base.GetCurrID()
	}

	key := notifyKey{
		Id:      id,
		Idx:     idx,
		ReqType: pld.Type,
	}
	return key, nil
}

func (h *ephemRepHandler) logDropRep(addr net.Addr, rep *sibra_mgmt.EphemRep, err error) {
	h.Error("Dropping request", "addr", addr, "rep", rep, "err", err)
}

var _ infra.Handler = (*ephemReqHandler)(nil)

type ephemReqHandler struct {
	whitelist *whitelist
	log.Logger
}

func (h *ephemReqHandler) Handle(r *infra.Request) {
	saddr := r.Peer.(*snet.Addr)
	pld := r.Message.(*sibra_mgmt.EphemReq)
	h.Debug("Received request", "addr", saddr)

	pkt, err := h.handle(saddr, pld)
	if err != nil {
		h.logDropReq(saddr, pld, err)
		return
	}
	if err := h.sendRep(r, saddr, pkt); err != nil {
		h.logDropReq(saddr, pld, err)
	}
}

func (h *ephemReqHandler) handle(saddr *snet.Addr, pld *sibra_mgmt.EphemReq) (*spkt.ScnPkt, error) {
	pkt, err := pld.ScnPkt()
	if err != nil {
		return nil, err
	}
	event, base, err := parseRep(pkt)
	if err != nil {
		return nil, err
	}
	log.Debug("Received event", "addr", saddr, "extn", event.extn, "pld", event.pld)
	ok, err := h.checkWhitelist(pkt.SrcIA, pkt.SrcHost.IP(), event.pld, base.CurrHop)
	if err != nil {
		return nil, err
	}
	if !ok {
		h.Debug("Denied request (client not on whitelist)", "addr", saddr, "req", pld)
	}
	// FIXME(roosd): remove when extension reversing is supported
	if _, err = event.extn.Reverse(); err != nil {
		return nil, err
	}
	if err = event.pld.Reverse(); err != nil {
		return nil, err
	}
	if err := pkt.Reverse(); err != nil {
		return nil, err
	}
	pkt.Pld = event.pld
	return pkt, nil
}

func (h *ephemReqHandler) checkWhitelist(ia addr.IA, ip net.IP, pld *sbreq.Pld,
	currHop int) (bool, error) {
	if r, ok := pld.Data.(*sbreq.EphemReq); ok {
		if !h.whitelist.isAllowed(ia, ip) {
			pld.Data = r.Fail(sbreq.ClientDenied, 0, currHop)
			pld.Accepted = false
			return false, nil
		}
		return true, nil
	}
	return false, common.NewBasicError("Not a ephemeral request", nil)
}

func (h *ephemReqHandler) sendRep(r *infra.Request, saddr *snet.Addr, pkt *spkt.ScnPkt) error {
	buf := make(common.RawBytes, pkt.TotalLen())
	if _, err := hpkt.WriteScnPkt(pkt, buf); err != nil {
		return err
	}
	rep := &sibra_mgmt.EphemRep{
		ExternalPkt: &sibra_mgmt.ExternalPkt{
			RawPkt: buf,
		},
	}
	msgr, ok := infra.MessengerFromContext(r.Context())
	if !ok {
		return common.NewBasicError("No messenger found", nil)
	}
	if err := msgr.SendSibraEphemRep(r.Context(), rep, saddr, r.ID); err != nil {
		return err
	}
	h.Debug("Sent reply", "saddr", saddr, "id", r.ID)
	return nil
}

func (h *ephemReqHandler) logDropReq(addr net.Addr, req *sibra_mgmt.EphemReq, err error) {
	h.Error("Dropping request", "addr", addr, "req", req, "err", err)
}

func parseRep(pkt *spkt.ScnPkt) (notifyEvent, *sbextn.Base, error) {
	exts := pkt.GetExtn(common.ExtnSIBRAType)
	if len(exts) < 1 {
		return notifyEvent{}, nil, common.NewBasicError("No SIBRA header found", nil)

	}
	ext := exts[0]
	var base *sbextn.Base
	switch sibraExtn := ext.(type) {
	case *sbextn.Steady:
		base = sibraExtn.Base
	case *sbextn.Ephemeral:
		base = sibraExtn.Base
	}
	pld, err := sbreq.PldFromRaw(pkt.Pld.(common.RawBytes))
	if err != nil {
		return notifyEvent{}, nil, common.NewBasicError("Unable to parse payload", err)
	}
	event := notifyEvent{
		extn: ext,
		pld:  pld,
	}
	return event, base, nil
}
