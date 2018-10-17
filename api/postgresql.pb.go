// Code generated by protoc-gen-go. DO NOT EDIT.
// source: postgresql.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PostgreSQLNode struct {
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostgreSQLNode) Reset()         { *m = PostgreSQLNode{} }
func (m *PostgreSQLNode) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLNode) ProtoMessage()    {}
func (*PostgreSQLNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_1f3feebc0db14cd3, []int{0}
}
func (m *PostgreSQLNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLNode.Unmarshal(m, b)
}
func (m *PostgreSQLNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLNode.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLNode.Merge(dst, src)
}
func (m *PostgreSQLNode) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLNode.Size(m)
}
func (m *PostgreSQLNode) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLNode.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLNode proto.InternalMessageInfo

func (m *PostgreSQLNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PostgreSQLService struct {
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Port                 uint32   `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	Engine               string   `protobuf:"bytes,6,opt,name=engine,proto3" json:"engine,omitempty"`
	EngineVersion        string   `protobuf:"bytes,7,opt,name=engine_version,json=engineVersion,proto3" json:"engine_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostgreSQLService) Reset()         { *m = PostgreSQLService{} }
func (m *PostgreSQLService) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLService) ProtoMessage()    {}
func (*PostgreSQLService) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_1f3feebc0db14cd3, []int{1}
}
func (m *PostgreSQLService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLService.Unmarshal(m, b)
}
func (m *PostgreSQLService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLService.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLService.Merge(dst, src)
}
func (m *PostgreSQLService) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLService.Size(m)
}
func (m *PostgreSQLService) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLService.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLService proto.InternalMessageInfo

func (m *PostgreSQLService) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PostgreSQLService) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *PostgreSQLService) GetEngine() string {
	if m != nil {
		return m.Engine
	}
	return ""
}

func (m *PostgreSQLService) GetEngineVersion() string {
	if m != nil {
		return m.EngineVersion
	}
	return ""
}

