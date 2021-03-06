// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

package user_proto

import (
	context "context"
	fmt "fmt"
	tars "github.com/TarsCloud/TarsGo/tars"
	model "github.com/TarsCloud/TarsGo/tars/model"
	requestf "github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	tools "github.com/TarsCloud/TarsGo/tars/util/tools"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request message containing the user's name.
type Request struct {
	Version              uint32   `protobuf:"fixed32,1,opt,name=version,proto3" json:"version,omitempty"`
	Servant              uint32   `protobuf:"fixed32,2,opt,name=servant,proto3" json:"servant,omitempty"`
	Seq                  uint32   `protobuf:"fixed32,3,opt,name=seq,proto3" json:"seq,omitempty"`
	Uid                  uint64   `protobuf:"fixed64,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Body                 []byte   `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Request) GetServant() uint32 {
	if m != nil {
		return m.Servant
	}
	return 0
}

func (m *Request) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Request) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// The response message containing the greetings
type Respond struct {
	Body                 []byte   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Extend               []byte   `protobuf:"bytes,2,opt,name=extend,proto3" json:"extend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Respond) Reset()         { *m = Respond{} }
func (m *Respond) String() string { return proto.CompactTextString(m) }
func (*Respond) ProtoMessage()    {}
func (*Respond) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{1}
}

func (m *Respond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Respond.Unmarshal(m, b)
}
func (m *Respond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Respond.Marshal(b, m, deterministic)
}
func (m *Respond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Respond.Merge(m, src)
}
func (m *Respond) XXX_Size() int {
	return xxx_messageInfo_Respond.Size(m)
}
func (m *Respond) XXX_DiscardUnknown() {
	xxx_messageInfo_Respond.DiscardUnknown(m)
}

var xxx_messageInfo_Respond proto.InternalMessageInfo

func (m *Respond) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Respond) GetExtend() []byte {
	if m != nil {
		return m.Extend
	}
	return nil
}

// The request message containing the user's name.
type MsgHead struct {
	BodyLen              uint32   `protobuf:"fixed32,1,opt,name=body_len,json=bodyLen,proto3" json:"body_len,omitempty"`
	Version              uint32   `protobuf:"fixed32,2,opt,name=version,proto3" json:"version,omitempty"`
	App                  uint32   `protobuf:"fixed32,3,opt,name=app,proto3" json:"app,omitempty"`
	Server               uint32   `protobuf:"fixed32,4,opt,name=server,proto3" json:"server,omitempty"`
	Servant              uint32   `protobuf:"fixed32,5,opt,name=servant,proto3" json:"servant,omitempty"`
	Seq                  uint32   `protobuf:"fixed32,6,opt,name=seq,proto3" json:"seq,omitempty"`
	RouteId              uint64   `protobuf:"fixed64,7,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
	Encrypt              uint32   `protobuf:"fixed32,8,opt,name=encrypt,proto3" json:"encrypt,omitempty"`
	CacheIs              uint32   `protobuf:"fixed32,9,opt,name=cache_is,json=cacheIs,proto3" json:"cache_is,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgHead) Reset()         { *m = MsgHead{} }
func (m *MsgHead) String() string { return proto.CompactTextString(m) }
func (*MsgHead) ProtoMessage()    {}
func (*MsgHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{2}
}

func (m *MsgHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgHead.Unmarshal(m, b)
}
func (m *MsgHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgHead.Marshal(b, m, deterministic)
}
func (m *MsgHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgHead.Merge(m, src)
}
func (m *MsgHead) XXX_Size() int {
	return xxx_messageInfo_MsgHead.Size(m)
}
func (m *MsgHead) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgHead.DiscardUnknown(m)
}

var xxx_messageInfo_MsgHead proto.InternalMessageInfo

func (m *MsgHead) GetBodyLen() uint32 {
	if m != nil {
		return m.BodyLen
	}
	return 0
}

func (m *MsgHead) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *MsgHead) GetApp() uint32 {
	if m != nil {
		return m.App
	}
	return 0
}

func (m *MsgHead) GetServer() uint32 {
	if m != nil {
		return m.Server
	}
	return 0
}

func (m *MsgHead) GetServant() uint32 {
	if m != nil {
		return m.Servant
	}
	return 0
}

func (m *MsgHead) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *MsgHead) GetRouteId() uint64 {
	if m != nil {
		return m.RouteId
	}
	return 0
}

func (m *MsgHead) GetEncrypt() uint32 {
	if m != nil {
		return m.Encrypt
	}
	return 0
}

func (m *MsgHead) GetCacheIs() uint32 {
	if m != nil {
		return m.CacheIs
	}
	return 0
}

type MsgBody struct {
	Body                 []byte   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Extend               []byte   `protobuf:"bytes,2,opt,name=extend,proto3" json:"extend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgBody) Reset()         { *m = MsgBody{} }
func (m *MsgBody) String() string { return proto.CompactTextString(m) }
func (*MsgBody) ProtoMessage()    {}
func (*MsgBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{3}
}

func (m *MsgBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgBody.Unmarshal(m, b)
}
func (m *MsgBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgBody.Marshal(b, m, deterministic)
}
func (m *MsgBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBody.Merge(m, src)
}
func (m *MsgBody) XXX_Size() int {
	return xxx_messageInfo_MsgBody.Size(m)
}
func (m *MsgBody) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBody.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBody proto.InternalMessageInfo

func (m *MsgBody) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *MsgBody) GetExtend() []byte {
	if m != nil {
		return m.Extend
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "user_proto.Request")
	proto.RegisterType((*Respond)(nil), "user_proto.Respond")
	proto.RegisterType((*MsgHead)(nil), "user_proto.MsgHead")
	proto.RegisterType((*MsgBody)(nil), "user_proto.MsgBody")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_2bc2336598a3f7e0) }

