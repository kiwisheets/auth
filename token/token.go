package token

import (
	"crypto/ecdsa"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/emvi/hide"
	"github.com/kiwisheets/auth"
	"github.com/kiwisheets/auth/permission"
)

type UserTokenParams struct {
	ID          hide.ID
	CompanyID   hide.ID
	Email       string
	Permissions []*permission.Permission
}

// ValidateTokenAndGetUserID verifies and returns the contents of a signed JWT
func ValidateTokenAndGetUserID(t string, publicKey *ecdsa.PublicKey) (hide.ID, error) {
	token, err := jwt.ParseWithClaims(t, &auth.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*auth.UserClaim); ok && token.Valid && token.Method == jwt.SigningMethodES256 {
		return claims.UserID, nil
	}
	return 0, err
}

// BuildAndSignToken signs and returned a JWT token from a User
func BuildAndSignToken(u UserTokenParams, privateKey *ecdsa.PrivateKey, expires time.Duration) (string, error) {
	claims := auth.UserClaim{
		UserID:    u.ID,
		CompanyID: u.CompanyID,
		Scopes:    u.Permissions,
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
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		log.Println(err)
	}

	return tokenString, err
}
