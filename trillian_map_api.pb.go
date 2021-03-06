// Code generated by protoc-gen-go.
// source: trillian_map_api.proto
// DO NOT EDIT!

package trillian

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

// MapLeaf represents the data behind Map leaves.
type MapLeaf struct {
	// index is the location of this leaf.
	// All indexes for a given Map must contain a constant number of bits.
	Index []byte `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	// leaf_hash is the tree hash of leaf_value.
	LeafHash []byte `protobuf:"bytes,2,opt,name=leaf_hash,json=leafHash,proto3" json:"leaf_hash,omitempty"`
	// leaf_value is the data the tree commits to.
	LeafValue []byte `protobuf:"bytes,3,opt,name=leaf_value,json=leafValue,proto3" json:"leaf_value,omitempty"`
	// extra_data holds related contextual data, but is not covered by any hash.
	ExtraData []byte `protobuf:"bytes,4,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
}

func (m *MapLeaf) Reset()                    { *m = MapLeaf{} }
func (m *MapLeaf) String() string            { return proto.CompactTextString(m) }
func (*MapLeaf) ProtoMessage()               {}
func (*MapLeaf) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MapLeaf) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *MapLeaf) GetLeafHash() []byte {
	if m != nil {
		return m.LeafHash
	}
	return nil
}

func (m *MapLeaf) GetLeafValue() []byte {
	if m != nil {
		return m.LeafValue
	}
	return nil
}

func (m *MapLeaf) GetExtraData() []byte {
	if m != nil {
		return m.ExtraData
	}
	return nil
}

type MapLeafInclusion struct {
	Leaf      *MapLeaf `protobuf:"bytes,1,opt,name=leaf" json:"leaf,omitempty"`
	Inclusion [][]byte `protobuf:"bytes,2,rep,name=inclusion,proto3" json:"inclusion,omitempty"`
}

func (m *MapLeafInclusion) Reset()                    { *m = MapLeafInclusion{} }
func (m *MapLeafInclusion) String() string            { return proto.CompactTextString(m) }
func (*MapLeafInclusion) ProtoMessage()               {}
func (*MapLeafInclusion) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *MapLeafInclusion) GetLeaf() *MapLeaf {
	if m != nil {
		return m.Leaf
	}
	return nil
}

func (m *MapLeafInclusion) GetInclusion() [][]byte {
	if m != nil {
		return m.Inclusion
	}
	return nil
}

type GetMapLeavesRequest struct {
	MapId    int64    `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
	Index    [][]byte `protobuf:"bytes,2,rep,name=index,proto3" json:"index,omitempty"`
	Revision int64    `protobuf:"varint,3,opt,name=revision" json:"revision,omitempty"`
}

func (m *GetMapLeavesRequest) Reset()                    { *m = GetMapLeavesRequest{} }
func (m *GetMapLeavesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMapLeavesRequest) ProtoMessage()               {}
func (*GetMapLeavesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetMapLeavesRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *GetMapLeavesRequest) GetIndex() [][]byte {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *GetMapLeavesRequest) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

type GetMapLeavesResponse struct {
	MapLeafInclusion []*MapLeafInclusion `protobuf:"bytes,2,rep,name=map_leaf_inclusion,json=mapLeafInclusion" json:"map_leaf_inclusion,omitempty"`
	MapRoot          *SignedMapRoot      `protobuf:"bytes,3,opt,name=map_root,json=mapRoot" json:"map_root,omitempty"`
}

func (m *GetMapLeavesResponse) Reset()                    { *m = GetMapLeavesResponse{} }
func (m *GetMapLeavesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMapLeavesResponse) ProtoMessage()               {}
func (*GetMapLeavesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetMapLeavesResponse) GetMapLeafInclusion() []*MapLeafInclusion {
	if m != nil {
		return m.MapLeafInclusion
	}
	return nil
}

func (m *GetMapLeavesResponse) GetMapRoot() *SignedMapRoot {
	if m != nil {
		return m.MapRoot
	}
	return nil
}

type SetMapLeavesRequest struct {
	MapId      int64           `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
	Leaves     []*MapLeaf      `protobuf:"bytes,2,rep,name=leaves" json:"leaves,omitempty"`
	MapperData *MapperMetadata `protobuf:"bytes,3,opt,name=mapper_data,json=mapperData" json:"mapper_data,omitempty"`
}

