/*
 * @Author: ww
 * @Date: 2022-07-12 03:26:20
 * @Description:
 * @FilePath: /easy-chat/model/subscribe.go
 */
package model

// 用户订阅
type Subscribe struct {
	Id       string       // 用户在聊天室中的ID
	Username string       // 用户名
	Pipe     <-chan Event // 事件接收通道 用户从这个通道接收消息
	EmitChn  chan Event   // 用户消息推送通道
	LeaveChn chan string     // 用户离开事件推送
}

func (s *Subscribe) Leave() {
	s.LeaveChn <- s.Id // 将用户从聊天室列表中移除
}

func (s *Subscribe) Say(message string) {
	s.EmitChn <- NewEvent(EventTypeMsg, s.Username, message)
}