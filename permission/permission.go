package permission

import (
	"log"
	"strings"

	"github.com/kiwisheets/auth/operation"
	"github.com/kiwisheets/auth/subject"
	orm "github.com/kiwisheets/orm/model"
)

// Permission model
type Permission struct {
	orm.Model
	Name        string
	Description string
	Subject     subject.Subject     `gorm:"type:integer"`
	Operation   operation.Operation `gorm:"type:integer"`
}

// CheckPermission check a full permission string (like ones from the graphql schema)
// And return whether or not the permission matches
func (p Permission) CheckPermission(sub subject.Subject, op operation.Operation) bool {
	return (p.Subject == subject.Any || p.Subject == sub) &&
		(p.Operation == operation.Any || p.Operation == op)
}

func (p Permission) CheckPermissionString(requestedPerm string) bool {
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

	return p.CheckPermission(sub, op)
}

// func (p Permission) CheckPermission(permString string) bool {
// 	strings := strings.Split(permString, ":")
// 	if len(strings) != 2 {
// 		return false
// 	}

// 	var sub subject.Subject
// 	var op operation.Operation

// 	sub.Scan(strings[0])
// 	op.Scan(strings[1])

// 	return (p.Subject == subject.Any || p.Subject == sub) &&
// 		(p.Operation == operation.Any || p.Operation == op)
// }
