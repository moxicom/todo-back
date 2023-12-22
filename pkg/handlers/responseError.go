package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	Message string `json:"message"`
}

func newResponseError(c *gin.Context, statusCode int, errorMessage string) {
	logrus.Error(errorMessage)
	c.AbortWithStatusJSON(statusCode, error{errorMessage})
}
