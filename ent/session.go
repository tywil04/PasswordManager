// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/session"
	"PasswordManager/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// N holds the value of the "n" field.
	N []byte `json:"-"`
	// E holds the value of the "e" field.
	E int `json:"e,omitempty"`
	// Expiry holds the value of the "expiry" field.
	Expiry time.Time `json:"expiry,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SessionQuery when eager-loading is set.
	Edges         SessionEdges `json:"edges"`
	user_sessions *uuid.UUID
}

// SessionEdges holds the relations/edges for other nodes in the graph.
type SessionEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SessionEdges) UserOrErr() (*User, error) {
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
func (*Session) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case session.FieldN:
			values[i] = new([]byte)
		case session.FieldE:
			values[i] = new(sql.NullInt64)
		case session.FieldExpiry:
			values[i] = new(sql.NullTime)
		case session.FieldID:
			values[i] = new(uuid.UUID)
		case session.ForeignKeys[0]: // user_sessions
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Session", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case session.FieldN:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field n", values[i])
			} else if value != nil {
				s.N = *value
			}
		case session.FieldE:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field e", values[i])
			} else if value.Valid {
				s.E = int(value.Int64)
			}
		case session.FieldExpiry:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiry", values[i])
			} else if value.Valid {
				s.Expiry = value.Time
			}
		case session.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_sessions", values[i])
			} else if value.Valid {
				s.user_sessions = new(uuid.UUID)
				*s.user_sessions = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Session entity.
func (s *Session) QueryUser() *UserQuery {
	return NewSessionClient(s.config).QueryUser(s)
}

// Update returns a builder for updating this Session.
// Note that you need to call Session.Unwrap() before calling this method if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return NewSessionClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Session entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Session is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("n=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("e=")
	builder.WriteString(fmt.Sprintf("%v", s.E))
	builder.WriteString(", ")
	builder.WriteString("expiry=")
	builder.WriteString(s.Expiry.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session
