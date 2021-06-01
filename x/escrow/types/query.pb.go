// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: iov/escrow/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryEscrowRequest is the request type for the Query/Escrow RPC method
type QueryEscrowRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryEscrowRequest) Reset()         { *m = QueryEscrowRequest{} }
func (m *QueryEscrowRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEscrowRequest) ProtoMessage()    {}
func (*QueryEscrowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e53d70ef3e4e87e, []int{0}
}
func (m *QueryEscrowRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEscrowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEscrowRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEscrowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEscrowRequest.Merge(m, src)
}
func (m *QueryEscrowRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEscrowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEscrowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEscrowRequest proto.InternalMessageInfo

func (m *QueryEscrowRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// QueryBalanceResponse is the response type for the Query/Escrow RPC method
type QueryEscrowResponse struct {
	Escrow *Escrow `protobuf:"bytes,1,opt,name=escrow,proto3" json:"escrow,omitempty"`
}

func (m *QueryEscrowResponse) Reset()         { *m = QueryEscrowResponse{} }
func (m *QueryEscrowResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEscrowResponse) ProtoMessage()    {}
func (*QueryEscrowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e53d70ef3e4e87e, []int{1}
}
func (m *QueryEscrowResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEscrowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEscrowResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEscrowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEscrowResponse.Merge(m, src)
}
func (m *QueryEscrowResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEscrowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEscrowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEscrowResponse proto.InternalMessageInfo

func (m *QueryEscrowResponse) GetEscrow() *Escrow {
	if m != nil {
		return m.Escrow
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryEscrowRequest)(nil), "starnamed.x.escrow.v1beta1.QueryEscrowRequest")
	proto.RegisterType((*QueryEscrowResponse)(nil), "starnamed.x.escrow.v1beta1.QueryEscrowResponse")
}

func init() { proto.RegisterFile("iov/escrow/v1beta1/query.proto", fileDescriptor_7e53d70ef3e4e87e) }

var fileDescriptor_7e53d70ef3e4e87e = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x9b, 0x82, 0x05, 0x23, 0x78, 0x88, 0x1e, 0x46, 0x91, 0x30, 0x8a, 0x07, 0x41, 0x4c,
	0xd8, 0xbc, 0x79, 0x14, 0xf6, 0x01, 0xb6, 0xa3, 0xb7, 0x74, 0x0d, 0x35, 0xe0, 0xf2, 0xba, 0x26,
	0xad, 0x1b, 0xe2, 0xc5, 0xab, 0x17, 0x41, 0xfc, 0x4e, 0x1e, 0x07, 0x5e, 0x3c, 0x4a, 0xeb, 0x07,
	0x91, 0xa5, 0x51, 0x10, 0xa7, 0x78, 0x7b, 0x24, 0xbf, 0xdf, 0xff, 0x3d, 0xfe, 0x98, 0x2a, 0xa8,
	0xb9, 0x34, 0xd3, 0x12, 0xae, 0x79, 0x3d, 0x48, 0xa5, 0x15, 0x03, 0x3e, 0xaf, 0x64, 0xb9, 0x64,
	0x45, 0x09, 0x16, 0x48, 0x6c, 0xac, 0x28, 0xb5, 0x98, 0xc9, 0x8c, 0x2d, 0x58, 0xc7, 0x31, 0xcf,
	0xc5, 0x07, 0x39, 0x40, 0x7e, 0x25, 0xb9, 0x28, 0x14, 0x17, 0x5a, 0x83, 0x15, 0x56, 0x81, 0x36,
	0x9d, 0x19, 0xef, 0xe7, 0x90, 0x83, 0x1b, 0xf9, 0x7a, 0xf2, 0xaf, 0x9b, 0xf6, 0xd9, 0x65, 0x21,
	0xbd, 0x95, 0x1c, 0x62, 0x32, 0x5e, 0xaf, 0x1f, 0x39, 0x64, 0x22, 0xe7, 0x95, 0x34, 0x96, 0xec,
	0xe2, 0x50, 0x65, 0x3d, 0xd4, 0x47, 0x47, 0xdb, 0x93, 0x50, 0x65, 0xc9, 0x18, 0xef, 0x7d, 0xa3,
	0x4c, 0x01, 0xda, 0x48, 0x72, 0x86, 0xa3, 0x2e, 0xda, 0xa1, 0x3b, 0xc3, 0x84, 0xfd, 0x7e, 0x3d,
	0xf3, 0xae, 0x37, 0x86, 0x4f, 0x08, 0x6f, 0xb9, 0x4c, 0x72, 0x8f, 0x70, 0xd4, 0x7d, 0x12, 0xf6,
	0x57, 0xc0, 0xcf, 0x3b, 0x63, 0xfe, 0x6f, 0xbe, 0xbb, 0x38, 0xe9, 0xdf, 0xbd, 0xbc, 0x3f, 0x86,
	0x31, 0xe9, 0xf1, 0x2f, 0x91, 0x2f, 0x3e, 0xfb, 0xb9, 0x51, 0xd9, 0xed, 0xf9, 0xe8, 0xb9, 0xa1,
	0x68, 0xd5, 0x50, 0xf4, 0xd6, 0x50, 0xf4, 0xd0, 0xd2, 0x60, 0xd5, 0xd2, 0xe0, 0xb5, 0xa5, 0xc1,
	0xc5, 0x71, 0xae, 0xec, 0x65, 0x95, 0xb2, 0x29, 0xcc, 0xb8, 0x82, 0xfa, 0x04, 0xb4, 0xdc, 0x94,
	0xe2, 0xda, 0x4d, 0x23, 0x57, 0xef, 0xe9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x67, 0xe9,
	0x9e, 0xf0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Escrow queries the Escrow by the specified hash lock
	Escrow(ctx context.Context, in *QueryEscrowRequest, opts ...grpc.CallOption) (*QueryEscrowResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Escrow(ctx context.Context, in *QueryEscrowRequest, opts ...grpc.CallOption) (*QueryEscrowResponse, error) {
	out := new(QueryEscrowResponse)
	err := c.cc.Invoke(ctx, "/starnamed.x.escrow.v1beta1.Query/Escrow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Escrow queries the Escrow by the specified hash lock
	Escrow(context.Context, *QueryEscrowRequest) (*QueryEscrowResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Escrow(ctx context.Context, req *QueryEscrowRequest) (*QueryEscrowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Escrow not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Escrow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEscrowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Escrow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starnamed.x.escrow.v1beta1.Query/Escrow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Escrow(ctx, req.(*QueryEscrowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "starnamed.x.escrow.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Escrow",
			Handler:    _Query_Escrow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iov/escrow/v1beta1/query.proto",
}

func (m *QueryEscrowRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEscrowRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEscrowRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryEscrowResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEscrowResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEscrowResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Escrow != nil {
		{
			size, err := m.Escrow.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryEscrowRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryEscrowResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Escrow != nil {
		l = m.Escrow.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryEscrowRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEscrowRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEscrowRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEscrowResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEscrowResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEscrowResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Escrow", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Escrow == nil {
				m.Escrow = &Escrow{}
			}
			if err := m.Escrow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)