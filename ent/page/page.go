// Code generated by ent, DO NOT EDIT.

package page

const (
	// Label holds the string label denoting the page type in the database.
	Label = "page"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldAuthorName holds the string denoting the author_name field in the database.
	FieldAuthorName = "author_name"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldViews holds the string denoting the views field in the database.
	FieldViews = "views"
	// FieldCanEdit holds the string denoting the can_edit field in the database.
	FieldCanEdit = "can_edit"
	// Table holds the table name of the page in the database.
	Table = "pages"
)

// Columns holds all SQL columns for page fields.
var Columns = []string{
	FieldID,
	FieldPath,
	FieldURL,
	FieldTitle,
	FieldDescription,
	FieldAuthorName,
	FieldImageURL,
	FieldContent,
	FieldViews,
	FieldCanEdit,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "pages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"account_pages",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultViews holds the default value on creation for the "views" field.
	DefaultViews int
	// DefaultCanEdit holds the default value on creation for the "can_edit" field.
	DefaultCanEdit bool
)
