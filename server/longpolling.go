/*
 * @Author: ww
 * @Date: 2022-07-15 02:49:57
 * @Description:
 * @FilePath: /easy-chat/server/longpolling.go
 */
package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uerax/easychat/model"
)

var Longpolling = &longpolling{}


type longpolling struct {
}

func (longpolling) Msg(c *gin.Context) {
	type req struct {
		Name string `json:"name"`
		Msg  string `json:"msg"`
	}
		form := req{}
		if err := c.BindJSON(&form); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		Room.MsgSay(form.Name, form.Msg)
		c.JSON(http.StatusOK, struct {
			Status int `json:"status"`
		}{200})
}

// 轮询获取指定时间戳之后的聊天记录
func (longpolling) Archive(c *gin.Context) {
		lastReceived, _ := strconv.ParseInt(c.Query("ts"), 10, 64)

		var events []model.Event
		// filter archive
		for _, event := range Room.GetArchive() {
			if event.Timestamp > lastReceived {
				events = append(events, event)
			}
		}
		c.JSON(http.StatusOK, events)
}

func (longpolling) Leave(c *gin.Context) {
		user := c.Query("name")
		Room.MsgLeave(user)
		c.Redirect(http.StatusMovedPermanently, "/")
}
