package requests

// PageReq 全局统一的分页
type PageReq struct {
	Page  int `form:"page,omitempty,default=1"`
	Limit int `form:"limit,omitempty,default=20"`
}
