package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/groupcache/singleflight"
	"time"
	"zcw-gin/global"
)

// CreateToken 创建token 生成access_token 和 refresh_token
func CreateToken(base global.UserInfo) (aToken, rToken string, err error) {
	fmt.Println(global.CONFIG.App.Jwt.SigningKey)

	TokenExpireDuration := time.Hour * time.Duration(global.CONFIG.App.Jwt.ExpiresTime) // 过期时长
	SigningKey := []byte(global.CONFIG.App.Jwt.SigningKey)
	claims := global.Claims{
		UserInfo: base,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    global.CONFIG.App.Jwt.Issuer,
		},
	}
	fmt.Println(SigningKey, "33333333333333")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token, "4444444444444")
	// access_token
	aToken, err = token.SignedString(SigningKey)
	fmt.Println(err, "22222222222")
	// refresh_token
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
		Issuer:    global.CONFIG.App.Jwt.Issuer,
	}).SignedString(SigningKey)
	fmt.Println(err, "1111111111111")
	return
}

// ParseToken 解析access_token
func ParseToken(tokenString string) (*global.UserInfo, error) {
	token, err := jwt.ParseWithClaims(tokenString, &global.Claims{}, checkToken)
	fmt.Println(err, "66666666666666666")
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token错误")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
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

func checkToken(*jwt.Token) (interface{}, error) {
	return []byte(global.CONFIG.App.Jwt.SigningKey), nil
}

// RefreshToken 刷新access_token和refresh_token
func RefreshToken(aToken, rToken string) (string, string, error) {
	// 检测refresh_token 时间 格式
	if _, err := jwt.Parse(rToken, checkToken); err != nil {
		return "", "", err
	}

	s := &singleflight.Group{}
	token, err := s.Do("JWT:"+aToken, func() (interface{}, error) {
		oToken, err := jwt.ParseWithClaims(aToken, &global.Claims{}, checkToken)
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			if claims, ok := oToken.Claims.(*global.Claims); ok && oToken.Valid {
				t1, t2, err := CreateToken(claims.UserInfo)
				if err == nil {
					return []string{t1, t2}, nil
				}
			}
		}
		return nil, err
	})
	tokens := token.([]string)
	if err != nil {
		return "", "", err
	}
	return tokens[0], tokens[1], nil
}
