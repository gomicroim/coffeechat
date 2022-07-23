// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/chat/chat.proto

package chat

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// 消息已读回复请求（我方）
type CIMMsgDataReadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 消息发送方
	SessionId   uint64         `protobuf:"varint,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	MsgId       uint64         `protobuf:"varint,3,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"` // 服务器消息ID，在该ID之前的所有消息被标记为已读
	SessionType CIMSessionType `protobuf:"varint,4,opt,name=session_type,json=sessionType,proto3,enum=coffeechat.CIMSessionType" json:"session_type,omitempty"`
}

func (x *CIMMsgDataReadReq) Reset() {
	*x = CIMMsgDataReadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMMsgDataReadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMMsgDataReadReq) ProtoMessage() {}

func (x *CIMMsgDataReadReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMMsgDataReadReq.ProtoReflect.Descriptor instead.
func (*CIMMsgDataReadReq) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *CIMMsgDataReadReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CIMMsgDataReadReq) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *CIMMsgDataReadReq) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *CIMMsgDataReadReq) GetSessionType() CIMSessionType {
	if x != nil {
		return x.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

// 消息已读回复响应
type CIMMsgDataReadRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CIMMsgDataReadRsp) Reset() {
	*x = CIMMsgDataReadRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMMsgDataReadRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMMsgDataReadRsp) ProtoMessage() {}

func (x *CIMMsgDataReadRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMMsgDataReadRsp.ProtoReflect.Descriptor instead.
func (*CIMMsgDataReadRsp) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{1}
}

// 最近聊天会话列表请求
type CIMRecentContactSessionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LatestUpdateTime uint32 `protobuf:"varint,2,opt,name=latest_update_time,json=latestUpdateTime,proto3" json:"latest_update_time,omitempty"` // 最后更新时间
}

func (x *CIMRecentContactSessionReq) Reset() {
	*x = CIMRecentContactSessionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMRecentContactSessionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMRecentContactSessionReq) ProtoMessage() {}

func (x *CIMRecentContactSessionReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMRecentContactSessionReq.ProtoReflect.Descriptor instead.
func (*CIMRecentContactSessionReq) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{2}
}

func (x *CIMRecentContactSessionReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CIMRecentContactSessionReq) GetLatestUpdateTime() uint32 {
	if x != nil {
		return x.LatestUpdateTime
	}
	return 0
}

type CIMRecentContactSessionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId             uint64                   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UnreadCounts       uint32                   `protobuf:"varint,2,opt,name=unread_counts,json=unreadCounts,proto3" json:"unread_counts,omitempty"`                    // 总未读数量
	ContactSessionList []*CIMContactSessionInfo `protobuf:"bytes,3,rep,name=contact_session_list,json=contactSessionList,proto3" json:"contact_session_list,omitempty"` // 会话列表
}

func (x *CIMRecentContactSessionRsp) Reset() {
	*x = CIMRecentContactSessionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMRecentContactSessionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMRecentContactSessionRsp) ProtoMessage() {}

func (x *CIMRecentContactSessionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMRecentContactSessionRsp.ProtoReflect.Descriptor instead.
func (*CIMRecentContactSessionRsp) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{3}
}

func (x *CIMRecentContactSessionRsp) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CIMRecentContactSessionRsp) GetUnreadCounts() uint32 {
	if x != nil {
		return x.UnreadCounts
	}
	return 0
}

func (x *CIMRecentContactSessionRsp) GetContactSessionList() []*CIMContactSessionInfo {
	if x != nil {
		return x.ContactSessionList
	}
	return nil
}

// 历史离线聊天消息请求
type CIMGetMsgListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cmd id:		0x0205
	UserId      uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionType CIMSessionType `protobuf:"varint,2,opt,name=session_type,json=sessionType,proto3,enum=coffeechat.CIMSessionType" json:"session_type,omitempty"`
	SessionId   uint64         `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	//   uint64 from_time_stamp = 4; // 起始时间点，单位：毫秒
	//   uint64 end_time_stamp = 5;  // 结束时间点，单位：毫秒
	EndMsgId   uint64 `protobuf:"varint,4,opt,name=end_msg_id,json=endMsgId,proto3" json:"end_msg_id,omitempty"`     // 结束服务器消息id(不包含在查询结果中)
	LimitCount uint32 `protobuf:"varint,6,opt,name=limit_count,json=limitCount,proto3" json:"limit_count,omitempty"` // 本次查询消息的条数上线(最多100条)
}

