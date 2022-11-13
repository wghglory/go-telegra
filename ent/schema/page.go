package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").NotEmpty(),
		field.String("url").NotEmpty(),
		field.String("title").NotEmpty().MinLen(1).MaxLen(256),
		field.String("description").Optional(),
		field.String("author_name").Optional().MaxLen(128),
		field.String("author_url").Optional().MaxLen(512),
		field.String("image_url").Optional(),
		field.String("content").NotEmpty(),
		field.Int("views").Default(0),
		field.Bool("can_edit").Default(false),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
		// and reference it to the "cars" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("author", Account.Type).
			Ref("pages").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
	}
}
