package permission

import (
	"log"
	"strings"

	"github.com/emvi/hide"
	"github.com/kiwisheets/auth/operation"
	"github.com/kiwisheets/auth/subject"
	orm "github.com/kiwisheets/orm/model"
)

// CustomRole is a group of permissions that is created for a specific company
type CustomRole struct {
	orm.SoftDelete
	Name        string
	Description string
	CompanyID   hide.ID
	Permissions []Permission `gorm:"many2many:customrole_permissions"`
}

// GetName returns the name of the role
func (r CustomRole) GetName() string {
	return r.Name
}

// GetDescription returns the description of the role
func (r CustomRole) GetDescription() string {
	return r.Name
}

// GetPermissions return the roles permissions
func (r CustomRole) GetPermissions() []Permission {
	return r.Permissions
}

// CheckPermission will compare a permission string (like from graphql schema)
func (r CustomRole) CheckPermission(requestedPerm string) bool {
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
