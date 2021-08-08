package system

import (
	"blogo/api"
	"blogo/models/response"

	"github.com/gin-gonic/gin"
)

type LoginRouter struct {

}

func (s *LoginRouter)InitLoginRouter(Router *gin.RouterGroup) {
	systemApi := Router.Group("api/system/auth")
	{
		systemApi.POST("/login", func (c *gin.Context) {
			api.ApiGroupApp.SystemApiGroup.LoginApi.Login(&response.GinContextE{C: c})
		})
	}
}