func (x *CIMGetMsgListReq) Reset() {
	*x = CIMGetMsgListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMGetMsgListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMGetMsgListReq) ProtoMessage() {}

func (x *CIMGetMsgListReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMGetMsgListReq.ProtoReflect.Descriptor instead.
func (*CIMGetMsgListReq) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{4}
}

func (x *CIMGetMsgListReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CIMGetMsgListReq) GetSessionType() CIMSessionType {
	if x != nil {
		return x.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (x *CIMGetMsgListReq) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *CIMGetMsgListReq) GetEndMsgId() uint64 {
	if x != nil {
		return x.EndMsgId
	}
	return 0
}

func (x *CIMGetMsgListReq) GetLimitCount() uint32 {
	if x != nil {
		return x.LimitCount
	}
	return 0
}

//对于群而言，如果消息数目返回的数值小于请求的cnt,则表示群的消息能拉取的到头了，更早的消息没有权限拉取。
//如果limit_count 和 msg_list.count 不一致，说明服务器消息有缺失，需要
//客户端做一个缺失标记，避免下次再次拉取。
type CIMGetMsgListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cmd id:		0x0206
	UserId      uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionType CIMSessionType `protobuf:"varint,2,opt,name=session_type,json=sessionType,proto3,enum=coffeechat.CIMSessionType" json:"session_type,omitempty"`
	SessionId   uint64         `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	EndMsgId    uint64         `protobuf:"varint,4,opt,name=end_msg_id,json=endMsgId,proto3" json:"end_msg_id,omitempty"` // 结束服务器消息id(不包含在查询结果中)
	//   uint64 from_time_stamp = 4;     // 起始时间点，单位：毫秒
	//   uint64 end_time_stamp = 5;      // 结束时间点，单位：毫秒
	MsgList []*CIMMsgInfo `protobuf:"bytes,6,rep,name=msg_list,json=msgList,proto3" json:"msg_list,omitempty"` // 消息列表
}

func (x *CIMGetMsgListRsp) Reset() {
	*x = CIMGetMsgListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMGetMsgListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMGetMsgListRsp) ProtoMessage() {}

func (x *CIMGetMsgListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMGetMsgListRsp.ProtoReflect.Descriptor instead.
func (*CIMGetMsgListRsp) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{5}
}

func (x *CIMGetMsgListRsp) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CIMGetMsgListRsp) GetSessionType() CIMSessionType {
	if x != nil {
		return x.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (x *CIMGetMsgListRsp) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *CIMGetMsgListRsp) GetEndMsgId() uint64 {
	if x != nil {
		return x.EndMsgId
	}
	return 0
}

func (x *CIMGetMsgListRsp) GetMsgList() []*CIMMsgInfo {
	if x != nil {
		return x.MsgList
	}
	return nil
}

// 发送消息
type CIMMsgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cmd id:		0x0301
	FromUserId   uint64         `protobuf:"varint,1,opt,name=from_user_id,json=fromUserId,proto3" json:"from_user_id,omitempty"`                                 // 消息发送方
	FromNickName string         `protobuf:"bytes,2,opt,name=from_nick_name,json=fromNickName,proto3" json:"from_nick_name,omitempty"`                            // 消息发送方昵称
	ToSessionId  uint64         `protobuf:"varint,3,opt,name=to_session_id,json=toSessionId,proto3" json:"to_session_id,omitempty"`                              // 消息接受方，单聊用户ID，群聊群ID
	ClientMsgId  string         `protobuf:"bytes,4,opt,name=client_msg_id,json=clientMsgId,proto3" json:"client_msg_id,omitempty"`                               // 客户端消息ID，唯一（UUID）
	ServerMsgId  uint64         `protobuf:"varint,5,opt,name=server_msg_id,json=serverMsgId,proto3" json:"server_msg_id,omitempty"`                              // 服务端生成的消息ID，顺序（客户端发送时无需设置）
	CreateTime   int32          `protobuf:"varint,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`                                   // 消息创建时间戳(秒)
	MsgType      CIMMsgType     `protobuf:"varint,7,opt,name=msg_type,json=msgType,proto3,enum=coffeechat.CIMMsgType" json:"msg_type,omitempty"`                 // 消息类型
	SessionType  CIMSessionType `protobuf:"varint,8,opt,name=session_type,json=sessionType,proto3,enum=coffeechat.CIMSessionType" json:"session_type,omitempty"` // 会话类型
	MsgData      []byte         `protobuf:"bytes,9,opt,name=msg_data,json=msgData,proto3" json:"msg_data,omitempty"`                                             // 消息内容
}

