package requests

type ArticleTagsListReq struct {
	PageReq
	Name string `form:"name,omitempty"` // 标签名称
}
type ArticleTagsListResp struct {
	ID        int64  `json:"id"`         // 标签id
	Name      string `json:"name"`       // 标签名称
	CreatedAt string `json:"created_at"` // 创建时间
}

type ArticleTagsDelReq struct {
	Ids []int64 `form:"ids"` // 标签id集合
}

type ArticleTagsSelectResp struct {
	ID   int64  `json:"id"`   // 标签id
	Name string `json:"name"` // 标签名称
}
