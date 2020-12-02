package seed

import (
	"github.com/kiwisheets/auth/operation"
	"github.com/kiwisheets/auth/permission"
	"github.com/kiwisheets/auth/subject"
	"gorm.io/gorm"
)

func seedRoles(db *gorm.DB) {
	createOrGetBuiltInRole(db,
		"Service Admin",
		"Default Service Admin Role with all permissions",
		[]permission.Permission{
			createOrGetPerm(db, subject.Any, operation.Any),
		},
	)

	createOrGetBuiltInRole(db,
		"Standard User",
		"Default User Role with basic read and create permissions",
		[]permission.Permission{
			createOrGetPerm(db, subject.Me, operation.Any),
			createOrGetPerm(db, subject.User, operation.Read),
			createOrGetPerm(db, subject.Company, operation.Read),
			createOrGetPerm(db, subject.Client, operation.Read),
			createOrGetPerm(db, subject.UserContact, operation.Read),
		},
	)
}
