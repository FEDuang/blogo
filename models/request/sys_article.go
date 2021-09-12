package request

type AddArticleForm struct {
	Title      string     `form:"title" json:"title" binding:"require"`
	CoverImage CoverImage `json:"coverImage" form:"coverImage"`
	Desc       string     `json:"desc" form:"desc"`
	Content    string     `json:"content" form:"content"`
	Tags       []string   `json:"tags" form:"tags"`
}

type PublishArticleForm struct {
	ArticleId uint `json:"articleId,string" valid:"require"`
}

type DeleteArticleForm struct {
	ArticleId uint `json:"articleId,string" valid:"require"`
}
