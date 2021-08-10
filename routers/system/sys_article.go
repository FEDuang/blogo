package system

import (
	"blogo/api"
	"blogo/models/response"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
}

func (s *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	systemApi := Router.Group("api/system/article")
	{
		systemApi.POST("/listArticles")
		systemApi.POST("/getArticleDetails")
		systemApi.POST("/searchArticle")
		systemApi.POST("/saveArticle")
		systemApi.POST("/publishArticle", func(c *gin.Context) {
			api.ApiGroupApp.SystemApiGroup.PublishArticle(&response.GinContextE{C: c})
		})
		systemApi.POST("/updateArticle")
		systemApi.POST("/recoverArticle")
		systemApi.POST("/listArticleTags")
	}
}
