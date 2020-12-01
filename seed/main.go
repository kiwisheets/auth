package seed

import (
	"log"

	"github.com/kiwisheets/auth/operation"
	"github.com/kiwisheets/auth/permission"
	"github.com/kiwisheets/auth/subject"
	"gorm.io/gorm"
)

// EnsurePermissions ensures that all permissions exist in the database
func EnsurePermissions(db *gorm.DB) {
	seedPermissions(db)
}

// EnsureBuiltinRoles ensure that all builtin roles exist in the database
func EnsureBuiltinRoles(db *gorm.DB) {
	seedRoles(db)
}

func createPermsForAllOperations(db *gorm.DB, subject subject.Subject) {
	for _, o := range operation.OperationNames() {
		operation, _ := operation.ParseOperation(o)
		createOrGetPermWithDesc(
			db,
			subject,
			operation,
			"Allow "+operation.String()+" operations on "+subject.String()+" resources",
		)
	}
}

func createOrGetPerm(db *gorm.DB, subject subject.Subject, operation operation.Operation) permission.Permission {
	return createOrGetPermWithDesc(
		db,
		subject,
		operation,
		"Allow "+operation.String()+" operations on "+subject.String()+" resources",
	)
}

func createOrGetPermWithDesc(db *gorm.DB, subject subject.Subject, operation operation.Operation, description string) permission.Permission {
	var perm permission.Permission
	db.FirstOrCreate(&perm, permission.Permission{
		Name:        subject.String() + ":" + operation.String(),
		Description: description,
		Subject:     subject,
		Operation:   operation,
	})

	return perm
}

func createOrGetBuiltInRole(db *gorm.DB, name string, description string, permissions []permission.Permission) permission.BuiltinRole {
	var perm permission.BuiltinRole
	err := db.Where(permission.BuiltinRole{
		Name: name,
	}).Attrs(permission.BuiltinRole{
		Description: description,
		Permissions: []permission.Permission{
			createOrGetPerm(db, subject.Me, operation.Any),
			createOrGetPerm(db, subject.User, operation.Read),
			createOrGetPerm(db, subject.Company, operation.Read),
			createOrGetPerm(db, subject.Client, operation.Read),
			createOrGetPerm(db, subject.UserContact, operation.Read),
		},
	}).FirstOrCreate(&perm).Error
	if err != nil {
		log.Println("failed to create built in role " + name)
		log.Println(err)
	}
	return perm
}
