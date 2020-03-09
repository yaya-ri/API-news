package controllers

import "github.com/gin-gonic/gin"

//NewsControllerInterface interface
type NewsControllerInterface interface {
	Store(context *gin.Context)
	Find(context *gin.Context)
}
