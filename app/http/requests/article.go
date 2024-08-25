package requests

type ArticleListReq struct {
	PageReq
	Title      string `form:"title,omitempty"`       // 标题
	CategoryId int64  `form:"category_id,omitempty"` // 分类id
	Status     int64  `form:"status,omitempty"`      // 状态，0：全部，1：正常，2：隐藏
	StartTime  string `form:"start_time,omitempty"`  // 开始时间，格式：2020-12-12 00:00:00
	EndTime    string `form:"end_time,omitempty"`    // 结束时间，格式：2020-12-12 00:00:00
}

type ArticleListResp struct {
	ID           int64    `json:"id"`            // 文章id
	UUID         string   `json:"uuid"`          // 文章唯一id
	Title        string   `json:"title"`         // 文章标题
	CategoryName string   `json:"category_name"` // 文章分类
	CoverImg     string   `json:"cover_img"`     // 封面
	PublishTime  string   `json:"publish_time"`  // 发布时间
	Status       int64    `json:"status"`        // 状态，1：正常，2：隐藏
	Tags         []string `json:"tags"`          // 文章标签
	ViewNum      int64    `json:"view_num"`      // 浏览量
	LikeNum      int64    `json:"like_num"`      // 点赞量
	CreatedAt    string   `json:"created_at"`    // 创建时间
}

type ArticleAddReq struct {
	CategoryId  int64   `json:"category_id"`  // 分类id
	Title       string  `json:"title"`        // 文章标题
	Describe    string  `json:"describe"`     // 描述
	CoverImg    string  `json:"cover_img"`    // 封面
	Tags        []int64 `json:"tags"`         // 标签
	Status      int64   `json:"status"`       // 状态，1：显示，2：隐藏
	PublishTime string  `json:"publish_time"` // 发布时间，格式2024-12-12 00:00:00
	ContentHtml string  `json:"content_html"` // 文章内容，html格式
	ContentMd   string  `json:"content_md"`   // 文章内容，md格式
}
