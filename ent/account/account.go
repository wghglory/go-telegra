// Code generated by ent, DO NOT EDIT.

package account

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldShortName holds the string denoting the short_name field in the database.
	FieldShortName = "short_name"
	// FieldAuthorName holds the string denoting the author_name field in the database.
	FieldAuthorName = "author_name"
	// FieldAuthorURL holds the string denoting the author_url field in the database.
	FieldAuthorURL = "author_url"
	// FieldAccessToken holds the string denoting the access_token field in the database.
	FieldAccessToken = "access_token"
	// FieldAuthURL holds the string denoting the auth_url field in the database.
	FieldAuthURL = "auth_url"
	// FieldPageCount holds the string denoting the page_count field in the database.
	FieldPageCount = "page_count"
	// EdgePages holds the string denoting the pages edge name in mutations.
	EdgePages = "pages"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// PagesTable is the table that holds the pages relation/edge.
	PagesTable = "pages"
	// PagesInverseTable is the table name for the Page entity.
	// It exists in this package in order to avoid circular dependency with the "page" package.
	PagesInverseTable = "pages"
	// PagesColumn is the table column denoting the pages relation/edge.
	PagesColumn = "account_pages"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldShortName,
	FieldAuthorName,
	FieldAuthorURL,
	FieldAccessToken,
	FieldAuthURL,
	FieldPageCount,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// ShortNameValidator is a validator for the "short_name" field. It is called by the builders before save.
	ShortNameValidator func(string) error
	// AuthorNameValidator is a validator for the "author_name" field. It is called by the builders before save.
	AuthorNameValidator func(string) error
	// AuthorURLValidator is a validator for the "author_url" field. It is called by the builders before save.
	AuthorURLValidator func(string) error
	// AccessTokenValidator is a validator for the "access_token" field. It is called by the builders before save.
	AccessTokenValidator func(string) error
	// AuthURLValidator is a validator for the "auth_url" field. It is called by the builders before save.
	AuthURLValidator func(string) error
	// DefaultPageCount holds the default value on creation for the "page_count" field.
	DefaultPageCount int
)
