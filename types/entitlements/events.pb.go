// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: halo/entitlements/v1/events.proto

package entitlements

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type PublicCapabilityUpdated struct {
	Method  string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Enabled bool   `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (m *PublicCapabilityUpdated) Reset()         { *m = PublicCapabilityUpdated{} }
func (m *PublicCapabilityUpdated) String() string { return proto.CompactTextString(m) }
func (*PublicCapabilityUpdated) ProtoMessage()    {}
func (*PublicCapabilityUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_f721b73b2cc60687, []int{0}
}
func (m *PublicCapabilityUpdated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PublicCapabilityUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PublicCapabilityUpdated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PublicCapabilityUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicCapabilityUpdated.Merge(m, src)
}
func (m *PublicCapabilityUpdated) XXX_Size() int {
	return m.Size()
}
func (m *PublicCapabilityUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicCapabilityUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_PublicCapabilityUpdated proto.InternalMessageInfo

func (m *PublicCapabilityUpdated) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *PublicCapabilityUpdated) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type RoleCapabilityUpdated struct {
	Role    Role   `protobuf:"varint,1,opt,name=role,proto3,enum=halo.entitlements.v1.Role" json:"role,omitempty"`
	Method  string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Enabled bool   `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (m *RoleCapabilityUpdated) Reset()         { *m = RoleCapabilityUpdated{} }
func (m *RoleCapabilityUpdated) String() string { return proto.CompactTextString(m) }
func (*RoleCapabilityUpdated) ProtoMessage()    {}
func (*RoleCapabilityUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_f721b73b2cc60687, []int{1}
}
func (m *RoleCapabilityUpdated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoleCapabilityUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoleCapabilityUpdated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RoleCapabilityUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleCapabilityUpdated.Merge(m, src)
}
func (m *RoleCapabilityUpdated) XXX_Size() int {
	return m.Size()
}
func (m *RoleCapabilityUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleCapabilityUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_RoleCapabilityUpdated proto.InternalMessageInfo

func (m *RoleCapabilityUpdated) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return ROLE_UNSPECIFIED
}

func (m *RoleCapabilityUpdated) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *RoleCapabilityUpdated) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type UserRoleUpdated struct {
	User    string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Role    Role   `protobuf:"varint,2,opt,name=role,proto3,enum=halo.entitlements.v1.Role" json:"role,omitempty"`
	Enabled bool   `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (m *UserRoleUpdated) Reset()         { *m = UserRoleUpdated{} }
func (m *UserRoleUpdated) String() string { return proto.CompactTextString(m) }
func (*UserRoleUpdated) ProtoMessage()    {}
func (*UserRoleUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_f721b73b2cc60687, []int{2}
}
func (m *UserRoleUpdated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserRoleUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserRoleUpdated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserRoleUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRoleUpdated.Merge(m, src)
}
func (m *UserRoleUpdated) XXX_Size() int {
	return m.Size()
}
func (m *UserRoleUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRoleUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_UserRoleUpdated proto.InternalMessageInfo

func (m *UserRoleUpdated) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserRoleUpdated) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return ROLE_UNSPECIFIED
}

func (m *UserRoleUpdated) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type Paused struct {
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *Paused) Reset()         { *m = Paused{} }
func (m *Paused) String() string { return proto.CompactTextString(m) }
func (*Paused) ProtoMessage()    {}
func (*Paused) Descriptor() ([]byte, []int) {
	return fileDescriptor_f721b73b2cc60687, []int{3}
}
func (m *Paused) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Paused) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Paused.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Paused) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paused.Merge(m, src)
}
func (m *Paused) XXX_Size() int {
	return m.Size()
}
func (m *Paused) XXX_DiscardUnknown() {
	xxx_messageInfo_Paused.DiscardUnknown(m)
}

var xxx_messageInfo_Paused proto.InternalMessageInfo

func (m *Paused) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type Unpaused struct {
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *Unpaused) Reset()         { *m = Unpaused{} }
func (m *Unpaused) String() string { return proto.CompactTextString(m) }
func (*Unpaused) ProtoMessage()    {}
func (*Unpaused) Descriptor() ([]byte, []int) {
	return fileDescriptor_f721b73b2cc60687, []int{4}
}
func (m *Unpaused) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Unpaused) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Unpaused.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Unpaused) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unpaused.Merge(m, src)
}
func (m *Unpaused) XXX_Size() int {
	return m.Size()
}
func (m *Unpaused) XXX_DiscardUnknown() {
	xxx_messageInfo_Unpaused.DiscardUnknown(m)
}

var xxx_messageInfo_Unpaused proto.InternalMessageInfo

func (m *Unpaused) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type OwnershipTransferred struct {
	PreviousOwner string `protobuf:"bytes,1,opt,name=previous_owner,json=previousOwner,proto3" json:"previous_owner,omitempty"`
	NewOwner      string `protobuf:"bytes,2,opt,name=new_owner,json=newOwner,proto3" json:"new_owner,omitempty"`
}

func (m *OwnershipTransferred) Reset()         { *m = OwnershipTransferred{} }
func (m *OwnershipTransferred) String() string { return proto.CompactTextString(m) }
func (*OwnershipTransferred) ProtoMessage()    {}
func (*OwnershipTransferred) Descriptor() ([]byte, []int) {
	return fileDescriptor_f721b73b2cc60687, []int{5}
}
func (m *OwnershipTransferred) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OwnershipTransferred) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OwnershipTransferred.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OwnershipTransferred) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnershipTransferred.Merge(m, src)
}
func (m *OwnershipTransferred) XXX_Size() int {
	return m.Size()
}
func (m *OwnershipTransferred) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnershipTransferred.DiscardUnknown(m)
}

var xxx_messageInfo_OwnershipTransferred proto.InternalMessageInfo

func (m *OwnershipTransferred) GetPreviousOwner() string {
	if m != nil {
		return m.PreviousOwner
	}
	return ""
}

func (m *OwnershipTransferred) GetNewOwner() string {
	if m != nil {
		return m.NewOwner
	}
	return ""
}

func init() {
	proto.RegisterType((*PublicCapabilityUpdated)(nil), "halo.entitlements.v1.PublicCapabilityUpdated")
	proto.RegisterType((*RoleCapabilityUpdated)(nil), "halo.entitlements.v1.RoleCapabilityUpdated")
	proto.RegisterType((*UserRoleUpdated)(nil), "halo.entitlements.v1.UserRoleUpdated")
	proto.RegisterType((*Paused)(nil), "halo.entitlements.v1.Paused")
	proto.RegisterType((*Unpaused)(nil), "halo.entitlements.v1.Unpaused")
	proto.RegisterType((*OwnershipTransferred)(nil), "halo.entitlements.v1.OwnershipTransferred")
}

func init() { proto.RegisterFile("halo/entitlements/v1/events.proto", fileDescriptor_f721b73b2cc60687) }

var fileDescriptor_f721b73b2cc60687 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x14, 0x45, 0x9b, 0x5a, 0x6a, 0x3b, 0x60, 0x85, 0x50, 0xb5, 0x54, 0x08, 0x35, 0x28, 0x76, 0x63,
	0x42, 0xdb, 0x3f, 0xd0, 0xa5, 0x82, 0x25, 0xd8, 0x4d, 0x37, 0x32, 0x49, 0x9e, 0x66, 0x60, 0x3a,
	0x33, 0xcc, 0x4c, 0x52, 0xfa, 0x17, 0x7e, 0x96, 0xcb, 0x2e, 0x5d, 0x4a, 0xfb, 0x23, 0x92, 0x31,
	0x91, 0x16, 0x5b, 0x74, 0xf7, 0x5e, 0x72, 0xee, 0xbb, 0x97, 0xe1, 0xa2, 0x8b, 0x04, 0x53, 0xee,
	0x03, 0xd3, 0x44, 0x53, 0x98, 0x01, 0xd3, 0xca, 0xcf, 0x06, 0x3e, 0x64, 0xf9, 0xe4, 0x09, 0xc9,
	0x35, 0xb7, 0xdb, 0x39, 0xe2, 0x6d, 0x22, 0x5e, 0x36, 0xe8, 0x5e, 0xef, 0x16, 0x6e, 0x52, 0x46,
	0xee, 0xde, 0xa3, 0xb3, 0x71, 0x1a, 0x52, 0x12, 0xdd, 0x61, 0x81, 0x43, 0x42, 0x89, 0x5e, 0x4c,
	0x44, 0x8c, 0x35, 0xc4, 0xf6, 0x29, 0xaa, 0xcf, 0x40, 0x27, 0x3c, 0xee, 0x58, 0x3d, 0xab, 0xdf,
	0x0c, 0x8a, 0xcd, 0xee, 0xa0, 0x43, 0x60, 0x38, 0xa4, 0x10, 0x77, 0xaa, 0x3d, 0xab, 0xdf, 0x08,
	0xca, 0xd5, 0x5d, 0xa0, 0x93, 0x80, 0x53, 0xf8, 0x7d, 0xca, 0x43, 0x35, 0xc9, 0x29, 0x98, 0x43,
	0xad, 0x61, 0xd7, 0xdb, 0x95, 0xd9, 0xcb, 0xa5, 0x81, 0xe1, 0x36, 0xac, 0xab, 0xfb, 0xac, 0x0f,
	0xb6, 0xad, 0x39, 0x3a, 0x9e, 0x28, 0x90, 0xf9, 0x8d, 0xd2, 0xd4, 0x46, 0xb5, 0x54, 0x81, 0x2c,
	0xd2, 0x9b, 0xf9, 0x27, 0x48, 0xf5, 0x9f, 0x41, 0xf6, 0x1b, 0xba, 0xa8, 0x3e, 0xc6, 0xa9, 0x02,
	0x13, 0x0a, 0x47, 0x11, 0x4f, 0x99, 0x2e, 0xac, 0xca, 0xd5, 0xbd, 0x44, 0x8d, 0x09, 0x13, 0x7f,
	0x51, 0x53, 0xd4, 0x7e, 0x9c, 0x33, 0x90, 0x2a, 0x21, 0xe2, 0x49, 0x62, 0xa6, 0x5e, 0x40, 0x4a,
	0x88, 0xed, 0x2b, 0xd4, 0x12, 0x12, 0x32, 0xc2, 0x53, 0xf5, 0xcc, 0x73, 0xa0, 0x10, 0x1e, 0x95,
	0x5f, 0x8d, 0xca, 0x3e, 0x47, 0x4d, 0x06, 0xf3, 0x82, 0xf8, 0x7e, 0xae, 0x06, 0x83, 0xb9, 0xf9,
	0x79, 0xfb, 0xf0, 0xbe, 0x72, 0xac, 0xe5, 0xca, 0xb1, 0x3e, 0x57, 0x8e, 0xf5, 0xb6, 0x76, 0x2a,
	0xcb, 0xb5, 0x53, 0xf9, 0x58, 0x3b, 0x95, 0xe9, 0xf0, 0x95, 0xe8, 0x24, 0x0d, 0xbd, 0x88, 0xcf,
	0x7c, 0xc6, 0x43, 0x0a, 0x37, 0x58, 0x29, 0xd0, 0xca, 0x37, 0xcd, 0xc9, 0x46, 0xbe, 0x5e, 0x08,
	0x50, 0x5b, 0x95, 0x09, 0xeb, 0xa6, 0x33, 0xa3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0x00,
	0xc3, 0xe8, 0x97, 0x02, 0x00, 0x00,
}

func (m *PublicCapabilityUpdated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublicCapabilityUpdated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PublicCapabilityUpdated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Method) > 0 {
		i -= len(m.Method)
		copy(dAtA[i:], m.Method)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Method)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RoleCapabilityUpdated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoleCapabilityUpdated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RoleCapabilityUpdated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Method) > 0 {
		i -= len(m.Method)
		copy(dAtA[i:], m.Method)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Method)))
		i--
		dAtA[i] = 0x12
	}
	if m.Role != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UserRoleUpdated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserRoleUpdated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserRoleUpdated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Role != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x10
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Paused) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Paused) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Paused) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Unpaused) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Unpaused) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Unpaused) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OwnershipTransferred) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OwnershipTransferred) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OwnershipTransferred) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewOwner) > 0 {
		i -= len(m.NewOwner)
		copy(dAtA[i:], m.NewOwner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NewOwner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PreviousOwner) > 0 {
		i -= len(m.PreviousOwner)
		copy(dAtA[i:], m.PreviousOwner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.PreviousOwner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PublicCapabilityUpdated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	return n
}

func (m *RoleCapabilityUpdated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Role != 0 {
		n += 1 + sovEvents(uint64(m.Role))
	}
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	return n
}

func (m *UserRoleUpdated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Role != 0 {
		n += 1 + sovEvents(uint64(m.Role))
	}
	if m.Enabled {
		n += 2
	}
	return n
}

func (m *Paused) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *Unpaused) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *OwnershipTransferred) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PreviousOwner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NewOwner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PublicCapabilityUpdated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: PublicCapabilityUpdated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublicCapabilityUpdated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *RoleCapabilityUpdated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: RoleCapabilityUpdated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoleCapabilityUpdated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= Role(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *UserRoleUpdated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: UserRoleUpdated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserRoleUpdated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= Role(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *Paused) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: Paused: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Paused: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *Unpaused) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: Unpaused: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Unpaused: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *OwnershipTransferred) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: OwnershipTransferred: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwnershipTransferred: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
