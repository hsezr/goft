package models

type ArticleModel struct {
	UserId   int    `json:"id" uri:"id" binding:"required,gt=0"`
	UserName string `json:"user_name"`
	UserView int    `json:"user_view"`
}

func (this *ArticleModel) String() string {
	return "ArticleModel"
}

func NewArticleModel() *ArticleModel {
	return &ArticleModel{}
}
