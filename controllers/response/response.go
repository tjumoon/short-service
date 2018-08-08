package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"short-service/utils/e"
)

type BasicResponse struct {
	succeed 	bool	`json:"successd"`
	errorCode   string	`json:"error_code"`
	errorMessage string	`json:"error_message"`
}

func J(c *gin.Context, data interface{})  {
	JSON(c, http.StatusOK, data)
}

func JSON(c *gin.Context, status int, data interface{})  {
	c.JSON(status, gin.H{
		"succeed": true,
		"errorCode": 0,
		"errorMessage": "",
		"data": data,
	})
}

func E(c *gin.Context, errorCode int)  {
	ErrorJSON(c, http.StatusOK, errorCode)
}

func ErrorJSON(c *gin.Context, status int, errorCode int)  {
	if errorCode == 0{
		errorCode = 999999
	}
	c.JSON(status, gin.H{
		"succeed": false,
		"errorCode": errorCode,
		"errorMessage": e.GetErrorMessage(errorCode),
	})
}

