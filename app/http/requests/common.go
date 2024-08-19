package requests

// 全局统一的分页
type PageReq struct {
	Page  int64 `form:"page,omitempty,default=1"`
	Limit int64 `form:"limit,omitempty,default=20"`
}

type PageResp struct {
	Total int64 `json:"total"` // 总条数
	Page  int64 `json:"page"`  // 当前页数
}
