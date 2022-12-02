package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("uuid"),
		field.String("name"),
		field.String("repository"),
		field.String("url"),
		field.String("domain").Nillable(),
		field.String("storeLocation"),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return nil
}
