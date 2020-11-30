//go:generate go-enum -f=$GOFILE --marshal --names --noprefix

package operation

// Operation is an enumeration of permission subject values
/*
ENUM(
None
Any
Create
Read
Update
Delete
)
*/
type Operation int64
