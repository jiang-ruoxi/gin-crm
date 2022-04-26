package lib

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// 创建一个jwt使用的密钥
var jwtSecret = []byte("51talk")

//Claims 结构体，创建JWT的结构，将jwt.StandardClaims作为嵌入式类型，提供对应的字段
type Claims struct {
	UserId   int64
	UserName string
	jwt.StandardClaims
}

//CreateTokenJWT 创建jwt Token
func CreateTokenJWT(userId int64, userName string, expireDuration time.Duration) (string, error) {
	//声明jwt toke令牌的到期时间
	expire := time.Now().Add(expireDuration)
	issuer := "jiang"

	//使用用于签名的算法和令牌，将userId，roleId，过期时间作为参数数据写入token中
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserId:   userId,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(), //到期时间
			Issuer:    issuer,        //签发者
		},
	})
	//创建jwt字符串
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		//如果创建jwt时候出错
		return tokenStr, err
	}

	return tokenStr, nil
}

//CheckTokenAuth 验证token信息是否有效
func CheckTokenAuth(token string) (*Claims, error) {
	//解析jwt字符串，并将结果信息存储在Claims中
	claimsToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claimsToken != nil {
		if claims, ok := claimsToken.Claims.(*Claims); ok && claimsToken.Valid {
			return claims, nil
		}
	}
	return nil, err
}
