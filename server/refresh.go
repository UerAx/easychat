/*
 * @Author: ww
 * @Date: 2022-07-13 01:54:16
 * @Description:
 * @FilePath: /easy-chat/server/refresh.go
 */
package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uerax/easychat/model"
)

var Refresh = &refresh{}

type refresh struct {
}

func (t *refresh) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Query("user")
		Room.MsgJoin(user)
		c.Redirect(http.StatusMovedPermanently, "/refresh/room?user="+user)
	}
}


func (refresh) Archive(c *gin.Context)  {
	type archive struct {
		User   string
		Events []model.Event
	}

		user := c.Query("user")

		c.HTML(http.StatusOK, "refresh.html", archive{
			User:   user,
			Events: Room.GetArchive(),
		})
}

func (refresh) Msg(c *gin.Context) {
		user := c.PostForm("user")
		message := c.PostForm("message")
		Room.MsgSay(user, message)
		c.Redirect(http.StatusMovedPermanently, "/refresh/room")
}

func (refresh) Leave(c *gin.Context) {
		user := c.Query("user")
		Room.MsgLeave(user)
		c.Redirect(http.StatusMovedPermanently, "/")
}
