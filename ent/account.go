// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"telegra/ent/account"

	"entgo.io/ent/dialect/sql"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ShortName holds the value of the "short_name" field.
	ShortName string `json:"short_name,omitempty"`
	// AuthorName holds the value of the "author_name" field.
	AuthorName string `json:"author_name,omitempty"`
	// AuthorURL holds the value of the "author_url" field.
	AuthorURL string `json:"author_url,omitempty"`
	// AccessToken holds the value of the "access_token" field.
	AccessToken string `json:"access_token,omitempty"`
	// AuthURL holds the value of the "auth_url" field.
	AuthURL string `json:"auth_url,omitempty"`
	// PageCount holds the value of the "page_count" field.
	PageCount int `json:"page_count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges AccountEdges `json:"edges"`
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// Pages holds the value of the pages edge.
	Pages []*Page `json:"pages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PagesOrErr returns the Pages value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) PagesOrErr() ([]*Page, error) {
	if e.loadedTypes[0] {
		return e.Pages, nil
	}
	return nil, &NotLoadedError{edge: "pages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldID, account.FieldPageCount:
			values[i] = new(sql.NullInt64)
		case account.FieldShortName, account.FieldAuthorName, account.FieldAuthorURL, account.FieldAccessToken, account.FieldAuthURL:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Account", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case account.FieldShortName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field short_name", values[i])
			} else if value.Valid {
				a.ShortName = value.String
			}
		case account.FieldAuthorName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author_name", values[i])
			} else if value.Valid {
				a.AuthorName = value.String
			}
		case account.FieldAuthorURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author_url", values[i])
			} else if value.Valid {
				a.AuthorURL = value.String
			}
		case account.FieldAccessToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_token", values[i])
			} else if value.Valid {
				a.AccessToken = value.String
			}
		case account.FieldAuthURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field auth_url", values[i])
			} else if value.Valid {
				a.AuthURL = value.String
			}
		case account.FieldPageCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field page_count", values[i])
			} else if value.Valid {
				a.PageCount = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPages queries the "pages" edge of the Account entity.
func (a *Account) QueryPages() *PageQuery {
	return (&AccountClient{config: a.config}).QueryPages(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return (&AccountClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("short_name=")
	builder.WriteString(a.ShortName)
	builder.WriteString(", ")
	builder.WriteString("author_name=")
	builder.WriteString(a.AuthorName)
	builder.WriteString(", ")
	builder.WriteString("author_url=")
	builder.WriteString(a.AuthorURL)
	builder.WriteString(", ")
	builder.WriteString("access_token=")
	builder.WriteString(a.AccessToken)
	builder.WriteString(", ")
	builder.WriteString("auth_url=")
	builder.WriteString(a.AuthURL)
	builder.WriteString(", ")
	builder.WriteString("page_count=")
	builder.WriteString(fmt.Sprintf("%v", a.PageCount))
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account

func (a Accounts) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
