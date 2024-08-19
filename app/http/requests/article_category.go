package requests

type ArticleCategoryListReq struct {
	PageReq
	Name string `form:"name,omitempty"`
}

type ArticleCategoryListResp struct {
	ID         int64  `json:"id"`          // 分类id
	Name       string `json:"name"`        // 分类名称
	Alias      string `json:"alias"`       // 分类别名
	ArticleNum int64  `json:"article_num"` // 文章数量
	CreatedAt  string `json:"created_at"`  // 创建时间
}

type ArticleCategoryAddReq struct {
	Name  string `json:"name"`            // 分类名称
	Alias string `json:"alias,omitempty"` // 分类别名
}

type ArticleCategoryEditReq struct {
	ID    int64  `json:"id"`              // 分类id
	Name  string `json:"name"`            // 分类名称
	Alias string `json:"alias,omitempty"` // 分类别名
}

type ArticleCategoryDelReq struct {
	ID int64 `json:"id"` // 分类id
}

type ArticleCategorySelectResp struct {
	ID   int64  `json:"id"`   // 分类id
	Name string `json:"name"` // 分类名称
}
