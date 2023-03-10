// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/user"
	"PasswordManager/ent/webauthnregisterchallenge"
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// WebAuthnRegisterChallenge is the model entity for the WebAuthnRegisterChallenge schema.
type WebAuthnRegisterChallenge struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// SdChallenge holds the value of the "sdChallenge" field.
	SdChallenge string `json:"sdChallenge,omitempty"`
	// UserId holds the value of the "userId" field.
	UserId []byte `json:"userId,omitempty"`
	// AllowedCredentialIds holds the value of the "allowedCredentialIds" field.
	AllowedCredentialIds [][]uint8 `json:"allowedCredentialIds,omitempty"`
	// UserVerification holds the value of the "userVerification" field.
	UserVerification string `json:"userVerification,omitempty"`
	// Extensions holds the value of the "extensions" field.
	Extensions map[string]interface{} `json:"extensions,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WebAuthnRegisterChallengeQuery when eager-loading is set.
	Edges                             WebAuthnRegisterChallengeEdges `json:"edges"`
	user_webauthn_register_challenges *uuid.UUID
}

// WebAuthnRegisterChallengeEdges holds the relations/edges for other nodes in the graph.
type WebAuthnRegisterChallengeEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WebAuthnRegisterChallengeEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WebAuthnRegisterChallenge) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case webauthnregisterchallenge.FieldUserId, webauthnregisterchallenge.FieldAllowedCredentialIds, webauthnregisterchallenge.FieldExtensions:
			values[i] = new([]byte)
		case webauthnregisterchallenge.FieldSdChallenge, webauthnregisterchallenge.FieldUserVerification:
			values[i] = new(sql.NullString)
		case webauthnregisterchallenge.FieldID:
			values[i] = new(uuid.UUID)
		case webauthnregisterchallenge.ForeignKeys[0]: // user_webauthn_register_challenges
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type WebAuthnRegisterChallenge", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WebAuthnRegisterChallenge fields.
func (warc *WebAuthnRegisterChallenge) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case webauthnregisterchallenge.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				warc.ID = *value
			}
		case webauthnregisterchallenge.FieldSdChallenge:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sdChallenge", values[i])
			} else if value.Valid {
				warc.SdChallenge = value.String
			}
		case webauthnregisterchallenge.FieldUserId:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value != nil {
				warc.UserId = *value
			}
		case webauthnregisterchallenge.FieldAllowedCredentialIds:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field allowedCredentialIds", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &warc.AllowedCredentialIds); err != nil {
					return fmt.Errorf("unmarshal field allowedCredentialIds: %w", err)
				}
			}
		case webauthnregisterchallenge.FieldUserVerification:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userVerification", values[i])
			} else if value.Valid {
				warc.UserVerification = value.String
			}
		case webauthnregisterchallenge.FieldExtensions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field extensions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &warc.Extensions); err != nil {
					return fmt.Errorf("unmarshal field extensions: %w", err)
				}
			}
		case webauthnregisterchallenge.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_webauthn_register_challenges", values[i])
			} else if value.Valid {
				warc.user_webauthn_register_challenges = new(uuid.UUID)
				*warc.user_webauthn_register_challenges = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the WebAuthnRegisterChallenge entity.
func (warc *WebAuthnRegisterChallenge) QueryUser() *UserQuery {
	return NewWebAuthnRegisterChallengeClient(warc.config).QueryUser(warc)
}

// Update returns a builder for updating this WebAuthnRegisterChallenge.
// Note that you need to call WebAuthnRegisterChallenge.Unwrap() before calling this method if this WebAuthnRegisterChallenge
// was returned from a transaction, and the transaction was committed or rolled back.
func (warc *WebAuthnRegisterChallenge) Update() *WebAuthnRegisterChallengeUpdateOne {
	return NewWebAuthnRegisterChallengeClient(warc.config).UpdateOne(warc)
}

// Unwrap unwraps the WebAuthnRegisterChallenge entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (warc *WebAuthnRegisterChallenge) Unwrap() *WebAuthnRegisterChallenge {
	_tx, ok := warc.config.driver.(*txDriver)
	if !ok {
		panic("ent: WebAuthnRegisterChallenge is not a transactional entity")
	}
	warc.config.driver = _tx.drv
	return warc
}

// String implements the fmt.Stringer.
func (warc *WebAuthnRegisterChallenge) String() string {
	var builder strings.Builder
	builder.WriteString("WebAuthnRegisterChallenge(")
	builder.WriteString(fmt.Sprintf("id=%v, ", warc.ID))
	builder.WriteString("sdChallenge=")
	builder.WriteString(warc.SdChallenge)
	builder.WriteString(", ")
	builder.WriteString("userId=")
	builder.WriteString(fmt.Sprintf("%v", warc.UserId))
	builder.WriteString(", ")
	builder.WriteString("allowedCredentialIds=")
	builder.WriteString(fmt.Sprintf("%v", warc.AllowedCredentialIds))
	builder.WriteString(", ")
	builder.WriteString("userVerification=")
	builder.WriteString(warc.UserVerification)
	builder.WriteString(", ")
	builder.WriteString("extensions=")
	builder.WriteString(fmt.Sprintf("%v", warc.Extensions))
	builder.WriteByte(')')
	return builder.String()
}

// WebAuthnRegisterChallenges is a parsable slice of WebAuthnRegisterChallenge.
type WebAuthnRegisterChallenges []*WebAuthnRegisterChallenge
