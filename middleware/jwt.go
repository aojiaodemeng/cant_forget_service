package middleware

import (
	"cant_forget/utils"
	"cant_forget/utils/status_code"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour) // 10h之后过期
	claims := MyClaims{
		Username: username, // token里不携带敏感信息，比如password
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(), // 发布时间
			Issuer:    utils.Issuer,      // 发布者
		},
	}
	// 生成token
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", status_code.ERROR
	}
	return token, status_code.SUCCESS
}

// 解析/验证token
func CheckToken(tokenStr string) (*MyClaims, int) {
	parsedToken, _ := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := parsedToken.Claims.(*MyClaims); parsedToken.Valid {
		return key, status_code.SUCCESS
	} else {
		return nil, status_code.ERROR
	}
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code := status_code.SUCCESS
		if tokenHeader == "" {
			code = status_code.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": status_code.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = status_code.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": status_code.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])

		if tCode == status_code.ERROR {
			code = status_code.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": status_code.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = status_code.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": status_code.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}
}
