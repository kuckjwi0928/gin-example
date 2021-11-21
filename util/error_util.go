package util

import "github.com/gin-gonic/gin"

func MakeErrorMessage(error error) map[string]interface{} {
	return gin.H{"message": error.Error()}
}
