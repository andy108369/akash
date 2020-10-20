// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/market/v1beta1/bid.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// State is an enum which refers to state of bid
type Bid_State int32

const (
	// Prefix should start with 0 in enum. So declaring dummy state
	BidStateInvalid Bid_State = 0
	// BidOpen denotes state for bid open
	BidOpen Bid_State = 1
	// BidMatched denotes state for bid open
	BidMatched Bid_State = 2
	// BidLost denotes state for bid lost
	BidLost Bid_State = 3
	// BidClosed denotes state for bid closed
	BidClosed Bid_State = 4
)

var Bid_State_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "matched",
	3: "lost",
	4: "closed",
}

var Bid_State_value = map[string]int32{
	"invalid": 0,
	"open":    1,
	"matched": 2,
	"lost":    3,
	"closed":  4,
}

func (x Bid_State) String() string {
	return proto.EnumName(Bid_State_name, int32(x))
}

func (Bid_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_057fd80e533b030c, []int{3, 0}
}

// MsgCreateBid defines an SDK message for creating Bid
type MsgCreateBid struct {
	Order    OrderID    `protobuf:"bytes,1,opt,name=order,proto3" json:"order" yaml:"order"`
	Provider string     `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider" yaml:"provider"`
	Price    types.Coin `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
}

func (m *MsgCreateBid) Reset()         { *m = MsgCreateBid{} }
func (m *MsgCreateBid) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBid) ProtoMessage()    {}
func (*MsgCreateBid) Descriptor() ([]byte, []int) {
	return fileDescriptor_057fd80e533b030c, []int{0}
}
func (m *MsgCreateBid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBid.Merge(m, src)
}
func (m *MsgCreateBid) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBid) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBid.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBid proto.InternalMessageInfo

func (m *MsgCreateBid) GetOrder() OrderID {
	if m != nil {
		return m.Order
	}
	return OrderID{}
}

func (m *MsgCreateBid) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *MsgCreateBid) GetPrice() types.Coin {
	if m != nil {
		return m.Price
	}
	return types.Coin{}
}

// MsgCloseBid defines an SDK message for closing bid
type MsgCloseBid struct {
	BidID BidID `protobuf:"bytes,1,opt,name=bid_id,json=bidId,proto3" json:"id" yaml:"id"`
}

func (m *MsgCloseBid) Reset()         { *m = MsgCloseBid{} }
func (m *MsgCloseBid) String() string { return proto.CompactTextString(m) }
func (*MsgCloseBid) ProtoMessage()    {}
func (*MsgCloseBid) Descriptor() ([]byte, []int) {
	return fileDescriptor_057fd80e533b030c, []int{1}
}
func (m *MsgCloseBid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCloseBid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCloseBid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCloseBid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCloseBid.Merge(m, src)
}
func (m *MsgCloseBid) XXX_Size() int {
	return m.Size()
}
func (m *MsgCloseBid) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCloseBid.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCloseBid proto.InternalMessageInfo

func (m *MsgCloseBid) GetBidID() BidID {
	if m != nil {
		return m.BidID
	}
	return BidID{}
}

// BidID stores owner and all other seq numbers
// A successful bid becomes a Lease(ID).
type BidID struct {
	Owner    string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	DSeq     uint64 `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	GSeq     uint32 `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
	OSeq     uint32 `protobuf:"varint,4,opt,name=oseq,proto3" json:"oseq" yaml:"oseq"`
	Provider string `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider" yaml:"provider"`
}

func (m *BidID) Reset()      { *m = BidID{} }
func (*BidID) ProtoMessage() {}
func (*BidID) Descriptor() ([]byte, []int) {
	return fileDescriptor_057fd80e533b030c, []int{2}
}
func (m *BidID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BidID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BidID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BidID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidID.Merge(m, src)
}
func (m *BidID) XXX_Size() int {
	return m.Size()
}
func (m *BidID) XXX_DiscardUnknown() {
	xxx_messageInfo_BidID.DiscardUnknown(m)
}

