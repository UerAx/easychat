/*
 * @Author: ww
 * @Date: 2022-07-12 03:23:23
 * @Description:
 * @FilePath: /easy-chat/model/event.go
 */
package model

import "time"

// 用户在聊天室中的唯一ID
type uid = string

const (
	EventTypeMsg    = "event-msg"    // 用户发言
	EventTypeSystem = "event-system" // 系统信息推送 如房间人数
	EventTypeJoin   = "event-join"   // 用户加入
	EventTypeTyping = "event-typing" // 用户正在输入
	EventTypeLeave  = "event-leave"  // 用户离开
	EventTypeImage  = "event-image"  // todo 消息图片
)

// 聊天室事件定义
type Event struct {
	Type      string `json:"type"`      // 事件类型
	User      string `json:"user"`      // 用户名
	Timestamp int64  `json:"timestamp"` // 时间戳
	Text      string `json:"text"`      // 事件内容
	Count     int    `json:"Count"`     // 房间用户数
}

func NewEvent(typ string, user, msg string) Event {
	return Event{
		Type:      typ,
		User:      user,
		Timestamp: time.Now().UnixNano() / 1e6,
		Text:      msg,
	}
}