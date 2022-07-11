/*
 * @Author: ww
 * @Date: 2022-06-30 10:52:15
 * @Description:
 * @FilePath: /easy-chat/cfg/config.go
 */
package cfg

import "github.com/uerax/goconf"

var Config *goconf.CfgFile

func Init() {
	Config = goconf.NewCfgFile()
	Config.ReadAll("../etc")
}