var fileDescriptor_2bc2336598a3f7e0 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xbf, 0x6e, 0xc2, 0x30,
	0x10, 0xc6, 0x1b, 0xfe, 0xc4, 0xf4, 0x84, 0x2a, 0xe4, 0x4a, 0x95, 0xe9, 0x84, 0x32, 0x31, 0x31,
	0xf4, 0xcf, 0xda, 0xa1, 0x13, 0x48, 0x65, 0xf1, 0x0b, 0x44, 0x26, 0x3e, 0xd1, 0x48, 0xc8, 0x36,
	0xb6, 0x83, 0xca, 0xfb, 0xf6, 0x41, 0xaa, 0x33, 0xa1, 0xa1, 0x55, 0x97, 0x6e, 0xdf, 0x77, 0x5f,
	0xe2, 0xbb, 0xdf, 0x1d, 0xdc, 0x38, 0x6f, 0xa3, 0xad, 0xec, 0x6e, 0x91, 0x04, 0x87, 0x26, 0xa0,
	0x2f, 0x93, 0x2e, 0x1a, 0x60, 0x12, 0xf7, 0x0d, 0x86, 0xc8, 0x05, 0xb0, 0x03, 0xfa, 0x50, 0x5b,
	0x23, 0xb2, 0x59, 0x36, 0x67, 0xf2, 0x6c, 0x29, 0x09, 0xe8, 0x0f, 0xca, 0x44, 0xd1, 0x3b, 0x25,
	0xad, 0xe5, 0x13, 0xe8, 0x07, 0xdc, 0x8b, 0x7e, 0xaa, 0x92, 0xa4, 0x4a, 0x53, 0x6b, 0x31, 0x98,
	0x65, 0xf3, 0x5c, 0x92, 0xe4, 0x1c, 0x06, 0x1b, 0xab, 0x8f, 0x62, 0x38, 0xcb, 0xe6, 0x63, 0x99,
	0x74, 0xf1, 0x4c, 0x6d, 0x83, 0xb3, 0xa6, 0x8b, 0xb3, 0x2e, 0xe6, 0x77, 0x90, 0xe3, 0x47, 0x44,
	0xa3, 0x53, 0xbf, 0xb1, 0x6c, 0x5d, 0xf1, 0x99, 0x01, 0x5b, 0x87, 0xed, 0x12, 0x95, 0xe6, 0x53,
	0x18, 0xd1, 0xb7, 0xe5, 0x0e, 0xbf, 0xe7, 0x25, 0xff, 0x86, 0xe6, 0x92, 0xa4, 0xf7, 0x93, 0x64,
	0x02, 0x7d, 0xe5, 0xdc, 0x79, 0x5e, 0xe5, 0x1c, 0xb5, 0x22, 0x18, 0xf4, 0x69, 0x64, 0x26, 0x5b,
	0x77, 0xc9, 0x3c, 0xfc, 0x93, 0x39, 0xef, 0x98, 0xa7, 0x30, 0xf2, 0xb6, 0x89, 0x58, 0xd6, 0x5a,
	0xb0, 0x04, 0xce, 0x92, 0x5f, 0x69, 0x7a, 0x06, 0x4d, 0xe5, 0x8f, 0x2e, 0x8a, 0xd1, 0xe9, 0x99,
	0xd6, 0xd2, 0x4f, 0x95, 0xaa, 0xde, 0xb1, 0xac, 0x83, 0xb8, 0x3e, 0x45, 0xc9, 0xaf, 0x02, 0x6d,
	0x67, 0x1d, 0xb6, 0xaf, 0xb4, 0x89, 0x7f, 0x6c, 0xe7, 0xe1, 0xe5, 0x8c, 0xc2, 0x9f, 0x20, 0x5f,
	0x2a, 0xa3, 0x77, 0xc8, 0x6f, 0x17, 0xdd, 0xb1, 0x17, 0xed, 0xa5, 0xef, 0x7f, 0x15, 0xd3, 0x1d,
	0x8a, 0xab, 0x4d, 0x9e, 0x0a, 0x8f, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xd3, 0xca, 0x5b,
	0x30, 0x02, 0x00, 0x00,
}

