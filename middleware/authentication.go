package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("maibokkrub")

const tokenDuration = 15 * time.Minute

type CustomClaim struct {
	UserID int `json:"userID"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Check if the token is provided
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		userID, err := GetUserID(tokenString)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		c.Set("userID", userID)
		c.Next()
	}
}

func GetUserID(tokenString string) (int, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*CustomClaim); ok && token.Valid {
		return claims.UserID, nil
	} else {
		return 0, errors.New("Invalid Token")
	}
}

func NewToken(userId int) (string, error) {
	// claims := jwt.MapClaims{
	// 	"userID": string(userId),
	// 	"exp":    time.Now().Add(tokenDuration).Unix(),
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaim{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenDuration)),
		},
	})
	return token.SignedString(secretKey)
}
