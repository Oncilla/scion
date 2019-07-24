// Code generated by capnpc-go. DO NOT EDIT.

package proto

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type DRKeyReq struct{ capnp.Struct }
type DRKeyReq_flags DRKeyReq

// DRKeyReq_TypeID is the unique identifier for the type DRKeyReq.
const DRKeyReq_TypeID = 0x9f50d21c9d4ce7ef

func NewDRKeyReq(s *capnp.Segment) (DRKeyReq, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1})
	return DRKeyReq{st}, err
}

func NewRootDRKeyReq(s *capnp.Segment) (DRKeyReq, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1})
	return DRKeyReq{st}, err
}

func ReadRootDRKeyReq(msg *capnp.Message) (DRKeyReq, error) {
	root, err := msg.RootPtr()
	return DRKeyReq{root.Struct()}, err
}

func (s DRKeyReq) String() string {
	str, _ := text.Marshal(0x9f50d21c9d4ce7ef, s.Struct)
	return str
}

func (s DRKeyReq) Isdas() uint64 {
	return s.Struct.Uint64(0)
}

func (s DRKeyReq) SetIsdas(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s DRKeyReq) Timestamp() uint32 {
	return s.Struct.Uint32(8)
}

func (s DRKeyReq) SetTimestamp(v uint32) {
	s.Struct.SetUint32(8, v)
}

func (s DRKeyReq) Signature() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s DRKeyReq) HasSignature() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s DRKeyReq) SetSignature(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s DRKeyReq) CertVer() uint32 {
	return s.Struct.Uint32(12)
}

func (s DRKeyReq) SetCertVer(v uint32) {
	s.Struct.SetUint32(12, v)
}

func (s DRKeyReq) TrcVer() uint32 {
	return s.Struct.Uint32(16)
}

func (s DRKeyReq) SetTrcVer(v uint32) {
	s.Struct.SetUint32(16, v)
}

func (s DRKeyReq) Flags() DRKeyReq_flags { return DRKeyReq_flags(s) }

func (s DRKeyReq_flags) Prefetch() bool {
	return s.Struct.Bit(160)
}

func (s DRKeyReq_flags) SetPrefetch(v bool) {
	s.Struct.SetBit(160, v)
}

// DRKeyReq_List is a list of DRKeyReq.
type DRKeyReq_List struct{ capnp.List }

// NewDRKeyReq creates a new list of DRKeyReq.
func NewDRKeyReq_List(s *capnp.Segment, sz int32) (DRKeyReq_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1}, sz)
	return DRKeyReq_List{l}, err
}

func (s DRKeyReq_List) At(i int) DRKeyReq { return DRKeyReq{s.List.Struct(i)} }

func (s DRKeyReq_List) Set(i int, v DRKeyReq) error { return s.List.SetStruct(i, v.Struct) }

func (s DRKeyReq_List) String() string {
	str, _ := text.MarshalList(0x9f50d21c9d4ce7ef, s.List)
	return str
}

// DRKeyReq_Promise is a wrapper for a DRKeyReq promised by a client call.
type DRKeyReq_Promise struct{ *capnp.Pipeline }

func (p DRKeyReq_Promise) Struct() (DRKeyReq, error) {
	s, err := p.Pipeline.Struct()
	return DRKeyReq{s}, err
}

func (p DRKeyReq_Promise) Flags() DRKeyReq_flags_Promise { return DRKeyReq_flags_Promise{p.Pipeline} }

// DRKeyReq_flags_Promise is a wrapper for a DRKeyReq_flags promised by a client call.
type DRKeyReq_flags_Promise struct{ *capnp.Pipeline }

func (p DRKeyReq_flags_Promise) Struct() (DRKeyReq_flags, error) {
	s, err := p.Pipeline.Struct()
	return DRKeyReq_flags{s}, err
}

type DRKeyRep struct{ capnp.Struct }

// DRKeyRep_TypeID is the unique identifier for the type DRKeyRep.
const DRKeyRep_TypeID = 0xc3fe25dd82681d64

func NewDRKeyRep(s *capnp.Segment) (DRKeyRep, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 2})
	return DRKeyRep{st}, err
}

func NewRootDRKeyRep(s *capnp.Segment) (DRKeyRep, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 2})
	return DRKeyRep{st}, err
}

