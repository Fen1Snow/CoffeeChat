package dao

import (
	"fmt"
	"github.com/CoffeeChat/server/src/api/cim"
	"github.com/CoffeeChat/server/src/pkg/db"
	"github.com/CoffeeChat/server/src/pkg/def"
	"github.com/CoffeeChat/server/src/pkg/logger"
	"time"
)

const kRedisKeyMsgId = "msg_id_group_"
const KRedisKeySingleMsgId = "msg_id_single"
const kIMMessageSendTableName = "im_message_send_"
const kIMMessageRecvTableName = "im_message_recv_"
const kIMMessageTableCount = 4

type Message struct {
}

var DefaultMessage = &Message{}

// if redis not connected then will failed
func (m *Message) GetMsgIdGroup(groupId uint64) (int64, error) {
	conn := DefaultRedisPool.GetMsgIdPool()
	key := fmt.Sprintf("%s_%d", kRedisKeyMsgId, groupId)
	return conn.Incr(key).Result()
}
func (m *Message) GetMsgIdSingle() (int64, error) {
	conn := DefaultRedisPool.GetMsgIdPool()
	return conn.Incr(KRedisKeySingleMsgId).Result()
}

func (m *Message) GetByMsgId(msgId string, tableName string) int64 {
	dbMaster := db.DefaultManager.GetDbMaster()
	if dbMaster != nil {
		row := dbMaster.QueryRow("select msg_id from %s where client_msg_id=%s", tableName, msgId)
		if row != nil {
			var msgId int64 = 0
			_ = row.Scan(&msgId)
			return msgId
		}
	}
	return 0
}

// 存储消息
func (m *Message) SaveMessage(fromId uint64, toId uint64, clientMsgId string, createTime int32,
	msgType cim.CIMMsgType, sessionType cim.CIMSessionType, msgData string) (uint64, error) {
	if sessionType == cim.CIMSessionType_kCIM_SESSION_TYPE_SINGLE {
		return m.saveSingleMessage(fromId, toId, clientMsgId, createTime, msgType, msgData)
	} else if sessionType == cim.CIMSessionType_kCIM_SESSION_TYPE_GROUP {
		return m.saveGroupMessage(fromId, toId, clientMsgId, createTime, msgType, msgData)
	} else {
		return 0, def.DefaultError
	}
}

func (m *Message) saveSingleMessage(fromId uint64, toId uint64, clientMsgId string, createTime int32, msgType cim.CIMMsgType, msgData string) (uint64, error) {
	dbMaster := db.DefaultManager.GetDbMaster()
	if dbMaster != nil {
		// Get MsgId
		msgId, err := m.GetMsgIdSingle()
		if err != nil {
			return 0, err
		}

		tableName := fmt.Sprintf("%s%d", kIMMessageSendTableName, fromId%kIMMessageTableCount)
		// 去重
		if existMsgId := m.GetByMsgId(clientMsgId, tableName); existMsgId > 0 {
			logger.Sugar.Info("msg already exist,msg_id=%s,from_id=%d,to_id=%d", clientMsgId, fromId, toId)
			return uint64(existMsgId), nil
		}

		// create session
		sessionModel := DefaultSession.Get(fromId, toId)
		sessionModel2 := DefaultSession.Get(toId, fromId)

		id1 := 0
		id2 := 0
		if sessionModel == nil || sessionModel2 == nil {
			id1, id2, err = DefaultSession.AddUserSession(fromId, toId, cim.CIMSessionType_kCIM_SESSION_TYPE_SINGLE,
				cim.CIMSessionStatusType_kCIM_SESSION_STATUS_OK, false)
			if err != nil {
				return 0, err
			}
		} else {
			id1 = sessionModel.Id
			id2 = sessionModel2.Id
		}
		// update latest session time
		timeStamp := time.Now().Unix()
		_ = DefaultSession.UpdateUpdated(id1, int(timeStamp))
		_ = DefaultSession.UpdateUpdated(id2, int(timeStamp))

		// save to im_message_send_x
		_, err = dbMaster.Exec("insert into %s(msg_id,client_msg_id,from_id,to_id,group_id,msg_type,msg_content,"+
			"msg_res_code,msg_feature,msg_status,created,updated) values(%d,%s,%d,%d,%d,%d,%s,%d,%d,%d,%d,%d)",
			tableName, msgId, clientMsgId, fromId, toId, 0, msgType, msgData, cim.CIMResCode_kCIM_RES_CODE_OK,
			cim.CIMMessageStatus_kCIM_MESSAGE_STATUS_NONE, timeStamp, timeStamp)
		if err != nil {
			return 0, err
		} else {
			return uint64(msgId), nil
		}
	} else {
		logger.Sugar.Error("no db connect for master")
	}
	return 0, def.DefaultError
}

// not support
func (m *Message) saveGroupMessage(fromId uint64, groupId uint64, clientMsgId string, createTime int32, msgType cim.CIMMsgType, msgData string) (uint64, error) {
	return 0, def.DefaultError
}