// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: authorization/v1/authorization.proto

package authorizationv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/sansan36/authorization-service/gen/authorization/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AuthorizationServiceName is the fully-qualified name of the AuthorizationService service.
	AuthorizationServiceName = "authorization.v1.AuthorizationService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthorizationServiceGetUserListProcedure is the fully-qualified name of the
	// AuthorizationService's GetUserList RPC.
	AuthorizationServiceGetUserListProcedure = "/authorization.v1.AuthorizationService/GetUserList"
	// AuthorizationServiceGetUserProcedure is the fully-qualified name of the AuthorizationService's
	// GetUser RPC.
	AuthorizationServiceGetUserProcedure = "/authorization.v1.AuthorizationService/GetUser"
	// AuthorizationServiceAddUserProcedure is the fully-qualified name of the AuthorizationService's
	// AddUser RPC.
	AuthorizationServiceAddUserProcedure = "/authorization.v1.AuthorizationService/AddUser"
	// AuthorizationServiceEditUserProcedure is the fully-qualified name of the AuthorizationService's
	// EditUser RPC.
	AuthorizationServiceEditUserProcedure = "/authorization.v1.AuthorizationService/EditUser"
	// AuthorizationServiceRemoveUserProcedure is the fully-qualified name of the AuthorizationService's
	// RemoveUser RPC.
	AuthorizationServiceRemoveUserProcedure = "/authorization.v1.AuthorizationService/RemoveUser"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	authorizationServiceServiceDescriptor           = v1.File_authorization_v1_authorization_proto.Services().ByName("AuthorizationService")
	authorizationServiceGetUserListMethodDescriptor = authorizationServiceServiceDescriptor.Methods().ByName("GetUserList")
	authorizationServiceGetUserMethodDescriptor     = authorizationServiceServiceDescriptor.Methods().ByName("GetUser")
	authorizationServiceAddUserMethodDescriptor     = authorizationServiceServiceDescriptor.Methods().ByName("AddUser")
	authorizationServiceEditUserMethodDescriptor    = authorizationServiceServiceDescriptor.Methods().ByName("EditUser")
	authorizationServiceRemoveUserMethodDescriptor  = authorizationServiceServiceDescriptor.Methods().ByName("RemoveUser")
)

// AuthorizationServiceClient is a client for the authorization.v1.AuthorizationService service.
type AuthorizationServiceClient interface {
	GetUserList(context.Context, *connect.Request[v1.GetUserListRequest]) (*connect.Response[v1.GetUserListResponse], error)
	GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error)
	AddUser(context.Context, *connect.Request[v1.AddUserRequest]) (*connect.Response[v1.AddUserResponse], error)
	EditUser(context.Context, *connect.Request[v1.EditUserRequest]) (*connect.Response[v1.EditUserResponse], error)
	RemoveUser(context.Context, *connect.Request[v1.RemoveUserRequest]) (*connect.Response[v1.RemoveUserResponse], error)
}

