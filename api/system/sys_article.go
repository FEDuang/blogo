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
	//参数绑定校验
	if err := c.C.ShouldBind(&publicArticleForm); err != nil {
		c.Response(http.StatusBadRequest, response.API_ERROR, "", nil)
		return
	}
	//获取文章
	article, err := article_service.GetArticle(publicArticleForm.ArticleId)
	if err != nil {
		c.FailWithDetailed(response.API_ERROR, err, "获取文章失败")
		return
	}
	//文章是否符合发布要求
	if article.Title == "" || article.ContentPath == "" || article.CoverImageURL == "" || len(article.Tags) == 0 {
		c.FailWithMessage(1, "发布失败")
		return
	}
	// 修改文章状态
	article.State = 1
	if err = article_service.EditArticle(article.ID, article); err != nil {
		c.FailWithMessage(1, "发布失败")
		return
	}
	c.Ok(0)
}

func (a *ArticleApi) DeleteArticle(c *response.GinContextE) {
	var deleteArticleForm request.DeleteArticleForm
	//绑定参数
	err := c.C.ShouldBindJSON(&deleteArticleForm)
	if err != nil {
		c.Response(http.StatusBadRequest, response.API_ERROR, "", nil)
		return
	}
	// 删除数据库文章
	if article_service.DeleteArticle(deleteArticleForm.ArticleId) != nil {
		c.FailWithData(1, nil)
		return
	}
	c.Ok(0)
}
