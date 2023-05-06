package user

import (
	"encoding/json"
	"io"

	"dubinyang1993/go-restful/handler"
	"dubinyang1993/go-restful/pkg/errno"
	"dubinyang1993/go-restful/service/user"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	// params
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		handler.SendResponse(c, errno.ApiErrorIORead.Code, err.Error())
		return
	}

	var params user.AddReq
	err = json.Unmarshal(data, &params)
	if err != nil {
		handler.SendResponse(c, errno.ApiErrorJson.Code, err.Error())
		return
	}

	// add user
	err = user.Add(params.Name, params.Phone)
	if err != nil {
		handler.SendResponse(c, errno.ApiErrorSystem.Code, err.Error())
		return
	}

	handler.SendResponse(c, 0, "")
}
