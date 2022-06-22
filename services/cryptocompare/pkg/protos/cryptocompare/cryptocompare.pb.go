// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cryptocompare.proto

package cryptocompare

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TopData struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	Rank                 int32    `protobuf:"varint,2,opt,name=Rank,proto3" json:"Rank,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopData) Reset()         { *m = TopData{} }
func (m *TopData) String() string { return proto.CompactTextString(m) }
func (*TopData) ProtoMessage()    {}
func (*TopData) Descriptor() ([]byte, []int) {
	return fileDescriptor_315e547c50f0688e, []int{0}
}

func (m *TopData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopData.Unmarshal(m, b)
}
func (m *TopData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopData.Marshal(b, m, deterministic)
}
func (m *TopData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopData.Merge(m, src)
}
func (m *TopData) XXX_Size() int {
	return xxx_messageInfo_TopData.Size(m)
}
func (m *TopData) XXX_DiscardUnknown() {
	xxx_messageInfo_TopData.DiscardUnknown(m)
}

var xxx_messageInfo_TopData proto.InternalMessageInfo

func (m *TopData) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *TopData) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

type TopResponse struct {
	Ranks                map[string]*TopData `protobuf:"bytes,1,rep,name=ranks,proto3" json:"ranks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TopResponse) Reset()         { *m = TopResponse{} }
func (m *TopResponse) String() string { return proto.CompactTextString(m) }
func (*TopResponse) ProtoMessage()    {}
func (*TopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_315e547c50f0688e, []int{1}
}

func (m *TopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopResponse.Unmarshal(m, b)
}
func (m *TopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopResponse.Marshal(b, m, deterministic)
}
func (m *TopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopResponse.Merge(m, src)
}
func (m *TopResponse) XXX_Size() int {
	return xxx_messageInfo_TopResponse.Size(m)
}
func (m *TopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TopResponse proto.InternalMessageInfo

func (m *TopResponse) GetRanks() map[string]*TopData {
	if m != nil {
		return m.Ranks
	}
	return nil
}

type TopRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopRequest) Reset()         { *m = TopRequest{} }
func (m *TopRequest) String() string { return proto.CompactTextString(m) }
func (*TopRequest) ProtoMessage()    {}
func (*TopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_315e547c50f0688e, []int{2}
}

func (m *TopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopRequest.Unmarshal(m, b)
}
func (m *TopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopRequest.Marshal(b, m, deterministic)
}
func (m *TopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopRequest.Merge(m, src)
}
func (m *TopRequest) XXX_Size() int {
	return xxx_messageInfo_TopRequest.Size(m)
}
func (m *TopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TopRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TopData)(nil), "TopData")
	proto.RegisterType((*TopResponse)(nil), "TopResponse")
	proto.RegisterMapType((map[string]*TopData)(nil), "TopResponse.RanksEntry")
	proto.RegisterType((*TopRequest)(nil), "TopRequest")
}

func init() { proto.RegisterFile("cryptocompare.proto", fileDescriptor_315e547c50f0688e) }

var fileDescriptor_315e547c50f0688e = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xc9, 0xae, 0xad, 0x3a, 0x59, 0x41, 0x46, 0xd0, 0xb2, 0x07, 0x29, 0xf5, 0xd2, 0x8b,
	0x39, 0x54, 0x05, 0xf1, 0xe8, 0x1f, 0xbc, 0xc7, 0x7e, 0x81, 0xec, 0x92, 0x53, 0x77, 0x33, 0x31,
	0x49, 0x85, 0xdc, 0xfc, 0xe8, 0x4b, 0xd3, 0x42, 0x7b, 0x7b, 0x79, 0x79, 0xc3, 0xef, 0xcd, 0xc0,
	0xcd, 0xde, 0x45, 0x1b, 0x68, 0x4f, 0x47, 0xab, 0x9c, 0x16, 0xd6, 0x51, 0xa0, 0xea, 0x05, 0xce,
	0x5b, 0xb2, 0x9f, 0x2a, 0x28, 0xbc, 0x85, 0xfc, 0x27, 0x1e, 0x77, 0x74, 0x28, 0x58, 0xc9, 0xea,
	0x4b, 0x39, 0xbd, 0x10, 0xe1, 0x4c, 0x2a, 0xd3, 0x15, 0xab, 0x92, 0xd5, 0x99, 0x4c, 0xba, 0xfa,
	0x67, 0xc0, 0x5b, 0xb2, 0x52, 0x7b, 0x4b, 0xc6, 0x6b, 0x7c, 0x84, 0xcc, 0x29, 0xd3, 0xf9, 0x82,
	0x95, 0xeb, 0x9a, 0x37, 0x77, 0x62, 0xf1, 0x29, 0x86, 0x09, 0xff, 0x65, 0x82, 0x8b, 0x72, 0x4c,
	0x6d, 0xdf, 0x01, 0x66, 0x13, 0xaf, 0x61, 0xdd, 0xe9, 0x38, 0x51, 0x07, 0x89, 0xf7, 0x90, 0xfd,
	0xa9, 0x43, 0xaf, 0x13, 0x93, 0x37, 0x17, 0x62, 0xea, 0x28, 0x47, 0xfb, 0x6d, 0xf5, 0xca, 0xaa,
	0x0d, 0x40, 0x82, 0xfc, 0xf6, 0xda, 0x87, 0xe6, 0x19, 0xae, 0x3e, 0x96, 0xeb, 0xe1, 0x03, 0xe4,
	0xdf, 0x3a, 0xb4, 0x64, 0x91, 0x8b, 0x39, 0xb7, 0xdd, 0x2c, 0x9b, 0xed, 0xf2, 0x74, 0x84, 0xa7,
	0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0x91, 0x5b, 0x22, 0x1b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CryptocompareClient is the client API for Cryptocompare service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CryptocompareClient interface {
	GetTop(ctx context.Context, in *TopRequest, opts ...grpc.CallOption) (*TopResponse, error)
}

type cryptocompareClient struct {
	cc *grpc.ClientConn
}

func NewCryptocompareClient(cc *grpc.ClientConn) CryptocompareClient {
	return &cryptocompareClient{cc}
}

func (c *cryptocompareClient) GetTop(ctx context.Context, in *TopRequest, opts ...grpc.CallOption) (*TopResponse, error) {
	out := new(TopResponse)
	err := c.cc.Invoke(ctx, "/Cryptocompare/GetTop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptocompareServer is the server API for Cryptocompare service.
type CryptocompareServer interface {
	GetTop(context.Context, *TopRequest) (*TopResponse, error)
}

// UnimplementedCryptocompareServer can be embedded to have forward compatible implementations.
type UnimplementedCryptocompareServer struct {
}

func (*UnimplementedCryptocompareServer) GetTop(ctx context.Context, req *TopRequest) (*TopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTop not implemented")
}

func RegisterCryptocompareServer(s *grpc.Server, srv CryptocompareServer) {
	s.RegisterService(&_Cryptocompare_serviceDesc, srv)
}

func _Cryptocompare_GetTop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptocompareServer).GetTop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cryptocompare/GetTop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptocompareServer).GetTop(ctx, req.(*TopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cryptocompare_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Cryptocompare",
	HandlerType: (*CryptocompareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTop",
			Handler:    _Cryptocompare_GetTop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cryptocompare.proto",
}
