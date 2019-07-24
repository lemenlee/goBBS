package data

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const accessTokenSecret string = "KvdevAccessTokenSecret"
const refreshTokenSecret string = "KvdevRefreshTokenSecret"
const name string = "kv"

//JwtToken jwt令牌
type JwtToken struct {
	RefreshToken
	AccessToken
}

//RefreshToken 更新令牌
type RefreshToken struct {
	RefreshTokenStr string `json:"refreshToken"`
}

//AccessToken 访问令牌
type AccessToken struct {
	AccessTokenStr string `json:"accessToken"`
}

//GenerateToken 生成Access 和 Refresh令牌
func (jwt *JwtToken) GenerateToken(user UserLogin) error {
	err := jwt.AccessToken.GenerateToken(user)
	err = jwt.RefreshToken.GenerateToken(user)
	return err
}

//GenerateToken 生成Refresh令牌
func (rf *RefreshToken) GenerateToken(user UserLogin) error {
	tokenStr, err := generateToken(user, "refresh", 30, refreshTokenSecret)
	rf.RefreshTokenStr = tokenStr
	return err
}

//ValidToken 验证Refresh令牌
func (rf *RefreshToken) ValidToken() string {
	return validToken(rf.RefreshTokenStr, refreshTokenSecret)
}

//GenerateToken 生成Access令牌
func (ac *AccessToken) GenerateToken(user UserLogin) error {
	tokenStr, err := generateToken(user, "access", 7, accessTokenSecret)
	ac.AccessTokenStr = tokenStr
	return err
}

//ValidToken access令牌验证
func (ac *AccessToken) ValidToken() string {
	return validToken(ac.AccessTokenStr, accessTokenSecret)
}

func generateToken(user UserLogin, sub string, exp time.Duration, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	rtClaims := token.Claims.(jwt.MapClaims)
	rtClaims["sub"] = sub
	rtClaims["username"] = user.Username
	rtClaims["name"] = name
	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * exp).Unix()
	return token.SignedString([]byte(secret))
}

func validToken(tokenStr string, secret string) string {
	if len(tokenStr) == 0 {
		return ""
	}
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["name"].(string) == name {
			return claims["username"].(string)
		}
		return ""
	}
	return ""
}
