// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package xuperp2p

import (
	fmt "fmt"
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

type XuperMessage_MessageType int32

const (
	XuperMessage_SENDBLOCK                XuperMessage_MessageType = 0
	XuperMessage_POSTTX                   XuperMessage_MessageType = 1
	XuperMessage_BATCHPOSTTX              XuperMessage_MessageType = 2
	XuperMessage_GET_BLOCK                XuperMessage_MessageType = 3
	XuperMessage_PING                     XuperMessage_MessageType = 4
	XuperMessage_GET_BLOCKCHAINSTATUS     XuperMessage_MessageType = 5
	XuperMessage_GET_BLOCK_RES            XuperMessage_MessageType = 6
	XuperMessage_GET_BLOCKCHAINSTATUS_RES XuperMessage_MessageType = 7
	// 向邻近确认区块是否为最新状态区块
	XuperMessage_CONFIRM_BLOCKCHAINSTATUS     XuperMessage_MessageType = 8
	XuperMessage_CONFIRM_BLOCKCHAINSTATUS_RES XuperMessage_MessageType = 9
	XuperMessage_MSG_TYPE_NONE                XuperMessage_MessageType = 10
	// 询问RPC端口信息
	XuperMessage_GET_RPC_PORT     XuperMessage_MessageType = 11
	XuperMessage_GET_RPC_PORT_RES XuperMessage_MessageType = 12
)

var XuperMessage_MessageType_name = map[int32]string{
	0:  "SENDBLOCK",
	1:  "POSTTX",
	2:  "BATCHPOSTTX",
	3:  "GET_BLOCK",
	4:  "PING",
	5:  "GET_BLOCKCHAINSTATUS",
	6:  "GET_BLOCK_RES",
	7:  "GET_BLOCKCHAINSTATUS_RES",
	8:  "CONFIRM_BLOCKCHAINSTATUS",
	9:  "CONFIRM_BLOCKCHAINSTATUS_RES",
	10: "MSG_TYPE_NONE",
	11: "GET_RPC_PORT",
	12: "GET_RPC_PORT_RES",
}

var XuperMessage_MessageType_value = map[string]int32{
	"SENDBLOCK":                    0,
	"POSTTX":                       1,
	"BATCHPOSTTX":                  2,
	"GET_BLOCK":                    3,
	"PING":                         4,
	"GET_BLOCKCHAINSTATUS":         5,
	"GET_BLOCK_RES":                6,
	"GET_BLOCKCHAINSTATUS_RES":     7,
	"CONFIRM_BLOCKCHAINSTATUS":     8,
	"CONFIRM_BLOCKCHAINSTATUS_RES": 9,
	"MSG_TYPE_NONE":                10,
	"GET_RPC_PORT":                 11,
	"GET_RPC_PORT_RES":             12,
}

func (x XuperMessage_MessageType) String() string {
	return proto.EnumName(XuperMessage_MessageType_name, int32(x))
}

func (XuperMessage_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

type XuperMessage_ErrorType int32

const (
	// success
	XuperMessage_SUCCESS XuperMessage_ErrorType = 0
	XuperMessage_NONE    XuperMessage_ErrorType = 1
	// common error
	XuperMessage_UNKNOW_ERROR             XuperMessage_ErrorType = 2
	XuperMessage_CHECK_SUM_ERROR          XuperMessage_ErrorType = 3
	XuperMessage_UNMARSHAL_MSG_BODY_ERROR XuperMessage_ErrorType = 4
	XuperMessage_CONNECT_REFUSE           XuperMessage_ErrorType = 5
	// block error
	XuperMessage_GET_BLOCKCHAIN_ERROR           XuperMessage_ErrorType = 6
	XuperMessage_BLOCKCHAIN_NOTEXIST            XuperMessage_ErrorType = 7
	XuperMessage_GET_BLOCK_ERROR                XuperMessage_ErrorType = 8
	XuperMessage_CONFIRM_BLOCKCHAINSTATUS_ERROR XuperMessage_ErrorType = 9
)

var XuperMessage_ErrorType_name = map[int32]string{
	0: "SUCCESS",
	1: "NONE",
	2: "UNKNOW_ERROR",
	3: "CHECK_SUM_ERROR",
	4: "UNMARSHAL_MSG_BODY_ERROR",
	5: "CONNECT_REFUSE",
	6: "GET_BLOCKCHAIN_ERROR",
	7: "BLOCKCHAIN_NOTEXIST",
	8: "GET_BLOCK_ERROR",
	9: "CONFIRM_BLOCKCHAINSTATUS_ERROR",
}

var XuperMessage_ErrorType_value = map[string]int32{
	"SUCCESS":                        0,
	"NONE":                           1,
	"UNKNOW_ERROR":                   2,
	"CHECK_SUM_ERROR":                3,
	"UNMARSHAL_MSG_BODY_ERROR":       4,
	"CONNECT_REFUSE":                 5,
	"GET_BLOCKCHAIN_ERROR":           6,
	"BLOCKCHAIN_NOTEXIST":            7,
	"GET_BLOCK_ERROR":                8,
	"CONFIRM_BLOCKCHAINSTATUS_ERROR": 9,
}

func (x XuperMessage_ErrorType) String() string {
	return proto.EnumName(XuperMessage_ErrorType_name, int32(x))
}

func (XuperMessage_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 1}
}

// XuperMessage is the message of Xuper p2p server
type XuperMessage struct {
	Header               *XuperMessage_MessageHeader `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Data                 *XuperMessage_MessageData   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *XuperMessage) Reset()         { *m = XuperMessage{} }
func (m *XuperMessage) String() string { return proto.CompactTextString(m) }
func (*XuperMessage) ProtoMessage()    {}
func (*XuperMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *XuperMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XuperMessage.Unmarshal(m, b)
}
func (m *XuperMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XuperMessage.Marshal(b, m, deterministic)
}
func (m *XuperMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XuperMessage.Merge(m, src)
}
func (m *XuperMessage) XXX_Size() int {
	return xxx_messageInfo_XuperMessage.Size(m)
}
func (m *XuperMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_XuperMessage.DiscardUnknown(m)
}

var xxx_messageInfo_XuperMessage proto.InternalMessageInfo

func (m *XuperMessage) GetHeader() *XuperMessage_MessageHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *XuperMessage) GetData() *XuperMessage_MessageData {
	if m != nil {
		return m.Data
	}
	return nil
}

// MessageHeader is the message header of Xuper p2p server
type XuperMessage_MessageHeader struct {
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// dataCheckSum is the message data checksum, it can be used check where the message have been received
	Logid                string                   `protobuf:"bytes,2,opt,name=logid,proto3" json:"logid,omitempty"`
	From                 string                   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	Bcname               string                   `protobuf:"bytes,4,opt,name=bcname,proto3" json:"bcname,omitempty"`
	Type                 XuperMessage_MessageType `protobuf:"varint,5,opt,name=type,proto3,enum=xuperp2p.XuperMessage_MessageType" json:"type,omitempty"`
	DataCheckSum         uint32                   `protobuf:"varint,6,opt,name=dataCheckSum,proto3" json:"dataCheckSum,omitempty"`
	ErrorType            XuperMessage_ErrorType   `protobuf:"varint,7,opt,name=errorType,proto3,enum=xuperp2p.XuperMessage_ErrorType" json:"errorType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *XuperMessage_MessageHeader) Reset()         { *m = XuperMessage_MessageHeader{} }
func (m *XuperMessage_MessageHeader) String() string { return proto.CompactTextString(m) }
func (*XuperMessage_MessageHeader) ProtoMessage()    {}
func (*XuperMessage_MessageHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

func (m *XuperMessage_MessageHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XuperMessage_MessageHeader.Unmarshal(m, b)
}
func (m *XuperMessage_MessageHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XuperMessage_MessageHeader.Marshal(b, m, deterministic)
}
func (m *XuperMessage_MessageHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XuperMessage_MessageHeader.Merge(m, src)
}
func (m *XuperMessage_MessageHeader) XXX_Size() int {
	return xxx_messageInfo_XuperMessage_MessageHeader.Size(m)
}
func (m *XuperMessage_MessageHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_XuperMessage_MessageHeader.DiscardUnknown(m)
}

var xxx_messageInfo_XuperMessage_MessageHeader proto.InternalMessageInfo

func (m *XuperMessage_MessageHeader) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *XuperMessage_MessageHeader) GetLogid() string {
	if m != nil {
		return m.Logid
	}
	return ""
}

func (m *XuperMessage_MessageHeader) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *XuperMessage_MessageHeader) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *XuperMessage_MessageHeader) GetType() XuperMessage_MessageType {
	if m != nil {
		return m.Type
	}
	return XuperMessage_SENDBLOCK
}

