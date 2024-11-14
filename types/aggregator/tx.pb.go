// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: halo/aggregator/v2/tx.proto

package aggregator

import (
	context "context"
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgTransmit implements the transmit (0xTODO) method.
type MsgTransmit struct {
	Signer    string                `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Answer    cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=answer,proto3,customtype=cosmossdk.io/math.Int" json:"answer"`
	UpdatedAt uint32                `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (m *MsgTransmit) Reset()         { *m = MsgTransmit{} }
func (m *MsgTransmit) String() string { return proto.CompactTextString(m) }
func (*MsgTransmit) ProtoMessage()    {}
func (*MsgTransmit) Descriptor() ([]byte, []int) {
	return fileDescriptor_d54c0eebbffa764b, []int{0}
}
func (m *MsgTransmit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransmit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransmit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransmit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransmit.Merge(m, src)
}
func (m *MsgTransmit) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransmit) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransmit.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransmit proto.InternalMessageInfo

// MsgTransmitResponse ...
type MsgTransmitResponse struct {
	RoundId uint64 `protobuf:"varint,1,opt,name=round_id,json=roundId,proto3" json:"round_id,omitempty"`
}

func (m *MsgTransmitResponse) Reset()         { *m = MsgTransmitResponse{} }
func (m *MsgTransmitResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTransmitResponse) ProtoMessage()    {}
func (*MsgTransmitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d54c0eebbffa764b, []int{1}
}
func (m *MsgTransmitResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransmitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransmitResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransmitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransmitResponse.Merge(m, src)
}
func (m *MsgTransmitResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransmitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransmitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransmitResponse proto.InternalMessageInfo

func (m *MsgTransmitResponse) GetRoundId() uint64 {
	if m != nil {
		return m.RoundId
	}
	return 0
}

// MsgSetNextPrice implements the setNextPrice (0xfeca6988) method.
type MsgSetNextPrice struct {
	Signer    string                `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	NextPrice cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=next_price,json=nextPrice,proto3,customtype=cosmossdk.io/math.Int" json:"next_price"`
}

func (m *MsgSetNextPrice) Reset()         { *m = MsgSetNextPrice{} }
func (m *MsgSetNextPrice) String() string { return proto.CompactTextString(m) }
func (*MsgSetNextPrice) ProtoMessage()    {}
func (*MsgSetNextPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_d54c0eebbffa764b, []int{2}
}
func (m *MsgSetNextPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetNextPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetNextPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetNextPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetNextPrice.Merge(m, src)
}
func (m *MsgSetNextPrice) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetNextPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetNextPrice.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetNextPrice proto.InternalMessageInfo

// MsgSetNextPriceResponse ...
type MsgSetNextPriceResponse struct {
}

func (m *MsgSetNextPriceResponse) Reset()         { *m = MsgSetNextPriceResponse{} }
func (m *MsgSetNextPriceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetNextPriceResponse) ProtoMessage()    {}
func (*MsgSetNextPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d54c0eebbffa764b, []int{3}
}
func (m *MsgSetNextPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetNextPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetNextPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetNextPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetNextPriceResponse.Merge(m, src)
}
func (m *MsgSetNextPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetNextPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetNextPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetNextPriceResponse proto.InternalMessageInfo

// MsgTransferOwnership implements the transferOwnership (0xf2fde38b) method.
type MsgTransferOwnership struct {
	Signer      string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	NewReporter string `protobuf:"bytes,2,opt,name=new_reporter,json=newReporter,proto3" json:"new_reporter,omitempty"`
}

func (m *MsgTransferOwnership) Reset()         { *m = MsgTransferOwnership{} }
func (m *MsgTransferOwnership) String() string { return proto.CompactTextString(m) }
func (*MsgTransferOwnership) ProtoMessage()    {}
func (*MsgTransferOwnership) Descriptor() ([]byte, []int) {
	return fileDescriptor_d54c0eebbffa764b, []int{4}
}
func (m *MsgTransferOwnership) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferOwnership) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferOwnership.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferOwnership) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferOwnership.Merge(m, src)
}
func (m *MsgTransferOwnership) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferOwnership) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferOwnership.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferOwnership proto.InternalMessageInfo

// MsgTransferOwnershipResponse ...
type MsgTransferOwnershipResponse struct {
}

