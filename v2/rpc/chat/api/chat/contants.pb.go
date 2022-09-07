// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: rpc/chat/api/chat/contants.proto

package chat

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 客户端类型
type CIMClientType int32

const (
	CIMClientType_kCIM_CLIENT_TYPE_DEFAULT    CIMClientType = 0 // unset
	CIMClientType_kCIM_CLIENT_TYPE_ANDROID    CIMClientType = 1 // 安卓
	CIMClientType_kCIM_CLIENT_TYPE_IOS        CIMClientType = 2 // iOS
	CIMClientType_kCIM_CLIENT_TYPE_WEB        CIMClientType = 3 // WebSocket
	CIMClientType_kCIM_CLIENT_TYPE_REST_API   CIMClientType = 4 // RestAPI
	CIMClientType_kCIM_CLIENT_TYPE_PC_WINDOWS CIMClientType = 5 // PC Windows
	CIMClientType_kCIM_CLIENT_TYPE_MAC_OS     CIMClientType = 6 // MAC
)

// Enum value maps for CIMClientType.
var (
	CIMClientType_name = map[int32]string{
		0: "kCIM_CLIENT_TYPE_DEFAULT",
		1: "kCIM_CLIENT_TYPE_ANDROID",
		2: "kCIM_CLIENT_TYPE_IOS",
		3: "kCIM_CLIENT_TYPE_WEB",
		4: "kCIM_CLIENT_TYPE_REST_API",
		5: "kCIM_CLIENT_TYPE_PC_WINDOWS",
		6: "kCIM_CLIENT_TYPE_MAC_OS",
	}
	CIMClientType_value = map[string]int32{
		"kCIM_CLIENT_TYPE_DEFAULT":    0,
		"kCIM_CLIENT_TYPE_ANDROID":    1,
		"kCIM_CLIENT_TYPE_IOS":        2,
		"kCIM_CLIENT_TYPE_WEB":        3,
		"kCIM_CLIENT_TYPE_REST_API":   4,
		"kCIM_CLIENT_TYPE_PC_WINDOWS": 5,
		"kCIM_CLIENT_TYPE_MAC_OS":     6,
	}
)

func (x CIMClientType) Enum() *CIMClientType {
	p := new(CIMClientType)
	*p = x
	return p
}

func (x CIMClientType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CIMClientType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_chat_api_chat_contants_proto_enumTypes[0].Descriptor()
}

func (CIMClientType) Type() protoreflect.EnumType {
	return &file_rpc_chat_api_chat_contants_proto_enumTypes[0]
}

