// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CIM.Def.proto

package grpc

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

// 外部消息：客户端 <-> 服务器消息定义
type CIMCmdID int32

const (
	CIMCmdID_kCIM_CID_UNKNOWN                         CIMCmdID = 0
	CIMCmdID_kCIM_CID_LOGIN_AUTH_TOKEN_REQ            CIMCmdID = 257
	CIMCmdID_kCIM_CID_LOGIN_AUTH_TOKEN_RSP            CIMCmdID = 258
	CIMCmdID_kCIM_CID_LOGIN_AUTH_LOGOUT_REQ           CIMCmdID = 259
	CIMCmdID_kCIM_CID_LOGIN_AUTH_LOGOUT_RSP           CIMCmdID = 260
	CIMCmdID_kCIM_CID_LOGIN_HEARTBEAT                 CIMCmdID = 261
	CIMCmdID_kCIM_CID_LIST_RECENT_CONTACT_SESSION_REQ CIMCmdID = 513
	CIMCmdID_kCIM_CID_LIST_RECENT_CONTACT_SESSION_RSP CIMCmdID = 514
	// kCIM_CID_LIST_UNREAD_CNT_REQ = 0x0203; // 未读消息计数列表请求
	// kCIM_CID_LIST_UNREAD_CNT_RSP = 0x0204;
	CIMCmdID_kCIM_CID_LIST_MSG_REQ              CIMCmdID = 517
	CIMCmdID_kCIM_CID_LIST_MSG_RSP              CIMCmdID = 518
	CIMCmdID_kCIM_CID_MSG_DATA                  CIMCmdID = 769
	CIMCmdID_kCIM_CID_MSG_DATA_ACK              CIMCmdID = 770
	CIMCmdID_kCIM_CID_MSG_READ_ACK              CIMCmdID = 771
	CIMCmdID_kCIM_CID_MSG_READ_NOTIFY           CIMCmdID = 772
	CIMCmdID_kCIM_CID_MSG_GET_LATEST_MSG_ID_REQ CIMCmdID = 773
	CIMCmdID_kCIM_CID_MSG_GET_LATEST_MSG_ID_RSP CIMCmdID = 774
	CIMCmdID_kCIM_CID_MSG_GET_BY_MSG_ID_REQ     CIMCmdID = 775
	CIMCmdID_kCIM_CID_MSG_GET_BY_MSG_ID_RSP     CIMCmdID = 776
)

var CIMCmdID_name = map[int32]string{
	0:   "kCIM_CID_UNKNOWN",
	257: "kCIM_CID_LOGIN_AUTH_TOKEN_REQ",
	258: "kCIM_CID_LOGIN_AUTH_TOKEN_RSP",
	259: "kCIM_CID_LOGIN_AUTH_LOGOUT_REQ",
	260: "kCIM_CID_LOGIN_AUTH_LOGOUT_RSP",
	261: "kCIM_CID_LOGIN_HEARTBEAT",
	513: "kCIM_CID_LIST_RECENT_CONTACT_SESSION_REQ",
	514: "kCIM_CID_LIST_RECENT_CONTACT_SESSION_RSP",
	517: "kCIM_CID_LIST_MSG_REQ",
	518: "kCIM_CID_LIST_MSG_RSP",
	769: "kCIM_CID_MSG_DATA",
	770: "kCIM_CID_MSG_DATA_ACK",
	771: "kCIM_CID_MSG_READ_ACK",
	772: "kCIM_CID_MSG_READ_NOTIFY",
	773: "kCIM_CID_MSG_GET_LATEST_MSG_ID_REQ",
	774: "kCIM_CID_MSG_GET_LATEST_MSG_ID_RSP",
	775: "kCIM_CID_MSG_GET_BY_MSG_ID_REQ",
	776: "kCIM_CID_MSG_GET_BY_MSG_ID_RSP",
}

var CIMCmdID_value = map[string]int32{
	"kCIM_CID_UNKNOWN":                         0,
	"kCIM_CID_LOGIN_AUTH_TOKEN_REQ":            257,
	"kCIM_CID_LOGIN_AUTH_TOKEN_RSP":            258,
	"kCIM_CID_LOGIN_AUTH_LOGOUT_REQ":           259,
	"kCIM_CID_LOGIN_AUTH_LOGOUT_RSP":           260,
	"kCIM_CID_LOGIN_HEARTBEAT":                 261,
	"kCIM_CID_LIST_RECENT_CONTACT_SESSION_REQ": 513,
	"kCIM_CID_LIST_RECENT_CONTACT_SESSION_RSP": 514,
	"kCIM_CID_LIST_MSG_REQ":                    517,
	"kCIM_CID_LIST_MSG_RSP":                    518,
	"kCIM_CID_MSG_DATA":                        769,
	"kCIM_CID_MSG_DATA_ACK":                    770,
	"kCIM_CID_MSG_READ_ACK":                    771,
	"kCIM_CID_MSG_READ_NOTIFY":                 772,
	"kCIM_CID_MSG_GET_LATEST_MSG_ID_REQ":       773,
	"kCIM_CID_MSG_GET_LATEST_MSG_ID_RSP":       774,
	"kCIM_CID_MSG_GET_BY_MSG_ID_REQ":           775,
	"kCIM_CID_MSG_GET_BY_MSG_ID_RSP":           776,
}

func (x CIMCmdID) String() string {
	return proto.EnumName(CIMCmdID_name, int32(x))
}

func (CIMCmdID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59f2b4dd9a8c1a4e, []int{0}
}

// 内部消息：服务器 <-> 服务器消息定义
type CIMIntenralCmdID int32

const (
	CIMIntenralCmdID_kCIM_SID_UNKNOWN         CIMIntenralCmdID = 0
	CIMIntenralCmdID_kCIM_SID_DB_VALIDATE_REQ CIMIntenralCmdID = 1793
	CIMIntenralCmdID_kCIM_SID_DB_VALIDATE_RSP CIMIntenralCmdID = 1794
)

