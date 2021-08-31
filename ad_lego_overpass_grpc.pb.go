// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ad_lego_overpass

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

// LegoServiceClient is the client API for LegoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LegoServiceClient interface {
	UploadImage(ctx context.Context, in *UploadImageRequest, opts ...grpc.CallOption) (*UploadImageResponse, error)
}

type legoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLegoServiceClient(cc grpc.ClientConnInterface) LegoServiceClient {
	return &legoServiceClient{cc}
}

func (c *legoServiceClient) UploadImage(ctx context.Context, in *UploadImageRequest, opts ...grpc.CallOption) (*UploadImageResponse, error) {
	out := new(UploadImageResponse)
	err := c.cc.Invoke(ctx, "/LegoService/UploadImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LegoServiceServer is the server API for LegoService service.
// All implementations must embed UnimplementedLegoServiceServer
// for forward compatibility
type LegoServiceServer interface {
	UploadImage(context.Context, *UploadImageRequest) (*UploadImageResponse, error)
	mustEmbedUnimplementedLegoServiceServer()
}

// UnimplementedLegoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLegoServiceServer struct {
}

func (UnimplementedLegoServiceServer) UploadImage(context.Context, *UploadImageRequest) (*UploadImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedLegoServiceServer) mustEmbedUnimplementedLegoServiceServer() {}

// UnsafeLegoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LegoServiceServer will
// result in compilation errors.
type UnsafeLegoServiceServer interface {
	mustEmbedUnimplementedLegoServiceServer()
}

func RegisterLegoServiceServer(s grpc.ServiceRegistrar, srv LegoServiceServer) {
	s.RegisterService(&LegoService_ServiceDesc, srv)
}

func _LegoService_UploadImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LegoServiceServer).UploadImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LegoService/UploadImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LegoServiceServer).UploadImage(ctx, req.(*UploadImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LegoService_ServiceDesc is the grpc.ServiceDesc for LegoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LegoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LegoService",
	HandlerType: (*LegoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadImage",
			Handler:    _LegoService_UploadImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ad_lego_overpass.proto",
}
