package Routes

import (
	"golang/Controllers/UserController"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Routes Group
	group0 := r.Group("api/user")
	{
		group0.GET("", UserController.GetListUser) // Using Method GET
		group0.POST("", UserController.CreateUser) // Using Method POST
		// group0.PUT("/:id", Controller.Data)			// Using Method PUT
		// group0.DELETE("/:id", Controller.Data)		// Using Method DELETE
	}

	return r
}
