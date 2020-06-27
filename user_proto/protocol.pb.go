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
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcb, 0x6a, 0xf3, 0x30,
	0x10, 0x85, 0x7f, 0xe5, 0x62, 0x85, 0x21, 0xfc, 0x04, 0x15, 0x8a, 0xd2, 0x55, 0xf0, 0x2a, 0x2b,
	0x2f, 0x7a, 0xd9, 0x76, 0xd1, 0x55, 0x0a, 0xcd, 0x46, 0x2f, 0x60, 0x9c, 0x68, 0x08, 0x06, 0x23,
	0x29, 0x92, 0x1c, 0x9a, 0x47, 0xec, 0x5b, 0x95, 0x51, 0xec, 0x3a, 0x2d, 0xdd, 0x74, 0x77, 0xce,
	0x8c, 0xd0, 0x9c, 0x6f, 0x06, 0xfe, 0x3b, 0x6f, 0xa3, 0xdd, 0xdb, 0xa6, 0x48, 0x42, 0x40, 0x1b,
	0xd0, 0x97, 0x49, 0xe7, 0x2d, 0x70, 0x85, 0xc7, 0x16, 0x43, 0x14, 0x12, 0xf8, 0x09, 0x7d, 0xa8,
	0xad, 0x91, 0x6c, 0xc5, 0xd6, 0x5c, 0xf5, 0x96, 0x3a, 0x01, 0xfd, 0xa9, 0x32, 0x51, 0x8e, 0x2e,
	0x9d, 0xce, 0x8a, 0x05, 0x8c, 0x03, 0x1e, 0xe5, 0x38, 0x55, 0x49, 0x52, 0xa5, 0xad, 0xb5, 0x9c,
	0xac, 0xd8, 0x3a, 0x53, 0x24, 0x85, 0x80, 0xc9, 0xce, 0xea, 0xb3, 0x9c, 0xae, 0xd8, 0x7a, 0xae,
	0x92, 0xce, 0x9f, 0x68, 0x6c, 0x70, 0xd6, 0x0c, 0x6d, 0x36, 0xb4, 0xc5, 0x2d, 0x64, 0xf8, 0x1e,
	0xd1, 0xe8, 0x34, 0x6f, 0xae, 0x3a, 0x97, 0x7f, 0x30, 0xe0, 0xdb, 0x70, 0xd8, 0x60, 0xa5, 0xc5,
	0x12, 0x66, 0xf4, 0xb6, 0x6c, 0xf0, 0x2b, 0x2f, 0xf9, 0x37, 0x34, 0xd7, 0x24, 0xa3, 0xef, 0x24,
	0x0b, 0x18, 0x57, 0xce, 0xf5, 0x79, 0x2b, 0xe7, 0x68, 0x14, 0xc1, 0xa0, 0x4f, 0x91, 0xb9, 0xea,
	0xdc, 0x35, 0xf3, 0xf4, 0x57, 0xe6, 0x6c, 0x60, 0x5e, 0xc2, 0xcc, 0xdb, 0x36, 0x62, 0x59, 0x6b,
	0xc9, 0x13, 0x38, 0x4f, 0xfe, 0x55, 0xd3, 0x37, 0x68, 0xf6, 0xfe, 0xec, 0xa2, 0x9c, 0x5d, 0xbe,
	0xe9, 0x2c, 0xad, 0x60, 0x1b, 0x0e, 0x2f, 0x84, 0xfb, 0x87, 0x15, 0xdc, 0x3f, 0xf7, 0x79, 0xc5,
	0x23, 0x64, 0x9b, 0xca, 0xe8, 0x06, 0xc5, 0x4d, 0x31, 0x5c, 0xb4, 0xe8, 0xce, 0x79, 0xf7, 0xa3,
	0x98, 0x96, 0x9d, 0xff, 0xdb, 0x65, 0xa9, 0xf0, 0xf0, 0x19, 0x00, 0x00, 0xff, 0xff, 0x09, 0xad,
	0x05, 0x52, 0x15, 0x02, 0x00, 0x00,
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
