package global

import "github.com/golang-jwt/jwt/v4"

// 全局结构体

// UserInfo token信息
type UserInfo struct {
	UID         int    `json:"uid"`          //用户id
	Username    string `json:"username"`     //用户名
	AuthorityId int    `json:"authority_id"` //权限等级
}

type Claims struct {
	UserInfo
	BufferTime int
	jwt.RegisteredClaims
}
