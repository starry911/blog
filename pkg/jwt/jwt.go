package jwt

import (
	"blog/pkg/config"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// Secret 定义Jwt密钥
var Secret = []byte(config.GetConf().Jwt.Secret)

// TTL 定义Jwt过期时间
var TTL = time.Second * time.Duration(config.GetConf().Jwt.TTL)

// Claims 用来生成token的结构体
type Claims struct {
	Param int64 `json:"param"`
	jwt.RegisteredClaims
}

// CreateToken 创建token
func CreateToken(UserId int64) (string, error) {
	c := Claims{
		UserId, // 自定义字段，一般是用户的唯一标识
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TTL)), // 过期时间
			Issuer:    "blog-jwt",                              // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Secret)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return 0, err
	}
	// 校验token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Param, nil
	}
	return 0, errors.New("令牌无效")
}
