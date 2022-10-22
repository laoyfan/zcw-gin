package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/groupcache/singleflight"
	"time"
	"zcw-admin-server/global"
)

const TokenExpireDuration = time.Hour * 24

// CreateToken 创建token
func CreateToken(base global.UserInfo) (string, error) {
	claims := global.Claims{
		UserInfo:   base,
		BufferTime: global.CONFIG.Jwt.BufferTime,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    global.CONFIG.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(global.CONFIG.Jwt.SigningKey))
}

// ParseToken 解析token
func ParseToken(tokenString string) (*global.UserInfo, error) {
	token, err := jwt.ParseWithClaims(tokenString, &global.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return global.CONFIG.Jwt.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token错误")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errors.New("token过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token未激活")
			} else {
				return nil, errors.New("token异常")
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*global.Claims); ok && token.Valid {
			return &claims.UserInfo, nil
		}
		return nil, errors.New("token无效")
	} else {
		return nil, errors.New("token异常")
	}
}

// CreateTokenByOldToken 以旧换新
func CreateTokenByOldToken(oldToken string, base global.UserInfo) (string, error) {
	s := &singleflight.Group{}
	token, err := s.Do("JWT:"+oldToken, func() (interface{}, error) {
		return CreateToken(base)
	})
	return token.(string), err
}
