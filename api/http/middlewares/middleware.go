package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	http2 "net/http"
	"time"
	"vulscan/api/http"
	"vulscan/configs"
	"vulscan/src/adapter/repositories"
	"vulscan/src/enums"
)

// 🤤🐒🤫
func RequireAccessToken(context *http.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method alg %s", token.Header["alg"])
				}
				return configs.Get().AuthSecretKet, nil
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
		userRepository := repositories.NewUserRepositoryWithDbConnection(context.DBConnection)
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
