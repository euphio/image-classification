// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: image_recognition.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ImageLabelService_GetLabels_FullMethodName = "/image_recognition.ImageLabelService/GetLabels"
)

// ImageLabelServiceClient is the client API for ImageLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageLabelServiceClient interface {
	GetLabels(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
}

type imageLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageLabelServiceClient(cc grpc.ClientConnInterface) ImageLabelServiceClient {
	return &imageLabelServiceClient{cc}
}

func (c *imageLabelServiceClient) GetLabels(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, ImageLabelService_GetLabels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageLabelServiceServer is the server API for ImageLabelService service.
// All implementations must embed UnimplementedImageLabelServiceServer
// for forward compatibility.
type ImageLabelServiceServer interface {
	GetLabels(context.Context, *ImageRequest) (*ImageResponse, error)
	mustEmbedUnimplementedImageLabelServiceServer()
}

// UnimplementedImageLabelServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedImageLabelServiceServer struct{}

func (UnimplementedImageLabelServiceServer) GetLabels(context.Context, *ImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabels not implemented")
}
func (UnimplementedImageLabelServiceServer) mustEmbedUnimplementedImageLabelServiceServer() {}
func (UnimplementedImageLabelServiceServer) testEmbeddedByValue()                           {}

// UnsafeImageLabelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageLabelServiceServer will
// result in compilation errors.
type UnsafeImageLabelServiceServer interface {
	mustEmbedUnimplementedImageLabelServiceServer()
}

func RegisterImageLabelServiceServer(s grpc.ServiceRegistrar, srv ImageLabelServiceServer) {
	// If the following call pancis, it indicates UnimplementedImageLabelServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ImageLabelService_ServiceDesc, srv)
}

func _ImageLabelService_GetLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageLabelServiceServer).GetLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageLabelService_GetLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageLabelServiceServer).GetLabels(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageLabelService_ServiceDesc is the grpc.ServiceDesc for ImageLabelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageLabelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "image_recognition.ImageLabelService",
	HandlerType: (*ImageLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLabels",
			Handler:    _ImageLabelService_GetLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image_recognition.proto",
}
