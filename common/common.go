/*
 * @Author: ww
 * @Date: 2022-07-12 03:35:45
 * @Description:
 * @FilePath: /easy-chat/common/common.go
 */
package common

import (
	"encoding/base64"
	"time"
)

func UUid() string {
	return base64.StdEncoding.EncodeToString([]byte(time.Now().GoString()))
}