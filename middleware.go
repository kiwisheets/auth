package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// Middleware decodes the authorization header
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userJSON := c.Request.Header.Get("user")
		if userJSON == "" {
			c.Next()
			return
		}

		userClaim := &UserClaim{
			StandardClaims: jwt.StandardClaims{},
		}
		err := json.Unmarshal([]byte(userJSON), userClaim)
		if err != nil {
			log.Println("failed to unmarshal user JSON: " + userJSON)
			log.Println(err)
			c.Next()
			return
		}

		ctx := context.WithValue(c.Request.Context(), userCtxKey, Context{
			UserID:    userClaim.UserID,
			CompanyID: userClaim.CompanyID,
			Scopes:    userClaim.Scopes,
		})

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// For find the user from the context. Middleware must have run
func For(ctx context.Context) Context {
	raw, _ := ctx.Value(userCtxKey).(Context)
	return raw
}

func splitToken(header string) (string, error) {
	splitToken := strings.Split(header, "Bearer")

	if len(splitToken) != 2 || len(splitToken[1]) < 2 {
		return "", fmt.Errorf("bad token format")
	}

	return strings.TrimSpace(splitToken[1]), nil
}