var xxx_messageInfo_BidID proto.InternalMessageInfo

func (m *BidID) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *BidID) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *BidID) GetGSeq() uint32 {
	if m != nil {
		return m.GSeq
	}
	return 0
}

func (m *BidID) GetOSeq() uint32 {
	if m != nil {
		return m.OSeq
	}
	return 0
}

func (m *BidID) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

// Bid stores BidID, state of bid and price
type Bid struct {
	BidID BidID      `protobuf:"bytes,1,opt,name=bid_id,json=bidId,proto3" json:"id" yaml:"id"`
	State Bid_State  `protobuf:"varint,2,opt,name=state,proto3,enum=akash.market.v1beta1.Bid_State" json:"state" yaml:"state"`
	Price types.Coin `protobuf:"bytes,3,opt,name=price,proto3" json:"price" yaml:"price"`
}

func (m *Bid) Reset()      { *m = Bid{} }
func (*Bid) ProtoMessage() {}
func (*Bid) Descriptor() ([]byte, []int) {
	return fileDescriptor_057fd80e533b030c, []int{3}
}
func (m *Bid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bid.Merge(m, src)
}
func (m *Bid) XXX_Size() int {
	return m.Size()
}
func (m *Bid) XXX_DiscardUnknown() {
	xxx_messageInfo_Bid.DiscardUnknown(m)
}

var xxx_messageInfo_Bid proto.InternalMessageInfo

func (m *Bid) GetBidID() BidID {
	if m != nil {
		return m.BidID
	}
	return BidID{}
}

func (m *Bid) GetState() Bid_State {
	if m != nil {
		return m.State
	}
	return BidStateInvalid
}

func (m *Bid) GetPrice() types.Coin {
	if m != nil {
		return m.Price
	}
	return types.Coin{}
}

// BidFilters defines flags for bid list filter
type BidFilters struct {
	Owner    string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	DSeq     uint64 `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	GSeq     uint32 `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
	OSeq     uint32 `protobuf:"varint,4,opt,name=oseq,proto3" json:"oseq" yaml:"oseq"`
	Provider string `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider" yaml:"provider"`
	State    string `protobuf:"bytes,6,opt,name=state,proto3" json:"state" yaml:"state"`
}

func (m *BidFilters) Reset()         { *m = BidFilters{} }
func (m *BidFilters) String() string { return proto.CompactTextString(m) }
func (*BidFilters) ProtoMessage()    {}
func (*BidFilters) Descriptor() ([]byte, []int) {
	return fileDescriptor_057fd80e533b030c, []int{4}
}
func (m *BidFilters) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BidFilters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BidFilters.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BidFilters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidFilters.Merge(m, src)
}
func (m *BidFilters) XXX_Size() int {
	return m.Size()
}
func (m *BidFilters) XXX_DiscardUnknown() {
	xxx_messageInfo_BidFilters.DiscardUnknown(m)
}

var xxx_messageInfo_BidFilters proto.InternalMessageInfo

func (m *BidFilters) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *BidFilters) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *BidFilters) GetGSeq() uint32 {
	if m != nil {
		return m.GSeq
	}
	return 0
}

func (m *BidFilters) GetOSeq() uint32 {
	if m != nil {
		return m.OSeq
	}
	return 0
}

func (m *BidFilters) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *BidFilters) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterEnum("akash.market.v1beta1.Bid_State", Bid_State_name, Bid_State_value)
	proto.RegisterType((*MsgCreateBid)(nil), "akash.market.v1beta1.MsgCreateBid")
	proto.RegisterType((*MsgCloseBid)(nil), "akash.market.v1beta1.MsgCloseBid")
	proto.RegisterType((*BidID)(nil), "akash.market.v1beta1.BidID")
	proto.RegisterType((*Bid)(nil), "akash.market.v1beta1.Bid")
	proto.RegisterType((*BidFilters)(nil), "akash.market.v1beta1.BidFilters")
}

func init() { proto.RegisterFile("akash/market/v1beta1/bid.proto", fileDescriptor_057fd80e533b030c) }

