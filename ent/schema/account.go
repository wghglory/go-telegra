package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("short_name"),
		field.String("author_name"),
		field.String("author_url"),
		field.String("access_token"),
		field.String("auth_url"),
		field.Int("page_count").Default(0),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