func (m *MsgTransferOwnershipResponse) Reset()         { *m = MsgTransferOwnershipResponse{} }
func (m *MsgTransferOwnershipResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTransferOwnershipResponse) ProtoMessage()    {}
func (*MsgTransferOwnershipResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d54c0eebbffa764b, []int{5}
}
func (m *MsgTransferOwnershipResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferOwnershipResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferOwnershipResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferOwnershipResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferOwnershipResponse.Merge(m, src)
}
func (m *MsgTransferOwnershipResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferOwnershipResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferOwnershipResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferOwnershipResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgTransmit)(nil), "halo.aggregator.v2.MsgTransmit")
	proto.RegisterType((*MsgTransmitResponse)(nil), "halo.aggregator.v2.MsgTransmitResponse")
	proto.RegisterType((*MsgSetNextPrice)(nil), "halo.aggregator.v2.MsgSetNextPrice")
	proto.RegisterType((*MsgSetNextPriceResponse)(nil), "halo.aggregator.v2.MsgSetNextPriceResponse")
	proto.RegisterType((*MsgTransferOwnership)(nil), "halo.aggregator.v2.MsgTransferOwnership")
	proto.RegisterType((*MsgTransferOwnershipResponse)(nil), "halo.aggregator.v2.MsgTransferOwnershipResponse")
}

func init() { proto.RegisterFile("halo/aggregator/v2/tx.proto", fileDescriptor_d54c0eebbffa764b) }