// NewAuthorizationServiceClient constructs a client for the authorization.v1.AuthorizationService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthorizationServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuthorizationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authorizationServiceClient{
		getUserList: connect.NewClient[v1.GetUserListRequest, v1.GetUserListResponse](
			httpClient,
			baseURL+AuthorizationServiceGetUserListProcedure,
			connect.WithSchema(authorizationServiceGetUserListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getUser: connect.NewClient[v1.GetUserRequest, v1.GetUserResponse](
			httpClient,
			baseURL+AuthorizationServiceGetUserProcedure,
			connect.WithSchema(authorizationServiceGetUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		addUser: connect.NewClient[v1.AddUserRequest, v1.AddUserResponse](
			httpClient,
			baseURL+AuthorizationServiceAddUserProcedure,
			connect.WithSchema(authorizationServiceAddUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		editUser: connect.NewClient[v1.EditUserRequest, v1.EditUserResponse](
			httpClient,
			baseURL+AuthorizationServiceEditUserProcedure,
			connect.WithSchema(authorizationServiceEditUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		removeUser: connect.NewClient[v1.RemoveUserRequest, v1.RemoveUserResponse](
			httpClient,
			baseURL+AuthorizationServiceRemoveUserProcedure,
			connect.WithSchema(authorizationServiceRemoveUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// authorizationServiceClient implements AuthorizationServiceClient.
type authorizationServiceClient struct {
	getUserList *connect.Client[v1.GetUserListRequest, v1.GetUserListResponse]
	getUser     *connect.Client[v1.GetUserRequest, v1.GetUserResponse]
	addUser     *connect.Client[v1.AddUserRequest, v1.AddUserResponse]
	editUser    *connect.Client[v1.EditUserRequest, v1.EditUserResponse]
	removeUser  *connect.Client[v1.RemoveUserRequest, v1.RemoveUserResponse]
}

// GetUserList calls authorization.v1.AuthorizationService.GetUserList.
func (c *authorizationServiceClient) GetUserList(ctx context.Context, req *connect.Request[v1.GetUserListRequest]) (*connect.Response[v1.GetUserListResponse], error) {
	return c.getUserList.CallUnary(ctx, req)
}

// GetUser calls authorization.v1.AuthorizationService.GetUser.
func (c *authorizationServiceClient) GetUser(ctx context.Context, req *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// AddUser calls authorization.v1.AuthorizationService.AddUser.
func (c *authorizationServiceClient) AddUser(ctx context.Context, req *connect.Request[v1.AddUserRequest]) (*connect.Response[v1.AddUserResponse], error) {
	return c.addUser.CallUnary(ctx, req)
}

// EditUser calls authorization.v1.AuthorizationService.EditUser.
func (c *authorizationServiceClient) EditUser(ctx context.Context, req *connect.Request[v1.EditUserRequest]) (*connect.Response[v1.EditUserResponse], error) {
	return c.editUser.CallUnary(ctx, req)
}

// RemoveUser calls authorization.v1.AuthorizationService.RemoveUser.
func (c *authorizationServiceClient) RemoveUser(ctx context.Context, req *connect.Request[v1.RemoveUserRequest]) (*connect.Response[v1.RemoveUserResponse], error) {
	return c.removeUser.CallUnary(ctx, req)
}

// AuthorizationServiceHandler is an implementation of the authorization.v1.AuthorizationService
// service.
type AuthorizationServiceHandler interface {
	GetUserList(context.Context, *connect.Request[v1.GetUserListRequest]) (*connect.Response[v1.GetUserListResponse], error)
	GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error)
	AddUser(context.Context, *connect.Request[v1.AddUserRequest]) (*connect.Response[v1.AddUserResponse], error)
	EditUser(context.Context, *connect.Request[v1.EditUserRequest]) (*connect.Response[v1.EditUserResponse], error)
	RemoveUser(context.Context, *connect.Request[v1.RemoveUserRequest]) (*connect.Response[v1.RemoveUserResponse], error)
}

// NewAuthorizationServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthorizationServiceHandler(svc AuthorizationServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	authorizationServiceGetUserListHandler := connect.NewUnaryHandler(
		AuthorizationServiceGetUserListProcedure,
		svc.GetUserList,
		connect.WithSchema(authorizationServiceGetUserListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authorizationServiceGetUserHandler := connect.NewUnaryHandler(
		AuthorizationServiceGetUserProcedure,
		svc.GetUser,
		connect.WithSchema(authorizationServiceGetUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authorizationServiceAddUserHandler := connect.NewUnaryHandler(
		AuthorizationServiceAddUserProcedure,
		svc.AddUser,
		connect.WithSchema(authorizationServiceAddUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authorizationServiceEditUserHandler := connect.NewUnaryHandler(
		AuthorizationServiceEditUserProcedure,
		svc.EditUser,
		connect.WithSchema(authorizationServiceEditUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	authorizationServiceRemoveUserHandler := connect.NewUnaryHandler(
		AuthorizationServiceRemoveUserProcedure,
		svc.RemoveUser,
		connect.WithSchema(authorizationServiceRemoveUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/authorization.v1.AuthorizationService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthorizationServiceGetUserListProcedure:
			authorizationServiceGetUserListHandler.ServeHTTP(w, r)
		case AuthorizationServiceGetUserProcedure:
			authorizationServiceGetUserHandler.ServeHTTP(w, r)
		case AuthorizationServiceAddUserProcedure:
			authorizationServiceAddUserHandler.ServeHTTP(w, r)
		case AuthorizationServiceEditUserProcedure:
			authorizationServiceEditUserHandler.ServeHTTP(w, r)
		case AuthorizationServiceRemoveUserProcedure:
			authorizationServiceRemoveUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthorizationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthorizationServiceHandler struct{}

func (UnimplementedAuthorizationServiceHandler) GetUserList(context.Context, *connect.Request[v1.GetUserListRequest]) (*connect.Response[v1.GetUserListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("authorization.v1.AuthorizationService.GetUserList is not implemented"))
}

func (UnimplementedAuthorizationServiceHandler) GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.GetUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("authorization.v1.AuthorizationService.GetUser is not implemented"))
}

func (UnimplementedAuthorizationServiceHandler) AddUser(context.Context, *connect.Request[v1.AddUserRequest]) (*connect.Response[v1.AddUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("authorization.v1.AuthorizationService.AddUser is not implemented"))
}

func (UnimplementedAuthorizationServiceHandler) EditUser(context.Context, *connect.Request[v1.EditUserRequest]) (*connect.Response[v1.EditUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("authorization.v1.AuthorizationService.EditUser is not implemented"))
}

func (UnimplementedAuthorizationServiceHandler) RemoveUser(context.Context, *connect.Request[v1.RemoveUserRequest]) (*connect.Response[v1.RemoveUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("authorization.v1.AuthorizationService.RemoveUser is not implemented"))
}
