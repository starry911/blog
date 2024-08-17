package requests

type TestReq struct {
	Name string `form:"name" json:"name"`
}

type TestResp struct {
	Name string `json:"name,omitempty"` // 绑定json参数
}
