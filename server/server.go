/*
 * @Author: ww
 * @Date: 2022-06-30 10:07:18
 * @Description:
 * @FilePath: /easy-chat/server/server.go
 */
package server

import (
	"github.com/gin-gonic/gin"
	"github.com/uerax/easychat/chat"
)

var Room = chat.NewRoom()

func Server() *gin.Engine {
	svr := gin.Default()
	
	// refresh
	svr.Static("/static", "static")
	svr.StaticFile("/", "web/index.html")
	svr.StaticFile("/refresh", "web/refresh.html")
	svr.StaticFile("/polling", "web/polling.html")
	svr.StaticFile("/ws", "web/ws.html")

	{
		// refresh
		svr.GET("/refresh/archive", Refresh.Archive)
		svr.POST("/refresh/msg", Refresh.Msg)
		svr.GET("/refresh/leave", Refresh.Leave)
	}
	{
		// polling
		svr.GET("/polling/archive", Longpolling.Archive)
		svr.POST("/polling/msg", Longpolling.Msg)
		svr.GET("/polling/leave", Longpolling.Leave)

	}

	{
		// websocket
		svr.GET("/ws/socket", Websocket.Handle())
	}
	return svr
}