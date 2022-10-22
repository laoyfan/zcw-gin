package entity

type Index struct {
	Id   int    `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

// LoginReq 登录参数
type LoginReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginResp struct {
	Username string `json:"username"`
	Access   string `json:"access"`
	Refresh  string `json:"refresh"`
}