// This following code was generated by tarsrpc
// Gernerated from protocol.proto
type Server struct {
	s model.Servant
}

//SetServant is required by the servant interface.
func (obj *Server) SetServant(s model.Servant) {
	obj.s = s
}

//AddServant is required by the servant interface
func (obj *Server) AddServant(imp impServer, objStr string) {
	tars.AddServant(obj, imp, objStr)
}

//TarsSetTimeout is required by the servant interface. t is the timeout in ms.
func (obj *Server) TarsSetTimeout(t int) {
	obj.s.TarsSetTimeout(t)
}

type impServer interface {
	Handle(input Request) (output Respond, err error)
}

//Dispatch is used to call the user implement of the defined method.
func (obj *Server) Dispatch(ctx context.Context, val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
	input := tools.Int8ToByte(req.SBuffer)
	var output []byte
	imp := val.(impServer)
	funcName := req.SFuncName
	switch funcName {

	case "Handle":
		inputDefine := Request{}
		if err = proto.Unmarshal(input, &inputDefine); err != nil {
			return err
		}
		res, err := imp.Handle(inputDefine)
		if err != nil {
			return err
		}
		output, err = proto.Marshal(&res)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var status map[string]string
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(output),
		Status:       status,
		SResultDesc:  "",
		Context:      req.Context,
	}
	return nil
}

// Handle is client rpc method as defined
func (obj *Server) Handle(input Request) (output Respond, err error) {
	var _status map[string]string
	var _context map[string]string
	var inputMarshal []byte
	inputMarshal, err = proto.Marshal(&input)
	if err != nil {
		return output, err
	}
	resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = obj.s.Tars_invoke(ctx, 0, "Handle", inputMarshal, _status, _context, resp)
	if err != nil {
		return output, err
	}
	if err = proto.Unmarshal(tools.Int8ToByte(resp.SBuffer), &output); err != nil {
		return output, err
	}
	return output, nil
}