var fileDescriptor_057fd80e533b030c = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x94, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xed, 0xc4, 0x4e, 0x9b, 0x4b, 0x7f, 0x44, 0xa6, 0x48, 0xad, 0xab, 0xfa, 0x2c, 0x0f,
	0x55, 0x27, 0x5b, 0x6d, 0xb7, 0xb2, 0xa0, 0x6b, 0x05, 0x8a, 0x44, 0x29, 0x4a, 0x99, 0x60, 0x40,
	0xb6, 0xef, 0xe4, 0x9e, 0x9a, 0xe4, 0x52, 0xfb, 0x08, 0xf4, 0x3f, 0x40, 0x99, 0x58, 0x90, 0x58,
	0x02, 0x95, 0xf8, 0x67, 0x3a, 0x76, 0x64, 0xb2, 0x50, 0xba, 0xa0, 0x8c, 0xf9, 0x03, 0x10, 0xba,
	0x3b, 0x27, 0x69, 0xa5, 0x82, 0x84, 0x04, 0x1b, 0x93, 0xf3, 0xbe, 0xef, 0x7d, 0xee, 0xde, 0x8f,
	0xcb, 0x03, 0x4e, 0x78, 0x1a, 0x66, 0x27, 0x41, 0x3b, 0x4c, 0x4f, 0x09, 0x0f, 0x7a, 0xdb, 0x11,
	0xe1, 0xe1, 0x76, 0x10, 0x51, 0xec, 0x77, 0x53, 0xc6, 0x99, 0xb5, 0x22, 0xfd, 0xbe, 0xf2, 0xfb,
	0x85, 0xdf, 0x5e, 0x49, 0x58, 0xc2, 0x64, 0x40, 0x20, 0x7e, 0xa9, 0x58, 0xdb, 0xbd, 0xf3, 0x2c,
	0x96, 0x62, 0x92, 0x16, 0x11, 0x4e, 0xcc, 0xb2, 0x36, 0xcb, 0x82, 0x28, 0xcc, 0xc8, 0x34, 0x20,
	0x66, 0xb4, 0xa3, 0xfc, 0xde, 0x0f, 0x1d, 0x2c, 0x1c, 0x66, 0xc9, 0x7e, 0x4a, 0x42, 0x4e, 0x10,
	0xc5, 0xd6, 0x4b, 0x60, 0x4a, 0x7e, 0x55, 0x77, 0xf5, 0xad, 0xda, 0xce, 0x86, 0x7f, 0x57, 0x3a,
	0xfe, 0x91, 0x08, 0x69, 0x1c, 0xa0, 0xcd, 0xcb, 0x1c, 0x6a, 0xc3, 0x1c, 0x9a, 0x52, 0x18, 0xe5,
	0x50, 0xc1, 0xe3, 0x1c, 0x2e, 0x9c, 0x87, 0xed, 0xd6, 0x9e, 0x27, 0x4d, 0xaf, 0xa9, 0x64, 0xeb,
	0x01, 0x98, 0xef, 0xa6, 0xac, 0x47, 0xc5, 0xf9, 0x25, 0x57, 0xdf, 0xaa, 0x22, 0x38, 0xca, 0xe1,
	0x54, 0x1b, 0xe7, 0x70, 0x59, 0x61, 0x13, 0xc5, 0x6b, 0x4e, 0x9d, 0xd6, 0x53, 0x60, 0x76, 0x53,
	0x1a, 0x93, 0xd5, 0xb2, 0xcc, 0x6c, 0xcd, 0x57, 0xa5, 0xf9, 0xa2, 0xb4, 0x69, 0x62, 0xfb, 0x8c,
	0x76, 0xd0, 0x86, 0xc8, 0x4a, 0x24, 0x23, 0xe3, 0x67, 0xc9, 0x48, 0xd3, 0x6b, 0x2a, 0x79, 0xcf,
	0xf8, 0x7e, 0x01, 0x35, 0x8f, 0x82, 0x9a, 0xa8, 0xbf, 0xc5, 0x32, 0x59, 0xfe, 0x73, 0x50, 0x89,
	0x28, 0x7e, 0x45, 0x71, 0x51, 0xff, 0xfa, 0xdd, 0xf5, 0x23, 0x8a, 0x1b, 0x07, 0xc8, 0x9d, 0x54,
	0x2f, 0xcd, 0x51, 0x0e, 0x4b, 0x14, 0x8f, 0x73, 0x58, 0x55, 0xb7, 0x51, 0xec, 0x35, 0xcd, 0x88,
	0xe2, 0x06, 0x2e, 0xae, 0xfa, 0x54, 0x02, 0x2a, 0xd2, 0x0a, 0x80, 0xc9, 0xde, 0x74, 0x8a, 0x26,
	0x57, 0xd1, 0x9a, 0x6c, 0x9c, 0x10, 0x6e, 0x34, 0x4e, 0x98, 0xa2, 0x71, 0xe2, 0x6b, 0xed, 0x02,
	0x03, 0x67, 0xe4, 0x4c, 0x36, 0xcd, 0x40, 0x70, 0x98, 0x43, 0xe3, 0xe0, 0x98, 0x9c, 0x8d, 0x72,
	0x28, 0xf5, 0x71, 0x0e, 0x6b, 0x0a, 0x13, 0x96, 0xd7, 0x94, 0xa2, 0x80, 0x12, 0x01, 0x89, 0x7e,
	0x2d, 0x2a, 0xe8, 0x71, 0x01, 0x25, 0xb7, 0xa0, 0x44, 0x41, 0x49, 0x01, 0x31, 0x01, 0x19, 0x33,
	0xe8, 0xa8, 0x80, 0xd8, 0x2d, 0x88, 0x29, 0x48, 0x7c, 0x6e, 0xcd, 0xd5, 0xfc, 0xc3, 0xb9, 0xee,
	0xcd, 0x7f, 0xbc, 0x80, 0x9a, 0x6c, 0xd0, 0xe7, 0x32, 0x28, 0xff, 0xb3, 0x21, 0x58, 0xcf, 0x80,
	0x99, 0xf1, 0x90, 0x13, 0xd9, 0xc4, 0xa5, 0x1d, 0xf8, 0xcb, 0x43, 0xfd, 0x63, 0x11, 0xa6, 0xa6,
	0x22, 0x89, 0xd9, 0x54, 0xa4, 0xe9, 0x35, 0x95, 0xfc, 0xb7, 0x5f, 0xa4, 0xf7, 0x41, 0x07, 0xa6,
	0xbc, 0xdb, 0x72, 0xc1, 0x1c, 0xed, 0xf4, 0xc2, 0x16, 0xc5, 0x75, 0xcd, 0xbe, 0xd7, 0x1f, 0xb8,
	0xcb, 0x88, 0x62, 0xe9, 0x6a, 0x28, 0xd9, 0xba, 0x0f, 0x0c, 0xd6, 0x25, 0x9d, 0xba, 0x6e, 0xd7,
	0xfa, 0x03, 0x77, 0x0e, 0x51, 0x7c, 0xd4, 0x25, 0x1d, 0x6b, 0x1d, 0xcc, 0xb5, 0x43, 0x1e, 0x9f,
	0x10, 0x5c, 0x2f, 0xd9, 0x4b, 0xfd, 0x81, 0x0b, 0x10, 0xc5, 0x87, 0x4a, 0x11, 0x4c, 0x8b, 0x65,
	0xbc, 0x5e, 0x9e, 0x32, 0x4f, 0x58, 0xc6, 0xad, 0x35, 0x50, 0x89, 0xc5, 0xfb, 0xc7, 0x75, 0xc3,
	0x5e, 0xec, 0x0f, 0xdc, 0x2a, 0xa2, 0x58, 0xfe, 0x21, 0xb0, 0x6d, 0xbc, 0xfb, 0xe2, 0x68, 0x37,
	0x26, 0x74, 0x55, 0x02, 0xe2, 0xc0, 0x47, 0xb4, 0xc5, 0x49, 0x9a, 0xfd, 0x7f, 0xc7, 0x37, 0xf7,
	0x53, 0x30, 0x79, 0x5f, 0x95, 0x59, 0x33, 0x7e, 0xf7, 0x7c, 0xd4, 0x56, 0x40, 0x0f, 0x2f, 0x87,
	0x8e, 0x7e, 0x35, 0x74, 0xf4, 0x6f, 0x43, 0x47, 0x7f, 0x7f, 0xed, 0x68, 0x57, 0xd7, 0x8e, 0xf6,
	0xf5, 0xda, 0xd1, 0x5e, 0x6c, 0x26, 0x94, 0x9f, 0xbc, 0x8e, 0xfc, 0x98, 0xb5, 0x03, 0xd6, 0x4b,
	0xe3, 0xd6, 0x69, 0xa0, 0xf6, 0xfd, 0xdb, 0xc9, 0xc6, 0xe7, 0xe7, 0x5d, 0x92, 0x45, 0x15, 0xb9,
	0xca, 0x77, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x33, 0x30, 0xee, 0x93, 0x5a, 0x06, 0x00, 0x00,
}

