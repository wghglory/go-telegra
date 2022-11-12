package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.String("path"),
		field.String("url"),
		field.String("title"),
		field.String("description"),
		field.String("author_name"),
		field.String("image_url"),
		field.String("content"),
		field.Int("views").Default(0),
		field.Bool("can_edit").Default(false),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return nil
}
