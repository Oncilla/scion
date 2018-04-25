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

package sibra

import "github.com/scionproto/scion/go/lib/common"

const (
	// infoLen is the reservation info length.
	infoLen = common.LineLen
)

// Info is the SIBRA reservation info field. It stores information about
// a (requested or active) reservation.
//
// 0B       1        2        3        4        5        6        7
// +--------+--------+--------+--------+--------+--------+--------+--------+
// | Expiration time (4B)              | BW Cls | RTT Cls|Idx|Type|Fail hop|
// +--------+--------+--------+--------+--------+--------+--------+--------+
//
// The bandwidth class (BW Cls) indicates the reserved bandwidth in an active
// reservation. In a steady request, it indicates the minimal bandwidth class
// reserved so far. In a ephemeral request, it indicates the bandwidth class
// that the source endhost is seeking to reserve.
//
// The round trip class (RTT Cls) allows for more granular control in the
// pending request garbage collection.
//
// The reservation index (Idx) is used to allow for multiple overlapping
// reservations within a single path, which enables renewal and changing the
// bandwidth requested.
//
// Type indicates which path type the reservation is for
//  000 := Down
//  001 := Up
//  010 := Peering-Down
//  011 := Peering-Up
//  100 := Ephemeral
//  101 := Core
//
// The fail hop field is normally set to 0, and ignored unless this reservation
// info is part of a denied request, in which case it is set to the number of
// the first hop to reject the reservation.
type Info struct {
	// ExpTick is the SIBRA tick when the reservation expires.
	ExpTick Tick
	// BwClass is the bandwidth class.
	BwClass BwClass
	// RttClass is the round trip class.
	RttClass RttClass
	// Index is the reservation index.
	Index Index
	// PathType is the path type.
	PathType PathType
	// FailHop is the fail hop. It is unset in an active reservation.
	FailHop uint8
}

func NewInfoFromRaw(raw common.RawBytes) *Info {
	return &Info{
		ExpTick:  Tick(common.Order.Uint32(raw[:4])),
		BwClass:  BwClass(raw[4]),
		RttClass: RttClass(raw[5]),
		Index:    parseIdx(raw[6]),
		PathType: parsePathType(raw[6]),
		FailHop:  raw[7],
	}
}

func (i *Info) Write(b common.RawBytes) error {
	if len(b) < i.Len() {
		return common.NewBasicError(BufferToShort, nil, "method", "SIBRAInfo.Write",
			"min", i.Len(), "actual", len(b))
	}
	common.Order.PutUint32(b[:4], uint32(i.ExpTick))
	b[4] = uint8(i.BwClass)
	b[5] = uint8(i.RttClass)
	b[6] = i.packIdxPathType()
	b[7] = i.FailHop
	return nil
}

func (i *Info) Len() int {
	return infoLen
}

func (i *Info) packIdxPathType() uint8 {
	return (uint8(i.Index) << 4) | uint8(i.PathType)
}

func parseIdx(field uint8) Index {
	return Index(field >> 4)
}

func parsePathType(field uint8) PathType {
	return PathType(field & 0x7)
}