func (x *CIMMsgData) Reset() {
	*x = CIMMsgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMMsgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMMsgData) ProtoMessage() {}

func (x *CIMMsgData) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMMsgData.ProtoReflect.Descriptor instead.
func (*CIMMsgData) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{6}
}

func (x *CIMMsgData) GetFromUserId() uint64 {
	if x != nil {
		return x.FromUserId
	}
	return 0
}

func (x *CIMMsgData) GetFromNickName() string {
	if x != nil {
		return x.FromNickName
	}
	return ""
}

func (x *CIMMsgData) GetToSessionId() uint64 {
	if x != nil {
		return x.ToSessionId
	}
	return 0
}

func (x *CIMMsgData) GetClientMsgId() string {
	if x != nil {
		return x.ClientMsgId
	}
	return ""
}

func (x *CIMMsgData) GetServerMsgId() uint64 {
	if x != nil {
		return x.ServerMsgId
	}
	return 0
}

func (x *CIMMsgData) GetCreateTime() int32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *CIMMsgData) GetMsgType() CIMMsgType {
	if x != nil {
		return x.MsgType
	}
	return CIMMsgType_kCIM_MSG_TYPE_UNKNOWN
}

func (x *CIMMsgData) GetSessionType() CIMSessionType {
	if x != nil {
		return x.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (x *CIMMsgData) GetMsgData() []byte {
	if x != nil {
		return x.MsgData
	}
	return nil
}

// 消息收到回复
type CIMMsgDataAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cmd id:		0x0302
	FromUserId  uint64         `protobuf:"varint,1,opt,name=from_user_id,json=fromUserId,proto3" json:"from_user_id,omitempty"`                                 // 消息发送方
	ToSessionId uint64         `protobuf:"varint,2,opt,name=to_session_id,json=toSessionId,proto3" json:"to_session_id,omitempty"`                              // 消息接受方，单聊用户ID，群聊群ID
	ClientMsgId string         `protobuf:"bytes,3,opt,name=client_msg_id,json=clientMsgId,proto3" json:"client_msg_id,omitempty"`                               // 客户端消息ID，唯一（UUID）
	ServerMsgId uint64         `protobuf:"varint,4,opt,name=server_msg_id,json=serverMsgId,proto3" json:"server_msg_id,omitempty"`                              // 服务端生成的消息ID，顺序
	ResCode     CIMResCode     `protobuf:"varint,5,opt,name=res_code,json=resCode,proto3,enum=coffeechat.CIMResCode" json:"res_code,omitempty"`                 // 错误码
	SessionType CIMSessionType `protobuf:"varint,6,opt,name=session_type,json=sessionType,proto3,enum=coffeechat.CIMSessionType" json:"session_type,omitempty"` // 会话类型
	//optional
	CreateTime int32 `protobuf:"varint,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"` // 创建时间戳(毫秒)
}

func (x *CIMMsgDataAck) Reset() {
	*x = CIMMsgDataAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMMsgDataAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMMsgDataAck) ProtoMessage() {}

func (x *CIMMsgDataAck) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMMsgDataAck.ProtoReflect.Descriptor instead.
func (*CIMMsgDataAck) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{7}
}

func (x *CIMMsgDataAck) GetFromUserId() uint64 {
	if x != nil {
		return x.FromUserId
	}
	return 0
}

func (x *CIMMsgDataAck) GetToSessionId() uint64 {
	if x != nil {
		return x.ToSessionId
	}
	return 0
}

