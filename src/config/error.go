package config

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (httpErr *HttpError) Error() string {
	rawBytes, err := json.Marshal(httpErr)
	if err != nil {
		return "{\"status\":\"500\", \"message\":\"Something Went Wrong\"}"
	}
	return string(rawBytes)
}

func NewHttpError(status int, message ...string) *HttpError {
	fallbackStatus := http.StatusInternalServerError
	fallbackMsg := http.StatusText(fallbackStatus)
	msg := ""
	statusCode := status

	if statusCode < 200 || statusCode > 599 {
		statusCode = fallbackStatus
	}

	if len(message) > 0 {
		msg = message[0]
	}
	if msg == "" && statusCode == fallbackStatus {
		msg = fallbackMsg
	} else {
		msg = http.StatusText(statusCode)
	}

	return &HttpError{
		Status:  statusCode,
		Message: msg,
	}
}

func IsHttpError(err error) bool {
	_, isHttpErr := err.(*HttpError)
	return isHttpErr
}

func HandleErr(ctx *gin.Context, err *HttpError) {
	log.Println(err.Error())
	ctx.AbortWithStatusJSON(err.Status, err)
}
