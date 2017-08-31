// Code generated by protoc-gen-go.
// source: box.proto
// DO NOT EDIT!

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	box.proto

It has these top-level messages:
	None
	Box
	Boxes
	Link
	Links
*/
package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type None struct {
}

func (m *None) Reset()                    { *m = None{} }
func (m *None) String() string            { return proto.CompactTextString(m) }
func (*None) ProtoMessage()               {}
func (*None) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Box struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *Box) Reset()                    { *m = Box{} }
func (m *Box) String() string            { return proto.CompactTextString(m) }
func (*Box) ProtoMessage()               {}
func (*Box) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Box) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Box) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Box) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Boxes struct {
	Boxes []*Box `protobuf:"bytes,1,rep,name=boxes" json:"boxes,omitempty"`
}

func (m *Boxes) Reset()                    { *m = Boxes{} }
func (m *Boxes) String() string            { return proto.CompactTextString(m) }
func (*Boxes) ProtoMessage()               {}
func (*Boxes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Boxes) GetBoxes() []*Box {
	if m != nil {
		return m.Boxes
	}
	return nil
}

type Link struct {
	Id          string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Url         string   `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Tags        []string `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
	BoxId       string   `protobuf:"bytes,6,opt,name=boxId" json:"boxId,omitempty"`
	CreatedOn   int64    `protobuf:"varint,7,opt,name=createdOn" json:"createdOn,omitempty"`
}

func (m *Link) Reset()                    { *m = Link{} }
func (m *Link) String() string            { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()               {}
func (*Link) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Link) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Link) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Link) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Link) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Link) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Link) GetBoxId() string {
	if m != nil {
		return m.BoxId
	}
	return ""
}

func (m *Link) GetCreatedOn() int64 {
	if m != nil {
		return m.CreatedOn
	}
	return 0
}

type Links struct {
	Links []*Link `protobuf:"bytes,1,rep,name=links" json:"links,omitempty"`
}

