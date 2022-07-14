/*
 * @Author: ww
 * @Date: 2022-07-12 03:35:45
 * @Description:
 * @FilePath: /easy-chat/common/common_test.go
 */
package common

import (
	"encoding/base64"
	"testing"
	"time"
)

func TestUUid(t *testing.T) {
	println(base64.StdEncoding.EncodeToString([]byte(time.Now().GoString())))
	print(base64.StdEncoding.EncodeToString([]byte(time.Now().GoString())))
}
