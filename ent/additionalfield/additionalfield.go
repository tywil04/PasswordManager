// Code generated by ent, DO NOT EDIT.

package additionalfield

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the additionalfield type in the database.
	Label = "additional_field"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldKeyIv holds the string denoting the keyiv field in the database.
	FieldKeyIv = "key_iv"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldValueIv holds the string denoting the valueiv field in the database.
	FieldValueIv = "value_iv"
	// EdgePassword holds the string denoting the password edge name in mutations.
	EdgePassword = "password"
	// Table holds the table name of the additionalfield in the database.
	Table = "additional_fields"
	// PasswordTable is the table that holds the password relation/edge.
	PasswordTable = "additional_fields"
	// PasswordInverseTable is the table name for the Password entity.
	// It exists in this package in order to avoid circular dependency with the "password" package.
	PasswordInverseTable = "passwords"
	// PasswordColumn is the table column denoting the password relation/edge.
	PasswordColumn = "password_additional_fields"
)

// Columns holds all SQL columns for additionalfield fields.
var Columns = []string{
	FieldID,
	FieldKey,
	FieldKeyIv,
	FieldValue,
	FieldValueIv,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "additional_fields"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"password_additional_fields",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// KeyValidator is a validator for the "key" field. It is called by the builders before save.
	KeyValidator func([]byte) error
	// KeyIvValidator is a validator for the "keyIv" field. It is called by the builders before save.
	KeyIvValidator func([]byte) error
	// ValueValidator is a validator for the "value" field. It is called by the builders before save.
	ValueValidator func([]byte) error
	// ValueIvValidator is a validator for the "valueIv" field. It is called by the builders before save.
	ValueIvValidator func([]byte) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
