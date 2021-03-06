// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package docker_downloadpb

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

// DockerDownloadServiceClient is the client API for DockerDownloadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DockerDownloadServiceClient interface {
	Download(ctx context.Context, in *DockerDownloadRequest, opts ...grpc.CallOption) (*DockerDownloadResponse, error)
}

type dockerDownloadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDockerDownloadServiceClient(cc grpc.ClientConnInterface) DockerDownloadServiceClient {
	return &dockerDownloadServiceClient{cc}
}

func (c *dockerDownloadServiceClient) Download(ctx context.Context, in *DockerDownloadRequest, opts ...grpc.CallOption) (*DockerDownloadResponse, error) {
	out := new(DockerDownloadResponse)
	err := c.cc.Invoke(ctx, "/docker_download.DockerDownloadService/Download", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DockerDownloadServiceServer is the server API for DockerDownloadService service.
// All implementations must embed UnimplementedDockerDownloadServiceServer
// for forward compatibility
type DockerDownloadServiceServer interface {
	Download(context.Context, *DockerDownloadRequest) (*DockerDownloadResponse, error)
	mustEmbedUnimplementedDockerDownloadServiceServer()
}

// UnimplementedDockerDownloadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDockerDownloadServiceServer struct {
}

func (UnimplementedDockerDownloadServiceServer) Download(context.Context, *DockerDownloadRequest) (*DockerDownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedDockerDownloadServiceServer) mustEmbedUnimplementedDockerDownloadServiceServer() {}

// UnsafeDockerDownloadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DockerDownloadServiceServer will
// result in compilation errors.
type UnsafeDockerDownloadServiceServer interface {
	mustEmbedUnimplementedDockerDownloadServiceServer()
}

func RegisterDockerDownloadServiceServer(s grpc.ServiceRegistrar, srv DockerDownloadServiceServer) {
	s.RegisterService(&DockerDownloadService_ServiceDesc, srv)
}

func _DockerDownloadService_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockerDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerDownloadServiceServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker_download.DockerDownloadService/Download",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerDownloadServiceServer).Download(ctx, req.(*DockerDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DockerDownloadService_ServiceDesc is the grpc.ServiceDesc for DockerDownloadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DockerDownloadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "docker_download.DockerDownloadService",
	HandlerType: (*DockerDownloadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Download",
			Handler:    _DockerDownloadService_Download_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "docker_download.proto",
}
