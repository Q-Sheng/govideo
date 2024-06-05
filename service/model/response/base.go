package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResSuccess(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK,
		ResData{
			code,
			msg,
			data,
		})
}

func ResFail(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, ResData{
		http.StatusInternalServerError,
		err.Error(),
		nil,
	})
}
