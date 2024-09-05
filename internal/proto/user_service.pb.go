// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user_service.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 请求的结构体
type UserRegisterRequest struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	PasswordConfirm      string   `protobuf:"bytes,4,opt,name=password_confirm,json=passwordConfirm,proto3" json:"password_confirm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterRequest) Reset()         { *m = UserRegisterRequest{} }
func (m *UserRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*UserRegisterRequest) ProtoMessage()    {}
func (*UserRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6d6fe680f72820c, []int{0}
}

func (m *UserRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterRequest.Unmarshal(m, b)
}
func (m *UserRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterRequest.Marshal(b, m, deterministic)
}
func (m *UserRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterRequest.Merge(m, src)
}
func (m *UserRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_UserRegisterRequest.Size(m)
}
func (m *UserRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterRequest proto.InternalMessageInfo

func (m *UserRegisterRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserRegisterRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserRegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserRegisterRequest) GetPasswordConfirm() string {
	if m != nil {
		return m.PasswordConfirm
	}
	return ""
}

// 返回的结构体
type RequestReplay struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestReplay) Reset()         { *m = RequestReplay{} }
func (m *RequestReplay) String() string { return proto.CompactTextString(m) }
func (*RequestReplay) ProtoMessage()    {}
func (*RequestReplay) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6d6fe680f72820c, []int{1}
}

func (m *RequestReplay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestReplay.Unmarshal(m, b)
}
func (m *RequestReplay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestReplay.Marshal(b, m, deterministic)
}
func (m *RequestReplay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestReplay.Merge(m, src)
}
func (m *RequestReplay) XXX_Size() int {
	return xxx_messageInfo_RequestReplay.Size(m)
}
func (m *RequestReplay) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestReplay.DiscardUnknown(m)
}

var xxx_messageInfo_RequestReplay proto.InternalMessageInfo

func (m *RequestReplay) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RequestReplay) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RequestReplay) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// 请求的结构体
type UserUpdateRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserUpdateRequest) Reset()         { *m = UserUpdateRequest{} }
func (m *UserUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UserUpdateRequest) ProtoMessage()    {}
func (*UserUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6d6fe680f72820c, []int{2}
}

func (m *UserUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserUpdateRequest.Unmarshal(m, b)
}
func (m *UserUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserUpdateRequest.Marshal(b, m, deterministic)
}
func (m *UserUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdateRequest.Merge(m, src)
}
func (m *UserUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UserUpdateRequest.Size(m)
}
func (m *UserUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdateRequest proto.InternalMessageInfo

func (m *UserUpdateRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserUpdateRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// 请求的结构体
type DeletetUserRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletetUserRequest) Reset()         { *m = DeletetUserRequest{} }
func (m *DeletetUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeletetUserRequest) ProtoMessage()    {}
func (*DeletetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6d6fe680f72820c, []int{3}
}

func (m *DeletetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletetUserRequest.Unmarshal(m, b)
}
func (m *DeletetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletetUserRequest.Marshal(b, m, deterministic)
}
func (m *DeletetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletetUserRequest.Merge(m, src)
}
func (m *DeletetUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeletetUserRequest.Size(m)
}
func (m *DeletetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletetUserRequest proto.InternalMessageInfo

func (m *DeletetUserRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *DeletetUserRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// 返回的结构体
type SelectUserRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectUserRequest) Reset()         { *m = SelectUserRequest{} }
func (m *SelectUserRequest) String() string { return proto.CompactTextString(m) }
func (*SelectUserRequest) ProtoMessage()    {}
func (*SelectUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6d6fe680f72820c, []int{4}
}

func (m *SelectUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectUserRequest.Unmarshal(m, b)
}
func (m *SelectUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectUserRequest.Marshal(b, m, deterministic)
}
func (m *SelectUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectUserRequest.Merge(m, src)
}
func (m *SelectUserRequest) XXX_Size() int {
	return xxx_messageInfo_SelectUserRequest.Size(m)
}
func (m *SelectUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SelectUserRequest proto.InternalMessageInfo

func (m *SelectUserRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SelectUserRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRegisterRequest)(nil), "proto.UserRegisterRequest")
	proto.RegisterType((*RequestReplay)(nil), "proto.RequestReplay")
	proto.RegisterType((*UserUpdateRequest)(nil), "proto.UserUpdateRequest")
	proto.RegisterType((*DeletetUserRequest)(nil), "proto.DeletetUserRequest")
	proto.RegisterType((*SelectUserRequest)(nil), "proto.SelectUserRequest")
}

func init() { proto.RegisterFile("proto/user_service.proto", fileDescriptor_d6d6fe680f72820c) }

var fileDescriptor_d6d6fe680f72820c = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0xcb, 0x4e, 0x32, 0x31,
	0x14, 0xfe, 0xb9, 0xfc, 0x5c, 0x0e, 0x18, 0xe5, 0x68, 0x62, 0x1d, 0x37, 0x66, 0x56, 0xba, 0xc1,
	0x44, 0xf7, 0xc6, 0xa8, 0x1b, 0x5c, 0x98, 0x38, 0x84, 0x35, 0xa9, 0xd3, 0x23, 0x99, 0x08, 0x33,
	0x63, 0x5b, 0x34, 0x3e, 0x87, 0xf1, 0x7d, 0x4d, 0xdb, 0x29, 0x0c, 0x41, 0xd8, 0xb0, 0x6a, 0xcf,
	0xed, 0xcb, 0x77, 0x01, 0x96, 0xcb, 0x4c, 0x67, 0x97, 0x73, 0x45, 0x72, 0xac, 0x48, 0x7e, 0x24,
	0x31, 0xf5, 0x6d, 0x0b, 0xff, 0xdb, 0x27, 0xfc, 0xa9, 0xc0, 0xe1, 0x48, 0x91, 0x8c, 0x68, 0x92,
	0x28, 0x6d, 0xde, 0xf7, 0x39, 0x29, 0x8d, 0x01, 0xb4, 0xd2, 0x24, 0x7e, 0x4b, 0xf9, 0x8c, 0x58,
	0xe5, 0xac, 0x72, 0xde, 0x8e, 0x16, 0x35, 0x9e, 0x42, 0xdb, 0x02, 0xda, 0x61, 0xd5, 0x0d, 0x4d,
	0xe3, 0xc9, 0x0c, 0x03, 0x68, 0xe5, 0x5c, 0xa9, 0xcf, 0x4c, 0x0a, 0x56, 0x73, 0x33, 0x5f, 0xe3,
	0x05, 0x1c, 0xf8, 0xff, 0x38, 0xce, 0xd2, 0xd7, 0x44, 0xce, 0x58, 0xdd, 0xee, 0xec, 0xfb, 0xfe,
	0xbd, 0x6b, 0x87, 0xcf, 0xb0, 0x57, 0x50, 0x89, 0x28, 0x9f, 0xf2, 0x2f, 0x44, 0xa8, 0xc7, 0x99,
	0x70, 0x64, 0x6a, 0x91, 0xfd, 0x23, 0x83, 0xe6, 0x8c, 0x94, 0xe2, 0x13, 0x4f, 0xc3, 0x97, 0x66,
	0x5b, 0x70, 0xcd, 0x2d, 0x83, 0x6e, 0x64, 0xff, 0xe1, 0x00, 0x7a, 0x46, 0xe9, 0x28, 0x17, 0x5c,
	0x93, 0xd7, 0x79, 0x0c, 0x4d, 0xab, 0x25, 0x11, 0x05, 0x72, 0xc3, 0x94, 0x03, 0xb1, 0x55, 0x64,
	0xf8, 0x08, 0xf8, 0x40, 0x53, 0xd2, 0xa4, 0x9d, 0x77, 0xbb, 0x60, 0x0d, 0xa0, 0x37, 0xa4, 0x29,
	0xc5, 0xbb, 0x43, 0x5d, 0x7d, 0x57, 0xa1, 0x63, 0x50, 0x86, 0x2e, 0x69, 0xbc, 0x83, 0x6e, 0x39,
	0x5b, 0x0c, 0x5c, 0xf6, 0xfd, 0x3f, 0x02, 0x0f, 0x8e, 0x8a, 0xd9, 0x8a, 0xeb, 0xe1, 0x3f, 0xbc,
	0x85, 0x4e, 0x49, 0x2a, 0x9e, 0x14, 0x6b, 0xeb, 0xf2, 0x37, 0x22, 0xdc, 0x00, 0x2c, 0x7d, 0x47,
	0x56, 0xe2, 0xb0, 0x12, 0xc5, 0xb6, 0xfb, 0xa5, 0x41, 0x8b, 0xfb, 0x35, 0xcf, 0x36, 0xdd, 0xbf,
	0x34, 0x6c, 0xfb, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x81, 0x9c, 0x27, 0x0c, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*RequestReplay, error)
	DeletetUser(ctx context.Context, in *DeletetUserRequest, opts ...grpc.CallOption) (*RequestReplay, error)
	UserUpdate(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*RequestReplay, error)
	SelectUser(ctx context.Context, in *SelectUserRequest, opts ...grpc.CallOption) (*RequestReplay, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*RequestReplay, error) {
	out := new(RequestReplay)
	err := c.cc.Invoke(ctx, "/proto.UserService/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeletetUser(ctx context.Context, in *DeletetUserRequest, opts ...grpc.CallOption) (*RequestReplay, error) {
	out := new(RequestReplay)
	err := c.cc.Invoke(ctx, "/proto.UserService/DeletetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserUpdate(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*RequestReplay, error) {
	out := new(RequestReplay)
	err := c.cc.Invoke(ctx, "/proto.UserService/UserUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SelectUser(ctx context.Context, in *SelectUserRequest, opts ...grpc.CallOption) (*RequestReplay, error) {
	out := new(RequestReplay)
	err := c.cc.Invoke(ctx, "/proto.UserService/SelectUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	UserRegister(context.Context, *UserRegisterRequest) (*RequestReplay, error)
	DeletetUser(context.Context, *DeletetUserRequest) (*RequestReplay, error)
	UserUpdate(context.Context, *UserUpdateRequest) (*RequestReplay, error)
	SelectUser(context.Context, *SelectUserRequest) (*RequestReplay, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) UserRegister(ctx context.Context, req *UserRegisterRequest) (*RequestReplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (*UnimplementedUserServiceServer) DeletetUser(ctx context.Context, req *DeletetUserRequest) (*RequestReplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletetUser not implemented")
}
func (*UnimplementedUserServiceServer) UserUpdate(ctx context.Context, req *UserUpdateRequest) (*RequestReplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdate not implemented")
}
func (*UnimplementedUserServiceServer) SelectUser(ctx context.Context, req *SelectUserRequest) (*RequestReplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRegister(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeletetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeletetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/DeletetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeletetUser(ctx, req.(*DeletetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/UserUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserUpdate(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SelectUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SelectUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/SelectUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SelectUser(ctx, req.(*SelectUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _UserService_UserRegister_Handler,
		},
		{
			MethodName: "DeletetUser",
			Handler:    _UserService_DeletetUser_Handler,
		},
		{
			MethodName: "UserUpdate",
			Handler:    _UserService_UserUpdate_Handler,
		},
		{
			MethodName: "SelectUser",
			Handler:    _UserService_SelectUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user_service.proto",
}
