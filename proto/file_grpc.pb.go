// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/file.proto

package pgenerated

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	SubscribeByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (ChatService_SubscribeByUserClient, error)
	SendMessageToChat(ctx context.Context, in *RequestSendMessage, opts ...grpc.CallOption) (*ResponseSendMessage, error)
	CreateChat(ctx context.Context, in *RequestCreateChat, opts ...grpc.CallOption) (*ResponseCreateChat, error)
	UpdateChat(ctx context.Context, in *RequestUpdateChat, opts ...grpc.CallOption) (*ResponseUpdateChat, error)
	GetChatsByUser(ctx context.Context, in *RequestGetChatsByUser, opts ...grpc.CallOption) (*ResponseGetChatsByUser, error)
	GetUsersByChat(ctx context.Context, in *RequestGetUsersByChat, opts ...grpc.CallOption) (*ResponseGetUsersByChat, error)
	GetMessagesByChat(ctx context.Context, in *RequestGetMessagesByChat, opts ...grpc.CallOption) (*ResponseGetMessagesByChat, error)
	AddUserToChat(ctx context.Context, in *RequestAddUserToChat, opts ...grpc.CallOption) (*ResponseAddUserToChat, error)
	RemoveUserFromChat(ctx context.Context, in *RequestRemoveUserFromChat, opts ...grpc.CallOption) (*ResponseRemoveUserFromChat, error)
	MakeUserAdmin(ctx context.Context, in *RequestMakeUserAdmin, opts ...grpc.CallOption) (*ResponseMakeUserAdmin, error)
	RemoveUserAdmin(ctx context.Context, in *RequestRemoveUserAdmin, opts ...grpc.CallOption) (*ResponseRemoveUserAdmin, error)
	RequestReactAMessage(ctx context.Context, in *RequestReactMessage, opts ...grpc.CallOption) (*ResponseReactMessage, error)
	RequestUnreactAMessage(ctx context.Context, in *RequestUnreactMessage, opts ...grpc.CallOption) (*ResponseUnreactMessage, error)
	Disconnect(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SubscribeByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (ChatService_SubscribeByUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/chatservice.ChatService/SubscribeByUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSubscribeByUserClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_SubscribeByUserClient interface {
	Recv() (*StreamByUser, error)
	grpc.ClientStream
}

type chatServiceSubscribeByUserClient struct {
	grpc.ClientStream
}

func (x *chatServiceSubscribeByUserClient) Recv() (*StreamByUser, error) {
	m := new(StreamByUser)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) SendMessageToChat(ctx context.Context, in *RequestSendMessage, opts ...grpc.CallOption) (*ResponseSendMessage, error) {
	out := new(ResponseSendMessage)
	err := c.cc.Invoke(ctx, "/chatservice.ChatService/SendMessageToChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CreateChat(ctx context.Context, in *RequestCreateChat, opts ...grpc.CallOption) (*ResponseCreateChat, error) {
	out := new(ResponseCreateChat)
	err := c.cc.Invoke(ctx, "/chatservice.ChatService/CreateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateChat(ctx context.Context, in *RequestUpdateChat, opts ...grpc.CallOption) (*ResponseUpdateChat, error) {
	out := new(ResponseUpdateChat)
	err := c.cc.Invoke(ctx, "/chatservice.ChatService/UpdateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChatsByUser(ctx context.Context, in *RequestGetChatsByUser, opts ...grpc.CallOption) (*ResponseGetChatsByUser, error) {
	out := new(ResponseGetChatsByUser)
	err := c.cc.Invoke(ctx, "/chatservice.ChatService/GetChatsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetUsersByChat(ctx context.Context, in *RequestGetUsersByChat, opts ...grpc.CallOption) (*ResponseGetUsersByChat, error) {
	out := new(ResponseGetUsersByChat)
	err := c.cc.Invoke(ctx, "/chatservice.ChatService/GetUsersByChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetMessagesByChat(ctx context.Context, in *RequestGetMessagesByChat, opts ...grpc.CallOption) (*ResponseGetMessagesByChat, error) {
	out := new(ResponseGetMessagesByChat)
	err := c.cc.Invoke(ctx, "/chatservice.ChatService/GetMessagesByChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) AddUserToChat(ctx context.Context, in *RequestAddUserToChat, opts ...grpc.CallOption) (*ResponseAddUserToChat, error) {
	out := new(ResponseAddUserToChat)
	err := c.cc.Invoke(ctx, "/chatservice.ChatService/AddUserToChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RemoveUserFromChat(ctx context.Context, in *RequestRemoveUserFromChat, opts ...grpc.CallOption) (*ResponseRemoveUserFromChat, error) {
	out := new(ResponseRemoveUserFromChat)
	err := c.cc.Invoke(ctx, "/chatservice.ChatService/RemoveUserFromChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) MakeUserAdmin(ctx context.Context, in *RequestMakeUserAdmin, opts ...grpc.CallOption) (*ResponseMakeUserAdmin, error) {
	out := new(ResponseMakeUserAdmin)
	err := c.cc.Invoke(ctx, "/chatservice.ChatService/MakeUserAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RemoveUserAdmin(ctx context.Context, in *RequestRemoveUserAdmin, opts ...grpc.CallOption) (*ResponseRemoveUserAdmin, error) {
	out := new(ResponseRemoveUserAdmin)
	err := c.cc.Invoke(ctx, "/chatservice.ChatService/RemoveUserAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RequestReactAMessage(ctx context.Context, in *RequestReactMessage, opts ...grpc.CallOption) (*ResponseReactMessage, error) {
	out := new(ResponseReactMessage)
	err := c.cc.Invoke(ctx, "/chatservice.ChatService/RequestReactAMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RequestUnreactAMessage(ctx context.Context, in *RequestUnreactMessage, opts ...grpc.CallOption) (*ResponseUnreactMessage, error) {
	out := new(ResponseUnreactMessage)
	err := c.cc.Invoke(ctx, "/chatservice.ChatService/RequestUnreactAMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Disconnect(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chatservice.ChatService/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	SubscribeByUser(*User, ChatService_SubscribeByUserServer) error
	SendMessageToChat(context.Context, *RequestSendMessage) (*ResponseSendMessage, error)
	CreateChat(context.Context, *RequestCreateChat) (*ResponseCreateChat, error)
	UpdateChat(context.Context, *RequestUpdateChat) (*ResponseUpdateChat, error)
	GetChatsByUser(context.Context, *RequestGetChatsByUser) (*ResponseGetChatsByUser, error)
	GetUsersByChat(context.Context, *RequestGetUsersByChat) (*ResponseGetUsersByChat, error)
	GetMessagesByChat(context.Context, *RequestGetMessagesByChat) (*ResponseGetMessagesByChat, error)
	AddUserToChat(context.Context, *RequestAddUserToChat) (*ResponseAddUserToChat, error)
	RemoveUserFromChat(context.Context, *RequestRemoveUserFromChat) (*ResponseRemoveUserFromChat, error)
	MakeUserAdmin(context.Context, *RequestMakeUserAdmin) (*ResponseMakeUserAdmin, error)
	RemoveUserAdmin(context.Context, *RequestRemoveUserAdmin) (*ResponseRemoveUserAdmin, error)
	RequestReactAMessage(context.Context, *RequestReactMessage) (*ResponseReactMessage, error)
	RequestUnreactAMessage(context.Context, *RequestUnreactMessage) (*ResponseUnreactMessage, error)
	Disconnect(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) SubscribeByUser(*User, ChatService_SubscribeByUserServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeByUser not implemented")
}
func (UnimplementedChatServiceServer) SendMessageToChat(context.Context, *RequestSendMessage) (*ResponseSendMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToChat not implemented")
}
func (UnimplementedChatServiceServer) CreateChat(context.Context, *RequestCreateChat) (*ResponseCreateChat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatServiceServer) UpdateChat(context.Context, *RequestUpdateChat) (*ResponseUpdateChat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChat not implemented")
}
func (UnimplementedChatServiceServer) GetChatsByUser(context.Context, *RequestGetChatsByUser) (*ResponseGetChatsByUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatsByUser not implemented")
}
func (UnimplementedChatServiceServer) GetUsersByChat(context.Context, *RequestGetUsersByChat) (*ResponseGetUsersByChat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersByChat not implemented")
}
func (UnimplementedChatServiceServer) GetMessagesByChat(context.Context, *RequestGetMessagesByChat) (*ResponseGetMessagesByChat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessagesByChat not implemented")
}
func (UnimplementedChatServiceServer) AddUserToChat(context.Context, *RequestAddUserToChat) (*ResponseAddUserToChat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToChat not implemented")
}
func (UnimplementedChatServiceServer) RemoveUserFromChat(context.Context, *RequestRemoveUserFromChat) (*ResponseRemoveUserFromChat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserFromChat not implemented")
}
func (UnimplementedChatServiceServer) MakeUserAdmin(context.Context, *RequestMakeUserAdmin) (*ResponseMakeUserAdmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeUserAdmin not implemented")
}
func (UnimplementedChatServiceServer) RemoveUserAdmin(context.Context, *RequestRemoveUserAdmin) (*ResponseRemoveUserAdmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserAdmin not implemented")
}
func (UnimplementedChatServiceServer) RequestReactAMessage(context.Context, *RequestReactMessage) (*ResponseReactMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestReactAMessage not implemented")
}
func (UnimplementedChatServiceServer) RequestUnreactAMessage(context.Context, *RequestUnreactMessage) (*ResponseUnreactMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestUnreactAMessage not implemented")
}
func (UnimplementedChatServiceServer) Disconnect(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_SubscribeByUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).SubscribeByUser(m, &chatServiceSubscribeByUserServer{stream})
}

type ChatService_SubscribeByUserServer interface {
	Send(*StreamByUser) error
	grpc.ServerStream
}

type chatServiceSubscribeByUserServer struct {
	grpc.ServerStream
}

func (x *chatServiceSubscribeByUserServer) Send(m *StreamByUser) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_SendMessageToChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSendMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendMessageToChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.ChatService/SendMessageToChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendMessageToChat(ctx, req.(*RequestSendMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreateChat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.ChatService/CreateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateChat(ctx, req.(*RequestCreateChat))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUpdateChat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.ChatService/UpdateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateChat(ctx, req.(*RequestUpdateChat))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChatsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetChatsByUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChatsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.ChatService/GetChatsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChatsByUser(ctx, req.(*RequestGetChatsByUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetUsersByChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetUsersByChat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetUsersByChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.ChatService/GetUsersByChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetUsersByChat(ctx, req.(*RequestGetUsersByChat))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetMessagesByChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetMessagesByChat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetMessagesByChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.ChatService/GetMessagesByChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetMessagesByChat(ctx, req.(*RequestGetMessagesByChat))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_AddUserToChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAddUserToChat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddUserToChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.ChatService/AddUserToChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddUserToChat(ctx, req.(*RequestAddUserToChat))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_RemoveUserFromChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRemoveUserFromChat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RemoveUserFromChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.ChatService/RemoveUserFromChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RemoveUserFromChat(ctx, req.(*RequestRemoveUserFromChat))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_MakeUserAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMakeUserAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).MakeUserAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.ChatService/MakeUserAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).MakeUserAdmin(ctx, req.(*RequestMakeUserAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_RemoveUserAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRemoveUserAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RemoveUserAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.ChatService/RemoveUserAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RemoveUserAdmin(ctx, req.(*RequestRemoveUserAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_RequestReactAMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestReactMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RequestReactAMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.ChatService/RequestReactAMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RequestReactAMessage(ctx, req.(*RequestReactMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_RequestUnreactAMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUnreactMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RequestUnreactAMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.ChatService/RequestUnreactAMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RequestUnreactAMessage(ctx, req.(*RequestUnreactMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.ChatService/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Disconnect(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatservice.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessageToChat",
			Handler:    _ChatService_SendMessageToChat_Handler,
		},
		{
			MethodName: "CreateChat",
			Handler:    _ChatService_CreateChat_Handler,
		},
		{
			MethodName: "UpdateChat",
			Handler:    _ChatService_UpdateChat_Handler,
		},
		{
			MethodName: "GetChatsByUser",
			Handler:    _ChatService_GetChatsByUser_Handler,
		},
		{
			MethodName: "GetUsersByChat",
			Handler:    _ChatService_GetUsersByChat_Handler,
		},
		{
			MethodName: "GetMessagesByChat",
			Handler:    _ChatService_GetMessagesByChat_Handler,
		},
		{
			MethodName: "AddUserToChat",
			Handler:    _ChatService_AddUserToChat_Handler,
		},
		{
			MethodName: "RemoveUserFromChat",
			Handler:    _ChatService_RemoveUserFromChat_Handler,
		},
		{
			MethodName: "MakeUserAdmin",
			Handler:    _ChatService_MakeUserAdmin_Handler,
		},
		{
			MethodName: "RemoveUserAdmin",
			Handler:    _ChatService_RemoveUserAdmin_Handler,
		},
		{
			MethodName: "RequestReactAMessage",
			Handler:    _ChatService_RequestReactAMessage_Handler,
		},
		{
			MethodName: "RequestUnreactAMessage",
			Handler:    _ChatService_RequestUnreactAMessage_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _ChatService_Disconnect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeByUser",
			Handler:       _ChatService_SubscribeByUser_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/file.proto",
}

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *RequestUserLogin, opts ...grpc.CallOption) (*User, error)
	Register(ctx context.Context, in *RequestUserRegister, opts ...grpc.CallOption) (*User, error)
	Logout(ctx context.Context, in *RequestUserLogout, opts ...grpc.CallOption) (*Empty, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *RequestUserLogin, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/chatservice.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Register(ctx context.Context, in *RequestUserRegister, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/chatservice.AuthService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *RequestUserLogout, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/chatservice.AuthService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Login(context.Context, *RequestUserLogin) (*User, error)
	Register(context.Context, *RequestUserRegister) (*User, error)
	Logout(context.Context, *RequestUserLogout) (*Empty, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Login(context.Context, *RequestUserLogin) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) Register(context.Context, *RequestUserRegister) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServiceServer) Logout(context.Context, *RequestUserLogout) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*RequestUserLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.AuthService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Register(ctx, req.(*RequestUserRegister))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserLogout)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatservice.AuthService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*RequestUserLogout))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatservice.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AuthService_Register_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthService_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/file.proto",
}
