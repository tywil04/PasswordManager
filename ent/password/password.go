// Code generated by ent, DO NOT EDIT.

package password

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the password type in the database.
	Label = "password"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNameIv holds the string denoting the nameiv field in the database.
	FieldNameIv = "name_iv"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldUsernameIv holds the string denoting the usernameiv field in the database.
	FieldUsernameIv = "username_iv"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldPasswordIv holds the string denoting the passwordiv field in the database.
	FieldPasswordIv = "password_iv"
	// FieldEmoji holds the string denoting the emoji field in the database.
	FieldEmoji = "emoji"
	// EdgeAdditionalFields holds the string denoting the additionalfields edge name in mutations.
	EdgeAdditionalFields = "additionalFields"
	// EdgeUrls holds the string denoting the urls edge name in mutations.
	EdgeUrls = "urls"
	// Table holds the table name of the password in the database.
	Table = "passwords"
	// AdditionalFieldsTable is the table that holds the additionalFields relation/edge.
	AdditionalFieldsTable = "additional_fields"
	// AdditionalFieldsInverseTable is the table name for the AdditionalField entity.
	// It exists in this package in order to avoid circular dependency with the "additionalfield" package.
	AdditionalFieldsInverseTable = "additional_fields"
	// AdditionalFieldsColumn is the table column denoting the additionalFields relation/edge.
	AdditionalFieldsColumn = "password_additional_fields"
	// UrlsTable is the table that holds the urls relation/edge.
	UrlsTable = "urls"
	// UrlsInverseTable is the table name for the Url entity.
	// It exists in this package in order to avoid circular dependency with the "url" package.
	UrlsInverseTable = "urls"
	// UrlsColumn is the table column denoting the urls relation/edge.
	UrlsColumn = "password_urls"
)

// Columns holds all SQL columns for password fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldNameIv,
	FieldUsername,
	FieldUsernameIv,
	FieldPassword,
	FieldPasswordIv,
	FieldEmoji,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func([]byte) error
	// NameIvValidator is a validator for the "nameIv" field. It is called by the builders before save.
	NameIvValidator func([]byte) error
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func([]byte) error
	// UsernameIvValidator is a validator for the "usernameIv" field. It is called by the builders before save.
	UsernameIvValidator func([]byte) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func([]byte) error
	// PasswordIvValidator is a validator for the "passwordIv" field. It is called by the builders before save.
	PasswordIvValidator func([]byte) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