var CIMIntenralCmdID_name = map[int32]string{
	0:    "kCIM_SID_UNKNOWN",
	1793: "kCIM_SID_DB_VALIDATE_REQ",
	1794: "kCIM_SID_DB_VALIDATE_RSP",
}

var CIMIntenralCmdID_value = map[string]int32{
	"kCIM_SID_UNKNOWN":         0,
	"kCIM_SID_DB_VALIDATE_REQ": 1793,
	"kCIM_SID_DB_VALIDATE_RSP": 1794,
}

func (x CIMIntenralCmdID) String() string {
	return proto.EnumName(CIMIntenralCmdID_name, int32(x))
}

func (CIMIntenralCmdID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59f2b4dd9a8c1a4e, []int{1}
}

type CIMErrorCode int32

const (
	// 通用错误码
	CIMErrorCode_kCIM_ERR_SUCCSSE CIMErrorCode = 0
	// 账号错误码
	CIMErrorCode_kCIM_ERR_LOGIN_DB_VALIDATE_FAILED CIMErrorCode = 1
	CIMErrorCode_kCIM_ERR_LOGIN_VERSION_TOO_OLD    CIMErrorCode = 2
	CIMErrorCode_kCIM_ERR_LOGIN_INVALID_USER_TOKEN CIMErrorCode = 3
)

var CIMErrorCode_name = map[int32]string{
	0: "kCIM_ERR_SUCCSSE",
	1: "kCIM_ERR_LOGIN_DB_VALIDATE_FAILED",
	2: "kCIM_ERR_LOGIN_VERSION_TOO_OLD",
	3: "kCIM_ERR_LOGIN_INVALID_USER_TOKEN",
}

var CIMErrorCode_value = map[string]int32{
	"kCIM_ERR_SUCCSSE":                  0,
	"kCIM_ERR_LOGIN_DB_VALIDATE_FAILED": 1,
	"kCIM_ERR_LOGIN_VERSION_TOO_OLD":    2,
	"kCIM_ERR_LOGIN_INVALID_USER_TOKEN": 3,
}

func (x CIMErrorCode) String() string {
	return proto.EnumName(CIMErrorCode_name, int32(x))
}

func (CIMErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59f2b4dd9a8c1a4e, []int{2}
}

// 客户端类型
type CIMClientType int32

const (
	CIMClientType_kCIM_CLIENT_TYPE_DEFAULT CIMClientType = 0
	CIMClientType_kCIM_CLIENT_TYPE_ANDROID CIMClientType = 1
	CIMClientType_kCIM_CLIENT_TYPE_IOS     CIMClientType = 2
	CIMClientType_kCIM_CLIENT_TYPE_WEB     CIMClientType = 3
)

var CIMClientType_name = map[int32]string{
	0: "kCIM_CLIENT_TYPE_DEFAULT",
	1: "kCIM_CLIENT_TYPE_ANDROID",
	2: "kCIM_CLIENT_TYPE_IOS",
	3: "kCIM_CLIENT_TYPE_WEB",
}

var CIMClientType_value = map[string]int32{
	"kCIM_CLIENT_TYPE_DEFAULT": 0,
	"kCIM_CLIENT_TYPE_ANDROID": 1,
	"kCIM_CLIENT_TYPE_IOS":     2,
	"kCIM_CLIENT_TYPE_WEB":     3,
}

func (x CIMClientType) String() string {
	return proto.EnumName(CIMClientType_name, int32(x))
}

func (CIMClientType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59f2b4dd9a8c1a4e, []int{3}
}

// 会话类型
type CIMSessionType int32

const (
	CIMSessionType_kCIM_SESSION_TYPE_Invalid CIMSessionType = 0
	CIMSessionType_kCIM_SESSION_TYPE_SINGLE  CIMSessionType = 1
	CIMSessionType_kCIM_SESSION_TYPE_GROUP   CIMSessionType = 2
)

var CIMSessionType_name = map[int32]string{
	0: "kCIM_SESSION_TYPE_Invalid",
	1: "kCIM_SESSION_TYPE_SINGLE",
	2: "kCIM_SESSION_TYPE_GROUP",
}

var CIMSessionType_value = map[string]int32{
	"kCIM_SESSION_TYPE_Invalid": 0,
	"kCIM_SESSION_TYPE_SINGLE":  1,
	"kCIM_SESSION_TYPE_GROUP":   2,
}

func (x CIMSessionType) String() string {
	return proto.EnumName(CIMSessionType_name, int32(x))
}

func (CIMSessionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59f2b4dd9a8c1a4e, []int{4}
}

// 消息类型
type CIMMsgType int32

const (
	CIMMsgType_kCIM_MSG_TYPE_UNKNOWN CIMMsgType = 0
	CIMMsgType_kCIM_MSG_TYPE_TEXT    CIMMsgType = 1
	CIMMsgType_kCIM_MSG_TYPE_IMAGE   CIMMsgType = 2
	CIMMsgType_kCIM_MSG_TYPE_Audio   CIMMsgType = 3
	// kCIM_MSG_TYPE_VIDEO = 4;        // 视频
	// kCIM_MSG_TYPE_LOCATION = 5;     // 位置
	// kCIM_MSG_TYPE_NOTIFACATION = 6; // 系统通知（包括入群出群通知等）
	// kCIM_MSG_TYPE_FILE = 7;         // 文件
	CIMMsgType_kCIM_MSG_TYPE_TIPS  CIMMsgType = 8
	CIMMsgType_kCIM_MSG_TYPE_Robot CIMMsgType = 9
)

var CIMMsgType_name = map[int32]string{
	0: "kCIM_MSG_TYPE_UNKNOWN",
	1: "kCIM_MSG_TYPE_TEXT",
	2: "kCIM_MSG_TYPE_IMAGE",
	3: "kCIM_MSG_TYPE_Audio",
	8: "kCIM_MSG_TYPE_TIPS",
	9: "kCIM_MSG_TYPE_Robot",
}

