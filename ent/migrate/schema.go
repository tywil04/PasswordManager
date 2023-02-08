// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdditionalFieldsColumns holds the columns for the "additional_fields" table.
	AdditionalFieldsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "key", Type: field.TypeBytes},
		{Name: "key_iv", Type: field.TypeBytes},
		{Name: "value", Type: field.TypeBytes},
		{Name: "value_iv", Type: field.TypeBytes},
		{Name: "password_additional_fields", Type: field.TypeUUID, Nullable: true},
	}
	// AdditionalFieldsTable holds the schema information for the "additional_fields" table.
	AdditionalFieldsTable = &schema.Table{
		Name:       "additional_fields",
		Columns:    AdditionalFieldsColumns,
		PrimaryKey: []*schema.Column{AdditionalFieldsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "additional_fields_passwords_additionalFields",
				Columns:    []*schema.Column{AdditionalFieldsColumns[5]},
				RefColumns: []*schema.Column{PasswordsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EmailChallengesColumns holds the columns for the "email_challenges" table.
	EmailChallengesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "code", Type: field.TypeString},
		{Name: "expiry", Type: field.TypeTime},
		{Name: "for", Type: field.TypeEnum, Enums: []string{"signup", "signin"}},
		{Name: "user_email_challenges", Type: field.TypeUUID},
	}
	// EmailChallengesTable holds the schema information for the "email_challenges" table.
	EmailChallengesTable = &schema.Table{
		Name:       "email_challenges",
		Columns:    EmailChallengesColumns,
		PrimaryKey: []*schema.Column{EmailChallengesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "email_challenges_users_emailChallenges",
				Columns:    []*schema.Column{EmailChallengesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PasswordsColumns holds the columns for the "passwords" table.
	PasswordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeBytes},
		{Name: "name_iv", Type: field.TypeBytes},
		{Name: "username", Type: field.TypeBytes},
		{Name: "username_iv", Type: field.TypeBytes},
		{Name: "password", Type: field.TypeBytes},
		{Name: "password_iv", Type: field.TypeBytes},
	}
	// PasswordsTable holds the schema information for the "passwords" table.
	PasswordsTable = &schema.Table{
		Name:       "passwords",
		Columns:    PasswordsColumns,
		PrimaryKey: []*schema.Column{PasswordsColumns[0]},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "n", Type: field.TypeBytes},
		{Name: "e", Type: field.TypeInt},
		{Name: "expiry", Type: field.TypeTime},
		{Name: "user_sessions", Type: field.TypeUUID, Nullable: true},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sessions_users_sessions",
				Columns:    []*schema.Column{SessionsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "strengthened_master_hash", Type: field.TypeBytes},
		{Name: "strengthened_master_hash_salt", Type: field.TypeBytes},
		{Name: "protected_database_key", Type: field.TypeBytes},
		{Name: "protected_database_key_iv", Type: field.TypeBytes},
		{Name: "default2fa", Type: field.TypeEnum, Enums: []string{"email", "webauthn"}, Default: "email"},
		{Name: "verified", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_email",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1]},
			},
		},
	}
	// WebAuthnChallengesColumns holds the columns for the "web_authn_challenges" table.
	WebAuthnChallengesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "challenge", Type: field.TypeString, Nullable: true},
		{Name: "user_id", Type: field.TypeBytes, Nullable: true},
		{Name: "allowed_credential_ids", Type: field.TypeJSON, Nullable: true},
		{Name: "user_verification", Type: field.TypeString, Nullable: true},
		{Name: "extensions", Type: field.TypeJSON, Nullable: true},
		{Name: "user_webauthn_challenges", Type: field.TypeUUID},
	}
	// WebAuthnChallengesTable holds the schema information for the "web_authn_challenges" table.
	WebAuthnChallengesTable = &schema.Table{
		Name:       "web_authn_challenges",
		Columns:    WebAuthnChallengesColumns,
		PrimaryKey: []*schema.Column{WebAuthnChallengesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "web_authn_challenges_users_webauthnChallenges",
				Columns:    []*schema.Column{WebAuthnChallengesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// WebAuthnCredentialsColumns holds the columns for the "web_authn_credentials" table.
	WebAuthnCredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "credential_id", Type: field.TypeBytes},
		{Name: "public_key", Type: field.TypeBytes},
		{Name: "attestation_type", Type: field.TypeString},
		{Name: "transport", Type: field.TypeJSON},
		{Name: "aaguid", Type: field.TypeBytes},
		{Name: "sign_count", Type: field.TypeUint32},
		{Name: "clone_warning", Type: field.TypeBool},
		{Name: "user_webauthn_credentials", Type: field.TypeUUID},
	}
	// WebAuthnCredentialsTable holds the schema information for the "web_authn_credentials" table.
	WebAuthnCredentialsTable = &schema.Table{
		Name:       "web_authn_credentials",
		Columns:    WebAuthnCredentialsColumns,
		PrimaryKey: []*schema.Column{WebAuthnCredentialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "web_authn_credentials_users_webauthnCredentials",
				Columns:    []*schema.Column{WebAuthnCredentialsColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdditionalFieldsTable,
		EmailChallengesTable,
		PasswordsTable,
		SessionsTable,
		UsersTable,
		WebAuthnChallengesTable,
		WebAuthnCredentialsTable,
	}
)

func init() {
	AdditionalFieldsTable.ForeignKeys[0].RefTable = PasswordsTable
	EmailChallengesTable.ForeignKeys[0].RefTable = UsersTable
	SessionsTable.ForeignKeys[0].RefTable = UsersTable
	WebAuthnChallengesTable.ForeignKeys[0].RefTable = UsersTable
	WebAuthnCredentialsTable.ForeignKeys[0].RefTable = UsersTable
}