func ReadRootDRKeyRep(msg *capnp.Message) (DRKeyRep, error) {
	root, err := msg.RootPtr()
	return DRKeyRep{root.Struct()}, err
}

func (s DRKeyRep) String() string {
	str, _ := text.Marshal(0xc3fe25dd82681d64, s.Struct)
	return str
}

func (s DRKeyRep) Isdas() uint64 {
	return s.Struct.Uint64(0)
}

func (s DRKeyRep) SetIsdas(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s DRKeyRep) Timestamp() uint32 {
	return s.Struct.Uint32(8)
}

func (s DRKeyRep) SetTimestamp(v uint32) {
	s.Struct.SetUint32(8, v)
}

func (s DRKeyRep) ExpTime() uint32 {
	return s.Struct.Uint32(12)
}

func (s DRKeyRep) SetExpTime(v uint32) {
	s.Struct.SetUint32(12, v)
}

func (s DRKeyRep) Cipher() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s DRKeyRep) HasCipher() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s DRKeyRep) SetCipher(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s DRKeyRep) Signature() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s DRKeyRep) HasSignature() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s DRKeyRep) SetSignature(v []byte) error {
	return s.Struct.SetData(1, v)
}

func (s DRKeyRep) CertVerSrc() uint32 {
	return s.Struct.Uint32(16)
}

func (s DRKeyRep) SetCertVerSrc(v uint32) {
	s.Struct.SetUint32(16, v)
}

func (s DRKeyRep) CertVerDst() uint32 {
	return s.Struct.Uint32(20)
}

func (s DRKeyRep) SetCertVerDst(v uint32) {
	s.Struct.SetUint32(20, v)
}

func (s DRKeyRep) TrcVer() uint32 {
	return s.Struct.Uint32(24)
}

func (s DRKeyRep) SetTrcVer(v uint32) {
	s.Struct.SetUint32(24, v)
}

// DRKeyRep_List is a list of DRKeyRep.
type DRKeyRep_List struct{ capnp.List }

// NewDRKeyRep creates a new list of DRKeyRep.
func NewDRKeyRep_List(s *capnp.Segment, sz int32) (DRKeyRep_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 32, PointerCount: 2}, sz)
	return DRKeyRep_List{l}, err
}

func (s DRKeyRep_List) At(i int) DRKeyRep { return DRKeyRep{s.List.Struct(i)} }

func (s DRKeyRep_List) Set(i int, v DRKeyRep) error { return s.List.SetStruct(i, v.Struct) }

func (s DRKeyRep_List) String() string {
	str, _ := text.MarshalList(0xc3fe25dd82681d64, s.List)
	return str
}

// DRKeyRep_Promise is a wrapper for a DRKeyRep promised by a client call.
type DRKeyRep_Promise struct{ *capnp.Pipeline }

func (p DRKeyRep_Promise) Struct() (DRKeyRep, error) {
	s, err := p.Pipeline.Struct()
	return DRKeyRep{s}, err
}

type DRKeyMgmt struct{ capnp.Struct }
type DRKeyMgmt_Which uint16

const (
	DRKeyMgmt_Which_unset    DRKeyMgmt_Which = 0
	DRKeyMgmt_Which_drkeyReq DRKeyMgmt_Which = 1
	DRKeyMgmt_Which_drkeyRep DRKeyMgmt_Which = 2
)

