package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/emvi/hide"
	"github.com/kiwisheets/auth/permission"
)

// UserClaim structure
type UserClaim struct {
	UserID    hide.ID                  `json:"userId"`
	CompanyID hide.ID                  `json:"companyId"`
	Scopes    []*permission.Permission `json:"scopes"`
	jwt.StandardClaims
}

// Context structure, passed through context
type Context struct {
	UserID    hide.ID
	CompanyID hide.ID
	Scopes    []*permission.Permission
	Secure    bool
}
