// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: halo/aggregator/v1/aggregator.proto

package aggregator

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type RoundData struct {
	Answer    cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=answer,proto3,customtype=cosmossdk.io/math.Int" json:"answer"`
	Balance   cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=balance,proto3,customtype=cosmossdk.io/math.Int" json:"balance"`
	Interest  cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=interest,proto3,customtype=cosmossdk.io/math.Int" json:"interest"`
	Supply    cosmossdk_io_math.Int `protobuf:"bytes,4,opt,name=supply,proto3,customtype=cosmossdk.io/math.Int" json:"supply"`
	UpdatedAt int64                 `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (m *RoundData) Reset()         { *m = RoundData{} }
func (m *RoundData) String() string { return proto.CompactTextString(m) }
func (*RoundData) ProtoMessage()    {}
func (*RoundData) Descriptor() ([]byte, []int) {
	return fileDescriptor_dab90eca84edcf6d, []int{0}
}
func (m *RoundData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoundData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RoundData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RoundData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundData.Merge(m, src)
}
func (m *RoundData) XXX_Size() int {
	return m.Size()
}
func (m *RoundData) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundData.DiscardUnknown(m)
}

var xxx_messageInfo_RoundData proto.InternalMessageInfo

func (m *RoundData) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*RoundData)(nil), "halo.aggregator.v1.RoundData")
}

func init() {
	proto.RegisterFile("halo/aggregator/v1/aggregator.proto", fileDescriptor_dab90eca84edcf6d)
}

var fileDescriptor_dab90eca84edcf6d = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0xbb, 0xa0, 0x28, 0x7b, 0xb3, 0xd1, 0xa4, 0x92, 0x58, 0x88, 0x5e, 0x88, 0x09, 0x5d,
	0x08, 0x4f, 0x20, 0xf1, 0x20, 0xc4, 0x13, 0x47, 0x2f, 0x64, 0x4a, 0x37, 0xa5, 0xb1, 0xdd, 0x69,
	0xba, 0x53, 0x0c, 0x6f, 0xe1, 0x63, 0x78, 0xf4, 0xe0, 0x43, 0x70, 0x24, 0x9e, 0x88, 0x07, 0x62,
	0xe0, 0xe0, 0x6b, 0x18, 0xda, 0x6a, 0x38, 0xf7, 0xb2, 0xd9, 0xff, 0xdf, 0x99, 0x6f, 0x33, 0x99,
	0x9f, 0xdf, 0xcc, 0x20, 0x44, 0x01, 0xbe, 0x9f, 0x48, 0x1f, 0x08, 0x13, 0x31, 0xef, 0x1d, 0x28,
	0x27, 0x4e, 0x90, 0xd0, 0x34, 0xf7, 0x45, 0xce, 0x81, 0x3d, 0xef, 0x35, 0xce, 0x20, 0x0a, 0x14,
	0x8a, 0xec, 0xcc, 0xcb, 0x1a, 0x97, 0x53, 0xd4, 0x11, 0xea, 0x49, 0xa6, 0x44, 0x2e, 0x8a, 0xa7,
	0x73, 0x1f, 0x7d, 0xcc, 0xfd, 0xfd, 0x2d, 0x77, 0xaf, 0xd7, 0x15, 0x5e, 0x1f, 0x63, 0xaa, 0xbc,
	0x7b, 0x20, 0x30, 0x1f, 0x78, 0x0d, 0x94, 0x7e, 0x91, 0x89, 0xc5, 0x5a, 0xac, 0x5d, 0x1f, 0x74,
	0x97, 0x9b, 0xa6, 0xf1, 0xb5, 0x69, 0x5e, 0xe4, 0x24, 0xed, 0x3d, 0x3b, 0x01, 0x8a, 0x08, 0x68,
	0xe6, 0x0c, 0x15, 0x7d, 0x7e, 0x74, 0x78, 0xf1, 0xc5, 0x50, 0xd1, 0xdb, 0xcf, 0xfb, 0x2d, 0x1b,
	0x17, 0xfd, 0xe6, 0x88, 0x9f, 0xb8, 0x10, 0x82, 0x9a, 0x4a, 0xab, 0x52, 0x12, 0xf5, 0x07, 0x30,
	0x1f, 0xf9, 0x69, 0xa0, 0x48, 0x26, 0x52, 0x93, 0x55, 0x2d, 0x09, 0xfb, 0x27, 0xec, 0x67, 0xd4,
	0x69, 0x1c, 0x87, 0x0b, 0xeb, 0xa8, 0xec, 0x8c, 0x79, 0xbf, 0x79, 0xc5, 0x79, 0x1a, 0x7b, 0x40,
	0xd2, 0x9b, 0x00, 0x59, 0xc7, 0x2d, 0xd6, 0xae, 0x8e, 0xeb, 0x85, 0x73, 0x47, 0x83, 0xd1, 0x72,
	0x6b, 0xb3, 0xd5, 0xd6, 0x66, 0xdf, 0x5b, 0x9b, 0xbd, 0xee, 0x6c, 0x63, 0xb5, 0xb3, 0x8d, 0xf5,
	0xce, 0x36, 0x9e, 0xba, 0x7e, 0x40, 0xb3, 0xd4, 0x75, 0xa6, 0x18, 0x09, 0x85, 0x6e, 0x28, 0x3b,
	0xa0, 0xb5, 0x24, 0x2d, 0xb2, 0x24, 0xcc, 0xfb, 0x82, 0x16, 0xb1, 0xd4, 0x07, 0x21, 0x70, 0x6b,
	0xd9, 0xb6, 0xfa, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xba, 0x59, 0xc0, 0x21, 0x2c, 0x02, 0x00,
	0x00,
}

func (m *RoundData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoundData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RoundData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdatedAt != 0 {
		i = encodeVarintAggregator(dAtA, i, uint64(m.UpdatedAt))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.Supply.Size()
		i -= size
		if _, err := m.Supply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAggregator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Interest.Size()
		i -= size
		if _, err := m.Interest.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAggregator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Balance.Size()
		i -= size
		if _, err := m.Balance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAggregator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Answer.Size()
		i -= size
		if _, err := m.Answer.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAggregator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAggregator(dAtA []byte, offset int, v uint64) int {
	offset -= sovAggregator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RoundData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Answer.Size()
	n += 1 + l + sovAggregator(uint64(l))
	l = m.Balance.Size()
	n += 1 + l + sovAggregator(uint64(l))
	l = m.Interest.Size()
	n += 1 + l + sovAggregator(uint64(l))
	l = m.Supply.Size()
	n += 1 + l + sovAggregator(uint64(l))
	if m.UpdatedAt != 0 {
		n += 1 + sovAggregator(uint64(m.UpdatedAt))
	}
	return n
}

func sovAggregator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAggregator(x uint64) (n int) {
	return sovAggregator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RoundData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAggregator
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
			return fmt.Errorf("proto: RoundData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoundData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Answer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregator
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
				return ErrInvalidLengthAggregator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAggregator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Answer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregator
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
				return ErrInvalidLengthAggregator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAggregator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregator
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
				return ErrInvalidLengthAggregator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAggregator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Interest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregator
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
				return ErrInvalidLengthAggregator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAggregator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Supply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			m.UpdatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAggregator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAggregator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAggregator
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
func skipAggregator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAggregator
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
					return 0, ErrIntOverflowAggregator
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
					return 0, ErrIntOverflowAggregator
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
				return 0, ErrInvalidLengthAggregator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAggregator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAggregator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAggregator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAggregator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAggregator = fmt.Errorf("proto: unexpected end of group")
)
