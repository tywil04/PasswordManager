// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/additionalfield"
	"PasswordManager/ent/password"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// AdditionalField is the model entity for the AdditionalField schema.
type AdditionalField struct {
	config `form:"-" json:"-" xml:"-"`
	// ID of the ent.
	ID uuid.UUID `form:"id" json:"id" xml:"id"`
	// Key holds the value of the "key" field.
	Key []byte `form:"key" json:"key" xml:"key"`
	// KeyIv holds the value of the "keyIv" field.
	KeyIv []byte `form:"keyIv" json:"keyIv" xml:"keyIv"`
	// Value holds the value of the "value" field.
	Value []byte `form:"value" json:"value" xml:"value"`
	// ValueIv holds the value of the "valueIv" field.
	ValueIv []byte `form:"valueIv" json:"valueIv" xml:"valueIv"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AdditionalFieldQuery when eager-loading is set.
	Edges                      AdditionalFieldEdges `json:"edges"`
	password_additional_fields *uuid.UUID
}

// AdditionalFieldEdges holds the relations/edges for other nodes in the graph.
type AdditionalFieldEdges struct {
	// Password holds the value of the password edge.
	Password *Password `json:"password,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PasswordOrErr returns the Password value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdditionalFieldEdges) PasswordOrErr() (*Password, error) {
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
func (*AdditionalField) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case additionalfield.FieldKey, additionalfield.FieldKeyIv, additionalfield.FieldValue, additionalfield.FieldValueIv:
			values[i] = new([]byte)
		case additionalfield.FieldID:
			values[i] = new(uuid.UUID)
		case additionalfield.ForeignKeys[0]: // password_additional_fields
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type AdditionalField", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AdditionalField fields.
func (af *AdditionalField) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case additionalfield.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				af.ID = *value
			}
		case additionalfield.FieldKey:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value != nil {
				af.Key = *value
			}
		case additionalfield.FieldKeyIv:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field keyIv", values[i])
			} else if value != nil {
				af.KeyIv = *value
			}
		case additionalfield.FieldValue:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value != nil {
				af.Value = *value
			}
		case additionalfield.FieldValueIv:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field valueIv", values[i])
			} else if value != nil {
				af.ValueIv = *value
			}
		case additionalfield.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field password_additional_fields", values[i])
			} else if value.Valid {
				af.password_additional_fields = new(uuid.UUID)
				*af.password_additional_fields = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryPassword queries the "password" edge of the AdditionalField entity.
func (af *AdditionalField) QueryPassword() *PasswordQuery {
	return NewAdditionalFieldClient(af.config).QueryPassword(af)
}

// Update returns a builder for updating this AdditionalField.
// Note that you need to call AdditionalField.Unwrap() before calling this method if this AdditionalField
// was returned from a transaction, and the transaction was committed or rolled back.
func (af *AdditionalField) Update() *AdditionalFieldUpdateOne {
	return NewAdditionalFieldClient(af.config).UpdateOne(af)
}

// Unwrap unwraps the AdditionalField entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (af *AdditionalField) Unwrap() *AdditionalField {
	_tx, ok := af.config.driver.(*txDriver)
	if !ok {
		panic("ent: AdditionalField is not a transactional entity")
	}
	af.config.driver = _tx.drv
	return af
}

// String implements the fmt.Stringer.
func (af *AdditionalField) String() string {
	var builder strings.Builder
	builder.WriteString("AdditionalField(")
	builder.WriteString(fmt.Sprintf("id=%v, ", af.ID))
	builder.WriteString("key=")
	builder.WriteString(fmt.Sprintf("%v", af.Key))
	builder.WriteString(", ")
	builder.WriteString("keyIv=")
	builder.WriteString(fmt.Sprintf("%v", af.KeyIv))
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", af.Value))
	builder.WriteString(", ")
	builder.WriteString("valueIv=")
	builder.WriteString(fmt.Sprintf("%v", af.ValueIv))
	builder.WriteByte(')')
	return builder.String()
}

// AdditionalFields is a parsable slice of AdditionalField.
type AdditionalFields []*AdditionalField
