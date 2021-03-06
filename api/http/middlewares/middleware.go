package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	http2 "net/http"
	"time"
	"vulscan/api/http/context"
	"vulscan/configs"
	"vulscan/src/adapter/repositories"
	"vulscan/src/enums"
)

// 🤤🐒🤫
func RequireAccessToken(context *context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method alg %s", token.Header["alg"])
				}
				return []byte(configs.Get().AuthSecretKet), nil
			})
		if err != nil || token == nil {
			unauthorized(c)
			return
		}
		mapClaims := token.Claims.(jwt.MapClaims)
		userID := mapClaims["uid"]
		exp := mapClaims["exp"]

		if !token.Valid || userID == "" || exp == nil {
			unauthorized(c)
		}
		userRepository := repositories.NewUserRepository(context.DBConnection)
		user, err := userRepository.FindByID(userID.(string))
		if err != nil {
			unauthorized(c)
		}
		if int64(exp.(float64)) < time.Now().Unix() {
			unauthorized(c)
			return
		}
		c.Set("user", user)
	}
}

func unauthorized(c *gin.Context) {
	c.JSON(http2.StatusUnauthorized, enums.ErrUnauthorized.GetMessage())
	return
}

func CrossBrowser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS, HEAD")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	if c.Request.Method == "OPTIONS" && len(c.GetHeader("X-Request-Method")) == 0 {
		c.AbortWithStatus(200)
		return
	}
	c.Next()
}
