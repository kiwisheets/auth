package token

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/emvi/hide"
	"github.com/kiwisheets/auth"
	"github.com/kiwisheets/util"
)

type UserTokenParams struct {
	ID    hide.ID
	Email string
}

func validateTokenAndGetUserID(t string, cfg *util.JWTConfig) (hide.ID, error) {
	token, err := jwt.ParseWithClaims(t, &auth.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return cfg.PublicKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*auth.UserClaim); ok && token.Valid && token.Method == jwt.SigningMethodES256 {
		return claims.UserID, nil
	}
	return 0, err
}

// buildAndSignToken signs and returned a JWT token from a User
func buildAndSignToken(u UserTokenParams, cfg *util.JWTConfig, expires time.Duration) (string, error) {
	claims := auth.UserClaim{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			Issuer:   "KiwiSheets",
			IssuedAt: time.Now().Unix(),
			Subject:  u.Email,
			// ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	if expires != 0 {
		claims.ExpiresAt = time.Now().Add(expires).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(cfg.PrivateKey)
	log.Println(tokenString)
	log.Println(err)

	userId, err := validateTokenAndGetUserID(tokenString, cfg)
	log.Println(userId)
	log.Println(err)

	return tokenString, err
}
