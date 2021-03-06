// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coins.proto

package coins

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

type CoinData struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	Rank                 int32    `protobuf:"varint,2,opt,name=Rank,proto3" json:"Rank,omitempty"`
	Score                float32  `protobuf:"fixed32,3,opt,name=Score,proto3" json:"Score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CoinData) Reset()         { *m = CoinData{} }
func (m *CoinData) String() string { return proto.CompactTextString(m) }
func (*CoinData) ProtoMessage()    {}
func (*CoinData) Descriptor() ([]byte, []int) {
	return fileDescriptor_da4483c99519c66a, []int{0}
}

func (m *CoinData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoinData.Unmarshal(m, b)
}
func (m *CoinData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoinData.Marshal(b, m, deterministic)
}
func (m *CoinData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinData.Merge(m, src)
}
func (m *CoinData) XXX_Size() int {
	return xxx_messageInfo_CoinData.Size(m)
}
func (m *CoinData) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinData.DiscardUnknown(m)
}

var xxx_messageInfo_CoinData proto.InternalMessageInfo

func (m *CoinData) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *CoinData) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *CoinData) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type TopCoinsResponse struct {
	Coins                []*CoinData `protobuf:"bytes,1,rep,name=coins,proto3" json:"coins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TopCoinsResponse) Reset()         { *m = TopCoinsResponse{} }
func (m *TopCoinsResponse) String() string { return proto.CompactTextString(m) }
func (*TopCoinsResponse) ProtoMessage()    {}
func (*TopCoinsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da4483c99519c66a, []int{1}
}

func (m *TopCoinsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopCoinsResponse.Unmarshal(m, b)
}
func (m *TopCoinsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopCoinsResponse.Marshal(b, m, deterministic)
}
func (m *TopCoinsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopCoinsResponse.Merge(m, src)
}
func (m *TopCoinsResponse) XXX_Size() int {
	return xxx_messageInfo_TopCoinsResponse.Size(m)
}
func (m *TopCoinsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TopCoinsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TopCoinsResponse proto.InternalMessageInfo

func (m *TopCoinsResponse) GetCoins() []*CoinData {
	if m != nil {
		return m.Coins
	}
	return nil
}

type CoinRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CoinRequest) Reset()         { *m = CoinRequest{} }
func (m *CoinRequest) String() string { return proto.CompactTextString(m) }
func (*CoinRequest) ProtoMessage()    {}
func (*CoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da4483c99519c66a, []int{2}
}

func (m *CoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoinRequest.Unmarshal(m, b)
}
func (m *CoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoinRequest.Marshal(b, m, deterministic)
}
func (m *CoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinRequest.Merge(m, src)
}
func (m *CoinRequest) XXX_Size() int {
	return xxx_messageInfo_CoinRequest.Size(m)
}
func (m *CoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CoinRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CoinData)(nil), "CoinData")
	proto.RegisterType((*TopCoinsResponse)(nil), "TopCoinsResponse")
	proto.RegisterType((*CoinRequest)(nil), "CoinRequest")
}

func init() { proto.RegisterFile("coins.proto", fileDescriptor_da4483c99519c66a) }

var fileDescriptor_da4483c99519c66a = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xce, 0xcf, 0xcc,
	0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xf2, 0xe1, 0xe2, 0x70, 0xce, 0xcf, 0xcc, 0x73,
	0x49, 0x2c, 0x49, 0x14, 0x12, 0xe3, 0x62, 0x0b, 0xae, 0xcc, 0x4d, 0xca, 0xcf, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x84, 0xb8, 0x58, 0x82, 0x12, 0xf3, 0xb2, 0x25, 0x98,
	0x14, 0x18, 0x35, 0x58, 0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xe0, 0xe4, 0xfc, 0xa2, 0x54,
	0x09, 0x66, 0x05, 0x46, 0x0d, 0xa6, 0x20, 0x08, 0x47, 0xc9, 0x98, 0x4b, 0x20, 0x24, 0xbf, 0x00,
	0x64, 0x60, 0x71, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x3c, 0x17, 0x2b, 0xd8,
	0x42, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x4e, 0x3d, 0x98, 0x7d, 0x41, 0x10, 0x71, 0x25,
	0x5e, 0x2e, 0x6e, 0x90, 0x50, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x91, 0x09, 0x17, 0x2b,
	0xd8, 0x00, 0x21, 0x6d, 0x2e, 0x0e, 0xf7, 0xd4, 0x12, 0x08, 0x9b, 0x47, 0x0f, 0x49, 0x89, 0x94,
	0xa0, 0x1e, 0xba, 0x2d, 0x49, 0x6c, 0x60, 0xef, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x60,
	0x34, 0x97, 0xf7, 0xdd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoinsClient is the client API for Coins service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoinsClient interface {
	GetCoins(ctx context.Context, in *CoinRequest, opts ...grpc.CallOption) (*TopCoinsResponse, error)
}

type coinsClient struct {
	cc *grpc.ClientConn
}

func NewCoinsClient(cc *grpc.ClientConn) CoinsClient {
	return &coinsClient{cc}
}

func (c *coinsClient) GetCoins(ctx context.Context, in *CoinRequest, opts ...grpc.CallOption) (*TopCoinsResponse, error) {
	out := new(TopCoinsResponse)
	err := c.cc.Invoke(ctx, "/Coins/GetCoins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoinsServer is the server API for Coins service.
type CoinsServer interface {
	GetCoins(context.Context, *CoinRequest) (*TopCoinsResponse, error)
}

// UnimplementedCoinsServer can be embedded to have forward compatible implementations.
type UnimplementedCoinsServer struct {
}

func (*UnimplementedCoinsServer) GetCoins(ctx context.Context, req *CoinRequest) (*TopCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoins not implemented")
}

func RegisterCoinsServer(s *grpc.Server, srv CoinsServer) {
	s.RegisterService(&_Coins_serviceDesc, srv)
}

func _Coins_GetCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinsServer).GetCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Coins/GetCoins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinsServer).GetCoins(ctx, req.(*CoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coins_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Coins",
	HandlerType: (*CoinsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCoins",
			Handler:    _Coins_GetCoins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coins.proto",
}
