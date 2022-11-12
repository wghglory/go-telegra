// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "short_name", Type: field.TypeString},
		{Name: "author_name", Type: field.TypeString},
		{Name: "author_url", Type: field.TypeString},
		{Name: "access_token", Type: field.TypeString},
		{Name: "auth_url", Type: field.TypeString},
		{Name: "page_count", Type: field.TypeInt, Default: 0},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
	}
	// PagesColumns holds the columns for the "pages" table.
	PagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "path", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "author_name", Type: field.TypeString},
		{Name: "image_url", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "views", Type: field.TypeInt, Default: 0},
		{Name: "can_edit", Type: field.TypeBool, Default: false},
	}
	// PagesTable holds the schema information for the "pages" table.
	PagesTable = &schema.Table{
		Name:       "pages",
		Columns:    PagesColumns,
		PrimaryKey: []*schema.Column{PagesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		PagesTable,
	}
)

func init() {
}