var CIMMsgType_value = map[string]int32{
	"kCIM_MSG_TYPE_UNKNOWN": 0,
	"kCIM_MSG_TYPE_TEXT":    1,
	"kCIM_MSG_TYPE_IMAGE":   2,
	"kCIM_MSG_TYPE_Audio":   3,
	"kCIM_MSG_TYPE_TIPS":    8,
	"kCIM_MSG_TYPE_Robot":   9,
}

func (x CIMMsgType) String() string {
	return proto.EnumName(CIMMsgType_name, int32(x))
}

func (CIMMsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59f2b4dd9a8c1a4e, []int{5}
}

// 消息状态
type CIMMessageStatus int32

const (
	CIMMessageStatus_kCIM_MESSAGE_STATUS_NONE    CIMMessageStatus = 0
	CIMMessageStatus_kCIM_MESSAGE_STATUS_UNREAD  CIMMessageStatus = 1
	CIMMessageStatus_kCIM_MESSAGE_STATUS_READ    CIMMessageStatus = 2
	CIMMessageStatus_kCIM_MESSAGE_STATUS_DELETED CIMMessageStatus = 3
	// kCIM_MESSAGE_STATUS_SENDING = 4;   // 发送中
	CIMMessageStatus_kCIM_MESSAGE_STATUS_SENT CIMMessageStatus = 5
)

var CIMMessageStatus_name = map[int32]string{
	0: "kCIM_MESSAGE_STATUS_NONE",
	1: "kCIM_MESSAGE_STATUS_UNREAD",
	2: "kCIM_MESSAGE_STATUS_READ",
	3: "kCIM_MESSAGE_STATUS_DELETED",
	5: "kCIM_MESSAGE_STATUS_SENT",
}

var CIMMessageStatus_value = map[string]int32{
	"kCIM_MESSAGE_STATUS_NONE":    0,
	"kCIM_MESSAGE_STATUS_UNREAD":  1,
	"kCIM_MESSAGE_STATUS_READ":    2,
	"kCIM_MESSAGE_STATUS_DELETED": 3,
	"kCIM_MESSAGE_STATUS_SENT":    5,
}

func (x CIMMessageStatus) String() string {
	return proto.EnumName(CIMMessageStatus_name, int32(x))
}

func (CIMMessageStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59f2b4dd9a8c1a4e, []int{6}
}

// 会话状态
type CIMSessionStatusType int32

const (
	CIMSessionStatusType_kCIM_SESSION_STATUS_OK CIMSessionStatusType = 0
)

var CIMSessionStatusType_name = map[int32]string{
	0: "kCIM_SESSION_STATUS_OK",
}

var CIMSessionStatusType_value = map[string]int32{
	"kCIM_SESSION_STATUS_OK": 0,
}

func (x CIMSessionStatusType) String() string {
	return proto.EnumName(CIMSessionStatusType_name, int32(x))
}

func (CIMSessionStatusType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59f2b4dd9a8c1a4e, []int{7}
}

// 消息属性
type CIMMessageFeature int32

const (
	CIMMessageFeature_kCIM_MESSAGE_FEATURE_DEFAULT CIMMessageFeature = 0
	// kCIM_MESSAGE_FEATURE_LEAVE_MSG = 1;      // 离线消息(和漫游消息有何区别)
	CIMMessageFeature_kCIM_MESSAGE_FEATURE_ROAM_MSG CIMMessageFeature = 2
)

var CIMMessageFeature_name = map[int32]string{
	0: "kCIM_MESSAGE_FEATURE_DEFAULT",
	2: "kCIM_MESSAGE_FEATURE_ROAM_MSG",
}

var CIMMessageFeature_value = map[string]int32{
	"kCIM_MESSAGE_FEATURE_DEFAULT":  0,
	"kCIM_MESSAGE_FEATURE_ROAM_MSG": 2,
}

func (x CIMMessageFeature) String() string {
	return proto.EnumName(CIMMessageFeature_name, int32(x))
}

func (CIMMessageFeature) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59f2b4dd9a8c1a4e, []int{8}
}

// 消息错误码
type CIMResCode int32

const (
	CIMResCode_kCIM_RES_CODE_OK CIMResCode = 0
)

var CIMResCode_name = map[int32]string{
	0: "kCIM_RES_CODE_OK",
}

var CIMResCode_value = map[string]int32{
	"kCIM_RES_CODE_OK": 0,
}

func (x CIMResCode) String() string {
	return proto.EnumName(CIMResCode_name, int32(x))
}

func (CIMResCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59f2b4dd9a8c1a4e, []int{9}
}

// 用户信息
type CIMUserInfo struct {
	UserId   uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NickName string `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	//optional
	AttachInfo           string   `protobuf:"bytes,11,opt,name=attach_info,json=attachInfo,proto3" json:"attach_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CIMUserInfo) Reset()         { *m = CIMUserInfo{} }
func (m *CIMUserInfo) String() string { return proto.CompactTextString(m) }
func (*CIMUserInfo) ProtoMessage()    {}
func (*CIMUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_59f2b4dd9a8c1a4e, []int{0}
}

func (m *CIMUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMUserInfo.Unmarshal(m, b)
}
func (m *CIMUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMUserInfo.Marshal(b, m, deterministic)
}
func (m *CIMUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMUserInfo.Merge(m, src)
}
func (m *CIMUserInfo) XXX_Size() int {
	return xxx_messageInfo_CIMUserInfo.Size(m)
}
func (m *CIMUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CIMUserInfo proto.InternalMessageInfo

func (m *CIMUserInfo) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMUserInfo) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *CIMUserInfo) GetAttachInfo() string {
	if m != nil {
		return m.AttachInfo
	}
	return ""
}

