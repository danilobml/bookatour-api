package auth_handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(gr *gin.RouterGroup) {
	group := gr.Group("/auth")
	group.POST("/signup", Signup)
	group.POST("/login", Login)
}
