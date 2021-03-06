// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/ad_group_audience_view_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
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

// Request message for [AdGroupAudienceViewService.GetAdGoupAudienceView][].
type GetAdGroupAudienceViewRequest struct {
	// The resource name of the ad group audience view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupAudienceViewRequest) Reset()         { *m = GetAdGroupAudienceViewRequest{} }
func (m *GetAdGroupAudienceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupAudienceViewRequest) ProtoMessage()    {}
func (*GetAdGroupAudienceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_audience_view_service_606af6b4aae1215b, []int{0}
}
func (m *GetAdGroupAudienceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupAudienceViewRequest.Unmarshal(m, b)
}
func (m *GetAdGroupAudienceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupAudienceViewRequest.Marshal(b, m, deterministic)
}
func (dst *GetAdGroupAudienceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupAudienceViewRequest.Merge(dst, src)
}
func (m *GetAdGroupAudienceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupAudienceViewRequest.Size(m)
}
func (m *GetAdGroupAudienceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupAudienceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupAudienceViewRequest proto.InternalMessageInfo

func (m *GetAdGroupAudienceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupAudienceViewRequest)(nil), "google.ads.googleads.v0.services.GetAdGroupAudienceViewRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdGroupAudienceViewServiceClient is the client API for AdGroupAudienceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupAudienceViewServiceClient interface {
	// Returns the requested ad group audience view in full detail.
	GetAdGroupAudienceView(ctx context.Context, in *GetAdGroupAudienceViewRequest, opts ...grpc.CallOption) (*resources.AdGroupAudienceView, error)
}

type adGroupAudienceViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdGroupAudienceViewServiceClient(cc *grpc.ClientConn) AdGroupAudienceViewServiceClient {
	return &adGroupAudienceViewServiceClient{cc}
}

func (c *adGroupAudienceViewServiceClient) GetAdGroupAudienceView(ctx context.Context, in *GetAdGroupAudienceViewRequest, opts ...grpc.CallOption) (*resources.AdGroupAudienceView, error) {
	out := new(resources.AdGroupAudienceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.AdGroupAudienceViewService/GetAdGroupAudienceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupAudienceViewServiceServer is the server API for AdGroupAudienceViewService service.
type AdGroupAudienceViewServiceServer interface {
	// Returns the requested ad group audience view in full detail.
	GetAdGroupAudienceView(context.Context, *GetAdGroupAudienceViewRequest) (*resources.AdGroupAudienceView, error)
}

func RegisterAdGroupAudienceViewServiceServer(s *grpc.Server, srv AdGroupAudienceViewServiceServer) {
	s.RegisterService(&_AdGroupAudienceViewService_serviceDesc, srv)
}

func _AdGroupAudienceViewService_GetAdGroupAudienceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupAudienceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAudienceViewServiceServer).GetAdGroupAudienceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.AdGroupAudienceViewService/GetAdGroupAudienceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAudienceViewServiceServer).GetAdGroupAudienceView(ctx, req.(*GetAdGroupAudienceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupAudienceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.AdGroupAudienceViewService",
	HandlerType: (*AdGroupAudienceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupAudienceView",
			Handler:    _AdGroupAudienceViewService_GetAdGroupAudienceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/ad_group_audience_view_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/ad_group_audience_view_service.proto", fileDescriptor_ad_group_audience_view_service_606af6b4aae1215b)
}

var fileDescriptor_ad_group_audience_view_service_606af6b4aae1215b = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x25, 0x7d, 0xf0, 0xe0, 0x85, 0xf7, 0x36, 0x59, 0x3c, 0x24, 0x28, 0x96, 0xea, 0x42, 0xba,
	0x98, 0x09, 0x0a, 0x45, 0x10, 0x2b, 0x29, 0x4a, 0x5c, 0x49, 0xa9, 0xd0, 0x85, 0x04, 0xc2, 0x98,
	0x5c, 0xc2, 0x40, 0x33, 0x13, 0xe7, 0x26, 0xe9, 0x42, 0xdc, 0xf8, 0x0b, 0xfe, 0x81, 0x4b, 0x3f,
	0xc5, 0x9d, 0xf8, 0x01, 0x6e, 0xfc, 0x00, 0x3f, 0x41, 0xd2, 0xc9, 0x14, 0xc4, 0xc6, 0xee, 0x0e,
	0x73, 0xef, 0x39, 0xe7, 0xde, 0x73, 0xc7, 0x3e, 0x4b, 0xa5, 0x4c, 0x67, 0x40, 0x59, 0x82, 0x54,
	0xc3, 0x1a, 0x55, 0x1e, 0x45, 0x50, 0x15, 0x8f, 0x01, 0x29, 0x4b, 0xa2, 0x54, 0xc9, 0x32, 0x8f,
	0x58, 0x99, 0x70, 0x10, 0x31, 0x44, 0x15, 0x87, 0x79, 0xd4, 0xd4, 0x49, 0xae, 0x64, 0x21, 0x9d,
	0xae, 0xe6, 0x12, 0x96, 0x20, 0x59, 0xca, 0x90, 0xca, 0x23, 0x46, 0xc6, 0x1d, 0xb6, 0x19, 0x29,
	0x40, 0x59, 0xaa, 0x76, 0x27, 0xed, 0xe0, 0x6e, 0x1a, 0x7e, 0xce, 0x29, 0x13, 0x42, 0x16, 0xac,
	0xe0, 0x52, 0xa0, 0xae, 0xf6, 0x4e, 0xed, 0xad, 0x00, 0x0a, 0x3f, 0x09, 0x6a, 0xbe, 0xdf, 0xd0,
	0xa7, 0x1c, 0xe6, 0x13, 0xb8, 0x29, 0x01, 0x0b, 0x67, 0xc7, 0xfe, 0x67, 0x8c, 0x22, 0xc1, 0x32,
	0xd8, 0xb0, 0xba, 0xd6, 0xde, 0x9f, 0xc9, 0x5f, 0xf3, 0x78, 0xc1, 0x32, 0xd8, 0xff, 0xb0, 0x6c,
	0x77, 0x85, 0xc6, 0xa5, 0xde, 0xc1, 0x79, 0xb1, 0xec, 0xff, 0xab, 0x5d, 0x9c, 0x13, 0xb2, 0x2e,
	0x00, 0xf2, 0xe3, 0x7c, 0xee, 0xa0, 0x55, 0x60, 0x99, 0x0f, 0x59, 0x41, 0xef, 0x0d, 0xef, 0x5f,
	0xdf, 0x1f, 0x3a, 0x87, 0xce, 0xa0, 0x8e, 0xf2, 0xf6, 0xcb, 0x8a, 0xc7, 0x71, 0x89, 0x85, 0xcc,
	0x40, 0x21, 0xed, 0x53, 0xf6, 0x9d, 0x8b, 0xb4, 0x7f, 0x37, 0x7a, 0xb3, 0xec, 0xdd, 0x58, 0x66,
	0x6b, 0xc7, 0x1f, 0x6d, 0xb7, 0x07, 0x33, 0xae, 0x4f, 0x30, 0xb6, 0xae, 0xce, 0x1b, 0x91, 0x54,
	0xce, 0x98, 0x48, 0x89, 0x54, 0x29, 0x4d, 0x41, 0x2c, 0x0e, 0x64, 0x4e, 0x9e, 0x73, 0x6c, 0xff,
	0x6a, 0x47, 0x06, 0x3c, 0x76, 0x7e, 0x05, 0xbe, 0xff, 0xd4, 0xe9, 0x06, 0x5a, 0xd0, 0x4f, 0x90,
	0x68, 0x58, 0xa3, 0xa9, 0x47, 0x1a, 0x63, 0x7c, 0x36, 0x2d, 0xa1, 0x9f, 0x60, 0xb8, 0x6c, 0x09,
	0xa7, 0x5e, 0x68, 0x5a, 0xae, 0x7f, 0x2f, 0x06, 0x38, 0xf8, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xf0,
	0x63, 0x70, 0x60, 0xea, 0x02, 0x00, 0x00,
}
