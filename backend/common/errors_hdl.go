package common

import (
	"time"

	"github.com/gin-gonic/gin"
)

// ErrorResponse struct to define structure for error responses
type ResponseHandle struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Error     string      `json:"error,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

func NewSuccessH(c *gin.Context, data interface{}) {
	var r ResponseHandle
	r.Status = 200
	r.Message = "Success"
	r.Data = data
	r.Timestamp = time.Now()

	c.JSON(r.Status, r)
}

func NewErrorH(c *gin.Context, app *AppError) {
	var r ResponseHandle
	r.Status = app.Code
	r.Message = app.Message
	r.Error = app.Err.Error()
	r.Timestamp = time.Now()

	c.JSON(r.Status, r)
}