func (w DRKeyMgmt_Which) String() string {
	const s = "unsetdrkeyReqdrkeyRep"
	switch w {
	case DRKeyMgmt_Which_unset:
		return s[0:5]
	case DRKeyMgmt_Which_drkeyReq:
		return s[5:13]
	case DRKeyMgmt_Which_drkeyRep:
		return s[13:21]

	}
	return "DRKeyMgmt_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// DRKeyMgmt_TypeID is the unique identifier for the type DRKeyMgmt.
const DRKeyMgmt_TypeID = 0xb1bdb7d6fb13f1ca

func NewDRKeyMgmt(s *capnp.Segment) (DRKeyMgmt, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return DRKeyMgmt{st}, err
}

func NewRootDRKeyMgmt(s *capnp.Segment) (DRKeyMgmt, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return DRKeyMgmt{st}, err
}

func ReadRootDRKeyMgmt(msg *capnp.Message) (DRKeyMgmt, error) {
	root, err := msg.RootPtr()
	return DRKeyMgmt{root.Struct()}, err
}

func (s DRKeyMgmt) String() string {
	str, _ := text.Marshal(0xb1bdb7d6fb13f1ca, s.Struct)
	return str
}

func (s DRKeyMgmt) Which() DRKeyMgmt_Which {
	return DRKeyMgmt_Which(s.Struct.Uint16(0))
}
func (s DRKeyMgmt) SetUnset() {
	s.Struct.SetUint16(0, 0)

}

func (s DRKeyMgmt) DrkeyReq() (DRKeyReq, error) {
	if s.Struct.Uint16(0) != 1 {
		panic("Which() != drkeyReq")
	}
	p, err := s.Struct.Ptr(0)
	return DRKeyReq{Struct: p.Struct()}, err
}

func (s DRKeyMgmt) HasDrkeyReq() bool {
	if s.Struct.Uint16(0) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s DRKeyMgmt) SetDrkeyReq(v DRKeyReq) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewDrkeyReq sets the drkeyReq field to a newly
// allocated DRKeyReq struct, preferring placement in s's segment.
func (s DRKeyMgmt) NewDrkeyReq() (DRKeyReq, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewDRKeyReq(s.Struct.Segment())
	if err != nil {
		return DRKeyReq{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s DRKeyMgmt) DrkeyRep() (DRKeyRep, error) {
	if s.Struct.Uint16(0) != 2 {
		panic("Which() != drkeyRep")
	}
	p, err := s.Struct.Ptr(0)
	return DRKeyRep{Struct: p.Struct()}, err
}

func (s DRKeyMgmt) HasDrkeyRep() bool {
	if s.Struct.Uint16(0) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s DRKeyMgmt) SetDrkeyRep(v DRKeyRep) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewDrkeyRep sets the drkeyRep field to a newly
// allocated DRKeyRep struct, preferring placement in s's segment.
func (s DRKeyMgmt) NewDrkeyRep() (DRKeyRep, error) {
	s.Struct.SetUint16(0, 2)
	ss, err := NewDRKeyRep(s.Struct.Segment())
	if err != nil {
		return DRKeyRep{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// DRKeyMgmt_List is a list of DRKeyMgmt.
type DRKeyMgmt_List struct{ capnp.List }

// NewDRKeyMgmt creates a new list of DRKeyMgmt.
func NewDRKeyMgmt_List(s *capnp.Segment, sz int32) (DRKeyMgmt_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return DRKeyMgmt_List{l}, err
}

func (s DRKeyMgmt_List) At(i int) DRKeyMgmt { return DRKeyMgmt{s.List.Struct(i)} }

func (s DRKeyMgmt_List) Set(i int, v DRKeyMgmt) error { return s.List.SetStruct(i, v.Struct) }

func (s DRKeyMgmt_List) String() string {
	str, _ := text.MarshalList(0xb1bdb7d6fb13f1ca, s.List)
	return str
}

// DRKeyMgmt_Promise is a wrapper for a DRKeyMgmt promised by a client call.
type DRKeyMgmt_Promise struct{ *capnp.Pipeline }

func (p DRKeyMgmt_Promise) Struct() (DRKeyMgmt, error) {
	s, err := p.Pipeline.Struct()
	return DRKeyMgmt{s}, err
}

func (p DRKeyMgmt_Promise) DrkeyReq() DRKeyReq_Promise {
	return DRKeyReq_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p DRKeyMgmt_Promise) DrkeyRep() DRKeyRep_Promise {
	return DRKeyRep_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

const schema_f85d2602085656c1 = "x\xda\x84\x92\xd1K\x14k\x18\xc6\x9f\xe7\xfbfv\x15" +
	"\\\xf0;3G\xe4p\x0e\x9e\x9b#\x9c\x1b+\xbd\x13" +
	"B\x09\x85,\x03\xbfU\x0c\xa2\x88e\x1d\xd7\xa5V\xc6" +
	"\x99\x11\xf2&\xa1\xbf\xa1\x9b\"\xc9\x8b\x02/\x8a\x12\x8a" +
	"\xba\xb0\x0b\x89\x08\xa2\x9b\x96\x0a\x89\x84\x02\xa5.\x92\x14" +
	"\x0a\x84\xd0\x89\xcfY\xd7%\x94\x98\xaby\xde\xe7{\xbf" +
	"\xf7\xf9~\xef\xe1u\xbb[\x1c\xb1\x0f\xb5\x02\x83\x0b\xb4" +
	"S\xf1\xd7O\xfd3\x7f\x97\x07nB+\xcaxqx" +
	"\xb8N\xb4\x9e\xdb\x84\xcd4\xe0l\xfdQv\xea\x9d4" +
	"\xd0a;\xa7\x09\xc6/6\x9c\x1fo\x1f=\x997n" +
	"\xee\xb9{\x99\x96\x80s\xd4}\xe7\xf4\xb9\xe6`\xaf{" +
	"\x0f\x8cG\xfe\x19\xbb\xb2\xfc\xdf\xf6S\xe3\xb6jz\x0b" +
	"cYu\xcb\xce\x861w\xac\xb9\xcfM\xef\x85\xd6\xff" +
	"\xc3\xcbkse\xe8\xbf(\xf7\xe6\xfa\x93i\x02\x1d\x8b" +
	"M\x82\xa0\xf3\xac\xa9\x0bq\xe5\xbb\x1e\x8f\x04\x17\xbc\xa9" +
	"\xf3\xa5\x82(Em\xf9\x9c?\xeew\xf6dOzS" +
	"Y\x8f\x13\x03\xa4n\x96\x16`\x11P\xd7\xda\x01}U" +
	"R\xcf\x0a*\x0a\x97F\x9c\xc9\x02\xfa\x86\xa4\x9e\x134" +
	"\x9a\x00\xd4m\xa3\xdd\x92\xd4\xf7\x05\x95\x94.%\xa0\xee" +
	"\x1e\x03\xf4\x9c\xa4~ \xa8,\xcb\xa5\x05\xa8\xf9N@" +
	"\xdf\x91\xd4\x8f\x05i\xb3f~\xf5\xb0\x1d\xa2\xa5\x18\x8e" +
	"\xe4B\xd6C\xb0\x1e\x8c\xa3b\xc9\x0b\xa3\\\x09\xf4Y" +
	"\x07\xc1:0\x0e\x8b\x85\xf1\\4\x19\x80\x1e3\x10\xcc" +
	"\x80\xd3y/\x88\x86\xbd`\xd7\xd3\x15\x05\xf9\x9a\xdf\x96" +
	"\xd1\x8b\xb9BxP\xecS\x05Y\x8aL\xee\x06i5" +
	"\xc4\xf1N\xf0^\x13\xbc[R\xf7\x0bf\xb8\x1d'\xc9" +
	"\xfbN\x00\xfa\xb8\xa4\x1e\x12\xcc\x88\xad8\xc9\xae\x8d:" +
	" \xa9\xcf\x0a\xb6L\x8e\x87^\x84TrW\xd6\x9b\x00" +
	"\xc0\xc6=* \x1b\x0d\xe1\xa4\xea'\xd5*\xef\xda\xea" +
	"\xbex|3\xe6\xbfU<\xaf\xcc\x94/%\xf5R\x0d" +
	"\x9e7\x06\xc5kI\xfdAP\x09\x99\xcc\xb8lP," +
	"I\xea\x15AV\xf0|4$\xdeK\xea\xcf\x06\x0f\x13" +
	"<\xab\xe6\xf4\x8a\xa4^\x17T\xb6\xe5\xd2\x06\xd4\xda\x19" +
	"@\x7f\x91\xd4\x9b\x82*e\xbbL\x01\xea\xbb\x11\xbfI" +
	"f)\xa8\xd2)\xd7l\xbe\xda2=7%\x07-\x0a" +
	"\xfe\x9e\xe5\xb4w\xc9\x1f*\x96\xbc*\xb7|\xd1\x1f\xf3" +
	"\x82]\xac\xfb\xa1\x8e+\xa8\x07!\x83|u'*b" +
	"\x0fd\x18\x1d\xb0\x04\xd5g\x95\xbf>\xebD\xdb\xce~" +
	"@[\xd2j\x9cM\"g\x0c\xd4\x06I\xdd,\x18\xfb" +
	"\x817\xeaE\xf91\x83\x8b\x10$\xf83\x00\x00\xff\xff" +
	"T\xf1\xfd\xa0"

func init() {
	schemas.Register(schema_f85d2602085656c1,
		0x9f50d21c9d4ce7ef,
		0xb1bdb7d6fb13f1ca,
		0xc3fe25dd82681d64,
		0xd2a8ed7e732926bc)
}