/*
 * @Author: ww
 * @Date: 2022-07-12 03:30:01
 * @Description:
 * @FilePath: /easy-chat/chat/room.go
 */
package chat

import (
	"container/list"

	"github.com/uerax/easychat/common"
	"github.com/uerax/easychat/model"
)

// 保存历史消息的条数
const archiveSize = 20
const chanSize = 10

const msgJoin = "[加入房间]"
const msgLeave = "[离开房间]"
const msgTyping = "[正在输入]"

// 聊天室
type Room struct {
	users       map[string]chan model.Event // 当前房间订阅者
	userCount   int                         // 当前房间总人数
	publishChn  chan model.Event            // 聊天室的消息推送入口
	archive     *list.List                  // 历史记录 todo 未持久化 重启失效
	archiveChan chan chan []model.Event     // 通过接受chan来同步聊天内容
	joinChn     chan chan model.Subscribe   // 接收订阅事件的通道 用户加入聊天室后要把历史事件推送给用户
	leaveChn    chan string                 // 用户取消订阅通道 把通道中的历史事件释放并把用户从聊天室用户列表中删除
}

func NewRoom() *Room {
	r := &Room{
		users:     map[string]chan model.Event{},
		userCount: 0,

		publishChn:  make(chan model.Event, chanSize),
		archiveChan: make(chan chan []model.Event, chanSize),
		archive:     list.New(),

		joinChn:  make(chan chan model.Subscribe, chanSize),
		leaveChn: make(chan string, chanSize),
	}

	go r.Serve()

	return r
}

// 用来向聊天室发送用户消息
// 这些接口供非websocket连接方式调用
func (r *Room) MsgJoin(user string) {
	r.publishChn <- model.NewEvent(model.EventTypeJoin, user, msgJoin)
}

func (r *Room) MsgSay(user, message string) {
	r.publishChn <- model.NewEvent(model.EventTypeMsg, user, message)
}

func (r *Room) MsgLeave(user string) {
	r.publishChn <- model.NewEvent(model.EventTypeMsg, user, msgLeave)
}

func (r *Room) Remove(id string) {
	r.leaveChn <- id // 将用户从聊天室列表中移除
}

// 用户订阅聊天室入口函数
// 返回用户订阅的对象，用户根据对象中的属性读取历史消息和即时消息
func (r *Room) Join(username string) model.Subscribe {
	resp := make(chan model.Subscribe)
	r.joinChn <- resp
	s := <-resp
	s.Username = username
	return s
}

func (r *Room) GetArchive() []model.Event {
	ch := make(chan []model.Event)
	r.archiveChan <- ch
	return <-ch
}

// 处理聊天室中的事件
func (r *Room) Serve() {
	for {
		select {
		// 用户加入房间
		case ch := <-r.joinChn:
			chn := make(chan model.Event, chanSize)
			r.userCount++
			uid := common.UUid()
			r.users[uid] = chn
			ch <- model.Subscribe{
				Id:       uid,
				Pipe:     chn,
				EmitChn:  r.publishChn,
				LeaveChn: r.leaveChn,
			}
			ev := model.NewEvent(model.EventTypeSystem, "", "")
			ev.Count = r.userCount
			for _, v := range r.users {
				v <- ev
			}
		case arch := <-r.archiveChan:
			var events []model.Event
			//历史事件
			for e := r.archive.Front(); e != nil; e = e.Next() {
				events = append(events, e.Value.(model.Event))
			}
			arch <- events
		// 有新的消息
		case event := <-r.publishChn:
			// 推送给所有用户
			event.Count = r.userCount
			for _, v := range r.users {
				v <- event
			}
			// 推送消息后，限制本地只保存指定条历史消息
			if r.archive.Len() >= archiveSize {
				r.archive.Remove(r.archive.Front())
			}
			r.archive.PushBack(event)
		// 用户退出房间
		case k := <-r.leaveChn:
			if _, ok := r.users[k]; ok {
				delete(r.users, k)
				r.userCount--
			}
			ev := model.NewEvent(model.EventTypeSystem, "", "")
			ev.Count = r.userCount
			for _, v := range r.users {
				v <- ev
			}
		}
	}
}
