package seed

import (
	"github.com/kiwisheets/auth/subject"
	"gorm.io/gorm"
)

func seedPermissions(db *gorm.DB) {
	for _, sub := range subject.SubjectNames() {
		subject, _ := subject.ParseSubject(sub)
		createPermsForAllOperations(db, subject)
	}
}
