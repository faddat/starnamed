// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: iov/escrow/v1beta1/test.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Escrow defines the struct of an escrow
type TestObject struct {
	Id                  uint64                                        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner               github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner,omitempty"`
	NumAllowedTransfers int64                                         `protobuf:"varint,3,opt,name=numAllowedTransfers,proto3" json:"numAllowedTransfers,omitempty"`
}

func (m *TestObject) Reset()         { *m = TestObject{} }
func (m *TestObject) String() string { return proto.CompactTextString(m) }
func (*TestObject) ProtoMessage()    {}
func (*TestObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_869357724f579c88, []int{0}
}
func (m *TestObject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestObject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestObject.Merge(m, src)
}
func (m *TestObject) XXX_Size() int {
	return m.Size()
}
func (m *TestObject) XXX_DiscardUnknown() {
	xxx_messageInfo_TestObject.DiscardUnknown(m)
}

var xxx_messageInfo_TestObject proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TestObject)(nil), "starnamed.x.escrow.v1beta1.TestObject")
}

func init() { proto.RegisterFile("iov/escrow/v1beta1/test.proto", fileDescriptor_869357724f579c88) }

var fileDescriptor_869357724f579c88 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4e, 0xeb, 0x40,
	0x10, 0x86, 0xbd, 0xc9, 0x7b, 0x14, 0x2b, 0x44, 0x61, 0x28, 0xa2, 0x48, 0x2c, 0x11, 0x55, 0x24,
	0xe4, 0x5d, 0x2c, 0x4e, 0xe0, 0x34, 0x88, 0x0a, 0xc9, 0x4a, 0x45, 0x67, 0xef, 0x0e, 0xc6, 0x10,
	0x7b, 0xa2, 0x9d, 0x8d, 0x1d, 0x6e, 0x41, 0xc7, 0x95, 0x52, 0xa6, 0xa4, 0x42, 0x60, 0xdf, 0x82,
	0x0a, 0x61, 0x07, 0x44, 0x41, 0x35, 0x23, 0xcd, 0xf7, 0x4b, 0xff, 0x37, 0xfc, 0x38, 0xc7, 0x4a,
	0x01, 0x69, 0x8b, 0xb5, 0xaa, 0xc2, 0x14, 0x5c, 0x12, 0x2a, 0x07, 0xe4, 0xe4, 0xd2, 0xa2, 0x43,
	0x7f, 0x4c, 0x2e, 0xb1, 0x65, 0x52, 0x80, 0x91, 0x6b, 0xd9, 0x63, 0x72, 0x87, 0x8d, 0x8f, 0x32,
	0xcc, 0xb0, 0xc3, 0xd4, 0xd7, 0xd6, 0x27, 0x4e, 0x9f, 0x19, 0xe7, 0x73, 0x20, 0x77, 0x9d, 0xde,
	0x83, 0x76, 0xfe, 0x01, 0x1f, 0xe4, 0x66, 0xc4, 0x26, 0x6c, 0xfa, 0x2f, 0x1e, 0xe4, 0xc6, 0xbf,
	0xe4, 0xff, 0xb1, 0x2e, 0xc1, 0x8e, 0x06, 0x13, 0x36, 0xdd, 0x9f, 0x85, 0x1f, 0xaf, 0x27, 0x41,
	0x96, 0xbb, 0xbb, 0x55, 0x2a, 0x35, 0x16, 0x4a, 0x23, 0x15, 0x48, 0xbb, 0x11, 0x90, 0x79, 0x50,
	0xee, 0x71, 0x09, 0x24, 0x23, 0xad, 0x23, 0x63, 0x2c, 0x10, 0xc5, 0x7d, 0xde, 0x3f, 0xe7, 0x87,
	0xe5, 0xaa, 0x88, 0x16, 0x0b, 0xac, 0xc1, 0xcc, 0x6d, 0x52, 0xd2, 0x2d, 0x58, 0x1a, 0x0d, 0x27,
	0x6c, 0x3a, 0x8c, 0xff, 0x3a, 0xcd, 0xae, 0x36, 0xef, 0xc2, 0xdb, 0x34, 0x82, 0x6d, 0x1b, 0xc1,
	0xde, 0x1a, 0xc1, 0x9e, 0x5a, 0xe1, 0x6d, 0x5b, 0xe1, 0xbd, 0xb4, 0xc2, 0xbb, 0x39, 0xfb, 0xd5,
	0x22, 0xc7, 0x2a, 0xc0, 0x12, 0xd4, 0x8f, 0xbc, 0x5a, 0x7f, 0xff, 0xa8, 0xab, 0x93, 0xee, 0x75,
	0xae, 0x17, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x63, 0x15, 0xde, 0x3e, 0x01, 0x00, 0x00,
}

func (m *TestObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestObject) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestObject) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumAllowedTransfers != 0 {
		i = encodeVarintTest(dAtA, i, uint64(m.NumAllowedTransfers))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTest(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTest(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTest(dAtA []byte, offset int, v uint64) int {
	offset -= sovTest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TestObject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTest(uint64(m.Id))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	if m.NumAllowedTransfers != 0 {
		n += 1 + sovTest(uint64(m.NumAllowedTransfers))
	}
	return n
}

func sovTest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTest(x uint64) (n int) {
	return sovTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TestObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: TestObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumAllowedTransfers", wireType)
			}
			m.NumAllowedTransfers = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumAllowedTransfers |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
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
func skipTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
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
				return 0, ErrInvalidLengthTest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTest = fmt.Errorf("proto: unexpected end of group")
)
