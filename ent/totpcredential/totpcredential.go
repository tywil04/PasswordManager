// Code generated by ent, DO NOT EDIT.

package totpcredential

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the totpcredential type in the database.
	Label = "totp_credential"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"
	// FieldValidated holds the string denoting the validated field in the database.
	FieldValidated = "validated"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeChallenge holds the string denoting the challenge edge name in mutations.
	EdgeChallenge = "challenge"
	// Table holds the table name of the totpcredential in the database.
	Table = "totp_credentials"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "totp_credentials"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_totp_credential"
	// ChallengeTable is the table that holds the challenge relation/edge.
	ChallengeTable = "totp_credentials"
	// ChallengeInverseTable is the table name for the Challenge entity.
	// It exists in this package in order to avoid circular dependency with the "challenge" package.
	ChallengeInverseTable = "challenges"
	// ChallengeColumn is the table column denoting the challenge relation/edge.
	ChallengeColumn = "challenge_totp_credential"
)

// Columns holds all SQL columns for totpcredential fields.
var Columns = []string{
	FieldID,
	FieldSecret,
	FieldValidated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "totp_credentials"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"challenge_totp_credential",
	"user_totp_credential",
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
	// SecretValidator is a validator for the "secret" field. It is called by the builders before save.
	SecretValidator func(string) error
	// DefaultValidated holds the default value on creation for the "validated" field.
	DefaultValidated bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
