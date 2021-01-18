// Code generated by go-enum
// DO NOT EDIT!

package subject

import (
	"fmt"
	"strings"
)

const (
	// None is a Subject of type None
	None Subject = iota
	// Any is a Subject of type Any
	Any
	// Me is a Subject of type Me
	Me
	// User is a Subject of type User
	User
	// UserContact is a Subject of type UserContact
	UserContact
	// OtherUser is a Subject of type OtherUser
	OtherUser
	// Company is a Subject of type Company
	Company
	// OtherCompany is a Subject of type OtherCompany
	OtherCompany
	// Client is a Subject of type Client
	Client
	// ClientContact is a Subject of type ClientContact
	ClientContact
	// Contact is a Subject of type Contact
	Contact
	// Invoice is a Subject of type Invoice
	Invoice
)

const _SubjectName = "NoneAnyMeUserUserContactOtherUserCompanyOtherCompanyClientClientContactContactInvoice"

var _SubjectNames = []string{
	_SubjectName[0:4],
	_SubjectName[4:7],
	_SubjectName[7:9],
	_SubjectName[9:13],
	_SubjectName[13:24],
	_SubjectName[24:33],
	_SubjectName[33:40],
	_SubjectName[40:52],
	_SubjectName[52:58],
	_SubjectName[58:71],
	_SubjectName[71:78],
	_SubjectName[78:85],
}

// SubjectNames returns a list of possible string values of Subject.
func SubjectNames() []string {
	tmp := make([]string, len(_SubjectNames))
	copy(tmp, _SubjectNames)
	return tmp
}

var _SubjectMap = map[Subject]string{
	0:  _SubjectName[0:4],
	1:  _SubjectName[4:7],
	2:  _SubjectName[7:9],
	3:  _SubjectName[9:13],
	4:  _SubjectName[13:24],
	5:  _SubjectName[24:33],
	6:  _SubjectName[33:40],
	7:  _SubjectName[40:52],
	8:  _SubjectName[52:58],
	9:  _SubjectName[58:71],
	10: _SubjectName[71:78],
	11: _SubjectName[78:85],
}

// String implements the Stringer interface.
func (x Subject) String() string {
	if str, ok := _SubjectMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Subject(%d)", x)
}

var _SubjectValue = map[string]Subject{
	_SubjectName[0:4]:   0,
	_SubjectName[4:7]:   1,
	_SubjectName[7:9]:   2,
	_SubjectName[9:13]:  3,
	_SubjectName[13:24]: 4,
	_SubjectName[24:33]: 5,
	_SubjectName[33:40]: 6,
	_SubjectName[40:52]: 7,
	_SubjectName[52:58]: 8,
	_SubjectName[58:71]: 9,
	_SubjectName[71:78]: 10,
	_SubjectName[78:85]: 11,
}

// ParseSubject attempts to convert a string to a Subject
func ParseSubject(name string) (Subject, error) {
	if x, ok := _SubjectValue[name]; ok {
		return x, nil
	}
	return Subject(0), fmt.Errorf("%s is not a valid Subject, try [%s]", name, strings.Join(_SubjectNames, ", "))
}

// MarshalText implements the text marshaller method
func (x Subject) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *Subject) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseSubject(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
