/*
 * @Author: ww
 * @Date: 2022-07-11 04:03:23
 * @Description:
 * @FilePath: /easy-chat/logger/logger.go
 */
package logger

import "github.com/UerAx/ulog/v2"

var Log *ulog.Ulog

func init() {
	Log = ulog.NewLog()
}