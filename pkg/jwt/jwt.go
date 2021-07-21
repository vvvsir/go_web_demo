// json web token 令牌
package jwt

import (
	"go_web_demo/models"
	"go_web_demo/models/response"
	"go_web_demo/pkg/logger"
	"go_web_demo/pkg/setting"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string             `json:"username"`
	Role     []response.CasRole `json:"role"`
	Auth     []models.Auth      `json:"auth"`
	Id       int                `json:"id"`
	RealName string             `json:"real_name"`
	Tel      string             `json:"tel"`
	jwt.StandardClaims
}

// 类型转换
var JwtKey = []byte(setting.JwtConf.Key)

// 生成令牌
func MakeToken(adminUser models.AdminUser) (string, error) {
	// 过期时间
	expTime := time.Now().Add(time.Duration(setting.JwtConf.ExpTime) * time.Hour)
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: adminUser.UserName,
		Id:       adminUser.Id,
		RealName: adminUser.RealName,
		Tel:      adminUser.Tel,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(), //过期时间
			Subject:   "go_web_demo",  //签名主题
		},
	})
	return tokenClaim.SignedString(JwtKey)
}

// 解析令牌
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	logger.Logger.Error("解析jwt出错 : ", err)
	return nil, err
}
