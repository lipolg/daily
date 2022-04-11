package api

import (
	"daily/service"
	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context) {
	var vs service.VerifyService
	//获取get请求参数
	vs.Signature = c.Query("signature")
	vs.Timestamp = c.Query("timestamp")
	vs.Nonce = c.Query("nonce")
	vs.Echostr = c.Query("echostr")
	res := vs.Verify()
	_, _ = c.Writer.Write([]byte(res))
}
