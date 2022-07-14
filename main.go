/*
 * @Author: ww
 * @Date: 2022-07-15 03:11:02
 * @Description:
 * @FilePath: /easy-chat/main.go
 */
package main

import "github.com/uerax/easychat/server"

func main() {
	s := server.Server()
	s.Run(":8089")
}