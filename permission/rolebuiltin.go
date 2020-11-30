package permission

import (
	"log"
	"strings"

	"github.com/kiwisheets/auth/operation"
	"github.com/kiwisheets/auth/subject"
	orm "github.com/kiwisheets/orm/model"
)

// BuiltinRole is a built in group of permissions
type BuiltinRole struct {
	orm.Model
	Name        string
	Description string
	Permissions []Permission `gorm:"many2many:builtinrole_permissions"`
}

// GetName returns the name of the role
func (r BuiltinRole) GetName() string {
	return r.Name
}

// GetDescription returns the description of the role
func (r BuiltinRole) GetDescription() string {
	return r.Name
}

// GetPermissions return the roles permissions
func (r BuiltinRole) GetPermissions() []Permission {
	return r.Permissions
}

// CheckPermission will compare a permission string (like from graphql schema)
func (r BuiltinRole) CheckPermission(requestedPerm string) bool {
	strings := strings.Split(requestedPerm, ":")
	if len(strings) != 2 {
		return false
	}

	sub, err := subject.ParseSubject(strings[0])
	if err != nil {
		log.Printf("failed to check permission. unable to parse subject: %s \n", strings[0])
		return false
	}
	op, err := operation.ParseOperation(strings[1])
	if err != nil {
		log.Printf("failed to check permission. unable to parse operation: %s \n", strings[1])
		return false
	}

	// Run through permissions to check for the permission
	for _, perm := range r.Permissions {
		if perm.CheckPermission(sub, op) {
			return true
		}
	}
	return false
}