func (x CIMClientType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CIMClientType.Descriptor instead.
func (CIMClientType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_chat_api_chat_contants_proto_rawDescGZIP(), []int{0}
}

// 会话类型
type IMSessionType int32

const (
	IMSessionType_kCIM_SESSION_TYPE_Invalid IMSessionType = 0 // 无效会话
	IMSessionType_kCIM_SESSION_TYPE_SINGLE  IMSessionType = 1 // 单聊
	IMSessionType_kCIM_SESSION_TYPE_GROUP   IMSessionType = 2 // 群聊
)

// Enum value maps for IMSessionType.
var (
	IMSessionType_name = map[int32]string{
		0: "kCIM_SESSION_TYPE_Invalid",
		1: "kCIM_SESSION_TYPE_SINGLE",
		2: "kCIM_SESSION_TYPE_GROUP",
	}
	IMSessionType_value = map[string]int32{
		"kCIM_SESSION_TYPE_Invalid": 0,
		"kCIM_SESSION_TYPE_SINGLE":  1,
		"kCIM_SESSION_TYPE_GROUP":   2,
	}
)

func (x IMSessionType) Enum() *IMSessionType {
	p := new(IMSessionType)
	*p = x
	return p
}

func (x IMSessionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IMSessionType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_chat_api_chat_contants_proto_enumTypes[1].Descriptor()
}

func (IMSessionType) Type() protoreflect.EnumType {
	return &file_rpc_chat_api_chat_contants_proto_enumTypes[1]
}

func (x IMSessionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IMSessionType.Descriptor instead.
func (IMSessionType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_chat_api_chat_contants_proto_rawDescGZIP(), []int{1}
}

// 消息类型
type IMMsgType int32

const (
	IMMsgType_kCIM_MSG_TYPE_UNKNOWN      IMMsgType = 0  // 未知类型消息，本地使用，发送时请勿使用
	IMMsgType_kCIM_MSG_TYPE_TEXT         IMMsgType = 1  // 文本
	IMMsgType_kCIM_MSG_TYPE_FILE         IMMsgType = 2  // 文件
	IMMsgType_kCIM_MSG_TYPE_IMAGE        IMMsgType = 3  // 图片
	IMMsgType_kCIM_MSG_TYPE_AUDIO        IMMsgType = 4  // 声音
	IMMsgType_kCIM_MSG_TYPE_VIDEO        IMMsgType = 5  // 视频
	IMMsgType_kCIM_MSG_TYPE_LOCATION     IMMsgType = 6  // 位置
	IMMsgType_kCIM_MSG_TYPE_ROBOT        IMMsgType = 7  // 图灵机器人消息
	IMMsgType_kCIM_MSG_TYPE_TIPS         IMMsgType = 8  // 提醒类型
	IMMsgType_kCIM_MSG_TYPE_NOTIFACATION IMMsgType = 9  // 系统通知（包括入群出群通知等）
	IMMsgType_kCIM_MSG_TYPE_AVCHAT       IMMsgType = 10 // 由音视频通话产生的消息
)

// Enum value maps for IMMsgType.
var (
	IMMsgType_name = map[int32]string{
		0:  "kCIM_MSG_TYPE_UNKNOWN",
		1:  "kCIM_MSG_TYPE_TEXT",
		2:  "kCIM_MSG_TYPE_FILE",
		3:  "kCIM_MSG_TYPE_IMAGE",
		4:  "kCIM_MSG_TYPE_AUDIO",
		5:  "kCIM_MSG_TYPE_VIDEO",
		6:  "kCIM_MSG_TYPE_LOCATION",
		7:  "kCIM_MSG_TYPE_ROBOT",
		8:  "kCIM_MSG_TYPE_TIPS",
		9:  "kCIM_MSG_TYPE_NOTIFACATION",
		10: "kCIM_MSG_TYPE_AVCHAT",
	}
	IMMsgType_value = map[string]int32{
		"kCIM_MSG_TYPE_UNKNOWN":      0,
		"kCIM_MSG_TYPE_TEXT":         1,
		"kCIM_MSG_TYPE_FILE":         2,
		"kCIM_MSG_TYPE_IMAGE":        3,
		"kCIM_MSG_TYPE_AUDIO":        4,
		"kCIM_MSG_TYPE_VIDEO":        5,
		"kCIM_MSG_TYPE_LOCATION":     6,
		"kCIM_MSG_TYPE_ROBOT":        7,
		"kCIM_MSG_TYPE_TIPS":         8,
		"kCIM_MSG_TYPE_NOTIFACATION": 9,
		"kCIM_MSG_TYPE_AVCHAT":       10,
	}
)

func (x IMMsgType) Enum() *IMMsgType {
	p := new(IMMsgType)
	*p = x
	return p
}

func (x IMMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IMMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_chat_api_chat_contants_proto_enumTypes[2].Descriptor()
}

func (IMMsgType) Type() protoreflect.EnumType {
	return &file_rpc_chat_api_chat_contants_proto_enumTypes[2]
}

func (x IMMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IMMsgType.Descriptor instead.
func (IMMsgType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_chat_api_chat_contants_proto_rawDescGZIP(), []int{2}
}

// 消息状态
type CIMMsgStatus int32

const (
	CIMMsgStatus_kCIM_MSG_STATUS_NONE      CIMMsgStatus = 0  // 默认，不能当查询条件，意义太多
	CIMMsgStatus_kCIM_MSG_STATUS_UNREAD    CIMMsgStatus = 1  // 收到消息，未读
	CIMMsgStatus_kCIM_MSG_STATUS_READ      CIMMsgStatus = 2  // 收到消息，已读
	CIMMsgStatus_kCIM_MSG_STATUS_DELETED   CIMMsgStatus = 3  // 已删
	CIMMsgStatus_kCIM_MSG_STATUS_SENDING   CIMMsgStatus = 4  // 发送中
	CIMMsgStatus_kCIM_MSG_STATUS_SENT      CIMMsgStatus = 5  // 已发送
	CIMMsgStatus_kCIM_MSG_STATUS_RECEIPT   CIMMsgStatus = 6  // 对方已读发送的内容
	CIMMsgStatus_kCIM_MSG_STATUS_DRAFT     CIMMsgStatus = 7  // 草稿
	CIMMsgStatus_kCIM_MSG_STATUS_SendCacel CIMMsgStatus = 8  // 发送取消
	CIMMsgStatus_kCIM_MSG_STATUS_REFUSED   CIMMsgStatus = 9  // 被对方拒绝，比如被对方加入黑名单等等
	CIMMsgStatus_kCIM_MSG_STATUS_FAILED    CIMMsgStatus = 10 // 发送失败，客户端使用
)

// Enum value maps for CIMMsgStatus.
var (
	CIMMsgStatus_name = map[int32]string{
		0:  "kCIM_MSG_STATUS_NONE",
		1:  "kCIM_MSG_STATUS_UNREAD",
		2:  "kCIM_MSG_STATUS_READ",
		3:  "kCIM_MSG_STATUS_DELETED",
		4:  "kCIM_MSG_STATUS_SENDING",
		5:  "kCIM_MSG_STATUS_SENT",
		6:  "kCIM_MSG_STATUS_RECEIPT",
		7:  "kCIM_MSG_STATUS_DRAFT",
		8:  "kCIM_MSG_STATUS_SendCacel",
		9:  "kCIM_MSG_STATUS_REFUSED",
		10: "kCIM_MSG_STATUS_FAILED",
	}
	CIMMsgStatus_value = map[string]int32{
		"kCIM_MSG_STATUS_NONE":      0,
		"kCIM_MSG_STATUS_UNREAD":    1,
		"kCIM_MSG_STATUS_READ":      2,
		"kCIM_MSG_STATUS_DELETED":   3,
		"kCIM_MSG_STATUS_SENDING":   4,
		"kCIM_MSG_STATUS_SENT":      5,
		"kCIM_MSG_STATUS_RECEIPT":   6,
		"kCIM_MSG_STATUS_DRAFT":     7,
		"kCIM_MSG_STATUS_SendCacel": 8,
		"kCIM_MSG_STATUS_REFUSED":   9,
		"kCIM_MSG_STATUS_FAILED":    10,
	}
)

func (x CIMMsgStatus) Enum() *CIMMsgStatus {
	p := new(CIMMsgStatus)
	*p = x
	return p
}

func (x CIMMsgStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CIMMsgStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_chat_api_chat_contants_proto_enumTypes[3].Descriptor()
}

func (CIMMsgStatus) Type() protoreflect.EnumType {
	return &file_rpc_chat_api_chat_contants_proto_enumTypes[3]
}

func (x CIMMsgStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CIMMsgStatus.Descriptor instead.
func (CIMMsgStatus) EnumDescriptor() ([]byte, []int) {
	return file_rpc_chat_api_chat_contants_proto_rawDescGZIP(), []int{3}
}

// 会话状态
type IMSessionStatusType int32

const (
	IMSessionStatusType_kCIM_SESSION_STATUS_UNKNOWN IMSessionStatusType = 0 // 未知
	IMSessionStatusType_kCIM_SESSION_STATUS_OK      IMSessionStatusType = 1
	IMSessionStatusType_kCIM_SESSION_STATUS_DELETE  IMSessionStatusType = 2
)

// Enum value maps for IMSessionStatusType.
var (
	IMSessionStatusType_name = map[int32]string{
		0: "kCIM_SESSION_STATUS_UNKNOWN",
		1: "kCIM_SESSION_STATUS_OK",
		2: "kCIM_SESSION_STATUS_DELETE",
	}
	IMSessionStatusType_value = map[string]int32{
		"kCIM_SESSION_STATUS_UNKNOWN": 0,
		"kCIM_SESSION_STATUS_OK":      1,
		"kCIM_SESSION_STATUS_DELETE":  2,
	}
)

func (x IMSessionStatusType) Enum() *IMSessionStatusType {
	p := new(IMSessionStatusType)
	*p = x
	return p
}

func (x IMSessionStatusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IMSessionStatusType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_chat_api_chat_contants_proto_enumTypes[4].Descriptor()
}

func (IMSessionStatusType) Type() protoreflect.EnumType {
	return &file_rpc_chat_api_chat_contants_proto_enumTypes[4]
}

func (x IMSessionStatusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IMSessionStatusType.Descriptor instead.
func (IMSessionStatusType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_chat_api_chat_contants_proto_rawDescGZIP(), []int{4}
}

// 消息属性
type IMMsgFeature int32

const (
	IMMsgFeature_kCIM_MSG_FEATURE_DEFAULT   IMMsgFeature = 0 // 默认
	IMMsgFeature_kCIM_MSG_FEATURE_LEAVE_MSG IMMsgFeature = 1 // 离线消息(和漫游消息有何区别)
	IMMsgFeature_kCIM_MSG_FEATURE_ROAM_MSG  IMMsgFeature = 2 // 漫游消息，多端同步接收，永久存储(或1年)
)

// Enum value maps for IMMsgFeature.
var (
	IMMsgFeature_name = map[int32]string{
		0: "kCIM_MSG_FEATURE_DEFAULT",
		1: "kCIM_MSG_FEATURE_LEAVE_MSG",
		2: "kCIM_MSG_FEATURE_ROAM_MSG",
	}
	IMMsgFeature_value = map[string]int32{
		"kCIM_MSG_FEATURE_DEFAULT":   0,
		"kCIM_MSG_FEATURE_LEAVE_MSG": 1,
		"kCIM_MSG_FEATURE_ROAM_MSG":  2,
	}
)

func (x IMMsgFeature) Enum() *IMMsgFeature {
	p := new(IMMsgFeature)
	*p = x
	return p
}

func (x IMMsgFeature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IMMsgFeature) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_chat_api_chat_contants_proto_enumTypes[5].Descriptor()
}

func (IMMsgFeature) Type() protoreflect.EnumType {
	return &file_rpc_chat_api_chat_contants_proto_enumTypes[5]
}

func (x IMMsgFeature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IMMsgFeature.Descriptor instead.
func (IMMsgFeature) EnumDescriptor() ([]byte, []int) {
	return file_rpc_chat_api_chat_contants_proto_rawDescGZIP(), []int{5}
}

// 消息错误码
type IMResCode int32

const (
	IMResCode_kCIM_RES_CODE_UNKNOWN IMResCode = 0 // unknown
	IMResCode_kCIM_RES_CODE_OK      IMResCode = 1 // 一切正常
)

// Enum value maps for IMResCode.
var (
	IMResCode_name = map[int32]string{
		0: "kCIM_RES_CODE_UNKNOWN",
		1: "kCIM_RES_CODE_OK",
	}
	IMResCode_value = map[string]int32{
		"kCIM_RES_CODE_UNKNOWN": 0,
		"kCIM_RES_CODE_OK":      1,
	}
)

func (x IMResCode) Enum() *IMResCode {
	p := new(IMResCode)
	*p = x
	return p
}

func (x IMResCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IMResCode) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_chat_api_chat_contants_proto_enumTypes[6].Descriptor()
}

func (IMResCode) Type() protoreflect.EnumType {
	return &file_rpc_chat_api_chat_contants_proto_enumTypes[6]
}

func (x IMResCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IMResCode.Descriptor instead.
func (IMResCode) EnumDescriptor() ([]byte, []int) {
	return file_rpc_chat_api_chat_contants_proto_rawDescGZIP(), []int{6}
}

var File_rpc_chat_api_chat_contants_proto protoreflect.FileDescriptor

var file_rpc_chat_api_chat_contants_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x2a, 0xdc, 0x01, 0x0a, 0x0d, 0x43, 0x49, 0x4d,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x6b, 0x43,
	0x49, 0x4d, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44,
	0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x6b, 0x43, 0x49, 0x4d,
	0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x44,
	0x52, 0x4f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x43,
	0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4f, 0x53, 0x10, 0x02,
	0x12, 0x18, 0x0a, 0x14, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x6b, 0x43,
	0x49, 0x4d, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52,
	0x45, 0x53, 0x54, 0x5f, 0x41, 0x50, 0x49, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x6b, 0x43, 0x49,
	0x4d, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x43,
	0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x6b, 0x43,
	0x49, 0x4d, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d,
	0x41, 0x43, 0x5f, 0x4f, 0x53, 0x10, 0x06, 0x2a, 0x69, 0x0a, 0x0d, 0x49, 0x4d, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x6b, 0x43, 0x49, 0x4d,
	0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x6b, 0x43, 0x49, 0x4d, 0x5f,
	0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x49, 0x4e,
	0x47, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x53, 0x45,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50,
	0x10, 0x02, 0x2a, 0xa8, 0x02, 0x0a, 0x09, 0x49, 0x4d, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x19, 0x0a, 0x15, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x6b,
	0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x58,
	0x54, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x6b,
	0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4d, 0x41,
	0x47, 0x45, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x10, 0x04, 0x12, 0x17, 0x0a,
	0x13, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56,
	0x49, 0x44, 0x45, 0x4f, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d,
	0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x52, 0x4f, 0x42, 0x4f, 0x54, 0x10, 0x07, 0x12, 0x16, 0x0a, 0x12, 0x6b,
	0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x50,
	0x53, 0x10, 0x08, 0x12, 0x1e, 0x0a, 0x1a, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x41, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x09, 0x12, 0x18, 0x0a, 0x14, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x56, 0x43, 0x48, 0x41, 0x54, 0x10, 0x0a, 0x2a, 0xc2, 0x02,
	0x0a, 0x0c, 0x43, 0x49, 0x4d, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x14, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x6b, 0x43, 0x49, 0x4d,
	0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x52, 0x45,
	0x41, 0x44, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x1b,
	0x0a, 0x17, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x6b,
	0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x6b, 0x43, 0x49, 0x4d,
	0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x45, 0x4e, 0x54,
	0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x50, 0x54, 0x10, 0x06, 0x12,
	0x19, 0x0a, 0x15, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x07, 0x12, 0x1d, 0x0a, 0x19, 0x6b, 0x43,
	0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x65,
	0x6e, 0x64, 0x43, 0x61, 0x63, 0x65, 0x6c, 0x10, 0x08, 0x12, 0x1b, 0x0a, 0x17, 0x6b, 0x43, 0x49,
	0x4d, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x46,
	0x55, 0x53, 0x45, 0x44, 0x10, 0x09, 0x12, 0x1a, 0x0a, 0x16, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d,
	0x53, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x0a, 0x2a, 0x72, 0x0a, 0x13, 0x49, 0x4d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x6b, 0x43, 0x49,
	0x4d, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x6b, 0x43,
	0x49, 0x4d, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x53,
	0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x2a, 0x6b, 0x0a, 0x0c, 0x49, 0x4d, 0x4d, 0x73, 0x67, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d,
	0x53, 0x47, 0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47,
	0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x5f, 0x4d,
	0x53, 0x47, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x4d, 0x53, 0x47,
	0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x52, 0x4f, 0x41, 0x4d, 0x5f, 0x4d, 0x53,
	0x47, 0x10, 0x02, 0x2a, 0x3c, 0x0a, 0x09, 0x49, 0x4d, 0x52, 0x65, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x0a, 0x15, 0x6b, 0x43, 0x49, 0x4d, 0x5f, 0x52, 0x45, 0x53, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x6b,
	0x43, 0x49, 0x4d, 0x5f, 0x52, 0x45, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x4b, 0x10,
	0x01, 0x42, 0x14, 0x5a, 0x12, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_chat_api_chat_contants_proto_rawDescOnce sync.Once
	file_rpc_chat_api_chat_contants_proto_rawDescData = file_rpc_chat_api_chat_contants_proto_rawDesc
)

func file_rpc_chat_api_chat_contants_proto_rawDescGZIP() []byte {
	file_rpc_chat_api_chat_contants_proto_rawDescOnce.Do(func() {
		file_rpc_chat_api_chat_contants_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_chat_api_chat_contants_proto_rawDescData)
	})
	return file_rpc_chat_api_chat_contants_proto_rawDescData
}

var file_rpc_chat_api_chat_contants_proto_enumTypes = make([]protoimpl.EnumInfo, 7)
var file_rpc_chat_api_chat_contants_proto_goTypes = []interface{}{
	(CIMClientType)(0),       // 0: chat.CIMClientType
	(IMSessionType)(0),       // 1: chat.IMSessionType
	(IMMsgType)(0),           // 2: chat.IMMsgType
	(CIMMsgStatus)(0),        // 3: chat.CIMMsgStatus
	(IMSessionStatusType)(0), // 4: chat.IMSessionStatusType
	(IMMsgFeature)(0),        // 5: chat.IMMsgFeature
	(IMResCode)(0),           // 6: chat.IMResCode
}
var file_rpc_chat_api_chat_contants_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_chat_api_chat_contants_proto_init() }
func file_rpc_chat_api_chat_contants_proto_init() {
	if File_rpc_chat_api_chat_contants_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_chat_api_chat_contants_proto_rawDesc,
			NumEnums:      7,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_chat_api_chat_contants_proto_goTypes,
		DependencyIndexes: file_rpc_chat_api_chat_contants_proto_depIdxs,
		EnumInfos:         file_rpc_chat_api_chat_contants_proto_enumTypes,
	}.Build()
	File_rpc_chat_api_chat_contants_proto = out.File
	file_rpc_chat_api_chat_contants_proto_rawDesc = nil
	file_rpc_chat_api_chat_contants_proto_goTypes = nil
	file_rpc_chat_api_chat_contants_proto_depIdxs = nil
}