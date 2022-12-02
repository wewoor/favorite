package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RelUserProject holds the schema definition for the RelUserProject entity.
type RelUserProject struct {
	ent.Schema
}

// Fields of the RelUserProject.
func (RelUserProject) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("uuid"),
		field.String("project"),
		field.String("user"),
	}
}

// Edges of the RelUserProject.
func (RelUserProject) Edges() []ent.Edge {
	return nil
}