func (m *SetMapLeavesRequest) Reset()                    { *m = SetMapLeavesRequest{} }
func (m *SetMapLeavesRequest) String() string            { return proto.CompactTextString(m) }
func (*SetMapLeavesRequest) ProtoMessage()               {}
func (*SetMapLeavesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *SetMapLeavesRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *SetMapLeavesRequest) GetLeaves() []*MapLeaf {
	if m != nil {
		return m.Leaves
	}
	return nil
}

func (m *SetMapLeavesRequest) GetMapperData() *MapperMetadata {
	if m != nil {
		return m.MapperData
	}
	return nil
}

type SetMapLeavesResponse struct {
	MapRoot *SignedMapRoot `protobuf:"bytes,2,opt,name=map_root,json=mapRoot" json:"map_root,omitempty"`
}

func (m *SetMapLeavesResponse) Reset()                    { *m = SetMapLeavesResponse{} }
func (m *SetMapLeavesResponse) String() string            { return proto.CompactTextString(m) }
func (*SetMapLeavesResponse) ProtoMessage()               {}
func (*SetMapLeavesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *SetMapLeavesResponse) GetMapRoot() *SignedMapRoot {
	if m != nil {
		return m.MapRoot
	}
	return nil
}

type GetSignedMapRootRequest struct {
	MapId int64 `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
}

func (m *GetSignedMapRootRequest) Reset()                    { *m = GetSignedMapRootRequest{} }
func (m *GetSignedMapRootRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSignedMapRootRequest) ProtoMessage()               {}
func (*GetSignedMapRootRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *GetSignedMapRootRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

type GetSignedMapRootByRevisionRequest struct {
	MapId    int64 `protobuf:"varint,1,opt,name=map_id,json=mapId" json:"map_id,omitempty"`
	Revision int64 `protobuf:"varint,2,opt,name=revision" json:"revision,omitempty"`
}

func (m *GetSignedMapRootByRevisionRequest) Reset()         { *m = GetSignedMapRootByRevisionRequest{} }
func (m *GetSignedMapRootByRevisionRequest) String() string { return proto.CompactTextString(m) }
func (*GetSignedMapRootByRevisionRequest) ProtoMessage()    {}
func (*GetSignedMapRootByRevisionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{7}
}

func (m *GetSignedMapRootByRevisionRequest) GetMapId() int64 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *GetSignedMapRootByRevisionRequest) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

type GetSignedMapRootResponse struct {
	MapRoot *SignedMapRoot `protobuf:"bytes,2,opt,name=map_root,json=mapRoot" json:"map_root,omitempty"`
}

func (m *GetSignedMapRootResponse) Reset()                    { *m = GetSignedMapRootResponse{} }
func (m *GetSignedMapRootResponse) String() string            { return proto.CompactTextString(m) }
func (*GetSignedMapRootResponse) ProtoMessage()               {}
func (*GetSignedMapRootResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *GetSignedMapRootResponse) GetMapRoot() *SignedMapRoot {
	if m != nil {
		return m.MapRoot
	}
	return nil
}

func init() {
	proto.RegisterType((*MapLeaf)(nil), "trillian.MapLeaf")
	proto.RegisterType((*MapLeafInclusion)(nil), "trillian.MapLeafInclusion")
	proto.RegisterType((*GetMapLeavesRequest)(nil), "trillian.GetMapLeavesRequest")
	proto.RegisterType((*GetMapLeavesResponse)(nil), "trillian.GetMapLeavesResponse")
	proto.RegisterType((*SetMapLeavesRequest)(nil), "trillian.SetMapLeavesRequest")
	proto.RegisterType((*SetMapLeavesResponse)(nil), "trillian.SetMapLeavesResponse")
	proto.RegisterType((*GetSignedMapRootRequest)(nil), "trillian.GetSignedMapRootRequest")
	proto.RegisterType((*GetSignedMapRootByRevisionRequest)(nil), "trillian.GetSignedMapRootByRevisionRequest")
	proto.RegisterType((*GetSignedMapRootResponse)(nil), "trillian.GetSignedMapRootResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TrillianMap service

type TrillianMapClient interface {
	GetLeaves(ctx context.Context, in *GetMapLeavesRequest, opts ...grpc.CallOption) (*GetMapLeavesResponse, error)
	SetLeaves(ctx context.Context, in *SetMapLeavesRequest, opts ...grpc.CallOption) (*SetMapLeavesResponse, error)
	GetSignedMapRoot(ctx context.Context, in *GetSignedMapRootRequest, opts ...grpc.CallOption) (*GetSignedMapRootResponse, error)
	GetSignedMapRootByRevision(ctx context.Context, in *GetSignedMapRootByRevisionRequest, opts ...grpc.CallOption) (*GetSignedMapRootResponse, error)
}

type trillianMapClient struct {
	cc *grpc.ClientConn
}

func NewTrillianMapClient(cc *grpc.ClientConn) TrillianMapClient {
	return &trillianMapClient{cc}
}

func (c *trillianMapClient) GetLeaves(ctx context.Context, in *GetMapLeavesRequest, opts ...grpc.CallOption) (*GetMapLeavesResponse, error) {
	out := new(GetMapLeavesResponse)
	err := grpc.Invoke(ctx, "/trillian.TrillianMap/GetLeaves", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapClient) SetLeaves(ctx context.Context, in *SetMapLeavesRequest, opts ...grpc.CallOption) (*SetMapLeavesResponse, error) {
	out := new(SetMapLeavesResponse)
	err := grpc.Invoke(ctx, "/trillian.TrillianMap/SetLeaves", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapClient) GetSignedMapRoot(ctx context.Context, in *GetSignedMapRootRequest, opts ...grpc.CallOption) (*GetSignedMapRootResponse, error) {
	out := new(GetSignedMapRootResponse)
	err := grpc.Invoke(ctx, "/trillian.TrillianMap/GetSignedMapRoot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapClient) GetSignedMapRootByRevision(ctx context.Context, in *GetSignedMapRootByRevisionRequest, opts ...grpc.CallOption) (*GetSignedMapRootResponse, error) {
	out := new(GetSignedMapRootResponse)
	err := grpc.Invoke(ctx, "/trillian.TrillianMap/GetSignedMapRootByRevision", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TrillianMap service

type TrillianMapServer interface {
	GetLeaves(context.Context, *GetMapLeavesRequest) (*GetMapLeavesResponse, error)
	SetLeaves(context.Context, *SetMapLeavesRequest) (*SetMapLeavesResponse, error)
	GetSignedMapRoot(context.Context, *GetSignedMapRootRequest) (*GetSignedMapRootResponse, error)
	GetSignedMapRootByRevision(context.Context, *GetSignedMapRootByRevisionRequest) (*GetSignedMapRootResponse, error)
}

func RegisterTrillianMapServer(s *grpc.Server, srv TrillianMapServer) {
	s.RegisterService(&_TrillianMap_serviceDesc, srv)
}

func _TrillianMap_GetLeaves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMapLeavesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianMapServer).GetLeaves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianMap/GetLeaves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianMapServer).GetLeaves(ctx, req.(*GetMapLeavesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianMap_SetLeaves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMapLeavesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianMapServer).SetLeaves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianMap/SetLeaves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianMapServer).SetLeaves(ctx, req.(*SetMapLeavesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianMap_GetSignedMapRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignedMapRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianMapServer).GetSignedMapRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianMap/GetSignedMapRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianMapServer).GetSignedMapRoot(ctx, req.(*GetSignedMapRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrillianMap_GetSignedMapRootByRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignedMapRootByRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrillianMapServer).GetSignedMapRootByRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trillian.TrillianMap/GetSignedMapRootByRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrillianMapServer).GetSignedMapRootByRevision(ctx, req.(*GetSignedMapRootByRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrillianMap_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trillian.TrillianMap",
	HandlerType: (*TrillianMapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLeaves",
			Handler:    _TrillianMap_GetLeaves_Handler,
		},
		{
			MethodName: "SetLeaves",
			Handler:    _TrillianMap_SetLeaves_Handler,
		},
		{
			MethodName: "GetSignedMapRoot",
			Handler:    _TrillianMap_GetSignedMapRoot_Handler,
		},
		{
			MethodName: "GetSignedMapRootByRevision",
			Handler:    _TrillianMap_GetSignedMapRootByRevision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trillian_map_api.proto",
}

func init() { proto.RegisterFile("trillian_map_api.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xb5, 0xe9, 0x7e, 0xb4, 0x37, 0x22, 0x75, 0xb6, 0xba, 0x31, 0xba, 0xb2, 0x1b, 0x10, 0x14,
	0xa1, 0x4a, 0x7c, 0xf2, 0xd1, 0x22, 0x74, 0x57, 0xb6, 0xb2, 0x24, 0xb2, 0x3e, 0x08, 0x96, 0xeb,
	0x66, 0xb6, 0x1d, 0x48, 0x32, 0x63, 0x32, 0x2d, 0xab, 0x3f, 0x43, 0xfc, 0xaf, 0xbe, 0xca, 0xcc,
	0xa4, 0xe9, 0xa6, 0x4d, 0x6b, 0xf1, 0xad, 0xb9, 0xe7, 0xdc, 0x73, 0xcf, 0xb9, 0x77, 0x28, 0x3c,
	0x94, 0x19, 0x8b, 0x63, 0x86, 0xe9, 0x28, 0x41, 0x31, 0x42, 0xc1, 0x7a, 0x22, 0xe3, 0x92, 0x93,
	0xd6, 0xbc, 0xee, 0xde, 0x9b, 0xff, 0x32, 0x88, 0xf7, 0x13, 0xf6, 0x87, 0x28, 0xce, 0x29, 0x5e,
	0x93, 0x2e, 0xec, 0xb2, 0x34, 0xa2, 0x37, 0x4e, 0xe3, 0xb8, 0xf1, 0xfc, 0x6e, 0x60, 0x3e, 0xc8,
	0x63, 0x68, 0xc7, 0x14, 0xaf, 0x47, 0x13, 0xcc, 0x27, 0x8e, 0xa5, 0x91, 0x96, 0x2a, 0x9c, 0x62,
	0x3e, 0x21, 0x47, 0x00, 0x1a, 0x9c, 0x61, 0x3c, 0xa5, 0x4e, 0x53, 0xa3, 0x9a, 0x7e, 0xa9, 0x0a,
	0x0a, 0xa6, 0x37, 0x32, 0xc3, 0x51, 0x84, 0x12, 0x9d, 0x1d, 0x03, 0xeb, 0xca, 0x7b, 0x94, 0xe8,
	0x7d, 0x86, 0x4e, 0x31, 0xfb, 0x2c, 0xbd, 0x8a, 0xa7, 0x39, 0xe3, 0x29, 0x79, 0x06, 0x3b, 0xaa,
	0x5f, 0x7b, 0xb0, 0xfd, 0xfb, 0xbd, 0xd2, 0x6e, 0xc1, 0x0c, 0x34, 0x4c, 0x9e, 0x40, 0x9b, 0xcd,
	0x7b, 0x1c, 0xeb, 0xb8, 0xa9, 0x84, 0xcb, 0x82, 0xf7, 0x15, 0x0e, 0x06, 0x54, 0x9a, 0x8e, 0x19,
	0xcd, 0x03, 0xfa, 0x7d, 0x4a, 0x73, 0x49, 0x1e, 0xc0, 0x9e, 0x5a, 0x0b, 0x8b, 0xb4, 0x7a, 0x33,
	0xd8, 0x4d, 0x50, 0x9c, 0x45, 0x8b, 0xdc, 0x46, 0xa7, 0xc8, 0xed, 0x42, 0x2b, 0xa3, 0x33, 0xa6,
	0x07, 0x34, 0x35, 0xbd, 0xfc, 0xf6, 0x7e, 0x37, 0xa0, 0x5b, 0x1d, 0x90, 0x0b, 0x9e, 0xe6, 0x94,
	0x9c, 0x02, 0x51, 0x13, 0xf4, 0x4e, 0xaa, 0xfe, 0x6c, 0xdf, 0x5d, 0xc9, 0x52, 0xa6, 0x0e, 0x3a,
	0xc9, 0xf2, 0x1e, 0x7c, 0x68, 0x29, 0xa5, 0x8c, 0x73, 0xa9, 0xc7, 0xdb, 0xfe, 0xe1, 0xa2, 0x3f,
	0x64, 0xe3, 0x94, 0x46, 0x43, 0x14, 0x01, 0xe7, 0x32, 0xd8, 0x4f, 0xcc, 0x0f, 0xef, 0x57, 0x03,
	0x0e, 0xc2, 0xed, 0x73, 0xbf, 0x80, 0xbd, 0x58, 0xf3, 0x0a, 0x83, 0x35, 0xcb, 0x2e, 0x08, 0xe4,
	0x2d, 0xd8, 0x09, 0x0a, 0x41, 0x33, 0x73, 0x49, 0x63, 0xc8, 0xa9, 0xf0, 0x05, 0xcd, 0x86, 0x54,
	0xa2, 0xc2, 0x03, 0x30, 0x64, 0x7d, 0xe4, 0x0f, 0xd0, 0x0d, 0xeb, 0x56, 0x75, 0x3b, 0xa0, 0xb5,
	0x65, 0xc0, 0xd7, 0x70, 0x38, 0xa0, 0xb2, 0x0a, 0x6e, 0xcc, 0xe8, 0x5d, 0xc2, 0xc9, 0x72, 0x47,
	0xff, 0x47, 0x50, 0xdc, 0xf1, 0x1f, 0xfb, 0xb9, 0xfd, 0x02, 0xac, 0xa5, 0x17, 0xf0, 0x11, 0x9c,
	0x55, 0x27, 0xff, 0x9f, 0xcc, 0xff, 0x63, 0x81, 0xfd, 0xa9, 0xe0, 0x0c, 0x51, 0x90, 0x73, 0x68,
	0x0f, 0xa8, 0x34, 0x2b, 0x23, 0x47, 0x8b, 0xf6, 0x9a, 0x67, 0xed, 0x3e, 0x5d, 0x07, 0x1b, 0x3f,
	0xde, 0x1d, 0xa5, 0x16, 0xd6, 0xa9, 0x85, 0x9b, 0xd5, 0xc2, 0x7a, 0xb5, 0x2f, 0xd0, 0x59, 0xce,
	0x4e, 0x4e, 0x2a, 0x1e, 0xea, 0x2e, 0xe4, 0x7a, 0x9b, 0x28, 0xa5, 0x38, 0x07, 0x77, 0xfd, 0xc1,
	0xc8, 0xcb, 0xf5, 0x1a, 0x2b, 0x67, 0xdd, 0x6e, 0x60, 0xff, 0x15, 0x3c, 0xba, 0xe2, 0x49, 0x6f,
	0xcc, 0xf9, 0x38, 0xa6, 0xbd, 0xea, 0xbf, 0x63, 0xbf, 0x33, 0xbf, 0xc9, 0x3b, 0xc1, 0x2e, 0x54,
	0xe5, 0xa2, 0xf1, 0x6d, 0x4f, 0x43, 0x6f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xb0, 0x74,
	0x5f, 0x6c, 0x05, 0x00, 0x00,
}