var fileDescriptor_d54c0eebbffa764b = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xb6, 0x5b, 0x28, 0xcd, 0xb5, 0x08, 0xd5, 0x04, 0x35, 0x31, 0xc5, 0x29, 0x66, 0x20, 0x0a,
	0x8a, 0x6d, 0x02, 0x0b, 0x61, 0x6a, 0x27, 0x82, 0x14, 0x8a, 0xdc, 0x4e, 0x2c, 0xe1, 0x12, 0x1f,
	0x17, 0x8b, 0xfa, 0xce, 0xba, 0xbb, 0xfc, 0x60, 0x43, 0x4c, 0x88, 0x89, 0x3f, 0xa1, 0x23, 0x63,
	0x86, 0xb2, 0x33, 0x56, 0x4c, 0x55, 0x27, 0xd4, 0xa1, 0x42, 0xc9, 0x10, 0xfe, 0x0c, 0xe4, 0x5f,
	0xc1, 0xa4, 0x0d, 0x54, 0x59, 0x6c, 0xdf, 0xbd, 0xef, 0xbb, 0xf7, 0xbe, 0xf7, 0x3e, 0x1f, 0xb8,
	0xdd, 0x86, 0xfb, 0xd4, 0x84, 0x18, 0x33, 0x84, 0xa1, 0xa0, 0xcc, 0xec, 0x56, 0x4c, 0xd1, 0x37,
	0x7c, 0x46, 0x05, 0x55, 0x94, 0x20, 0x68, 0xfc, 0x09, 0x1a, 0xdd, 0x8a, 0xba, 0x06, 0x3d, 0x97,
	0x50, 0x33, 0x7c, 0x46, 0x30, 0x75, 0xbd, 0x45, 0xb9, 0x47, 0xb9, 0xe9, 0x71, 0x6c, 0x76, 0x1f,
	0x06, 0xaf, 0x38, 0x90, 0x8f, 0x02, 0x8d, 0x70, 0x65, 0x46, 0x8b, 0x38, 0x94, 0xc5, 0x14, 0xd3,
	0x68, 0x3f, 0xf8, 0x8a, 0x76, 0xf5, 0x53, 0x19, 0xac, 0xd4, 0x39, 0xde, 0x63, 0x90, 0x70, 0xcf,
	0x15, 0x8a, 0x05, 0x96, 0xb8, 0x8b, 0x09, 0x62, 0x39, 0x79, 0x53, 0x2e, 0x66, 0xb6, 0x73, 0x27,
	0x87, 0xe5, 0x6c, 0x7c, 0xce, 0x96, 0xe3, 0x30, 0xc4, 0xf9, 0xae, 0x60, 0x2e, 0xc1, 0x76, 0x8c,
	0x53, 0x9e, 0x81, 0x25, 0x48, 0x78, 0x0f, 0xb1, 0xdc, 0x42, 0xc8, 0xb0, 0x8e, 0xce, 0x0a, 0xd2,
	0xe9, 0x59, 0xe1, 0x56, 0xc4, 0xe2, 0xce, 0x5b, 0xc3, 0xa5, 0xa6, 0x07, 0x45, 0xdb, 0xa8, 0x11,
	0x71, 0x72, 0x58, 0x06, 0xf1, 0x71, 0x35, 0x22, 0xbe, 0x8c, 0x07, 0x25, 0xd9, 0x8e, 0xf9, 0xca,
	0x1d, 0x00, 0x3a, 0xbe, 0x03, 0x05, 0x72, 0x1a, 0x50, 0xe4, 0x16, 0x37, 0xe5, 0xe2, 0x75, 0x3b,
	0x13, 0xef, 0x6c, 0x89, 0xaa, 0xf5, 0xf1, 0xa0, 0x20, 0xfd, 0x3a, 0x28, 0x48, 0x1f, 0xc6, 0x83,
	0x52, 0x9c, 0xfd, 0xd3, 0x78, 0x50, 0xca, 0x4d, 0x37, 0x34, 0x11, 0xa3, 0x5b, 0xe0, 0x66, 0x4a,
	0x9b, 0x8d, 0xb8, 0x4f, 0x09, 0x47, 0x4a, 0x1e, 0x2c, 0x33, 0xda, 0x21, 0x4e, 0xc3, 0x75, 0x42,
	0x95, 0x57, 0xec, 0x6b, 0xe1, 0xba, 0xe6, 0xe8, 0xdf, 0x65, 0x70, 0xa3, 0xce, 0xf1, 0x2e, 0x12,
	0x2f, 0x50, 0x5f, 0xbc, 0x64, 0x6e, 0x0b, 0xcd, 0xd1, 0x92, 0x1d, 0x00, 0x08, 0xea, 0x8b, 0x86,
	0x1f, 0xf0, 0xe7, 0x6e, 0x4b, 0x86, 0x24, 0x25, 0x54, 0x1f, 0xcf, 0x90, 0xbe, 0x31, 0x2d, 0x3d,
	0x5d, 0xb8, 0x9e, 0x07, 0xeb, 0x53, 0x5a, 0x92, 0x16, 0xe8, 0xdf, 0x64, 0x90, 0x4d, 0x5a, 0xf3,
	0x06, 0xb1, 0x9d, 0x1e, 0x41, 0x8c, 0xb7, 0x5d, 0x7f, 0x0e, 0xb1, 0x4f, 0xc1, 0x2a, 0x41, 0xbd,
	0x06, 0x43, 0x3e, 0x65, 0x62, 0xe2, 0x82, 0xd9, 0xbc, 0x15, 0x82, 0x7a, 0x76, 0x0c, 0xae, 0x3e,
	0x99, 0x21, 0xec, 0xee, 0x85, 0x33, 0x4d, 0x57, 0xaa, 0x6b, 0x60, 0xe3, 0x22, 0x05, 0x89, 0xc4,
	0xca, 0xd7, 0x05, 0xb0, 0x58, 0xe7, 0x58, 0xd9, 0x03, 0xcb, 0x13, 0x77, 0x17, 0x8c, 0xf3, 0xff,
	0x97, 0x91, 0xb2, 0x88, 0x7a, 0xff, 0x3f, 0x80, 0x89, 0x87, 0x5e, 0x83, 0xd5, 0xbf, 0x4c, 0x72,
	0x6f, 0x06, 0x31, 0x0d, 0x52, 0x1f, 0x5c, 0x02, 0x34, 0xc9, 0x40, 0xc1, 0xda, 0xf9, 0xf1, 0x14,
	0xff, 0x55, 0x5f, 0x1a, 0xa9, 0x5a, 0x97, 0x45, 0x26, 0x09, 0xd5, 0xab, 0xef, 0x03, 0xdb, 0x6d,
	0x3f, 0x3f, 0x1a, 0x6a, 0xf2, 0xf1, 0x50, 0x93, 0x7f, 0x0e, 0x35, 0xf9, 0xf3, 0x48, 0x93, 0x8e,
	0x47, 0x9a, 0xf4, 0x63, 0xa4, 0x49, 0xaf, 0x2c, 0xec, 0x8a, 0x76, 0xa7, 0x69, 0xb4, 0xa8, 0x67,
	0x12, 0xda, 0xdc, 0x47, 0x65, 0xc8, 0x39, 0x12, 0xdc, 0x0c, 0x87, 0x15, 0x5c, 0x63, 0xef, 0x7c,
	0xc4, 0x53, 0x53, 0x6b, 0x2e, 0x85, 0x97, 0xcc, 0xa3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x82,
	0xf6, 0xce, 0x3a, 0xf4, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Transmit(ctx context.Context, in *MsgTransmit, opts ...grpc.CallOption) (*MsgTransmitResponse, error)
	SetNextPrice(ctx context.Context, in *MsgSetNextPrice, opts ...grpc.CallOption) (*MsgSetNextPriceResponse, error)
	TransferOwnership(ctx context.Context, in *MsgTransferOwnership, opts ...grpc.CallOption) (*MsgTransferOwnershipResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Transmit(ctx context.Context, in *MsgTransmit, opts ...grpc.CallOption) (*MsgTransmitResponse, error) {
	out := new(MsgTransmitResponse)
	err := c.cc.Invoke(ctx, "/halo.aggregator.v2.Msg/Transmit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetNextPrice(ctx context.Context, in *MsgSetNextPrice, opts ...grpc.CallOption) (*MsgSetNextPriceResponse, error) {
	out := new(MsgSetNextPriceResponse)
	err := c.cc.Invoke(ctx, "/halo.aggregator.v2.Msg/SetNextPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferOwnership(ctx context.Context, in *MsgTransferOwnership, opts ...grpc.CallOption) (*MsgTransferOwnershipResponse, error) {
	out := new(MsgTransferOwnershipResponse)
	err := c.cc.Invoke(ctx, "/halo.aggregator.v2.Msg/TransferOwnership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Transmit(context.Context, *MsgTransmit) (*MsgTransmitResponse, error)
	SetNextPrice(context.Context, *MsgSetNextPrice) (*MsgSetNextPriceResponse, error)
	TransferOwnership(context.Context, *MsgTransferOwnership) (*MsgTransferOwnershipResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Transmit(ctx context.Context, req *MsgTransmit) (*MsgTransmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transmit not implemented")
}
func (*UnimplementedMsgServer) SetNextPrice(ctx context.Context, req *MsgSetNextPrice) (*MsgSetNextPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNextPrice not implemented")
}
func (*UnimplementedMsgServer) TransferOwnership(ctx context.Context, req *MsgTransferOwnership) (*MsgTransferOwnershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferOwnership not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Transmit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransmit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Transmit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halo.aggregator.v2.Msg/Transmit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Transmit(ctx, req.(*MsgTransmit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetNextPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetNextPrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetNextPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halo.aggregator.v2.Msg/SetNextPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetNextPrice(ctx, req.(*MsgSetNextPrice))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferOwnership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferOwnership)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferOwnership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halo.aggregator.v2.Msg/TransferOwnership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferOwnership(ctx, req.(*MsgTransferOwnership))
	}
	return interceptor(ctx, in, info, handler)
}

var Msg_serviceDesc = _Msg_serviceDesc
var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "halo.aggregator.v2.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transmit",
			Handler:    _Msg_Transmit_Handler,
		},
		{
			MethodName: "SetNextPrice",
			Handler:    _Msg_SetNextPrice_Handler,
		},
		{
			MethodName: "TransferOwnership",
			Handler:    _Msg_TransferOwnership_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "halo/aggregator/v2/tx.proto",
}

func (m *MsgTransmit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransmit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransmit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdatedAt != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Answer.Size()
		i -= size
		if _, err := m.Answer.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgTransmitResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransmitResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransmitResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RoundId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.RoundId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetNextPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetNextPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetNextPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.NextPrice.Size()
		i -= size
		if _, err := m.NextPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetNextPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetNextPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetNextPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgTransferOwnership) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferOwnership) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferOwnership) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewReporter) > 0 {
		i -= len(m.NewReporter)
		copy(dAtA[i:], m.NewReporter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.NewReporter)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgTransferOwnershipResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferOwnershipResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferOwnershipResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgTransmit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Answer.Size()
	n += 1 + l + sovTx(uint64(l))
	if m.UpdatedAt != 0 {
		n += 1 + sovTx(uint64(m.UpdatedAt))
	}
	return n
}

func (m *MsgTransmitResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RoundId != 0 {
		n += 1 + sovTx(uint64(m.RoundId))
	}
	return n
}

func (m *MsgSetNextPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.NextPrice.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSetNextPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgTransferOwnership) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.NewReporter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgTransferOwnershipResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgTransmit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTransmit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransmit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Answer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Answer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgTransmitResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTransmitResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransmitResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoundId", wireType)
			}
			m.RoundId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoundId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetNextPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetNextPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetNextPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NextPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetNextPriceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetNextPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetNextPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgTransferOwnership) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTransferOwnership: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferOwnership: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewReporter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewReporter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgTransferOwnershipResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTransferOwnershipResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferOwnershipResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
