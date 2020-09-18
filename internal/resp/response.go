package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RetResponse struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func JsonResponse(c *gin.Context, code int, data ...interface{}) {
	resp := &RetResponse{
		Code: code,
		Msg:  GetCodeMsg(code),
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
	return
}

func GetCodeMsg(code int) string {
	var (
		msg    string
		exists bool
	)

	if msg, exists = MsgCode[code]; exists {
		return msg
	}
	return MsgCode[Fail]
}