func (m *Links) Reset()                    { *m = Links{} }
func (m *Links) String() string            { return proto.CompactTextString(m) }
func (*Links) ProtoMessage()               {}
func (*Links) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Links) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func init() {
	proto.RegisterType((*None)(nil), "msg.None")
	proto.RegisterType((*Box)(nil), "msg.Box")
	proto.RegisterType((*Boxes)(nil), "msg.Boxes")
	proto.RegisterType((*Link)(nil), "msg.Link")
	proto.RegisterType((*Links)(nil), "msg.Links")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BoxService service

type BoxServiceClient interface {
	NewBox(ctx context.Context, in *Box, opts ...grpc.CallOption) (*Box, error)
	SaveBox(ctx context.Context, in *Box, opts ...grpc.CallOption) (*Box, error)
	GetBoxById(ctx context.Context, in *Box, opts ...grpc.CallOption) (*Box, error)
	GetBoxes(ctx context.Context, in *None, opts ...grpc.CallOption) (*Boxes, error)
	NewLink(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Link, error)
	SaveLink(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Link, error)
	GetLinkById(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Link, error)
	GetLinks(ctx context.Context, in *None, opts ...grpc.CallOption) (*Links, error)
	GetLinksByBoxId(ctx context.Context, in *Box, opts ...grpc.CallOption) (*Links, error)
}

type boxServiceClient struct {
	cc *grpc.ClientConn
}

func NewBoxServiceClient(cc *grpc.ClientConn) BoxServiceClient {
	return &boxServiceClient{cc}
}

func (c *boxServiceClient) NewBox(ctx context.Context, in *Box, opts ...grpc.CallOption) (*Box, error) {
	out := new(Box)
	err := grpc.Invoke(ctx, "/msg.BoxService/NewBox", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) SaveBox(ctx context.Context, in *Box, opts ...grpc.CallOption) (*Box, error) {
	out := new(Box)
	err := grpc.Invoke(ctx, "/msg.BoxService/SaveBox", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) GetBoxById(ctx context.Context, in *Box, opts ...grpc.CallOption) (*Box, error) {
	out := new(Box)
	err := grpc.Invoke(ctx, "/msg.BoxService/GetBoxById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) GetBoxes(ctx context.Context, in *None, opts ...grpc.CallOption) (*Boxes, error) {
	out := new(Boxes)
	err := grpc.Invoke(ctx, "/msg.BoxService/GetBoxes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) NewLink(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Link, error) {
	out := new(Link)
	err := grpc.Invoke(ctx, "/msg.BoxService/NewLink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) SaveLink(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Link, error) {
	out := new(Link)
	err := grpc.Invoke(ctx, "/msg.BoxService/SaveLink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) GetLinkById(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Link, error) {
	out := new(Link)
	err := grpc.Invoke(ctx, "/msg.BoxService/GetLinkById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) GetLinks(ctx context.Context, in *None, opts ...grpc.CallOption) (*Links, error) {
	out := new(Links)
	err := grpc.Invoke(ctx, "/msg.BoxService/GetLinks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) GetLinksByBoxId(ctx context.Context, in *Box, opts ...grpc.CallOption) (*Links, error) {
	out := new(Links)
	err := grpc.Invoke(ctx, "/msg.BoxService/GetLinksByBoxId", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BoxService service

type BoxServiceServer interface {
	NewBox(context.Context, *Box) (*Box, error)
	SaveBox(context.Context, *Box) (*Box, error)
	GetBoxById(context.Context, *Box) (*Box, error)
	GetBoxes(context.Context, *None) (*Boxes, error)
	NewLink(context.Context, *Link) (*Link, error)
	SaveLink(context.Context, *Link) (*Link, error)
	GetLinkById(context.Context, *Link) (*Link, error)
	GetLinks(context.Context, *None) (*Links, error)
	GetLinksByBoxId(context.Context, *Box) (*Links, error)
}

func RegisterBoxServiceServer(s *grpc.Server, srv BoxServiceServer) {
	s.RegisterService(&_BoxService_serviceDesc, srv)
}

func _BoxService_NewBox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Box)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).NewBox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.BoxService/NewBox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).NewBox(ctx, req.(*Box))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_SaveBox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Box)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).SaveBox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.BoxService/SaveBox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).SaveBox(ctx, req.(*Box))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_GetBoxById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Box)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).GetBoxById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.BoxService/GetBoxById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).GetBoxById(ctx, req.(*Box))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_GetBoxes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).GetBoxes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.BoxService/GetBoxes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).GetBoxes(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_NewLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Link)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).NewLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.BoxService/NewLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).NewLink(ctx, req.(*Link))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_SaveLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Link)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).SaveLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.BoxService/SaveLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).SaveLink(ctx, req.(*Link))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_GetLinkById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Link)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).GetLinkById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.BoxService/GetLinkById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).GetLinkById(ctx, req.(*Link))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_GetLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(None)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).GetLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.BoxService/GetLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).GetLinks(ctx, req.(*None))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_GetLinksByBoxId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Box)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).GetLinksByBoxId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.BoxService/GetLinksByBoxId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).GetLinksByBoxId(ctx, req.(*Box))
	}
	return interceptor(ctx, in, info, handler)
}

var _BoxService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "msg.BoxService",
	HandlerType: (*BoxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewBox",
			Handler:    _BoxService_NewBox_Handler,
		},
		{
			MethodName: "SaveBox",
			Handler:    _BoxService_SaveBox_Handler,
		},
		{
			MethodName: "GetBoxById",
			Handler:    _BoxService_GetBoxById_Handler,
		},
		{
			MethodName: "GetBoxes",
			Handler:    _BoxService_GetBoxes_Handler,
		},
		{
			MethodName: "NewLink",
			Handler:    _BoxService_NewLink_Handler,
		},
		{
			MethodName: "SaveLink",
			Handler:    _BoxService_SaveLink_Handler,
		},
		{
			MethodName: "GetLinkById",
			Handler:    _BoxService_GetLinkById_Handler,
		},
		{
			MethodName: "GetLinks",
			Handler:    _BoxService_GetLinks_Handler,
		},
		{
			MethodName: "GetLinksByBoxId",
			Handler:    _BoxService_GetLinksByBoxId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "box.proto",
}

