package utils

import (
	"blogo/global"
	"blogo/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(global.CONF.Jwt.Secret) // jwtSecret

func GetToken(userID int, userName string) (string, int64, error) { // 生成Token，返回Encoded jwt，和过期时间
	expireTime := time.Now().Add(time.Duration(global.CONF.Jwt.ExpireTime) * time.Second)
	claims := &models.Claims{
		UserID:   userID, // FIXME: 无法取得UserID
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			Issuer:    "blogo",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	return token, expireTime.Unix(), err
}

func VerifyToken(tokenString string) (*models.Claims, error) { //验证Token，接收一个Encoded jwt，返回UserID和UserName
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

/*
Use example
func yourfunc() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
        token, err := getToken(123, "root")
        if err != nil {
            c.JSON(401, gin.H{
                "Error": "获取Token失败",
                "Info": err,
            })
            return
        }
		c.JSON(200, gin.H{
			"JwtToken": token,
		})
	})
    r.GET("/ver", func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        userID, userName, err := verifyToken(tokenString)
        if err != nil {
            fmt.Println(err)
            c.JSON(401, gin.H{
                "Error": "验证失败",
                "Info": err,
            })
            return
        }
        c.JSON(200, gin.H{
            "Info": "验证成功",
        })
        fmt.Println(userID, userName)
    })
    r.Run(":8080")
}
*/
