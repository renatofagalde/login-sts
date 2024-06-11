package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.F) {

	r.GET("/:id", userController.FindUserByID)
}
