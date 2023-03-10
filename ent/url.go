// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/password"
	"PasswordManager/ent/url"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Url is the model entity for the Url schema.
type Url struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `form:"id" json:"id" xml:"id"`
	// URL holds the value of the "url" field.
	URL []byte `json:"url,omitempty"`
	// UrlIv holds the value of the "urlIv" field.
	UrlIv []byte `json:"urlIv,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UrlQuery when eager-loading is set.
	Edges         UrlEdges `json:"edges"`
	password_urls *uuid.UUID
}

// UrlEdges holds the relations/edges for other nodes in the graph.
type UrlEdges struct {
	// Password holds the value of the password edge.
	Password *Password `json:"password,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PasswordOrErr returns the Password value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UrlEdges) PasswordOrErr() (*Password, error) {
	if e.loadedTypes[0] {
		if e.Password == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: password.Label}
		}
		return e.Password, nil
	}
	return nil, &NotLoadedError{edge: "password"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Url) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case url.FieldURL, url.FieldUrlIv:
			values[i] = new([]byte)
		case url.FieldID:
			values[i] = new(uuid.UUID)
		case url.ForeignKeys[0]: // password_urls
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Url", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Url fields.
func (u *Url) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case url.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case url.FieldURL:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value != nil {
				u.URL = *value
			}
		case url.FieldUrlIv:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field urlIv", values[i])
			} else if value != nil {
				u.UrlIv = *value
			}
		case url.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field password_urls", values[i])
			} else if value.Valid {
				u.password_urls = new(uuid.UUID)
				*u.password_urls = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryPassword queries the "password" edge of the Url entity.
func (u *Url) QueryPassword() *PasswordQuery {
	return NewURLClient(u.config).QueryPassword(u)
}

// Update returns a builder for updating this Url.
// Note that you need to call Url.Unwrap() before calling this method if this Url
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Url) Update() *URLUpdateOne {
	return NewURLClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the Url entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Url) Unwrap() *Url {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Url is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Url) String() string {
	var builder strings.Builder
	builder.WriteString("Url(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("url=")
	builder.WriteString(fmt.Sprintf("%v", u.URL))
	builder.WriteString(", ")
	builder.WriteString("urlIv=")
	builder.WriteString(fmt.Sprintf("%v", u.UrlIv))
	builder.WriteByte(')')
	return builder.String()
}

// Urls is a parsable slice of Url.
type Urls []*Url
