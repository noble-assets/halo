// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: halo/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	aggregator "github.com/noble-assets/halo/v3/types/aggregator"
	entitlements "github.com/noble-assets/halo/v3/types/entitlements"
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

type GenesisState struct {
	// aggregator_state is the genesis state of the aggregator submodule.
	AggregatorState aggregator.GenesisState `protobuf:"bytes,1,opt,name=aggregator_state,json=aggregatorState,proto3" json:"aggregator_state"`
	// entitlements_state is the genesis state of the entitlements submodule.
	EntitlementsState entitlements.GenesisState `protobuf:"bytes,2,opt,name=entitlements_state,json=entitlementsState,proto3" json:"entitlements_state"`
	// owner is the address that can control this module.
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	// nonces contains the withdrawal nonce per user.
	Nonces map[string]uint64 `protobuf:"bytes,4,rep,name=nonces,proto3" json:"nonces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afc1cd6dd6a46b0, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetAggregatorState() aggregator.GenesisState {
	if m != nil {
		return m.AggregatorState
	}
	return aggregator.GenesisState{}
}

func (m *GenesisState) GetEntitlementsState() entitlements.GenesisState {
	if m != nil {
		return m.EntitlementsState
	}
	return entitlements.GenesisState{}
}

func (m *GenesisState) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *GenesisState) GetNonces() map[string]uint64 {
	if m != nil {
		return m.Nonces
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "halo.v1.GenesisState")
	proto.RegisterMapType((map[string]uint64)(nil), "halo.v1.GenesisState.NoncesEntry")
}

func init() { proto.RegisterFile("halo/v1/genesis.proto", fileDescriptor_7afc1cd6dd6a46b0) }

var fileDescriptor_7afc1cd6dd6a46b0 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6a, 0xfa, 0x40,
	0x10, 0xc7, 0x13, 0xff, 0xfd, 0x70, 0xfd, 0x41, 0x6d, 0xb0, 0x90, 0xe6, 0x90, 0xa6, 0x42, 0xc1,
	0x8b, 0x1b, 0xd4, 0x4b, 0xed, 0xa5, 0x54, 0x28, 0xbd, 0x15, 0x1a, 0x0f, 0x85, 0x5e, 0x24, 0xea,
	0xb0, 0x86, 0xc6, 0x5d, 0xd9, 0x5d, 0x53, 0x7c, 0x8b, 0x3e, 0x4c, 0x1f, 0xc2, 0xa3, 0xf4, 0xd4,
	0x53, 0x11, 0x7d, 0x91, 0x92, 0xdd, 0x80, 0x29, 0xd2, 0xdb, 0xcc, 0x77, 0xbe, 0xf3, 0xd9, 0xd9,
	0xd9, 0x45, 0x67, 0xb3, 0x30, 0x66, 0x7e, 0xd2, 0xf1, 0x09, 0x50, 0x10, 0x91, 0xc0, 0x0b, 0xce,
	0x24, 0xb3, 0xfe, 0xa5, 0x32, 0x4e, 0x3a, 0xce, 0xf9, 0x84, 0x89, 0x39, 0x13, 0x23, 0x25, 0xfb,
	0x3a, 0xd1, 0x1e, 0xa7, 0x41, 0x18, 0x61, 0x5a, 0x4f, 0xa3, 0x4c, 0xf5, 0x14, 0x30, 0x24, 0x84,
	0x03, 0x09, 0x25, 0xe3, 0x47, 0x6c, 0xa7, 0xa9, 0x1c, 0x40, 0x65, 0x24, 0x63, 0x98, 0x03, 0x95,
	0xe2, 0xc8, 0xd3, 0xdc, 0x16, 0xd0, 0xff, 0x07, 0xad, 0x0c, 0x65, 0x28, 0xc1, 0x7a, 0x42, 0xf5,
	0x03, 0x73, 0x24, 0x52, 0xcd, 0x36, 0x3d, 0xb3, 0x55, 0xeb, 0x7a, 0x58, 0xcd, 0x7a, 0xa8, 0xe2,
	0xa4, 0x83, 0xf3, 0xbd, 0x83, 0xd2, 0xfa, 0xfb, 0xc2, 0x08, 0x4e, 0x0e, 0x0e, 0x8d, 0x7c, 0x46,
	0x56, 0x7e, 0x88, 0x0c, 0x5a, 0x50, 0xd0, 0xa6, 0x86, 0xe6, 0xeb, 0x7f, 0x60, 0x4f, 0xf3, 0x1e,
	0x0d, 0xc6, 0xa8, 0xcc, 0xde, 0x28, 0x70, 0xbb, 0xe8, 0x99, 0xad, 0xea, 0xc0, 0xfe, 0xfc, 0x68,
	0x37, 0xb2, 0xcd, 0xdd, 0x4d, 0xa7, 0x1c, 0x84, 0x18, 0x4a, 0x1e, 0x51, 0x12, 0x68, 0x9b, 0xd5,
	0x47, 0x15, 0xca, 0xe8, 0x04, 0x84, 0x5d, 0xf2, 0x8a, 0xad, 0x5a, 0xf7, 0x12, 0x67, 0xdb, 0xff,
	0x75, 0x1e, 0x7e, 0x54, 0x9e, 0x7b, 0x2a, 0xf9, 0x2a, 0xc8, 0x1a, 0x9c, 0x3e, 0xaa, 0xe5, 0x64,
	0xab, 0x8e, 0x8a, 0xaf, 0xb0, 0x52, 0x8b, 0xa9, 0x06, 0x69, 0x68, 0x35, 0x50, 0x39, 0x09, 0xe3,
	0xa5, 0xbe, 0x57, 0x29, 0xd0, 0xc9, 0x4d, 0xe1, 0xda, 0x1c, 0xdc, 0xae, 0x77, 0xae, 0xb9, 0xd9,
	0xb9, 0xe6, 0x76, 0xe7, 0x9a, 0xef, 0x7b, 0xd7, 0xd8, 0xec, 0x5d, 0xe3, 0x6b, 0xef, 0x1a, 0x2f,
	0x57, 0x24, 0x92, 0xb3, 0xe5, 0x18, 0x4f, 0xd8, 0xdc, 0xa7, 0x6c, 0x1c, 0x43, 0x3b, 0x14, 0x02,
	0xa4, 0xf0, 0xf5, 0x5f, 0xe9, 0xf9, 0x72, 0xb5, 0x00, 0x31, 0xae, 0xa8, 0xa7, 0xea, 0xfd, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xc7, 0xab, 0x62, 0x1d, 0x43, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nonces) > 0 {
		for k := range m.Nonces {
			v := m.Nonces[k]
			baseI := i
			i = encodeVarintGenesis(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenesis(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.EntitlementsState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.AggregatorState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AggregatorState.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.EntitlementsState.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Nonces) > 0 {
		for k, v := range m.Nonces {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenesis(uint64(len(k))) + 1 + sovGenesis(uint64(v))
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregatorState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AggregatorState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntitlementsState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EntitlementsState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Nonces == nil {
				m.Nonces = make(map[string]uint64)
			}
			var mapkey string
			var mapvalue uint64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Nonces[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
