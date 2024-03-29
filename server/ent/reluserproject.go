// Code generated by ent, DO NOT EDIT.

package ent

import (
	"favorite/ent/reluserproject"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// RelUserProject is the model entity for the RelUserProject schema.
type RelUserProject struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty"`
	// Project holds the value of the "project" field.
	Project string `json:"project,omitempty"`
	// User holds the value of the "user" field.
	User string `json:"user,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RelUserProject) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case reluserproject.FieldID, reluserproject.FieldUUID, reluserproject.FieldProject, reluserproject.FieldUser:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RelUserProject", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RelUserProject fields.
func (rup *RelUserProject) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case reluserproject.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				rup.ID = value.String
			}
		case reluserproject.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				rup.UUID = value.String
			}
		case reluserproject.FieldProject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field project", values[i])
			} else if value.Valid {
				rup.Project = value.String
			}
		case reluserproject.FieldUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user", values[i])
			} else if value.Valid {
				rup.User = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this RelUserProject.
// Note that you need to call RelUserProject.Unwrap() before calling this method if this RelUserProject
// was returned from a transaction, and the transaction was committed or rolled back.
func (rup *RelUserProject) Update() *RelUserProjectUpdateOne {
	return (&RelUserProjectClient{config: rup.config}).UpdateOne(rup)
}

// Unwrap unwraps the RelUserProject entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rup *RelUserProject) Unwrap() *RelUserProject {
	_tx, ok := rup.config.driver.(*txDriver)
	if !ok {
		panic("ent: RelUserProject is not a transactional entity")
	}
	rup.config.driver = _tx.drv
	return rup
}

// String implements the fmt.Stringer.
func (rup *RelUserProject) String() string {
	var builder strings.Builder
	builder.WriteString("RelUserProject(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rup.ID))
	builder.WriteString("uuid=")
	builder.WriteString(rup.UUID)
	builder.WriteString(", ")
	builder.WriteString("project=")
	builder.WriteString(rup.Project)
	builder.WriteString(", ")
	builder.WriteString("user=")
	builder.WriteString(rup.User)
	builder.WriteByte(')')
	return builder.String()
}

// RelUserProjects is a parsable slice of RelUserProject.
type RelUserProjects []*RelUserProject

func (rup RelUserProjects) config(cfg config) {
	for _i := range rup {
		rup[_i].config = cfg
	}
}
