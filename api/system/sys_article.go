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
	if exist, err := article_service.ExistArticleByID(publicArticleForm.ArticleId); err != nil || exist == false {
		c.FailWithMessage(response.API_ERROR, "获取文章失败")
		return
	}
	//获取文章
	article, err := article_service.GetArticle(publicArticleForm.ArticleId)
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

func (a *ArticleApi) ListArticles(c *response.GinContextE) {
	var listArticlesForm request.ListArticlesForm
	//参数校验与绑定
	if err := c.C.ShouldBindJSON(&listArticlesForm); err != nil {
		c.Response(http.StatusBadRequest, response.API_ERROR, "", nil)
		return
	}
	//获取文章并返回
	if articles, err := article_service.GetArticles(listArticlesForm.Offset, listArticlesForm.Size, map[string]interface{}{}); err == nil {
		c.OkWithData(0, response.ListArticleResponse{
			Articles:   articles,
			TotalCount: len(articles),
		})
	} else {
		c.Fail(1)
		println(err.Error())
	}
}
