// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: userserver.proto

package grpcconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	_go "github.com/sweetloveinyourheart/planning-pocker/proto/code/userserver/go"
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
	// UserServerName is the fully-qualified name of the UserServer service.
	UserServerName = "com.sweetloveinyourheart.pocker.users.UserServer"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServerGetUserProcedure is the fully-qualified name of the UserServer's GetUser RPC.
	UserServerGetUserProcedure = "/com.sweetloveinyourheart.pocker.users.UserServer/GetUser"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	userServerServiceDescriptor       = _go.File_userserver_proto.Services().ByName("UserServer")
	userServerGetUserMethodDescriptor = userServerServiceDescriptor.Methods().ByName("GetUser")
)

// UserServerClient is a client for the com.sweetloveinyourheart.pocker.users.UserServer service.
type UserServerClient interface {
	// Get a user by user_id
	GetUser(context.Context, *connect.Request[_go.GetUserRequest]) (*connect.Response[_go.GetUserResponse], error)
}

// NewUserServerClient constructs a client for the com.sweetloveinyourheart.pocker.users.UserServer
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServerClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserServerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServerClient{
		getUser: connect.NewClient[_go.GetUserRequest, _go.GetUserResponse](
			httpClient,
			baseURL+UserServerGetUserProcedure,
			connect.WithSchema(userServerGetUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// userServerClient implements UserServerClient.
type userServerClient struct {
	getUser *connect.Client[_go.GetUserRequest, _go.GetUserResponse]
}

// GetUser calls com.sweetloveinyourheart.pocker.users.UserServer.GetUser.
func (c *userServerClient) GetUser(ctx context.Context, req *connect.Request[_go.GetUserRequest]) (*connect.Response[_go.GetUserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// UserServerHandler is an implementation of the com.sweetloveinyourheart.pocker.users.UserServer
// service.
type UserServerHandler interface {
	// Get a user by user_id
	GetUser(context.Context, *connect.Request[_go.GetUserRequest]) (*connect.Response[_go.GetUserResponse], error)
}

// NewUserServerHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServerHandler(svc UserServerHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServerGetUserHandler := connect.NewUnaryHandler(
		UserServerGetUserProcedure,
		svc.GetUser,
		connect.WithSchema(userServerGetUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/com.sweetloveinyourheart.pocker.users.UserServer/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServerGetUserProcedure:
			userServerGetUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServerHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServerHandler struct{}

func (UnimplementedUserServerHandler) GetUser(context.Context, *connect.Request[_go.GetUserRequest]) (*connect.Response[_go.GetUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("com.sweetloveinyourheart.pocker.users.UserServer.GetUser is not implemented"))
}