func (m *MsgCreateBid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintBid(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Order.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgCloseBid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCloseBid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCloseBid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BidID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BidID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BidID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BidID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintBid(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x2a
	}
	if m.OSeq != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.OSeq))
		i--
		dAtA[i] = 0x20
	}
	if m.GSeq != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.GSeq))
		i--
		dAtA[i] = 0x18
	}
	if m.DSeq != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintBid(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Bid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.State != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.BidID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BidFilters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BidFilters) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BidFilters) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintBid(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintBid(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x2a
	}
	if m.OSeq != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.OSeq))
		i--
		dAtA[i] = 0x20
	}
	if m.GSeq != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.GSeq))
		i--
		dAtA[i] = 0x18
	}
	if m.DSeq != 0 {
		i = encodeVarintBid(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintBid(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBid(dAtA []byte, offset int, v uint64) int {
	offset -= sovBid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateBid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Order.Size()
	n += 1 + l + sovBid(uint64(l))
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovBid(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovBid(uint64(l))
	return n
}

func (m *MsgCloseBid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BidID.Size()
	n += 1 + l + sovBid(uint64(l))
	return n
}

func (m *BidID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovBid(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovBid(uint64(m.DSeq))
	}
	if m.GSeq != 0 {
		n += 1 + sovBid(uint64(m.GSeq))
	}
	if m.OSeq != 0 {
		n += 1 + sovBid(uint64(m.OSeq))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovBid(uint64(l))
	}
	return n
}

func (m *Bid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BidID.Size()
	n += 1 + l + sovBid(uint64(l))
	if m.State != 0 {
		n += 1 + sovBid(uint64(m.State))
	}
	l = m.Price.Size()
	n += 1 + l + sovBid(uint64(l))
	return n
}

func (m *BidFilters) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovBid(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovBid(uint64(m.DSeq))
	}
	if m.GSeq != 0 {
		n += 1 + sovBid(uint64(m.GSeq))
	}
	if m.OSeq != 0 {
		n += 1 + sovBid(uint64(m.OSeq))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovBid(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovBid(uint64(l))
	}
	return n
}

func sovBid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBid(x uint64) (n int) {
	return sovBid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateBid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateBid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Order.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCloseBid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCloseBid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCloseBid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BidID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BidID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BidID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BidID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GSeq", wireType)
			}
			m.GSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSeq", wireType)
			}
			m.OSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Bid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Bid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BidID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Bid_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BidFilters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BidFilters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BidFilters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GSeq", wireType)
			}
			m.GSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSeq", wireType)
			}
			m.OSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBid
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBid
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBid
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthBid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBid = fmt.Errorf("proto: unexpected end of group")
)