func (m *XuperMessage_MessageHeader) GetDataCheckSum() uint32 {
	if m != nil {
		return m.DataCheckSum
	}
	return 0
}

func (m *XuperMessage_MessageHeader) GetErrorType() XuperMessage_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return XuperMessage_SUCCESS
}

// MessageData is the message data of Xuper p2p server
type XuperMessage_MessageData struct {
	// msgInfo is the message infomation, use protobuf coding style
	MsgInfo              []byte   `protobuf:"bytes,3,opt,name=msgInfo,proto3" json:"msgInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XuperMessage_MessageData) Reset()         { *m = XuperMessage_MessageData{} }
func (m *XuperMessage_MessageData) String() string { return proto.CompactTextString(m) }
func (*XuperMessage_MessageData) ProtoMessage()    {}
func (*XuperMessage_MessageData) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 1}
}

func (m *XuperMessage_MessageData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XuperMessage_MessageData.Unmarshal(m, b)
}
func (m *XuperMessage_MessageData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XuperMessage_MessageData.Marshal(b, m, deterministic)
}
func (m *XuperMessage_MessageData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XuperMessage_MessageData.Merge(m, src)
}
func (m *XuperMessage_MessageData) XXX_Size() int {
	return xxx_messageInfo_XuperMessage_MessageData.Size(m)
}
func (m *XuperMessage_MessageData) XXX_DiscardUnknown() {
	xxx_messageInfo_XuperMessage_MessageData.DiscardUnknown(m)
}

var xxx_messageInfo_XuperMessage_MessageData proto.InternalMessageInfo

func (m *XuperMessage_MessageData) GetMsgInfo() []byte {
	if m != nil {
		return m.MsgInfo
	}
	return nil
}

func init() {
	proto.RegisterEnum("xuperp2p.XuperMessage_MessageType", XuperMessage_MessageType_name, XuperMessage_MessageType_value)
	proto.RegisterEnum("xuperp2p.XuperMessage_ErrorType", XuperMessage_ErrorType_name, XuperMessage_ErrorType_value)
	proto.RegisterType((*XuperMessage)(nil), "xuperp2p.XuperMessage")
	proto.RegisterType((*XuperMessage_MessageHeader)(nil), "xuperp2p.XuperMessage.MessageHeader")
	proto.RegisterType((*XuperMessage_MessageData)(nil), "xuperp2p.XuperMessage.MessageData")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x51, 0x6e, 0xda, 0x40,
	0x10, 0x86, 0x63, 0x62, 0x6c, 0x3c, 0x40, 0xb2, 0x9d, 0x44, 0xad, 0x15, 0x45, 0x15, 0xb2, 0x2a,
	0x35, 0x4f, 0x3c, 0xa4, 0x52, 0x9f, 0xaa, 0x4a, 0x64, 0xb3, 0x01, 0x44, 0x58, 0xa3, 0x5d, 0x5b,
	0x25, 0x4f, 0x96, 0x13, 0x1c, 0x1a, 0xb5, 0x60, 0xcb, 0x90, 0xaa, 0x39, 0x41, 0xaf, 0xd0, 0xc3,
	0xf4, 0x3e, 0xbd, 0x46, 0xb5, 0x6b, 0x43, 0x88, 0x52, 0xda, 0x27, 0x98, 0x7f, 0xbe, 0x7f, 0xfe,
	0xd9, 0x91, 0x0c, 0xcd, 0x59, 0xb2, 0x58, 0xc4, 0xd3, 0xa4, 0x9d, 0xe5, 0xe9, 0x32, 0xc5, 0xda,
	0xf7, 0xfb, 0x2c, 0xc9, 0xb3, 0xd3, 0xcc, 0xfb, 0x65, 0x43, 0x63, 0xac, 0x8a, 0x61, 0x01, 0xe0,
	0x07, 0xb0, 0x7a, 0x49, 0x3c, 0x49, 0x72, 0xd7, 0x68, 0x19, 0x27, 0xf5, 0xd3, 0x37, 0xed, 0x15,
	0xdb, 0xde, 0xe4, 0xda, 0xe5, 0x6f, 0xc1, 0x8a, 0xd2, 0x83, 0xef, 0xc1, 0x3c, 0x8f, 0x97, 0xb1,
	0x5b, 0xd1, 0x5e, 0xef, 0xdf, 0x5e, 0x45, 0x0a, 0xcd, 0x1f, 0xfd, 0xa8, 0x40, 0xf3, 0xc9, 0x44,
	0x74, 0xc1, 0xfe, 0x96, 0xe4, 0x8b, 0xbb, 0x74, 0xae, 0x17, 0x71, 0xc4, 0xaa, 0xc4, 0x43, 0xa8,
	0x7e, 0x4d, 0xa7, 0x77, 0x13, 0x1d, 0xe2, 0x88, 0xa2, 0x40, 0x04, 0xf3, 0x36, 0x4f, 0x67, 0xee,
	0xae, 0x16, 0xf5, 0x7f, 0x7c, 0x09, 0xd6, 0xf5, 0xcd, 0x3c, 0x9e, 0x25, 0xae, 0xa9, 0xd5, 0xb2,
	0x52, 0x5b, 0x2e, 0x1f, 0xb2, 0xc4, 0xad, 0xb6, 0x8c, 0x93, 0xbd, 0xff, 0x6d, 0x19, 0x3c, 0x64,
	0x89, 0xd0, 0x3c, 0x7a, 0xd0, 0x98, 0xc4, 0xcb, 0x98, 0x7e, 0x4e, 0x6e, 0xbe, 0xc8, 0xfb, 0x99,
	0x6b, 0xb5, 0x8c, 0x93, 0xa6, 0x78, 0xa2, 0xe1, 0x47, 0x70, 0x92, 0x3c, 0x4f, 0x73, 0x65, 0x73,
	0x6d, 0x1d, 0xd0, 0xda, 0x12, 0xc0, 0x56, 0x9c, 0x78, 0xb4, 0x1c, 0xbd, 0x85, 0xfa, 0xc6, 0x79,
	0xd4, 0x19, 0x66, 0x8b, 0x69, 0x7f, 0x7e, 0x9b, 0xea, 0x97, 0x35, 0xc4, 0xaa, 0xf4, 0x7e, 0x56,
	0xd6, 0xa4, 0x32, 0x62, 0x13, 0x1c, 0xc9, 0xf8, 0xf9, 0xd9, 0xa5, 0x4f, 0x07, 0x64, 0x07, 0x01,
	0xac, 0x91, 0x2f, 0x83, 0x60, 0x4c, 0x0c, 0xdc, 0x87, 0xfa, 0x59, 0x27, 0xa0, 0xbd, 0x52, 0xa8,
	0x28, 0xb6, 0xcb, 0x82, 0xa8, 0x60, 0x77, 0xb1, 0x06, 0xe6, 0xa8, 0xcf, 0xbb, 0xc4, 0x44, 0x17,
	0x0e, 0xd7, 0x0d, 0xda, 0xeb, 0xf4, 0xb9, 0x0c, 0x3a, 0x41, 0x28, 0x49, 0x15, 0x5f, 0x40, 0x73,
	0xdd, 0x89, 0x04, 0x93, 0xc4, 0xc2, 0x63, 0x70, 0xff, 0x06, 0xeb, 0xae, 0xad, 0xba, 0xd4, 0xe7,
	0x17, 0x7d, 0x31, 0x7c, 0x3e, 0xae, 0x86, 0x2d, 0x38, 0xde, 0xd6, 0xd5, 0x7e, 0x47, 0x05, 0x0e,
	0x65, 0x37, 0x0a, 0xae, 0x46, 0x2c, 0xe2, 0x3e, 0x67, 0x04, 0x90, 0x40, 0x43, 0x05, 0x8a, 0x11,
	0x8d, 0x46, 0xbe, 0x08, 0x48, 0x1d, 0x0f, 0x81, 0x6c, 0x2a, 0xda, 0xda, 0xf0, 0x7e, 0x1b, 0xe0,
	0xac, 0x8f, 0x8b, 0x75, 0xb0, 0x65, 0x48, 0x29, 0x93, 0x92, 0xec, 0xa8, 0xa7, 0xea, 0x61, 0x86,
	0x1a, 0x16, 0xf2, 0x01, 0xf7, 0x3f, 0x45, 0x4c, 0x08, 0x5f, 0x90, 0x0a, 0x1e, 0xc0, 0x3e, 0xed,
	0x31, 0x3a, 0x88, 0x64, 0x38, 0x2c, 0xc5, 0x5d, 0xf5, 0x8c, 0x90, 0x0f, 0x3b, 0x42, 0xf6, 0x3a,
	0x97, 0x91, 0x5a, 0xe8, 0xcc, 0x3f, 0xbf, 0x2a, 0xbb, 0x26, 0x22, 0xec, 0x51, 0x9f, 0x73, 0x46,
	0x55, 0xf4, 0x45, 0x28, 0x19, 0xa9, 0x3e, 0xbf, 0x61, 0x49, 0x5b, 0xf8, 0x0a, 0x0e, 0x36, 0x54,
	0xee, 0x07, 0x6c, 0xdc, 0x97, 0x01, 0xb1, 0x55, 0xf2, 0xe3, 0x71, 0x0b, 0xba, 0x86, 0x1e, 0xbc,
	0xde, 0x7a, 0xa2, 0x82, 0x71, 0xae, 0x2d, 0xfd, 0x3d, 0xbf, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff,
	0xf4, 0xad, 0x25, 0xe6, 0xe0, 0x03, 0x00, 0x00,
}
