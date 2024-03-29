package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type responseId struct {
	Id int `json:"id"`
}

type errorMsg struct {
	Message string `json:"message"`
}

func newResponseError(c *gin.Context, statusCode int, errorMessage string) {
	logrus.Error(errorMessage)
	c.AbortWithStatusJSON(statusCode, errorMsg{errorMessage})
}
