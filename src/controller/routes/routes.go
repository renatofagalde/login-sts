package routes

import (
	"github.com/gin-gonic/gin"
	controller "main/src/controller/user"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/email/:email", userController.FindUserByEmail)
	r.POST("/login", userController.Login)
	r.Any("/token_verify", userController.TokenVerify)
}
