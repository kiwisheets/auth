//go:generate go-enum -f=$GOFILE --marshal --names --noprefix

package subject

// Subject is an enumeration of permission subject values
/*
ENUM(
None
Any

Me

User
UserContact
OtherUser

Company
OtherCompany

Client
ClientContact

Contact

Invoice
)
*/
type Subject int64