func (x *CIMMsgDataAck) GetClientMsgId() string {
	if x != nil {
		return x.ClientMsgId
	}
	return ""
}

func (x *CIMMsgDataAck) GetServerMsgId() uint64 {
	if x != nil {
		return x.ServerMsgId
	}
	return 0
}

func (x *CIMMsgDataAck) GetResCode() CIMResCode {
	if x != nil {
		return x.ResCode
	}
	return CIMResCode_kCIM_RES_CODE_UNKNOWN
}

func (x *CIMMsgDataAck) GetSessionType() CIMSessionType {
	if x != nil {
		return x.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (x *CIMMsgDataAck) GetCreateTime() int32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

var File_api_chat_chat_proto protoreflect.FileDescriptor

var file_api_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61,
	0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01,
	0x0a, 0x11, 0x43, 0x49, 0x4d, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6d,
	0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x73, 0x67,
	0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65,
	0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x49, 0x4d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x49, 0x4d, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x73, 0x70, 0x22, 0x63, 0x0a, 0x1a, 0x43, 0x49, 0x4d, 0x52, 0x65, 0x63,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x12, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x1a,
	0x43, 0x49, 0x4d, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x75, 0x6e, 0x72, 0x65,
	0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x53, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x49, 0x4d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xc8, 0x01,
	0x0a, 0x10, 0x43, 0x49, 0x4d, 0x47, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43,
	0x49, 0x4d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x65, 0x6e, 0x64,
	0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x65,
	0x6e, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xda, 0x01, 0x0a, 0x10, 0x43, 0x49, 0x4d,
	0x47, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63,
	0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x49, 0x4d, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67,
	0x49, 0x64, 0x12, 0x31, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x43, 0x49, 0x4d, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x6d, 0x73,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xee, 0x02, 0x0a, 0x0a, 0x43, 0x49, 0x4d, 0x4d, 0x73, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6e,
	0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x66, 0x72, 0x6f, 0x6d, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d,
	0x74, 0x6f, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x6f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d,
	0x73, 0x67, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6d,
	0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x6d, 0x73, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f,
	0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x49, 0x4d, 0x4d, 0x73, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x0c,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x49, 0x4d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x73, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d,
	0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x22, 0xb0, 0x02, 0x0a, 0x0d, 0x43, 0x49, 0x4d, 0x4d, 0x73,
	0x67, 0x44, 0x61, 0x74, 0x61, 0x41, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x6f,
	0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x74, 0x6f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65,
	0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x49, 0x4d, 0x52, 0x65, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x49, 0x4d,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xa0, 0x03, 0x0a, 0x04, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x7e, 0x0a, 0x14, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x66,
	0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x49, 0x4d, 0x52, 0x65, 0x63, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x49, 0x4d, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x10, 0x12, 0x0e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x60, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x49,
	0x4d, 0x47, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c,
	0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x49, 0x4d, 0x47,
	0x65, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x22, 0x16, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x2e, 0x63,
	0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x49, 0x4d, 0x4d, 0x73, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x43, 0x49, 0x4d, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x41, 0x63, 0x6b, 0x22,
	0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x73,
	0x65, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x64, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x61,
	0x64, 0x41, 0x63, 0x6b, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x43, 0x49, 0x4d, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x43, 0x49, 0x4d, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x73, 0x70, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x63, 0x68, 0x61,
	0x74, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x73, 0x67, 0x3a, 0x01, 0x2a, 0x42, 0x33, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5a, 0x18, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x3b, 0x63, 0x68, 0x61,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_chat_chat_proto_rawDescOnce sync.Once
	file_api_chat_chat_proto_rawDescData = file_api_chat_chat_proto_rawDesc
)

func file_api_chat_chat_proto_rawDescGZIP() []byte {
	file_api_chat_chat_proto_rawDescOnce.Do(func() {
		file_api_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_chat_chat_proto_rawDescData)
	})
	return file_api_chat_chat_proto_rawDescData
}

