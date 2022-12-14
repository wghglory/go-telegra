package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("short_name").NotEmpty().MinLen(1).MaxLen(32),
		field.String("author_name").MaxLen(128),
		field.String("author_url").MaxLen(512),
		field.String("access_token").NotEmpty(),
		field.String("auth_url").NotEmpty(),
		field.Int("page_count").Default(0),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		// 1 account : N pages
		edge.To("pages", Page.Type),
	}
}
