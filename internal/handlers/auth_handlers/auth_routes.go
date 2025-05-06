package auth_handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(gr *gin.RouterGroup) {
	authGroup := gr.Group("/")
	authGroup.POST("/signup", signup)
	authGroup.POST("/login", login)
}