// 会话信息
type CIMContactSessionInfo struct {
	SessionId     uint64               `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	SessionType   CIMSessionType       `protobuf:"varint,2,opt,name=session_type,json=sessionType,proto3,enum=CIM.Def.CIMSessionType" json:"session_type,omitempty"`
	SessionStatus CIMSessionStatusType `protobuf:"varint,3,opt,name=session_status,json=sessionStatus,proto3,enum=CIM.Def.CIMSessionStatusType" json:"session_status,omitempty"`
	UnreadCnt     uint32               `protobuf:"varint,4,opt,name=unread_cnt,json=unreadCnt,proto3" json:"unread_cnt,omitempty"`
	UpdatedTime   uint32               `protobuf:"varint,5,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	MsgId         string               `protobuf:"bytes,6,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	MsgTimeStamp  uint64               `protobuf:"varint,7,opt,name=msg_time_stamp,json=msgTimeStamp,proto3" json:"msg_time_stamp,omitempty"`
	MsgData       []byte               `protobuf:"bytes,8,opt,name=msg_data,json=msgData,proto3" json:"msg_data,omitempty"`
	MsgType       CIMMsgType           `protobuf:"varint,9,opt,name=msg_type,json=msgType,proto3,enum=CIM.Def.CIMMsgType" json:"msg_type,omitempty"`
	MsgFromUserId uint64               `protobuf:"varint,10,opt,name=msg_from_user_id,json=msgFromUserId,proto3" json:"msg_from_user_id,omitempty"`
	MsgStatus     CIMMessageStatus     `protobuf:"varint,11,opt,name=msg_status,json=msgStatus,proto3,enum=CIM.Def.CIMMessageStatus" json:"msg_status,omitempty"`
	//optional
	MsgAttach string `protobuf:"bytes,12,opt,name=msg_attach,json=msgAttach,proto3" json:"msg_attach,omitempty"`
	//optional
	ExtendData string `protobuf:"bytes,13,opt,name=extend_data,json=extendData,proto3" json:"extend_data,omitempty"`
	//optional
	IsRobotSession       bool     `protobuf:"varint,14,opt,name=is_robot_session,json=isRobotSession,proto3" json:"is_robot_session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CIMContactSessionInfo) Reset()         { *m = CIMContactSessionInfo{} }
func (m *CIMContactSessionInfo) String() string { return proto.CompactTextString(m) }
func (*CIMContactSessionInfo) ProtoMessage()    {}
func (*CIMContactSessionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_59f2b4dd9a8c1a4e, []int{1}
}

func (m *CIMContactSessionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMContactSessionInfo.Unmarshal(m, b)
}
func (m *CIMContactSessionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMContactSessionInfo.Marshal(b, m, deterministic)
}
func (m *CIMContactSessionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMContactSessionInfo.Merge(m, src)
}
func (m *CIMContactSessionInfo) XXX_Size() int {
	return xxx_messageInfo_CIMContactSessionInfo.Size(m)
}
func (m *CIMContactSessionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMContactSessionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CIMContactSessionInfo proto.InternalMessageInfo

func (m *CIMContactSessionInfo) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *CIMContactSessionInfo) GetSessionType() CIMSessionType {
	if m != nil {
		return m.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (m *CIMContactSessionInfo) GetSessionStatus() CIMSessionStatusType {
	if m != nil {
		return m.SessionStatus
	}
	return CIMSessionStatusType_kCIM_SESSION_STATUS_OK
}

func (m *CIMContactSessionInfo) GetUnreadCnt() uint32 {
	if m != nil {
		return m.UnreadCnt
	}
	return 0
}

func (m *CIMContactSessionInfo) GetUpdatedTime() uint32 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

func (m *CIMContactSessionInfo) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *CIMContactSessionInfo) GetMsgTimeStamp() uint64 {
	if m != nil {
		return m.MsgTimeStamp
	}
	return 0
}

func (m *CIMContactSessionInfo) GetMsgData() []byte {
	if m != nil {
		return m.MsgData
	}
	return nil
}

func (m *CIMContactSessionInfo) GetMsgType() CIMMsgType {
	if m != nil {
		return m.MsgType
	}
	return CIMMsgType_kCIM_MSG_TYPE_UNKNOWN
}

func (m *CIMContactSessionInfo) GetMsgFromUserId() uint64 {
	if m != nil {
		return m.MsgFromUserId
	}
	return 0
}

func (m *CIMContactSessionInfo) GetMsgStatus() CIMMessageStatus {
	if m != nil {
		return m.MsgStatus
	}
	return CIMMessageStatus_kCIM_MESSAGE_STATUS_NONE
}

func (m *CIMContactSessionInfo) GetMsgAttach() string {
	if m != nil {
		return m.MsgAttach
	}
	return ""
}

func (m *CIMContactSessionInfo) GetExtendData() string {
	if m != nil {
		return m.ExtendData
	}
	return ""
}

func (m *CIMContactSessionInfo) GetIsRobotSession() bool {
	if m != nil {
		return m.IsRobotSession
	}
	return false
}

// 消息信息
type CIMMsgInfo struct {
	ClientMsgId string            `protobuf:"bytes,1,opt,name=client_msg_id,json=clientMsgId,proto3" json:"client_msg_id,omitempty"`
	ServerMsgId uint64            `protobuf:"varint,2,opt,name=server_msg_id,json=serverMsgId,proto3" json:"server_msg_id,omitempty"`
	MsgResCode  CIMResCode        `protobuf:"varint,3,opt,name=msg_res_code,json=msgResCode,proto3,enum=CIM.Def.CIMResCode" json:"msg_res_code,omitempty"`
	MsgFeature  CIMMessageFeature `protobuf:"varint,4,opt,name=msg_feature,json=msgFeature,proto3,enum=CIM.Def.CIMMessageFeature" json:"msg_feature,omitempty"`
	SessionType CIMSessionType    `protobuf:"varint,5,opt,name=session_type,json=sessionType,proto3,enum=CIM.Def.CIMSessionType" json:"session_type,omitempty"`
	FromUserId  uint64            `protobuf:"varint,6,opt,name=from_user_id,json=fromUserId,proto3" json:"from_user_id,omitempty"`
	ToSessionId uint64            `protobuf:"varint,7,opt,name=to_session_id,json=toSessionId,proto3" json:"to_session_id,omitempty"`
	CreateTime  uint64            `protobuf:"varint,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	MsgType     CIMMsgType        `protobuf:"varint,9,opt,name=msg_type,json=msgType,proto3,enum=CIM.Def.CIMMsgType" json:"msg_type,omitempty"`
	MsgStatus   CIMMessageStatus  `protobuf:"varint,10,opt,name=msg_status,json=msgStatus,proto3,enum=CIM.Def.CIMMessageStatus" json:"msg_status,omitempty"`
	MsgData     []byte            `protobuf:"bytes,11,opt,name=msg_data,json=msgData,proto3" json:"msg_data,omitempty"`
	//optional
	Attach               string        `protobuf:"bytes,12,opt,name=attach,proto3" json:"attach,omitempty"`
	SenderClientType     CIMClientType `protobuf:"varint,13,opt,name=sender_client_type,json=senderClientType,proto3,enum=CIM.Def.CIMClientType" json:"sender_client_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CIMMsgInfo) Reset()         { *m = CIMMsgInfo{} }
func (m *CIMMsgInfo) String() string { return proto.CompactTextString(m) }
func (*CIMMsgInfo) ProtoMessage()    {}
func (*CIMMsgInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_59f2b4dd9a8c1a4e, []int{2}
}

func (m *CIMMsgInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMMsgInfo.Unmarshal(m, b)
}
func (m *CIMMsgInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMMsgInfo.Marshal(b, m, deterministic)
}
func (m *CIMMsgInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMMsgInfo.Merge(m, src)
}
func (m *CIMMsgInfo) XXX_Size() int {
	return xxx_messageInfo_CIMMsgInfo.Size(m)
}
func (m *CIMMsgInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMMsgInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CIMMsgInfo proto.InternalMessageInfo

func (m *CIMMsgInfo) GetClientMsgId() string {
	if m != nil {
		return m.ClientMsgId
	}
	return ""
}

func (m *CIMMsgInfo) GetServerMsgId() uint64 {
	if m != nil {
		return m.ServerMsgId
	}
	return 0
}

func (m *CIMMsgInfo) GetMsgResCode() CIMResCode {
	if m != nil {
		return m.MsgResCode
	}
	return CIMResCode_kCIM_RES_CODE_OK
}

func (m *CIMMsgInfo) GetMsgFeature() CIMMessageFeature {
	if m != nil {
		return m.MsgFeature
	}
	return CIMMessageFeature_kCIM_MESSAGE_FEATURE_DEFAULT
}

func (m *CIMMsgInfo) GetSessionType() CIMSessionType {
	if m != nil {
		return m.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (m *CIMMsgInfo) GetFromUserId() uint64 {
	if m != nil {
		return m.FromUserId
	}
	return 0
}

func (m *CIMMsgInfo) GetToSessionId() uint64 {
	if m != nil {
		return m.ToSessionId
	}
	return 0
}

func (m *CIMMsgInfo) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *CIMMsgInfo) GetMsgType() CIMMsgType {
	if m != nil {
		return m.MsgType
	}
	return CIMMsgType_kCIM_MSG_TYPE_UNKNOWN
}

func (m *CIMMsgInfo) GetMsgStatus() CIMMessageStatus {
	if m != nil {
		return m.MsgStatus
	}
	return CIMMessageStatus_kCIM_MESSAGE_STATUS_NONE
}

func (m *CIMMsgInfo) GetMsgData() []byte {
	if m != nil {
		return m.MsgData
	}
	return nil
}

func (m *CIMMsgInfo) GetAttach() string {
	if m != nil {
		return m.Attach
	}
	return ""
}

func (m *CIMMsgInfo) GetSenderClientType() CIMClientType {
	if m != nil {
		return m.SenderClientType
	}
	return CIMClientType_kCIM_CLIENT_TYPE_DEFAULT
}

func init() {
	proto.RegisterEnum("CIM.Def.CIMCmdID", CIMCmdID_name, CIMCmdID_value)
	proto.RegisterEnum("CIM.Def.CIMIntenralCmdID", CIMIntenralCmdID_name, CIMIntenralCmdID_value)
	proto.RegisterEnum("CIM.Def.CIMErrorCode", CIMErrorCode_name, CIMErrorCode_value)
	proto.RegisterEnum("CIM.Def.CIMClientType", CIMClientType_name, CIMClientType_value)
	proto.RegisterEnum("CIM.Def.CIMSessionType", CIMSessionType_name, CIMSessionType_value)
	proto.RegisterEnum("CIM.Def.CIMMsgType", CIMMsgType_name, CIMMsgType_value)
	proto.RegisterEnum("CIM.Def.CIMMessageStatus", CIMMessageStatus_name, CIMMessageStatus_value)
	proto.RegisterEnum("CIM.Def.CIMSessionStatusType", CIMSessionStatusType_name, CIMSessionStatusType_value)
	proto.RegisterEnum("CIM.Def.CIMMessageFeature", CIMMessageFeature_name, CIMMessageFeature_value)
	proto.RegisterEnum("CIM.Def.CIMResCode", CIMResCode_name, CIMResCode_value)
	proto.RegisterType((*CIMUserInfo)(nil), "CIM.Def.CIMUserInfo")
	proto.RegisterType((*CIMContactSessionInfo)(nil), "CIM.Def.CIMContactSessionInfo")
	proto.RegisterType((*CIMMsgInfo)(nil), "CIM.Def.CIMMsgInfo")
}

func init() { proto.RegisterFile("CIM.Def.proto", fileDescriptor_59f2b4dd9a8c1a4e) }

var fileDescriptor_59f2b4dd9a8c1a4e = []byte{
	// 1285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x36, 0x65, 0x59, 0xb6, 0x47, 0x96, 0xb0, 0xd9, 0x24, 0xb6, 0xec, 0xc4, 0x89, 0xad, 0xb6,
	0x88, 0x20, 0xa0, 0x3e, 0xa4, 0x28, 0x50, 0xb4, 0x27, 0x9a, 0x5c, 0x2b, 0x0b, 0x4b, 0xa4, 0x4a,
	0x52, 0xf9, 0xe9, 0x65, 0xc1, 0x88, 0x2b, 0x45, 0x4d, 0x48, 0x1a, 0x24, 0x15, 0x34, 0xa7, 0x22,
	0x7f, 0x6d, 0x9f, 0xa0, 0x0f, 0xd0, 0x7b, 0x5f, 0xa5, 0x7d, 0x97, 0x3e, 0x41, 0xb1, 0xbb, 0x94,
	0x25, 0x46, 0x32, 0x92, 0xdc, 0xc4, 0xef, 0xfb, 0x66, 0x66, 0x67, 0xe6, 0x5b, 0x8a, 0x50, 0x33,
	0x68, 0xef, 0xc4, 0xe4, 0xa3, 0x93, 0x8b, 0x24, 0xce, 0x62, 0xbc, 0x99, 0x3f, 0x36, 0x03, 0xa8,
	0x1a, 0xb4, 0x37, 0x48, 0x79, 0x42, 0xa3, 0x51, 0x8c, 0xf7, 0x60, 0x73, 0x9a, 0xf2, 0x84, 0x4d,
	0x82, 0x86, 0x76, 0xa4, 0xb5, 0xca, 0x4e, 0x45, 0x3c, 0xd2, 0x00, 0xdf, 0x82, 0xed, 0x68, 0x32,
	0x7c, 0xce, 0x22, 0x3f, 0xe4, 0x8d, 0xd2, 0x91, 0xd6, 0xda, 0x76, 0xb6, 0x04, 0x60, 0xf9, 0x21,
	0xc7, 0x77, 0xa1, 0xea, 0x67, 0x99, 0x3f, 0x7c, 0xc6, 0x26, 0xd1, 0x28, 0x6e, 0x54, 0x25, 0x0d,
	0x0a, 0x12, 0x69, 0x9b, 0xff, 0x96, 0xe1, 0xa6, 0x41, 0x7b, 0x46, 0x1c, 0x65, 0xfe, 0x30, 0x73,
	0x79, 0x9a, 0x4e, 0xe2, 0x48, 0x16, 0x3c, 0x04, 0x48, 0xd5, 0xe3, 0xbc, 0xe6, 0x76, 0x8e, 0xd0,
	0x00, 0x7f, 0x0f, 0x3b, 0x33, 0x3a, 0x7b, 0x75, 0xa1, 0x2a, 0xd7, 0xef, 0xef, 0x9d, 0xcc, 0xba,
	0x31, 0x68, 0x2f, 0xcf, 0xe6, 0xbd, 0xba, 0xe0, 0x4e, 0x35, 0x9d, 0x3f, 0x60, 0x13, 0xea, 0xb3,
	0xd8, 0x34, 0xf3, 0xb3, 0x69, 0xda, 0x58, 0x97, 0xd1, 0x87, 0x2b, 0xa2, 0x5d, 0x29, 0x90, 0x39,
	0x6a, 0xe9, 0x22, 0x24, 0x0e, 0x38, 0x8d, 0x12, 0xee, 0x07, 0x6c, 0x18, 0x65, 0x8d, 0xf2, 0x91,
	0xd6, 0xaa, 0x39, 0xdb, 0x0a, 0x31, 0xa2, 0x0c, 0x1f, 0xc3, 0xce, 0xf4, 0x22, 0xf0, 0x33, 0x1e,
	0xb0, 0x6c, 0x12, 0xf2, 0xc6, 0x86, 0x14, 0x54, 0x73, 0xcc, 0x9b, 0x84, 0x1c, 0xdf, 0x84, 0x4a,
	0x98, 0x8e, 0x45, 0x7b, 0x15, 0x39, 0x98, 0x8d, 0x30, 0x1d, 0xd3, 0x00, 0x7f, 0x09, 0x75, 0x01,
	0x8b, 0x28, 0x71, 0xbe, 0xf0, 0xa2, 0xb1, 0x29, 0xbb, 0xdf, 0x09, 0xd3, 0xb1, 0x88, 0x73, 0x05,
	0x86, 0xf7, 0x61, 0x4b, 0xa8, 0x02, 0x3f, 0xf3, 0x1b, 0x5b, 0x47, 0x5a, 0x6b, 0xc7, 0xd9, 0x0c,
	0xd3, 0xb1, 0xe9, 0x67, 0x3e, 0x3e, 0x51, 0x94, 0x9c, 0xcb, 0xb6, 0xec, 0xec, 0xfa, 0x62, 0x67,
	0xbd, 0x74, 0x2c, 0xfb, 0x11, 0x7a, 0x39, 0x8f, 0x7b, 0x80, 0x84, 0x7e, 0x94, 0xc4, 0x21, 0x9b,
	0x2d, 0x19, 0x64, 0xc9, 0x5a, 0x98, 0x8e, 0xcf, 0x92, 0x38, 0x1c, 0xa8, 0x5d, 0x7f, 0x07, 0x20,
	0x84, 0xf9, 0xd0, 0xaa, 0x32, 0xf5, 0x7e, 0x21, 0x35, 0x4f, 0x53, 0x7f, 0xcc, 0xd5, 0x84, 0x9c,
	0xed, 0x30, 0x1d, 0xcf, 0x87, 0x25, 0x22, 0xd5, 0xe6, 0x1b, 0x3b, 0xb2, 0x5d, 0x41, 0xeb, 0x12,
	0x10, 0x3e, 0xe1, 0xbf, 0x64, 0x3c, 0x0a, 0x54, 0x3f, 0x35, 0xe5, 0x13, 0x05, 0xc9, 0x96, 0x5a,
	0x80, 0x26, 0x29, 0x4b, 0xe2, 0xa7, 0x71, 0xc6, 0xf2, 0x35, 0x34, 0xea, 0x47, 0x5a, 0x6b, 0xcb,
	0xa9, 0x4f, 0x52, 0x47, 0xc0, 0xf9, 0xbe, 0x9a, 0xff, 0x94, 0x01, 0x54, 0x93, 0xd2, 0x46, 0x4d,
	0xa8, 0x0d, 0x5f, 0x4c, 0x78, 0x94, 0xb1, 0x7c, 0xd4, 0x9a, 0xcc, 0x5d, 0x55, 0x60, 0x4f, 0x0e,
	0xbc, 0x09, 0xb5, 0x94, 0x27, 0x2f, 0x79, 0x32, 0xd3, 0x94, 0x64, 0xf3, 0x55, 0x05, 0x2a, 0xcd,
	0xb7, 0x20, 0xc6, 0xcf, 0x12, 0x9e, 0xb2, 0x61, 0x1c, 0xf0, 0xdc, 0x31, 0x85, 0xb9, 0x3a, 0x3c,
	0x35, 0xe2, 0x80, 0x3b, 0xa2, 0xd3, 0xfc, 0x37, 0xfe, 0x01, 0xaa, 0x72, 0xb4, 0xdc, 0xcf, 0xa6,
	0x09, 0x97, 0x2e, 0xa9, 0xdf, 0x3f, 0x58, 0x31, 0xb2, 0x33, 0xa5, 0x90, 0xc1, 0xf9, 0xef, 0x25,
	0x8f, 0x6f, 0x7c, 0x86, 0xc7, 0x8f, 0x60, 0xa7, 0xb0, 0xcf, 0x8a, 0x6c, 0x09, 0x46, 0xf3, 0x65,
	0x36, 0xa1, 0x96, 0xc5, 0x6c, 0xe1, 0x8e, 0x29, 0x97, 0x55, 0xb3, 0xd8, 0xbd, 0xbc, 0x65, 0x77,
	0xa1, 0x3a, 0x4c, 0xb8, 0x9f, 0x71, 0xe5, 0xe1, 0x2d, 0x95, 0x44, 0x41, 0xd2, 0xc2, 0x9f, 0x6b,
	0xb5, 0xa2, 0x83, 0xe0, 0x33, 0x1c, 0xb4, 0xe8, 0xf7, 0x6a, 0xd1, 0xef, 0xbb, 0x50, 0x29, 0x18,
	0x2b, 0x7f, 0xc2, 0x26, 0xe0, 0x94, 0x47, 0x01, 0x4f, 0x58, 0x6e, 0x01, 0x79, 0xcc, 0x9a, 0x2c,
	0xba, 0xbb, 0x58, 0xd4, 0x90, 0xb4, 0x3c, 0x29, 0x52, 0x11, 0x73, 0xa4, 0xfd, 0x5f, 0x19, 0xb6,
	0x84, 0x26, 0x0c, 0xa8, 0x89, 0x6f, 0x00, 0x7a, 0x6e, 0xd0, 0x1e, 0x33, 0xa8, 0xc9, 0x06, 0xd6,
	0xb9, 0x65, 0x3f, 0xb2, 0xd0, 0x1a, 0x6e, 0xc2, 0xe1, 0x25, 0xda, 0xb5, 0x3b, 0xd4, 0x62, 0xfa,
	0xc0, 0x7b, 0xc0, 0x3c, 0xfb, 0x9c, 0x58, 0xcc, 0x21, 0x3f, 0xa2, 0xd7, 0xa5, 0x8f, 0x68, 0xdc,
	0x3e, 0x7a, 0x53, 0xc2, 0x5f, 0xc0, 0x9d, 0x55, 0x9a, 0xae, 0xdd, 0xb1, 0x07, 0x9e, 0x4c, 0xf4,
	0xf6, 0xa3, 0x22, 0xb7, 0x8f, 0xde, 0x95, 0xf0, 0x21, 0x34, 0x3e, 0x10, 0x3d, 0x20, 0xba, 0xe3,
	0x9d, 0x12, 0xdd, 0x43, 0xef, 0x4b, 0xf8, 0x6b, 0x68, 0xcd, 0x69, 0xea, 0x8a, 0xdc, 0x06, 0xb1,
	0x3c, 0x66, 0xd8, 0x96, 0xa7, 0x1b, 0x1e, 0x73, 0x89, 0xeb, 0x52, 0x3b, 0x3f, 0x7b, 0xf9, 0xd3,
	0xe5, 0xa2, 0x8d, 0x32, 0x3e, 0x80, 0x9b, 0x45, 0x79, 0xcf, 0xed, 0xc8, 0x54, 0xef, 0xaf, 0xe2,
	0xdc, 0x3e, 0xfa, 0xad, 0x8c, 0x77, 0xe1, 0xda, 0x25, 0x27, 0x60, 0x53, 0xf7, 0x74, 0xf4, 0xba,
	0x52, 0x88, 0x99, 0xe1, 0x4c, 0x37, 0xce, 0xd1, 0x9b, 0x65, 0xce, 0x21, 0xba, 0x29, 0xb9, 0xb7,
	0x95, 0xc2, 0x10, 0x2e, 0x39, 0xcb, 0xf6, 0xe8, 0xd9, 0x13, 0xf4, 0xae, 0x82, 0xef, 0x41, 0xb3,
	0x40, 0x77, 0x88, 0xc7, 0xba, 0xba, 0x47, 0xf2, 0x43, 0x51, 0x53, 0x9d, 0xf9, 0x93, 0x84, 0xa2,
	0x81, 0x4a, 0x61, 0x35, 0x33, 0xe1, 0xe9, 0x93, 0xc5, 0x6c, 0xbf, 0x7f, 0x54, 0xe4, 0xf6, 0xd1,
	0x1f, 0x95, 0xf6, 0x08, 0x90, 0x41, 0x7b, 0x34, 0xca, 0x78, 0x94, 0xf8, 0x2f, 0x8a, 0xde, 0x73,
	0x0b, 0xde, 0x9b, 0x35, 0x29, 0x50, 0xf3, 0x94, 0x3d, 0xd4, 0xbb, 0xd4, 0xd4, 0x3d, 0xa2, 0x56,
	0x57, 0xbf, 0x9a, 0x16, 0xab, 0xaa, 0xb7, 0xff, 0xd4, 0x60, 0xc7, 0xa0, 0x3d, 0x92, 0x24, 0x71,
	0x22, 0x5f, 0x58, 0xb3, 0x22, 0xc4, 0x71, 0x98, 0x3b, 0x30, 0x0c, 0xd7, 0x25, 0x68, 0x0d, 0x7f,
	0x05, 0xc7, 0x97, 0xa8, 0xb2, 0xd3, 0x62, 0xae, 0x33, 0x9d, 0x76, 0x89, 0x89, 0x34, 0xdc, 0xcc,
	0x5b, 0x9b, 0xcb, 0x1e, 0x12, 0x47, 0x3a, 0xc3, 0xb3, 0x6d, 0x66, 0x77, 0x4d, 0x54, 0x5a, 0x91,
	0x8a, 0x5a, 0x32, 0x13, 0x1b, 0xb8, 0xc4, 0x51, 0xf7, 0x01, 0xad, 0xb7, 0x7f, 0x95, 0x1f, 0x26,
	0xf3, 0x6b, 0x88, 0x6f, 0xcf, 0x96, 0xd9, 0xa5, 0xc2, 0x7a, 0xde, 0x93, 0x3e, 0x61, 0x26, 0x39,
	0xd3, 0x07, 0x5d, 0x0f, 0xad, 0xad, 0x64, 0x75, 0xcb, 0x74, 0x6c, 0x2a, 0xce, 0xd5, 0x80, 0x1b,
	0x4b, 0x2c, 0xb5, 0x5d, 0x54, 0x5a, 0xc9, 0x3c, 0x22, 0xa7, 0x68, 0xbd, 0xfd, 0x33, 0xd4, 0x8b,
	0xef, 0x57, 0x7c, 0x08, 0xfb, 0x6a, 0x94, 0xb9, 0xdb, 0x55, 0x9a, 0xe8, 0xa5, 0xff, 0x62, 0x12,
	0x2c, 0x1c, 0xa1, 0x40, 0xbb, 0xd4, 0xea, 0x74, 0x09, 0xd2, 0xf0, 0x2d, 0xd8, 0x5b, 0x66, 0x3b,
	0x8e, 0x3d, 0xe8, 0xa3, 0x52, 0xfb, 0x2f, 0x6d, 0xf6, 0x9f, 0x25, 0x0b, 0xed, 0xe7, 0x9e, 0x16,
	0x96, 0x90, 0xba, 0xf9, 0xb6, 0x77, 0x01, 0x17, 0x29, 0x8f, 0x3c, 0xf6, 0x90, 0x86, 0xf7, 0xe0,
	0x7a, 0x11, 0xa7, 0x3d, 0xbd, 0x43, 0x50, 0x69, 0x99, 0xd0, 0xa7, 0xc1, 0x24, 0x46, 0xeb, 0x2b,
	0x32, 0xd1, 0xbe, 0x8b, 0xb6, 0x96, 0x03, 0xe4, 0xbf, 0x2b, 0xda, 0x6e, 0xff, 0xad, 0x49, 0x4f,
	0x16, 0x5e, 0xd0, 0x97, 0x4d, 0xf7, 0x88, 0xeb, 0xea, 0x1d, 0xc2, 0x5c, 0x4f, 0xf7, 0x06, 0x2e,
	0xb3, 0x6c, 0x4b, 0xd8, 0xe6, 0x0e, 0x1c, 0xac, 0x62, 0x07, 0x96, 0xb8, 0x89, 0x48, 0xbb, 0x2a,
	0x5a, 0xb2, 0x25, 0x7c, 0x17, 0x6e, 0xad, 0x62, 0x4d, 0xd2, 0x25, 0x1e, 0x31, 0xd1, 0xfa, 0x55,
	0xe1, 0x2e, 0xb1, 0x3c, 0xb4, 0xd1, 0xbe, 0x0f, 0x37, 0x56, 0x7d, 0xc6, 0xe1, 0x03, 0xd8, 0x2d,
	0x6c, 0x22, 0x8f, 0xb2, 0xcf, 0xd1, 0x5a, 0xfb, 0x31, 0x5c, 0x5b, 0xfa, 0x4b, 0xc6, 0x47, 0x70,
	0xbb, 0x50, 0xe6, 0x8c, 0xe8, 0xde, 0xc0, 0x59, 0x74, 0xdf, 0x71, 0xfe, 0x6e, 0xff, 0x50, 0xe1,
	0xd8, 0xba, 0x1c, 0x24, 0x2a, 0xb5, 0x9b, 0x72, 0xc3, 0xb3, 0xcf, 0x82, 0xd9, 0x2d, 0x73, 0x88,
	0xcb, 0x0c, 0xdb, 0x24, 0xb2, 0xfa, 0xe9, 0x31, 0xec, 0x0d, 0xe3, 0xf0, 0x64, 0x18, 0x8f, 0x46,
	0x9c, 0x0f, 0x9f, 0xf9, 0x99, 0xfa, 0x26, 0x7f, 0x3a, 0x1d, 0x3d, 0x58, 0xff, 0xa9, 0x3c, 0x4e,
	0x2e, 0x86, 0x4f, 0x2b, 0x12, 0xf9, 0xe6, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xe4, 0x4c,
	0x95, 0xb6, 0x0b, 0x00, 0x00,
}