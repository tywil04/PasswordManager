// Code generated by ent, DO NOT EDIT.

package vault

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the vault type in the database.
	Label = "vault"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldColour holds the string denoting the colour field in the database.
	FieldColour = "colour"
	// EdgePasswords holds the string denoting the passwords edge name in mutations.
	EdgePasswords = "passwords"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the vault in the database.
	Table = "vaults"
	// PasswordsTable is the table that holds the passwords relation/edge.
	PasswordsTable = "passwords"
	// PasswordsInverseTable is the table name for the Password entity.
	// It exists in this package in order to avoid circular dependency with the "password" package.
	PasswordsInverseTable = "passwords"
	// PasswordsColumn is the table column denoting the passwords relation/edge.
	PasswordsColumn = "vault_passwords"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "vaults"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_vaults"
)

// Columns holds all SQL columns for vault fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldName,
	FieldColour,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "vaults"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_vaults",
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
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)