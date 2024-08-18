package requests

// 登录
type BackLoginReq struct {
	Account  string `json:"account"`  // 用户账号
	Password string `json:"password"` // 用户密码，经过MD5加密的
}

type BackLoginResp struct {
	AccessToken string `json:"access_token"` // 登录凭证
	ExpiresTime int64  `json:"expires_time"` // 过期时间
	AdminId     int64  `json:"admin_id"`     // 用户id
	Account     string `json:"account"`      // 登录账号
	Nickname    string `json:"nickname"`     // 昵称
	CoverImg    string `json:"cover_img"`    // 头像
	LastIp      string `json:"last_ip"`      // 最后登录IP
	LastTime    string `json:"last_time"`    // 最后登录时间
}

// 获取用户信息
type UserInfoResp struct {
	AdminId  int64  `json:"admin_id"`  // 用户id
	Account  string `json:"account"`   // 登录账号
	Nickname string `json:"nickname"`  // 昵称
	CoverImg string `json:"cover_img"` // 头像
	LastIp   string `json:"last_ip"`   // 最后登录IP
	LastTime string `json:"last_time"` // 最后登录时间
}

// 修改用户信息
type SetUserInfoReq struct {
	Nickname string `json:"nickname,omitempty"`  // 昵称
	CoverImg string `json:"cover_img,omitempty"` // 头像
}

// 修改密码
type SetUserPasswordReq struct {
	OldPassword string `json:"old_password"` // 原密码
	NewPassword string `json:"new_password"` // 新密码
}
