// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.tl.handshake_service.proto

package mtproto

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

// /////////////////////////////////////////////////////////////////////////////
// req_pq#60469778 nonce:int128 = ResPQ;
type TLReqPq struct {
	Nonce []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *TLReqPq) Reset()                    { *m = TLReqPq{} }
func (m *TLReqPq) String() string            { return proto.CompactTextString(m) }
func (*TLReqPq) ProtoMessage()               {}
func (*TLReqPq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *TLReqPq) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// req_DH_params#d712e4be nonce:int128 server_nonce:int128 p:string q:string public_key_fingerprint:long encrypted_data:string = Server_DH_Params;
type TLReq_DHParams struct {
	Nonce                []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce          []byte `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	P                    string `protobuf:"bytes,3,opt,name=p" json:"p,omitempty"`
	Q                    string `protobuf:"bytes,4,opt,name=q" json:"q,omitempty"`
	PublicKeyFingerprint int64  `protobuf:"varint,5,opt,name=public_key_fingerprint,json=publicKeyFingerprint" json:"public_key_fingerprint,omitempty"`
	EncryptedData        string `protobuf:"bytes,6,opt,name=encrypted_data,json=encryptedData" json:"encrypted_data,omitempty"`
}

func (m *TLReq_DHParams) Reset()                    { *m = TLReq_DHParams{} }
func (m *TLReq_DHParams) String() string            { return proto.CompactTextString(m) }
func (*TLReq_DHParams) ProtoMessage()               {}
func (*TLReq_DHParams) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *TLReq_DHParams) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *TLReq_DHParams) GetServerNonce() []byte {
	if m != nil {
		return m.ServerNonce
	}
	return nil
}

func (m *TLReq_DHParams) GetP() string {
	if m != nil {
		return m.P
	}
	return ""
}

func (m *TLReq_DHParams) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *TLReq_DHParams) GetPublicKeyFingerprint() int64 {
	if m != nil {
		return m.PublicKeyFingerprint
	}
	return 0
}

func (m *TLReq_DHParams) GetEncryptedData() string {
	if m != nil {
		return m.EncryptedData
	}
	return ""
}

// /////////////////////////////////////////////////////////////////////////////
// set_client_DH_params#f5045f1f nonce:int128 server_nonce:int128 encrypted_data:string = Set_client_DH_params_answer;
type TLSetClient_DHParams struct {
	Nonce         []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ServerNonce   []byte `protobuf:"bytes,2,opt,name=server_nonce,json=serverNonce,proto3" json:"server_nonce,omitempty"`
	EncryptedData string `protobuf:"bytes,3,opt,name=encrypted_data,json=encryptedData" json:"encrypted_data,omitempty"`
}

func (m *TLSetClient_DHParams) Reset()                    { *m = TLSetClient_DHParams{} }
func (m *TLSetClient_DHParams) String() string            { return proto.CompactTextString(m) }
func (*TLSetClient_DHParams) ProtoMessage()               {}
func (*TLSetClient_DHParams) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *TLSetClient_DHParams) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *TLSetClient_DHParams) GetServerNonce() []byte {
	if m != nil {
		return m.ServerNonce
	}
	return nil
}

func (m *TLSetClient_DHParams) GetEncryptedData() string {
	if m != nil {
		return m.EncryptedData
	}
	return ""
}

// /////////////////////////////////////////////////////////////////////////////
// destroy_auth_key#d1435160 = DestroyAuthKeyRes;
type TLDestroyAuthKey struct {
}

func (m *TLDestroyAuthKey) Reset()                    { *m = TLDestroyAuthKey{} }
func (m *TLDestroyAuthKey) String() string            { return proto.CompactTextString(m) }
func (*TLDestroyAuthKey) ProtoMessage()               {}
func (*TLDestroyAuthKey) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func init() {
	proto.RegisterType((*TLReqPq)(nil), "mtproto.TL_req_pq")
	proto.RegisterType((*TLReq_DHParams)(nil), "mtproto.TL_req_DH_params")
	proto.RegisterType((*TLSetClient_DHParams)(nil), "mtproto.TL_set_client_DH_params")
	proto.RegisterType((*TLDestroyAuthKey)(nil), "mtproto.TL_destroy_auth_key")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RPCAuthKey service

type RPCAuthKeyClient interface {
	// req_pq#60469778 nonce:int128 = ResPQ;
	ReqPq(ctx context.Context, in *TLReqPq, opts ...grpc.CallOption) (*ResPQ, error)
	// req_DH_params#d712e4be nonce:int128 server_nonce:int128 p:string q:string public_key_fingerprint:long encrypted_data:string = Server_DH_Params;
	Req_DHParams(ctx context.Context, in *TLReq_DHParams, opts ...grpc.CallOption) (*Server_DH_Params, error)
	// set_client_DH_params#f5045f1f nonce:int128 server_nonce:int128 encrypted_data:string = Set_client_DH_params_answer;
	SetClient_DHParams(ctx context.Context, in *TLSetClient_DHParams, opts ...grpc.CallOption) (*SetClient_DHParamsAnswer, error)
	// destroy_auth_key#d1435160 = DestroyAuthKeyRes;
	DestroyAuthKey(ctx context.Context, in *TLDestroyAuthKey, opts ...grpc.CallOption) (*DestroyAuthKeyRes, error)
}

type rPCAuthKeyClient struct {
	cc *grpc.ClientConn
}

func NewRPCAuthKeyClient(cc *grpc.ClientConn) RPCAuthKeyClient {
	return &rPCAuthKeyClient{cc}
}

func (c *rPCAuthKeyClient) ReqPq(ctx context.Context, in *TLReqPq, opts ...grpc.CallOption) (*ResPQ, error) {
	out := new(ResPQ)
	err := grpc.Invoke(ctx, "/mtproto.RPCAuthKey/req_pq", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthKeyClient) Req_DHParams(ctx context.Context, in *TLReq_DHParams, opts ...grpc.CallOption) (*Server_DH_Params, error) {
	out := new(Server_DH_Params)
	err := grpc.Invoke(ctx, "/mtproto.RPCAuthKey/req_DH_params", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthKeyClient) SetClient_DHParams(ctx context.Context, in *TLSetClient_DHParams, opts ...grpc.CallOption) (*SetClient_DHParamsAnswer, error) {
	out := new(SetClient_DHParamsAnswer)
	err := grpc.Invoke(ctx, "/mtproto.RPCAuthKey/set_client_DH_params", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCAuthKeyClient) DestroyAuthKey(ctx context.Context, in *TLDestroyAuthKey, opts ...grpc.CallOption) (*DestroyAuthKeyRes, error) {
	out := new(DestroyAuthKeyRes)
	err := grpc.Invoke(ctx, "/mtproto.RPCAuthKey/destroy_auth_key", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPCAuthKey service

type RPCAuthKeyServer interface {
	// req_pq#60469778 nonce:int128 = ResPQ;
	ReqPq(context.Context, *TLReqPq) (*ResPQ, error)
	// req_DH_params#d712e4be nonce:int128 server_nonce:int128 p:string q:string public_key_fingerprint:long encrypted_data:string = Server_DH_Params;
	Req_DHParams(context.Context, *TLReq_DHParams) (*Server_DH_Params, error)
	// set_client_DH_params#f5045f1f nonce:int128 server_nonce:int128 encrypted_data:string = Set_client_DH_params_answer;
	SetClient_DHParams(context.Context, *TLSetClient_DHParams) (*SetClient_DHParamsAnswer, error)
	// destroy_auth_key#d1435160 = DestroyAuthKeyRes;
	DestroyAuthKey(context.Context, *TLDestroyAuthKey) (*DestroyAuthKeyRes, error)
}

func RegisterRPCAuthKeyServer(s *grpc.Server, srv RPCAuthKeyServer) {
	s.RegisterService(&_RPCAuthKey_serviceDesc, srv)
}

func _RPCAuthKey_ReqPq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLReqPq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthKeyServer).ReqPq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCAuthKey/ReqPq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthKeyServer).ReqPq(ctx, req.(*TLReqPq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthKey_Req_DHParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLReq_DHParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthKeyServer).Req_DHParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCAuthKey/Req_DHParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthKeyServer).Req_DHParams(ctx, req.(*TLReq_DHParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthKey_SetClient_DHParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSetClient_DHParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthKeyServer).SetClient_DHParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCAuthKey/SetClient_DHParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthKeyServer).SetClient_DHParams(ctx, req.(*TLSetClient_DHParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCAuthKey_DestroyAuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLDestroyAuthKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCAuthKeyServer).DestroyAuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCAuthKey/DestroyAuthKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCAuthKeyServer).DestroyAuthKey(ctx, req.(*TLDestroyAuthKey))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCAuthKey_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mtproto.RPCAuthKey",
	HandlerType: (*RPCAuthKeyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "req_pq",
			Handler:    _RPCAuthKey_ReqPq_Handler,
		},
		{
			MethodName: "req_DH_params",
			Handler:    _RPCAuthKey_Req_DHParams_Handler,
		},
		{
			MethodName: "set_client_DH_params",
			Handler:    _RPCAuthKey_SetClient_DHParams_Handler,
		},
		{
			MethodName: "destroy_auth_key",
			Handler:    _RPCAuthKey_DestroyAuthKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema.tl.handshake_service.proto",
}

func init() { proto.RegisterFile("schema.tl.handshake_service.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x9b, 0x96, 0x75, 0xda, 0xa1, 0x9b, 0x26, 0x53, 0x20, 0x0b, 0x5c, 0xb4, 0x11, 0x48,
	0xbd, 0xf2, 0x05, 0xf0, 0x02, 0x8c, 0x0a, 0x2a, 0xad, 0x4c, 0x25, 0xf4, 0xde, 0xb8, 0xc9, 0x61,
	0x89, 0x96, 0x38, 0x8e, 0xed, 0x80, 0xf2, 0x64, 0xbc, 0x00, 0x0f, 0x86, 0x12, 0x87, 0xac, 0xb4,
	0xb9, 0xe4, 0xce, 0xe7, 0x7c, 0xbf, 0x8f, 0xfd, 0x9f, 0x1f, 0xe6, 0x3a, 0x8c, 0x31, 0xe3, 0xd4,
	0xa4, 0x34, 0xe6, 0x22, 0xd2, 0x31, 0xbf, 0x47, 0xa6, 0x51, 0xfd, 0x48, 0x42, 0xa4, 0x52, 0xe5,
	0x26, 0x27, 0xa7, 0x99, 0x69, 0x0e, 0xde, 0x55, 0x8f, 0xd6, 0x6a, 0xfc, 0x39, 0x9c, 0x6d, 0xd7,
	0x4c, 0x61, 0xc1, 0x64, 0x41, 0xa6, 0x70, 0x22, 0x72, 0x11, 0xa2, 0xeb, 0xcc, 0x9c, 0xc5, 0x24,
	0xb0, 0x85, 0xff, 0xdb, 0x81, 0xcb, 0x56, 0xb3, 0x5c, 0x31, 0xc9, 0x15, 0xcf, 0x74, 0xbf, 0x94,
	0xcc, 0x61, 0x52, 0x7f, 0x01, 0x15, 0xb3, 0x70, 0xd8, 0xc0, 0xc7, 0xb6, 0x77, 0xdb, 0x48, 0x26,
	0xe0, 0x48, 0x77, 0x34, 0x73, 0x16, 0x67, 0x81, 0x23, 0xeb, 0xaa, 0x70, 0x1f, 0xd9, 0xaa, 0x20,
	0xef, 0xe0, 0x99, 0x2c, 0x77, 0x69, 0x12, 0xb2, 0x7b, 0xac, 0xd8, 0xf7, 0x44, 0xdc, 0xa1, 0x92,
	0x2a, 0x11, 0xc6, 0x3d, 0x99, 0x39, 0x8b, 0x51, 0x30, 0xb5, 0xf4, 0x06, 0xab, 0x8f, 0x0f, 0x8c,
	0xbc, 0x86, 0x0b, 0x14, 0xa1, 0xaa, 0xa4, 0xc1, 0x88, 0x45, 0xdc, 0x70, 0x77, 0xdc, 0x0c, 0x3c,
	0xef, 0xba, 0x4b, 0x6e, 0xb8, 0x5f, 0xc1, 0xf3, 0xed, 0x9a, 0x69, 0x34, 0x2c, 0x4c, 0x13, 0x14,
	0xe6, 0x7f, 0x98, 0x39, 0x7e, 0x7a, 0xd4, 0xf7, 0xf4, 0x53, 0x78, 0xb2, 0x5d, 0xb3, 0x08, 0xb5,
	0x51, 0x79, 0xc5, 0x78, 0x69, 0xe2, 0xda, 0xe0, 0x9b, 0x5f, 0x43, 0x80, 0x60, 0xf3, 0xe1, 0x7d,
	0x69, 0xe2, 0x1b, 0xac, 0x08, 0x85, 0x71, 0x9b, 0x03, 0xa1, 0x6d, 0x72, 0xb4, 0xcb, 0xc6, 0xbb,
	0xe8, 0x7a, 0x01, 0xea, 0xcd, 0x17, 0x7f, 0x40, 0x3e, 0xc1, 0xf9, 0xbf, 0x99, 0x5c, 0x1d, 0x5e,
	0xeb, 0x90, 0xf7, 0x80, 0xbe, 0x5a, 0x4b, 0xcb, 0x15, 0xdb, 0x34, 0xc8, 0x1f, 0x90, 0x6f, 0x30,
	0xed, 0x5d, 0xcb, 0x6c, 0x7f, 0x5e, 0x9f, 0xc2, 0x7b, 0xb5, 0x37, 0xf6, 0x18, 0x33, 0x2e, 0xf4,
	0x4f, 0x54, 0xfe, 0x80, 0xdc, 0xc2, 0xe5, 0xa1, 0x7b, 0xf2, 0x72, 0x7f, 0xfa, 0x21, 0xf5, 0xbc,
	0x8e, 0x2e, 0x2d, 0x6a, 0xb7, 0x14, 0xa0, 0xf6, 0x07, 0xd7, 0x0b, 0x78, 0x11, 0xe6, 0x19, 0x15,
	0xb8, 0x2b, 0x53, 0x9e, 0x64, 0x14, 0xc5, 0x5d, 0x22, 0xf0, 0xef, 0x95, 0xeb, 0xd3, 0xcf, 0xdb,
	0x4d, 0x7d, 0x58, 0x0d, 0x77, 0xe3, 0xa6, 0xf3, 0xf6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a,
	0xe3, 0x52, 0xef, 0x2f, 0x03, 0x00, 0x00,
}
