package system

import (
	"blogo/models/request"
	"blogo/models/response"
	"blogo/services/article_service"
	"net/http"
)

type ArticleApi struct {
}

func (a *ArticleApi) PublishArticle(c *response.GinContextE) {
	var publicArticleForm request.PublishArticleForm
	err := c.C.ShouldBind(&publicArticleForm)
	if err != nil {
		//参数绑定失败
		c.Response(http.StatusBadRequest, response.API_ERROR, "", nil)
	}
	article, err := article_service.GetArticle(publicArticleForm.ArticleId)
	if err != nil {
		c.FailWithDetailed(response.API_ERROR, err, "获取文章失败")
	}
	if article.Title == "" || article.ContentPath == "" || article.CoverImageURL == "" || len(article.Tags) == 0 {
		c.FailWithMessage(1, "发布失败")
	}
	article.State = 1
	err = article_service.EditArticle(article.ID, article)
	if err != nil {
		c.FailWithMessage(1, "发布失败")
	}
	c.Ok(0)
}
