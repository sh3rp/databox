// Code generated by protoc-gen-go. DO NOT EDIT.
// source: box.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	box.proto

It has these top-level messages:
	None
	Search
	Box
	Boxes
	Link
	Links
	Note
	Notes
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

type Search struct {
	Term  string `protobuf:"bytes,1,opt,name=term" json:"term,omitempty"`
	Count int32  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Page  int32  `protobuf:"varint,3,opt,name=page" json:"page,omitempty"`
}

func (m *Search) Reset()                    { *m = Search{} }
func (m *Search) String() string            { return proto.CompactTextString(m) }
func (*Search) ProtoMessage()               {}
func (*Search) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Search) GetTerm() string {
	if m != nil {
		return m.Term
	}
	return ""
}

func (m *Search) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Search) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type Box struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *Box) Reset()                    { *m = Box{} }
func (m *Box) String() string            { return proto.CompactTextString(m) }
func (*Box) ProtoMessage()               {}
func (*Box) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
func (*Boxes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
func (*Link) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
func (*Links) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Links) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

type Note struct {
	Id           string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title        string   `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Text         []byte   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Tags         []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	CreatedOn    int64    `protobuf:"varint,5,opt,name=createdOn" json:"createdOn,omitempty"`
	LastModified int64    `protobuf:"varint,6,opt,name=lastModified" json:"lastModified,omitempty"`
}

func (m *Note) Reset()                    { *m = Note{} }
func (m *Note) String() string            { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()               {}
func (*Note) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Note) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Note) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Note) GetText() []byte {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *Note) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Note) GetCreatedOn() int64 {
	if m != nil {
		return m.CreatedOn
	}
	return 0
}

func (m *Note) GetLastModified() int64 {
	if m != nil {
		return m.LastModified
	}
	return 0
}

type Notes struct {
	Notes []*Note `protobuf:"bytes,1,rep,name=notes" json:"notes,omitempty"`
}

func (m *Notes) Reset()                    { *m = Notes{} }
func (m *Notes) String() string            { return proto.CompactTextString(m) }
func (*Notes) ProtoMessage()               {}
func (*Notes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Notes) GetNotes() []*Note {
	if m != nil {
		return m.Notes
	}
	return nil
}

func init() {
	proto.RegisterType((*None)(nil), "msg.None")
	proto.RegisterType((*Search)(nil), "msg.Search")
	proto.RegisterType((*Box)(nil), "msg.Box")
	proto.RegisterType((*Boxes)(nil), "msg.Boxes")
	proto.RegisterType((*Link)(nil), "msg.Link")
	proto.RegisterType((*Links)(nil), "msg.Links")
	proto.RegisterType((*Note)(nil), "msg.Note")
	proto.RegisterType((*Notes)(nil), "msg.Notes")
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
	GetLinksByBoxId(ctx context.Context, in *Box, opts ...grpc.CallOption) (*Links, error)
	SearchLinks(ctx context.Context, in *Search, opts ...grpc.CallOption) (*Links, error)
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

func (c *boxServiceClient) GetLinksByBoxId(ctx context.Context, in *Box, opts ...grpc.CallOption) (*Links, error) {
	out := new(Links)
	err := grpc.Invoke(ctx, "/msg.BoxService/GetLinksByBoxId", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) SearchLinks(ctx context.Context, in *Search, opts ...grpc.CallOption) (*Links, error) {
	out := new(Links)
	err := grpc.Invoke(ctx, "/msg.BoxService/SearchLinks", in, out, c.cc, opts...)
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
	GetLinksByBoxId(context.Context, *Box) (*Links, error)
	SearchLinks(context.Context, *Search) (*Links, error)
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

func _BoxService_SearchLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Search)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).SearchLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.BoxService/SearchLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).SearchLinks(ctx, req.(*Search))
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
			MethodName: "GetLinksByBoxId",
			Handler:    _BoxService_GetLinksByBoxId_Handler,
		},
		{
			MethodName: "SearchLinks",
			Handler:    _BoxService_SearchLinks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "box.proto",
}

func init() { proto.RegisterFile("box.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x5e, 0x9a, 0x38, 0x6b, 0x4e, 0x26, 0x40, 0xd6, 0x2e, 0xa2, 0x0a, 0x6d, 0xc1, 0x08, 0x11,
	0x6e, 0x7a, 0x31, 0xde, 0x20, 0x17, 0x4c, 0x13, 0x50, 0x24, 0xf7, 0x09, 0xd2, 0xe4, 0x10, 0x2c,
	0xda, 0xb8, 0x8a, 0xbd, 0x2d, 0x7b, 0x11, 0x1e, 0x81, 0x77, 0xe2, 0x6d, 0x90, 0x8f, 0xfb, 0xb3,
	0x6e, 0x2a, 0xe2, 0xee, 0x3b, 0x9f, 0xbf, 0xc4, 0xdf, 0x77, 0xce, 0x31, 0x24, 0x0b, 0x3d, 0x4c,
	0xd7, 0xbd, 0xb6, 0x9a, 0x87, 0x2b, 0xd3, 0x8a, 0x18, 0xa2, 0x99, 0xee, 0x50, 0x7c, 0x82, 0x78,
	0x8e, 0x55, 0x5f, 0xff, 0xe0, 0x1c, 0x22, 0x8b, 0xfd, 0x2a, 0x0b, 0xf2, 0xa0, 0x48, 0x24, 0x61,
	0x7e, 0x0e, 0xac, 0xd6, 0xb7, 0x9d, 0xcd, 0x46, 0x79, 0x50, 0x30, 0xe9, 0x0b, 0xa7, 0x5c, 0x57,
	0x2d, 0x66, 0x21, 0x91, 0x84, 0xc5, 0x67, 0x08, 0x4b, 0x3d, 0xf0, 0x17, 0x30, 0x52, 0xcd, 0xe6,
	0x17, 0x23, 0xd5, 0x38, 0x69, 0x57, 0xad, 0x90, 0xbe, 0x4f, 0x24, 0x61, 0x9e, 0x43, 0xda, 0xa0,
	0xa9, 0x7b, 0xb5, 0xb6, 0x4a, 0x77, 0xf4, 0x97, 0x44, 0x3e, 0xa6, 0xc4, 0x7b, 0x60, 0xa5, 0x1e,
	0xd0, 0xf0, 0x0b, 0x60, 0x0b, 0x07, 0xb2, 0x20, 0x0f, 0x8b, 0xf4, 0x6a, 0x3c, 0x5d, 0x99, 0x76,
	0x5a, 0xea, 0x41, 0x7a, 0x5a, 0xfc, 0x0e, 0x20, 0xfa, 0xa2, 0xba, 0x9f, 0xff, 0x75, 0xef, 0x2b,
	0x08, 0x6f, 0xfb, 0xe5, 0xe6, 0x3e, 0x07, 0x9f, 0x3a, 0x89, 0x9e, 0x39, 0xa1, 0xa6, 0x54, 0xad,
	0xc9, 0x58, 0x1e, 0x52, 0x53, 0xaa, 0xd6, 0xb8, 0xa6, 0x2c, 0xf4, 0x70, 0xd3, 0x64, 0x31, 0xe9,
	0x7d, 0xc1, 0x5f, 0x43, 0x52, 0xf7, 0x58, 0x59, 0x6c, 0xbe, 0x75, 0xd9, 0x69, 0x1e, 0x14, 0xa1,
	0xdc, 0x13, 0xa2, 0x00, 0xe6, 0x7c, 0x1a, 0x7e, 0x09, 0x6c, 0xe9, 0xc0, 0x26, 0x51, 0x42, 0x89,
	0xdc, 0x91, 0xf4, 0xbc, 0xf8, 0x15, 0xb8, 0xc9, 0x58, 0x7c, 0x16, 0xe9, 0x1c, 0x98, 0x55, 0x76,
	0xb9, 0xcd, 0xe4, 0x0b, 0x3f, 0xb5, 0xc1, 0x52, 0xaa, 0x33, 0x49, 0x78, 0x67, 0x3a, 0x7a, 0x64,
	0xfa, 0xc0, 0x1e, 0x7b, 0x62, 0x8f, 0x0b, 0x38, 0x5b, 0x56, 0xc6, 0x7e, 0xd5, 0x8d, 0xfa, 0xae,
	0xd0, 0x27, 0x0b, 0xe5, 0x01, 0xe7, 0x22, 0x38, 0x5f, 0x14, 0xa1, 0x73, 0xe0, 0x20, 0x82, 0x3b,
	0x92, 0x9e, 0xbf, 0xfa, 0x33, 0x02, 0x28, 0xf5, 0x30, 0xc7, 0xfe, 0x4e, 0xd5, 0xc8, 0x2f, 0x20,
	0x9e, 0xe1, 0xbd, 0xdb, 0x8e, 0xdd, 0xfc, 0x26, 0x3b, 0x24, 0x4e, 0xf8, 0x25, 0x9c, 0xce, 0xab,
	0x3b, 0x3c, 0x2e, 0x10, 0x00, 0xd7, 0x68, 0x4b, 0x3d, 0x94, 0x0f, 0x37, 0xcd, 0x11, 0xcd, 0x5b,
	0x18, 0x7b, 0x0d, 0x1a, 0xbe, 0x75, 0xd4, 0xe1, 0x04, 0xb6, 0x12, 0x34, 0xe2, 0x84, 0xbf, 0x81,
	0xd3, 0x19, 0xde, 0xd3, 0xc2, 0xec, 0x1b, 0x3f, 0xd9, 0x43, 0xba, 0x6b, 0xec, 0xcc, 0xfc, 0x53,
	0xf3, 0x0e, 0xd2, 0x6b, 0xb4, 0xae, 0x20, 0x43, 0xc7, 0x64, 0x1f, 0xe0, 0xe5, 0x46, 0x66, 0xca,
	0x87, 0x92, 0x96, 0x64, 0xef, 0x1d, 0x76, 0x4a, 0x67, 0xac, 0x80, 0xd4, 0xbf, 0x42, 0xbf, 0x24,
	0x29, 0x1d, 0x7a, 0xe6, 0x50, 0xb9, 0x88, 0xe9, 0x0d, 0x7f, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x1e, 0xd5, 0x5c, 0x87, 0xd0, 0x03, 0x00, 0x00,
}