var file_api_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_chat_chat_proto_goTypes = []interface{}{
	(*CIMMsgDataReadReq)(nil),          // 0: coffeechat.CIMMsgDataReadReq
	(*CIMMsgDataReadRsp)(nil),          // 1: coffeechat.CIMMsgDataReadRsp
	(*CIMRecentContactSessionReq)(nil), // 2: coffeechat.CIMRecentContactSessionReq
	(*CIMRecentContactSessionRsp)(nil), // 3: coffeechat.CIMRecentContactSessionRsp
	(*CIMGetMsgListReq)(nil),           // 4: coffeechat.CIMGetMsgListReq
	(*CIMGetMsgListRsp)(nil),           // 5: coffeechat.CIMGetMsgListRsp
	(*CIMMsgData)(nil),                 // 6: coffeechat.CIMMsgData
	(*CIMMsgDataAck)(nil),              // 7: coffeechat.CIMMsgDataAck
	(CIMSessionType)(0),                // 8: coffeechat.CIMSessionType
	(*CIMContactSessionInfo)(nil),      // 9: coffeechat.CIMContactSessionInfo
	(*CIMMsgInfo)(nil),                 // 10: coffeechat.CIMMsgInfo
	(CIMMsgType)(0),                    // 11: coffeechat.CIMMsgType
	(CIMResCode)(0),                    // 12: coffeechat.CIMResCode
}
var file_api_chat_chat_proto_depIdxs = []int32{
	8,  // 0: coffeechat.CIMMsgDataReadReq.session_type:type_name -> coffeechat.CIMSessionType
	9,  // 1: coffeechat.CIMRecentContactSessionRsp.contact_session_list:type_name -> coffeechat.CIMContactSessionInfo
	8,  // 2: coffeechat.CIMGetMsgListReq.session_type:type_name -> coffeechat.CIMSessionType
	8,  // 3: coffeechat.CIMGetMsgListRsp.session_type:type_name -> coffeechat.CIMSessionType
	10, // 4: coffeechat.CIMGetMsgListRsp.msg_list:type_name -> coffeechat.CIMMsgInfo
	11, // 5: coffeechat.CIMMsgData.msg_type:type_name -> coffeechat.CIMMsgType
	8,  // 6: coffeechat.CIMMsgData.session_type:type_name -> coffeechat.CIMSessionType
	12, // 7: coffeechat.CIMMsgDataAck.res_code:type_name -> coffeechat.CIMResCode
	8,  // 8: coffeechat.CIMMsgDataAck.session_type:type_name -> coffeechat.CIMSessionType
	2,  // 9: coffeechat.Chat.RecentContactSession:input_type -> coffeechat.CIMRecentContactSessionReq
	4,  // 10: coffeechat.Chat.GetMsgList:input_type -> coffeechat.CIMGetMsgListReq
	6,  // 11: coffeechat.Chat.Send:input_type -> coffeechat.CIMMsgData
	0,  // 12: coffeechat.Chat.MsgReadAck:input_type -> coffeechat.CIMMsgDataReadReq
	3,  // 13: coffeechat.Chat.RecentContactSession:output_type -> coffeechat.CIMRecentContactSessionRsp
	5,  // 14: coffeechat.Chat.GetMsgList:output_type -> coffeechat.CIMGetMsgListRsp
	7,  // 15: coffeechat.Chat.Send:output_type -> coffeechat.CIMMsgDataAck
	1,  // 16: coffeechat.Chat.MsgReadAck:output_type -> coffeechat.CIMMsgDataReadRsp
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_chat_chat_proto_init() }
func file_api_chat_chat_proto_init() {
	if File_api_chat_chat_proto != nil {
		return
	}
	file_api_chat_message_proto_init()
	file_api_chat_contants_proto_init()
	file_api_chat_error_reason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMMsgDataReadReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMMsgDataReadRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMRecentContactSessionReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_chat_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMRecentContactSessionRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_chat_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMGetMsgListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_chat_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMGetMsgListRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_chat_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMMsgData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_chat_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMMsgDataAck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_chat_chat_proto_goTypes,
		DependencyIndexes: file_api_chat_chat_proto_depIdxs,
		MessageInfos:      file_api_chat_chat_proto_msgTypes,
	}.Build()
	File_api_chat_chat_proto = out.File
	file_api_chat_chat_proto_rawDesc = nil
	file_api_chat_chat_proto_goTypes = nil
	file_api_chat_chat_proto_depIdxs = nil
}