func init() { proto.RegisterFile("box.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x6a, 0xea, 0x40,
	0x10, 0xc6, 0xcd, 0x5f, 0xcd, 0x08, 0xe7, 0x1c, 0x86, 0x73, 0x11, 0xa4, 0x68, 0xba, 0xa5, 0x34,
	0xbd, 0xf1, 0xc2, 0xbe, 0xc1, 0xde, 0x88, 0xb4, 0x58, 0x88, 0x4f, 0x10, 0xcd, 0x20, 0xa1, 0x9a,
	0x95, 0x6c, 0xaa, 0xf1, 0x85, 0xfa, 0x44, 0x7d, 0xa0, 0xb2, 0xbb, 0x6a, 0xda, 0x8a, 0xa5, 0x77,
	0xdf, 0x7e, 0xfb, 0xb1, 0xf3, 0x9b, 0x99, 0x85, 0x60, 0x2e, 0xea, 0xe1, 0xa6, 0x14, 0x95, 0x40,
	0x67, 0x2d, 0x97, 0xcc, 0x07, 0x77, 0x2a, 0x0a, 0x62, 0x8f, 0xe0, 0x70, 0x51, 0xe3, 0x1f, 0xb0,
	0xf3, 0x2c, 0xb4, 0x22, 0x2b, 0x0e, 0x12, 0x3b, 0xcf, 0x10, 0xc1, 0x2d, 0xd2, 0x35, 0x85, 0xb6,
	0x76, 0xb4, 0xc6, 0x08, 0xba, 0x19, 0xc9, 0x45, 0x99, 0x6f, 0xaa, 0x5c, 0x14, 0xa1, 0xa3, 0xaf,
	0x3e, 0x5b, 0xec, 0x0e, 0x3c, 0x2e, 0x6a, 0x92, 0xd8, 0x07, 0x6f, 0xae, 0x44, 0x68, 0x45, 0x4e,
	0xdc, 0x1d, 0x75, 0x86, 0x6b, 0xb9, 0x1c, 0x72, 0x51, 0x27, 0xc6, 0x66, 0x6f, 0x16, 0xb8, 0x4f,
	0x79, 0xf1, 0xf2, 0xab, 0xba, 0xff, 0xc0, 0x79, 0x2d, 0x57, 0x87, 0x7a, 0x4a, 0x7e, 0x27, 0x71,
	0xcf, 0x48, 0xd4, 0x3b, 0x55, 0xba, 0x94, 0xa1, 0x17, 0x39, 0xea, 0x1d, 0xa5, 0xf1, 0xbf, 0x86,
	0x9a, 0x64, 0xa1, 0xaf, 0xf3, 0xe6, 0x80, 0x57, 0x10, 0x2c, 0x4a, 0x4a, 0x2b, 0xca, 0x9e, 0x8b,
	0xb0, 0x1d, 0x59, 0xb1, 0x93, 0x34, 0x06, 0x8b, 0xc1, 0x53, 0x9c, 0x12, 0x07, 0xe0, 0xad, 0x94,
	0x38, 0x74, 0x14, 0xe8, 0x8e, 0xd4, 0x55, 0x62, 0xfc, 0xd1, 0xbb, 0x0d, 0xc0, 0x45, 0x3d, 0xa3,
	0x72, 0x9b, 0x2f, 0x08, 0xfb, 0xe0, 0x4f, 0x69, 0xa7, 0x46, 0x7b, 0x6a, 0xbe, 0x77, 0x52, 0xac,
	0x85, 0x03, 0x68, 0xcf, 0xd2, 0x2d, 0x5d, 0x0e, 0x30, 0x80, 0x31, 0x55, 0x5c, 0xd4, 0x7c, 0x3f,
	0xc9, 0x2e, 0x64, 0x6e, 0xa0, 0x63, 0x32, 0x24, 0xd1, 0x10, 0xa9, 0x9d, 0xf6, 0xe0, 0x18, 0x21,
	0xc9, 0x5a, 0x78, 0x0d, 0xed, 0x29, 0xed, 0xf4, 0xb4, 0x1b, 0xea, 0x5e, 0x23, 0x75, 0xad, 0x8e,
	0x82, 0xf9, 0x31, 0x73, 0x0b, 0xdd, 0x31, 0x55, 0xea, 0xa0, 0x81, 0x2e, 0xc5, 0x0c, 0x92, 0x99,
	0xd9, 0x19, 0x92, 0xb6, 0x59, 0x0b, 0xef, 0xe1, 0xef, 0x31, 0xc4, 0xf7, 0x5c, 0xaf, 0xa1, 0x69,
	0xf0, 0x4b, 0x74, 0xee, 0xeb, 0x3f, 0xfb, 0xf0, 0x11, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xfe, 0xbd,
	0xd3, 0xc0, 0x02, 0x00, 0x00,
}
