package dao

type InterfaceDao interface {
	IAdminUsers
	IArticle
	IArticleCategory
	IArticleContent
	IArticleTags
}

var IDao InterfaceDao

type Dao struct {
}

func New() {
	IDao = &Dao{}
}
