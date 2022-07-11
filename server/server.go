/*
 * @Author: ww
 * @Date: 2022-06-30 10:07:18
 * @Description:
 * @FilePath: /easy-chat/server/server.go
 */
package server

import "github.com/gin-gonic/gin"

func Init() {
	svr := gin.Default()
	svr.Run(":8080")
}

