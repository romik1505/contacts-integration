// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/contact_service/contact_service.proto

package contact

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ContactService_AuthIntegration_FullMethodName         = "/contact_service.ContactService/AuthIntegration"
	ContactService_ListContacts_FullMethodName            = "/contact_service.ContactService/ListContacts"
	ContactService_ListAccounts_FullMethodName            = "/contact_service.ContactService/ListAccounts"
	ContactService_ListAccountIntegrations_FullMethodName = "/contact_service.ContactService/ListAccountIntegrations"
	ContactService_GetAccount_FullMethodName              = "/contact_service.ContactService/GetAccount"
	ContactService_UnsubAccount_FullMethodName            = "/contact_service.ContactService/UnsubAccount"
)

// ContactServiceClient is the client API for ContactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContactServiceClient interface {
	// Добавление виджета
	AuthIntegration(ctx context.Context, in *AuthIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Список контактов в amoCRM
	ListContacts(ctx context.Context, in *ListContactsRequest, opts ...grpc.CallOption) (*ListContactsResponse, error)
	// Список учетных записей
	ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error)
	// Список интеграций аккаунта
	ListAccountIntegrations(ctx context.Context, in *ListAccountIntegrationsRequest, opts ...grpc.CallOption) (*ListAccountIntegrationsResponse, error)
	// Информация об аккаунте
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	// Отписка учетной записи
	UnsubAccount(ctx context.Context, in *UnsubAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type contactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContactServiceClient(cc grpc.ClientConnInterface) ContactServiceClient {
	return &contactServiceClient{cc}
}

func (c *contactServiceClient) AuthIntegration(ctx context.Context, in *AuthIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContactService_AuthIntegration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) ListContacts(ctx context.Context, in *ListContactsRequest, opts ...grpc.CallOption) (*ListContactsResponse, error) {
	out := new(ListContactsResponse)
	err := c.cc.Invoke(ctx, ContactService_ListContacts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error) {
	out := new(ListAccountsResponse)
	err := c.cc.Invoke(ctx, ContactService_ListAccounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) ListAccountIntegrations(ctx context.Context, in *ListAccountIntegrationsRequest, opts ...grpc.CallOption) (*ListAccountIntegrationsResponse, error) {
	out := new(ListAccountIntegrationsResponse)
	err := c.cc.Invoke(ctx, ContactService_ListAccountIntegrations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, ContactService_GetAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) UnsubAccount(ctx context.Context, in *UnsubAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContactService_UnsubAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServiceServer is the server API for ContactService service.
// All implementations should embed UnimplementedContactServiceServer
// for forward compatibility
type ContactServiceServer interface {
	// Добавление виджета
	AuthIntegration(context.Context, *AuthIntegrationRequest) (*emptypb.Empty, error)
	// Список контактов в amoCRM
	ListContacts(context.Context, *ListContactsRequest) (*ListContactsResponse, error)
	// Список учетных записей
	ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error)
	// Список интеграций аккаунта
	ListAccountIntegrations(context.Context, *ListAccountIntegrationsRequest) (*ListAccountIntegrationsResponse, error)
	// Информация об аккаунте
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	// Отписка учетной записи
	UnsubAccount(context.Context, *UnsubAccountRequest) (*emptypb.Empty, error)
}

// UnimplementedContactServiceServer should be embedded to have forward compatible implementations.
type UnimplementedContactServiceServer struct {
}

func (UnimplementedContactServiceServer) AuthIntegration(context.Context, *AuthIntegrationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthIntegration not implemented")
}
func (UnimplementedContactServiceServer) ListContacts(context.Context, *ListContactsRequest) (*ListContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContacts not implemented")
}
func (UnimplementedContactServiceServer) ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}
func (UnimplementedContactServiceServer) ListAccountIntegrations(context.Context, *ListAccountIntegrationsRequest) (*ListAccountIntegrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccountIntegrations not implemented")
}
func (UnimplementedContactServiceServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedContactServiceServer) UnsubAccount(context.Context, *UnsubAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsubAccount not implemented")
}

// UnsafeContactServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContactServiceServer will
// result in compilation errors.
type UnsafeContactServiceServer interface {
	mustEmbedUnimplementedContactServiceServer()
}

func RegisterContactServiceServer(s grpc.ServiceRegistrar, srv ContactServiceServer) {
	s.RegisterService(&ContactService_ServiceDesc, srv)
}

func _ContactService_AuthIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).AuthIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_AuthIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).AuthIntegration(ctx, req.(*AuthIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_ListContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).ListContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_ListContacts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).ListContacts(ctx, req.(*ListContactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_ListAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).ListAccounts(ctx, req.(*ListAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_ListAccountIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountIntegrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).ListAccountIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_ListAccountIntegrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).ListAccountIntegrations(ctx, req.(*ListAccountIntegrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_UnsubAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).UnsubAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContactService_UnsubAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).UnsubAccount(ctx, req.(*UnsubAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContactService_ServiceDesc is the grpc.ServiceDesc for ContactService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContactService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contact_service.ContactService",
	HandlerType: (*ContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthIntegration",
			Handler:    _ContactService_AuthIntegration_Handler,
		},
		{
			MethodName: "ListContacts",
			Handler:    _ContactService_ListContacts_Handler,
		},
		{
			MethodName: "ListAccounts",
			Handler:    _ContactService_ListAccounts_Handler,
		},
		{
			MethodName: "ListAccountIntegrations",
			Handler:    _ContactService_ListAccountIntegrations_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _ContactService_GetAccount_Handler,
		},
		{
			MethodName: "UnsubAccount",
			Handler:    _ContactService_UnsubAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/contact_service/contact_service.proto",
}
