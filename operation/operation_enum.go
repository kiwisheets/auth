// Code generated by go-enum
// DO NOT EDIT!

package operation

import (
	"fmt"
	"strings"
)

const (
	// None is a Operation of type None
	None Operation = iota
	// Any is a Operation of type Any
	Any
	// Create is a Operation of type Create
	Create
	// Read is a Operation of type Read
	Read
	// Update is a Operation of type Update
	Update
	// Delete is a Operation of type Delete
	Delete
)

const _OperationName = "NoneAnyCreateReadUpdateDelete"

var _OperationNames = []string{
	_OperationName[0:4],
	_OperationName[4:7],
	_OperationName[7:13],
	_OperationName[13:17],
	_OperationName[17:23],
	_OperationName[23:29],
}

// OperationNames returns a list of possible string values of Operation.
func OperationNames() []string {
	tmp := make([]string, len(_OperationNames))
	copy(tmp, _OperationNames)
	return tmp
}

var _OperationMap = map[Operation]string{
	0: _OperationName[0:4],
	1: _OperationName[4:7],
	2: _OperationName[7:13],
	3: _OperationName[13:17],
	4: _OperationName[17:23],
	5: _OperationName[23:29],
}

// String implements the Stringer interface.
func (x Operation) String() string {
	if str, ok := _OperationMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Operation(%d)", x)
}

var _OperationValue = map[string]Operation{
	_OperationName[0:4]:   0,
	_OperationName[4:7]:   1,
	_OperationName[7:13]:  2,
	_OperationName[13:17]: 3,
	_OperationName[17:23]: 4,
	_OperationName[23:29]: 5,
}

// ParseOperation attempts to convert a string to a Operation
func ParseOperation(name string) (Operation, error) {
	if x, ok := _OperationValue[name]; ok {
		return x, nil
	}
	return Operation(0), fmt.Errorf("%s is not a valid Operation, try [%s]", name, strings.Join(_OperationNames, ", "))
}

// MarshalText implements the text marshaller method
func (x Operation) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *Operation) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseOperation(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
