package response

type ArticleDetails struct {
	ArticleId      int          `json:"article_id,string"`
	Title          string       `json:"title"`
	Desc           string       `json:"desc"`
	Content        string       `json:"content"`
	State          int          `json:"state,string"`
	CoverImage     []CoverImage `json:"cover_image,string"`
	Views          int          `json:"views,string"`
	EffectiveViews int          `json:"effective_views,string"`
	Tags           []string     `json:"tags"`
	CreatedTime    int          `json:"created_time,string"`
	UpdateTime     int          `json:"update_time,string"`
	DeletedTime    int          `json:"deleted_time,string"`
}
