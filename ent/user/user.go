// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldStrengthenedMasterHash holds the string denoting the strengthenedmasterhash field in the database.
	FieldStrengthenedMasterHash = "strengthened_master_hash"
	// FieldStrengthenedMasterHashSalt holds the string denoting the strengthenedmasterhashsalt field in the database.
	FieldStrengthenedMasterHashSalt = "strengthened_master_hash_salt"
	// FieldProtectedDatabaseKey holds the string denoting the protecteddatabasekey field in the database.
	FieldProtectedDatabaseKey = "protected_database_key"
	// FieldProtectedDatabaseKeyIv holds the string denoting the protecteddatabasekeyiv field in the database.
	FieldProtectedDatabaseKeyIv = "protected_database_key_iv"
	// FieldDefault2FA holds the string denoting the default2fa field in the database.
	FieldDefault2FA = "default2fa"
	// FieldVerified holds the string denoting the verified field in the database.
	FieldVerified = "verified"
	// EdgeEmailChallenges holds the string denoting the emailchallenges edge name in mutations.
	EdgeEmailChallenges = "emailChallenges"
	// EdgeTotpCredential holds the string denoting the totpcredential edge name in mutations.
	EdgeTotpCredential = "totpCredential"
	// EdgeWebauthnCredentials holds the string denoting the webauthncredentials edge name in mutations.
	EdgeWebauthnCredentials = "webauthnCredentials"
	// EdgeWebauthnChallenges holds the string denoting the webauthnchallenges edge name in mutations.
	EdgeWebauthnChallenges = "webauthnChallenges"
	// EdgePasswords holds the string denoting the passwords edge name in mutations.
	EdgePasswords = "passwords"
	// EdgeSessions holds the string denoting the sessions edge name in mutations.
	EdgeSessions = "sessions"
	// Table holds the table name of the user in the database.
	Table = "users"
	// EmailChallengesTable is the table that holds the emailChallenges relation/edge.
	EmailChallengesTable = "email_challenges"
	// EmailChallengesInverseTable is the table name for the EmailChallenge entity.
	// It exists in this package in order to avoid circular dependency with the "emailchallenge" package.
	EmailChallengesInverseTable = "email_challenges"
	// EmailChallengesColumn is the table column denoting the emailChallenges relation/edge.
	EmailChallengesColumn = "user_email_challenges"
	// TotpCredentialTable is the table that holds the totpCredential relation/edge.
	TotpCredentialTable = "totp_credentials"
	// TotpCredentialInverseTable is the table name for the TotpCredential entity.
	// It exists in this package in order to avoid circular dependency with the "totpcredential" package.
	TotpCredentialInverseTable = "totp_credentials"
	// TotpCredentialColumn is the table column denoting the totpCredential relation/edge.
	TotpCredentialColumn = "user_totp_credential"
	// WebauthnCredentialsTable is the table that holds the webauthnCredentials relation/edge.
	WebauthnCredentialsTable = "web_authn_credentials"
	// WebauthnCredentialsInverseTable is the table name for the WebAuthnCredential entity.
	// It exists in this package in order to avoid circular dependency with the "webauthncredential" package.
	WebauthnCredentialsInverseTable = "web_authn_credentials"
	// WebauthnCredentialsColumn is the table column denoting the webauthnCredentials relation/edge.
	WebauthnCredentialsColumn = "user_webauthn_credentials"
	// WebauthnChallengesTable is the table that holds the webauthnChallenges relation/edge.
	WebauthnChallengesTable = "web_authn_challenges"
	// WebauthnChallengesInverseTable is the table name for the WebAuthnChallenge entity.
	// It exists in this package in order to avoid circular dependency with the "webauthnchallenge" package.
	WebauthnChallengesInverseTable = "web_authn_challenges"
	// WebauthnChallengesColumn is the table column denoting the webauthnChallenges relation/edge.
	WebauthnChallengesColumn = "user_webauthn_challenges"
	// PasswordsTable is the table that holds the passwords relation/edge.
	PasswordsTable = "passwords"
	// PasswordsInverseTable is the table name for the Password entity.
	// It exists in this package in order to avoid circular dependency with the "password" package.
	PasswordsInverseTable = "passwords"
	// PasswordsColumn is the table column denoting the passwords relation/edge.
	PasswordsColumn = "user_passwords"
	// SessionsTable is the table that holds the sessions relation/edge.
	SessionsTable = "sessions"
	// SessionsInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionsInverseTable = "sessions"
	// SessionsColumn is the table column denoting the sessions relation/edge.
	SessionsColumn = "user_sessions"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldStrengthenedMasterHash,
	FieldStrengthenedMasterHashSalt,
	FieldProtectedDatabaseKey,
	FieldProtectedDatabaseKeyIv,
	FieldDefault2FA,
	FieldVerified,
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
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// StrengthenedMasterHashValidator is a validator for the "strengthenedMasterHash" field. It is called by the builders before save.
	StrengthenedMasterHashValidator func([]byte) error
	// StrengthenedMasterHashSaltValidator is a validator for the "strengthenedMasterHashSalt" field. It is called by the builders before save.
	StrengthenedMasterHashSaltValidator func([]byte) error
	// ProtectedDatabaseKeyValidator is a validator for the "protectedDatabaseKey" field. It is called by the builders before save.
	ProtectedDatabaseKeyValidator func([]byte) error
	// ProtectedDatabaseKeyIvValidator is a validator for the "protectedDatabaseKeyIv" field. It is called by the builders before save.
	ProtectedDatabaseKeyIvValidator func([]byte) error
	// DefaultVerified holds the default value on creation for the "verified" field.
	DefaultVerified bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Default2FA defines the type for the "default2FA" enum field.
type Default2FA string

// Default2FAEmail is the default value of the Default2FA enum.
const DefaultDefault2FA = Default2FAEmail

// Default2FA values.
const (
	Default2FAEmail    Default2FA = "email"
	Default2FAWebauthn Default2FA = "webauthn"
	Default2FATotp     Default2FA = "totp"
)

func (d Default2FA) String() string {
	return string(d)
}

// Default2FAValidator is a validator for the "default2FA" field enum values. It is called by the builders before save.
func Default2FAValidator(d Default2FA) error {
	switch d {
	case Default2FAEmail, Default2FAWebauthn, Default2FATotp:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for default2FA field: %q", d)
	}
}