type PostgreSQLInstance struct {
	Node                 *PostgreSQLNode    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Service              *PostgreSQLService `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PostgreSQLInstance) Reset()         { *m = PostgreSQLInstance{} }
func (m *PostgreSQLInstance) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLInstance) ProtoMessage()    {}
func (*PostgreSQLInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_1f3feebc0db14cd3, []int{2}
}
func (m *PostgreSQLInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLInstance.Unmarshal(m, b)
}
func (m *PostgreSQLInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLInstance.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLInstance.Merge(dst, src)
}
func (m *PostgreSQLInstance) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLInstance.Size(m)
}
func (m *PostgreSQLInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLInstance.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLInstance proto.InternalMessageInfo

func (m *PostgreSQLInstance) GetNode() *PostgreSQLNode {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *PostgreSQLInstance) GetService() *PostgreSQLService {
	if m != nil {
		return m.Service
	}
	return nil
}

type PostgreSQLListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostgreSQLListRequest) Reset()         { *m = PostgreSQLListRequest{} }
func (m *PostgreSQLListRequest) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLListRequest) ProtoMessage()    {}
func (*PostgreSQLListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_1f3feebc0db14cd3, []int{3}
}
func (m *PostgreSQLListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLListRequest.Unmarshal(m, b)
}
func (m *PostgreSQLListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLListRequest.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLListRequest.Merge(dst, src)
}
func (m *PostgreSQLListRequest) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLListRequest.Size(m)
}
func (m *PostgreSQLListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLListRequest proto.InternalMessageInfo

type PostgreSQLListResponse struct {
	Instances            []*PostgreSQLInstance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PostgreSQLListResponse) Reset()         { *m = PostgreSQLListResponse{} }
func (m *PostgreSQLListResponse) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLListResponse) ProtoMessage()    {}
func (*PostgreSQLListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_1f3feebc0db14cd3, []int{4}
}
func (m *PostgreSQLListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLListResponse.Unmarshal(m, b)
}
func (m *PostgreSQLListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLListResponse.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLListResponse.Merge(dst, src)
}
func (m *PostgreSQLListResponse) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLListResponse.Size(m)
}
func (m *PostgreSQLListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLListResponse proto.InternalMessageInfo

func (m *PostgreSQLListResponse) GetInstances() []*PostgreSQLInstance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type PostgreSQLAddRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Username             string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostgreSQLAddRequest) Reset()         { *m = PostgreSQLAddRequest{} }
func (m *PostgreSQLAddRequest) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLAddRequest) ProtoMessage()    {}
func (*PostgreSQLAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_1f3feebc0db14cd3, []int{5}
}
func (m *PostgreSQLAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLAddRequest.Unmarshal(m, b)
}
func (m *PostgreSQLAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLAddRequest.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLAddRequest.Merge(dst, src)
}
func (m *PostgreSQLAddRequest) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLAddRequest.Size(m)
}
func (m *PostgreSQLAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLAddRequest proto.InternalMessageInfo

func (m *PostgreSQLAddRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PostgreSQLAddRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PostgreSQLAddRequest) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *PostgreSQLAddRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *PostgreSQLAddRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type PostgreSQLAddResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostgreSQLAddResponse) Reset()         { *m = PostgreSQLAddResponse{} }
func (m *PostgreSQLAddResponse) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLAddResponse) ProtoMessage()    {}
func (*PostgreSQLAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_1f3feebc0db14cd3, []int{6}
}
func (m *PostgreSQLAddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLAddResponse.Unmarshal(m, b)
}
func (m *PostgreSQLAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLAddResponse.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLAddResponse.Merge(dst, src)
}
func (m *PostgreSQLAddResponse) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLAddResponse.Size(m)
}
func (m *PostgreSQLAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLAddResponse proto.InternalMessageInfo

func (m *PostgreSQLAddResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PostgreSQLRemoveRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostgreSQLRemoveRequest) Reset()         { *m = PostgreSQLRemoveRequest{} }
func (m *PostgreSQLRemoveRequest) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLRemoveRequest) ProtoMessage()    {}
func (*PostgreSQLRemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_1f3feebc0db14cd3, []int{7}
}
func (m *PostgreSQLRemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLRemoveRequest.Unmarshal(m, b)
}
func (m *PostgreSQLRemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLRemoveRequest.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLRemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLRemoveRequest.Merge(dst, src)
}
func (m *PostgreSQLRemoveRequest) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLRemoveRequest.Size(m)
}
func (m *PostgreSQLRemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLRemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLRemoveRequest proto.InternalMessageInfo

func (m *PostgreSQLRemoveRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PostgreSQLRemoveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostgreSQLRemoveResponse) Reset()         { *m = PostgreSQLRemoveResponse{} }
func (m *PostgreSQLRemoveResponse) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLRemoveResponse) ProtoMessage()    {}
func (*PostgreSQLRemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_1f3feebc0db14cd3, []int{8}
}
func (m *PostgreSQLRemoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLRemoveResponse.Unmarshal(m, b)
}
func (m *PostgreSQLRemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLRemoveResponse.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLRemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLRemoveResponse.Merge(dst, src)
}
func (m *PostgreSQLRemoveResponse) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLRemoveResponse.Size(m)
}
func (m *PostgreSQLRemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLRemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLRemoveResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PostgreSQLNode)(nil), "api.PostgreSQLNode")
	proto.RegisterType((*PostgreSQLService)(nil), "api.PostgreSQLService")
	proto.RegisterType((*PostgreSQLInstance)(nil), "api.PostgreSQLInstance")
	proto.RegisterType((*PostgreSQLListRequest)(nil), "api.PostgreSQLListRequest")
	proto.RegisterType((*PostgreSQLListResponse)(nil), "api.PostgreSQLListResponse")
	proto.RegisterType((*PostgreSQLAddRequest)(nil), "api.PostgreSQLAddRequest")
	proto.RegisterType((*PostgreSQLAddResponse)(nil), "api.PostgreSQLAddResponse")
	proto.RegisterType((*PostgreSQLRemoveRequest)(nil), "api.PostgreSQLRemoveRequest")
	proto.RegisterType((*PostgreSQLRemoveResponse)(nil), "api.PostgreSQLRemoveResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostgreSQLClient is the client API for PostgreSQL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostgreSQLClient interface {
	List(ctx context.Context, in *PostgreSQLListRequest, opts ...grpc.CallOption) (*PostgreSQLListResponse, error)
	Add(ctx context.Context, in *PostgreSQLAddRequest, opts ...grpc.CallOption) (*PostgreSQLAddResponse, error)
	Remove(ctx context.Context, in *PostgreSQLRemoveRequest, opts ...grpc.CallOption) (*PostgreSQLRemoveResponse, error)
}

type postgreSQLClient struct {
	cc *grpc.ClientConn
}

func NewPostgreSQLClient(cc *grpc.ClientConn) PostgreSQLClient {
	return &postgreSQLClient{cc}
}

func (c *postgreSQLClient) List(ctx context.Context, in *PostgreSQLListRequest, opts ...grpc.CallOption) (*PostgreSQLListResponse, error) {
	out := new(PostgreSQLListResponse)
	err := c.cc.Invoke(ctx, "/api.PostgreSQL/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgreSQLClient) Add(ctx context.Context, in *PostgreSQLAddRequest, opts ...grpc.CallOption) (*PostgreSQLAddResponse, error) {
	out := new(PostgreSQLAddResponse)
	err := c.cc.Invoke(ctx, "/api.PostgreSQL/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgreSQLClient) Remove(ctx context.Context, in *PostgreSQLRemoveRequest, opts ...grpc.CallOption) (*PostgreSQLRemoveResponse, error) {
	out := new(PostgreSQLRemoveResponse)
	err := c.cc.Invoke(ctx, "/api.PostgreSQL/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostgreSQLServer is the server API for PostgreSQL service.
type PostgreSQLServer interface {
	List(context.Context, *PostgreSQLListRequest) (*PostgreSQLListResponse, error)
	Add(context.Context, *PostgreSQLAddRequest) (*PostgreSQLAddResponse, error)
	Remove(context.Context, *PostgreSQLRemoveRequest) (*PostgreSQLRemoveResponse, error)
}

func RegisterPostgreSQLServer(s *grpc.Server, srv PostgreSQLServer) {
	s.RegisterService(&_PostgreSQL_serviceDesc, srv)
}

func _PostgreSQL_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgreSQLListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgreSQLServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PostgreSQL/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgreSQLServer).List(ctx, req.(*PostgreSQLListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgreSQL_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgreSQLAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgreSQLServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PostgreSQL/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgreSQLServer).Add(ctx, req.(*PostgreSQLAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgreSQL_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgreSQLRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgreSQLServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PostgreSQL/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgreSQLServer).Remove(ctx, req.(*PostgreSQLRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostgreSQL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.PostgreSQL",
	HandlerType: (*PostgreSQLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _PostgreSQL_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _PostgreSQL_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _PostgreSQL_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "postgresql.proto",
}

func init() { proto.RegisterFile("postgresql.proto", fileDescriptor_postgresql_1f3feebc0db14cd3) }

var fileDescriptor_postgresql_1f3feebc0db14cd3 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4b, 0x6e, 0xdb, 0x30,
	0x10, 0x05, 0x25, 0xc5, 0x9f, 0x09, 0x62, 0xa8, 0x74, 0x63, 0x33, 0x4a, 0x0a, 0x18, 0x04, 0x8a,
	0xb8, 0x59, 0x58, 0x86, 0x8b, 0x6e, 0xba, 0xcb, 0xb2, 0x46, 0xd0, 0x8f, 0x02, 0x34, 0xcb, 0x82,
	0x09, 0x09, 0x83, 0x40, 0x42, 0x2a, 0xa2, 0xec, 0x2e, 0x8a, 0x6e, 0x7a, 0x82, 0xa2, 0xbd, 0x41,
	0xaf, 0xd4, 0x2b, 0xf4, 0x20, 0x85, 0x49, 0xd9, 0x94, 0x55, 0x77, 0xc7, 0xd1, 0x7b, 0xf3, 0xe6,
	0xcd, 0x3c, 0x08, 0xe2, 0x5c, 0x9b, 0x72, 0x51, 0x08, 0xf3, 0x78, 0x3f, 0xc9, 0x0b, 0x5d, 0x6a,
	0x1c, 0xb2, 0x5c, 0x26, 0x67, 0x0b, 0xad, 0x17, 0xf7, 0x22, 0x65, 0xb9, 0x4c, 0x99, 0x52, 0xba,
	0x64, 0xa5, 0xd4, 0xca, 0x38, 0x0a, 0x9d, 0x42, 0xef, 0xbd, 0x6b, 0xbb, 0xfe, 0x70, 0xf5, 0x56,
	0x73, 0x81, 0x31, 0x44, 0x8a, 0x3d, 0x08, 0x12, 0x8e, 0xd0, 0xb8, 0x9b, 0xd9, 0xf7, 0x3c, 0xea,
	0xa0, 0x38, 0x98, 0x47, 0x9d, 0x20, 0x0e, 0xe9, 0x0f, 0x04, 0x4f, 0x7c, 0xcb, 0xb5, 0x28, 0x56,
	0xf2, 0x4e, 0x60, 0x02, 0x6d, 0xc6, 0x79, 0x21, 0x8c, 0x21, 0x91, 0x6d, 0xdc, 0x94, 0x6b, 0xbd,
	0x5c, 0x17, 0x25, 0x39, 0x18, 0xa1, 0xf1, 0x51, 0x66, 0xdf, 0x78, 0x00, 0x2d, 0xa1, 0x16, 0x52,
	0x09, 0xd2, 0xb2, 0xe4, 0xaa, 0xc2, 0xcf, 0xa1, 0xe7, 0x5e, 0x9f, 0x56, 0xa2, 0x30, 0x52, 0x2b,
	0xd2, 0xb6, 0xf8, 0x91, 0xfb, 0xfa, 0xd1, 0x7d, 0xac, 0xdb, 0x99, 0x47, 0x9d, 0x30, 0x8e, 0xa8,
	0x06, 0xec, 0x3d, 0xbd, 0x51, 0xa6, 0x64, 0xea, 0x4e, 0xe0, 0x73, 0x88, 0x94, 0xe6, 0x82, 0xa0,
	0x11, 0x1a, 0x1f, 0xce, 0xfa, 0x13, 0x96, 0xcb, 0xc9, 0xee, 0xb6, 0x99, 0x25, 0xe0, 0x29, 0xb4,
	0x8d, 0x5b, 0x84, 0x04, 0x96, 0x3b, 0x68, 0x70, 0xab, 0x35, 0xb3, 0x0d, 0x8d, 0x0e, 0xe1, 0xd8,
	0xa3, 0x57, 0xd2, 0x94, 0x99, 0x78, 0x5c, 0x0a, 0x53, 0xd2, 0x77, 0x30, 0x68, 0x02, 0x26, 0xd7,
	0xca, 0x08, 0xfc, 0x0a, 0xba, 0xb2, 0x72, 0x66, 0x08, 0x1a, 0x85, 0xe3, 0xc3, 0xd9, 0xb0, 0x31,
	0x66, 0xe3, 0x3c, 0xf3, 0x4c, 0xfa, 0x1d, 0xc1, 0x53, 0xcf, 0xb8, 0xe4, 0xbc, 0x9a, 0xb4, 0x0d,
	0x0a, 0xf9, 0xa0, 0xea, 0x31, 0x04, 0xfb, 0x63, 0x08, 0x6b, 0x31, 0x24, 0xd0, 0x59, 0x1a, 0x51,
	0x58, 0x15, 0x97, 0xda, 0xb6, 0x5e, 0x63, 0x39, 0x33, 0xe6, 0xb3, 0x2e, 0xb8, 0x8d, 0xae, 0x9b,
	0x6d, 0x6b, 0x7a, 0x5e, 0x5f, 0xde, 0x3a, 0xaa, 0x56, 0xec, 0x41, 0x20, 0xb9, 0x35, 0x74, 0x90,
	0x05, 0x92, 0xd3, 0x17, 0x30, 0xf4, 0xc4, 0x4c, 0x3c, 0xe8, 0x95, 0xd8, 0xb8, 0x6f, 0x52, 0x13,
	0x20, 0xff, 0x52, 0x9d, 0xec, 0xec, 0x57, 0x00, 0xe0, 0x41, 0x7c, 0x03, 0xd1, 0xfa, 0xb0, 0x38,
	0x69, 0x5c, 0xaf, 0x16, 0x43, 0x72, 0xba, 0x17, 0x73, 0x7a, 0x74, 0xf0, 0xed, 0xf7, 0x9f, 0x9f,
	0x41, 0x8c, 0x7b, 0xe9, 0x6a, 0x9a, 0xfa, 0xbf, 0x06, 0xdf, 0x40, 0x78, 0xc9, 0x39, 0x3e, 0x69,
	0xf4, 0xfa, 0x9b, 0x27, 0xc9, 0x3e, 0xa8, 0x52, 0x3d, 0xb1, 0xaa, 0x7d, 0xda, 0x50, 0x7d, 0x8d,
	0x2e, 0xf0, 0x2d, 0xb4, 0xdc, 0x4a, 0xf8, 0xac, 0x21, 0xb0, 0x73, 0x94, 0xe4, 0xd9, 0x7f, 0xd0,
	0x6a, 0xc2, 0xa9, 0x9d, 0x70, 0x7c, 0xd1, 0xdf, 0x9d, 0x90, 0x7e, 0x91, 0xfc, 0xeb, 0x6d, 0xcb,
	0xfe, 0xd0, 0x2f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x66, 0x44, 0x32, 0x07, 0x04, 0x00,
	0x00,
}
