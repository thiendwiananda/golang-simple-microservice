// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// MovieHandlerClient is the client API for MovieHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieHandlerClient interface {
	GetMovie(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*Movie, error)
	SearchMovie(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*MovieList, error)
}

type movieHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieHandlerClient(cc grpc.ClientConnInterface) MovieHandlerClient {
	return &movieHandlerClient{cc}
}

func (c *movieHandlerClient) GetMovie(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, "/grpc.MovieHandler/GetMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieHandlerClient) SearchMovie(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*MovieList, error) {
	out := new(MovieList)
	err := c.cc.Invoke(ctx, "/grpc.MovieHandler/SearchMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieHandlerServer is the server API for MovieHandler service.
// All implementations must embed UnimplementedMovieHandlerServer
// for forward compatibility
type MovieHandlerServer interface {
	GetMovie(context.Context, *SingleRequest) (*Movie, error)
	SearchMovie(context.Context, *SearchRequest) (*MovieList, error)
	mustEmbedUnimplementedMovieHandlerServer()
}

// UnimplementedMovieHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedMovieHandlerServer struct {
}

func (UnimplementedMovieHandlerServer) GetMovie(context.Context, *SingleRequest) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}
func (UnimplementedMovieHandlerServer) SearchMovie(context.Context, *SearchRequest) (*MovieList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMovie not implemented")
}
func (UnimplementedMovieHandlerServer) mustEmbedUnimplementedMovieHandlerServer() {}

// UnsafeMovieHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieHandlerServer will
// result in compilation errors.
type UnsafeMovieHandlerServer interface {
	mustEmbedUnimplementedMovieHandlerServer()
}

func RegisterMovieHandlerServer(s grpc.ServiceRegistrar, srv MovieHandlerServer) {
	s.RegisterService(&MovieHandler_ServiceDesc, srv)
}

func _MovieHandler_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieHandlerServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MovieHandler/GetMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieHandlerServer).GetMovie(ctx, req.(*SingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieHandler_SearchMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieHandlerServer).SearchMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MovieHandler/SearchMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieHandlerServer).SearchMovie(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieHandler_ServiceDesc is the grpc.ServiceDesc for MovieHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.MovieHandler",
	HandlerType: (*MovieHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovie",
			Handler:    _MovieHandler_GetMovie_Handler,
		},
		{
			MethodName: "SearchMovie",
			Handler:    _MovieHandler_SearchMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie.proto",
}
