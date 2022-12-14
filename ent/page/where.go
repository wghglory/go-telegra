// Code generated by ent, DO NOT EDIT.

package page

import (
	"telegra/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// AuthorName applies equality check predicate on the "author_name" field. It's identical to AuthorNameEQ.
func AuthorName(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthorName), v))
	})
}

// AuthorURL applies equality check predicate on the "author_url" field. It's identical to AuthorURLEQ.
func AuthorURL(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthorURL), v))
	})
}

// ImageURL applies equality check predicate on the "image_url" field. It's identical to ImageURLEQ.
func ImageURL(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImageURL), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// Views applies equality check predicate on the "views" field. It's identical to ViewsEQ.
func Views(v int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldViews), v))
	})
}

// CanEdit applies equality check predicate on the "can_edit" field. It's identical to CanEditEQ.
func CanEdit(v bool) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCanEdit), v))
	})
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPath), v))
	})
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPath), v))
	})
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPath), v...))
	})
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPath), v...))
	})
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPath), v))
	})
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPath), v))
	})
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPath), v))
	})
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPath), v))
	})
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPath), v))
	})
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPath), v))
	})
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPath), v))
	})
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPath), v))
	})
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPath), v))
	})
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURL), v))
	})
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldURL), v...))
	})
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldURL), v...))
	})
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURL), v))
	})
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURL), v))
	})
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURL), v))
	})
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURL), v))
	})
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURL), v))
	})
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURL), v))
	})
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURL), v))
	})
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURL), v))
	})
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURL), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// AuthorNameEQ applies the EQ predicate on the "author_name" field.
func AuthorNameEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthorName), v))
	})
}

// AuthorNameNEQ applies the NEQ predicate on the "author_name" field.
func AuthorNameNEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthorName), v))
	})
}

// AuthorNameIn applies the In predicate on the "author_name" field.
func AuthorNameIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAuthorName), v...))
	})
}

// AuthorNameNotIn applies the NotIn predicate on the "author_name" field.
func AuthorNameNotIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAuthorName), v...))
	})
}

// AuthorNameGT applies the GT predicate on the "author_name" field.
func AuthorNameGT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAuthorName), v))
	})
}

// AuthorNameGTE applies the GTE predicate on the "author_name" field.
func AuthorNameGTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAuthorName), v))
	})
}

// AuthorNameLT applies the LT predicate on the "author_name" field.
func AuthorNameLT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAuthorName), v))
	})
}

// AuthorNameLTE applies the LTE predicate on the "author_name" field.
func AuthorNameLTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAuthorName), v))
	})
}

// AuthorNameContains applies the Contains predicate on the "author_name" field.
func AuthorNameContains(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAuthorName), v))
	})
}

// AuthorNameHasPrefix applies the HasPrefix predicate on the "author_name" field.
func AuthorNameHasPrefix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAuthorName), v))
	})
}

// AuthorNameHasSuffix applies the HasSuffix predicate on the "author_name" field.
func AuthorNameHasSuffix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAuthorName), v))
	})
}

// AuthorNameIsNil applies the IsNil predicate on the "author_name" field.
func AuthorNameIsNil() predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAuthorName)))
	})
}

// AuthorNameNotNil applies the NotNil predicate on the "author_name" field.
func AuthorNameNotNil() predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAuthorName)))
	})
}

// AuthorNameEqualFold applies the EqualFold predicate on the "author_name" field.
func AuthorNameEqualFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAuthorName), v))
	})
}

// AuthorNameContainsFold applies the ContainsFold predicate on the "author_name" field.
func AuthorNameContainsFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAuthorName), v))
	})
}

// AuthorURLEQ applies the EQ predicate on the "author_url" field.
func AuthorURLEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthorURL), v))
	})
}

// AuthorURLNEQ applies the NEQ predicate on the "author_url" field.
func AuthorURLNEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthorURL), v))
	})
}

// AuthorURLIn applies the In predicate on the "author_url" field.
func AuthorURLIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAuthorURL), v...))
	})
}

// AuthorURLNotIn applies the NotIn predicate on the "author_url" field.
func AuthorURLNotIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAuthorURL), v...))
	})
}

// AuthorURLGT applies the GT predicate on the "author_url" field.
func AuthorURLGT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAuthorURL), v))
	})
}

// AuthorURLGTE applies the GTE predicate on the "author_url" field.
func AuthorURLGTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAuthorURL), v))
	})
}

// AuthorURLLT applies the LT predicate on the "author_url" field.
func AuthorURLLT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAuthorURL), v))
	})
}

// AuthorURLLTE applies the LTE predicate on the "author_url" field.
func AuthorURLLTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAuthorURL), v))
	})
}

// AuthorURLContains applies the Contains predicate on the "author_url" field.
func AuthorURLContains(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAuthorURL), v))
	})
}

// AuthorURLHasPrefix applies the HasPrefix predicate on the "author_url" field.
func AuthorURLHasPrefix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAuthorURL), v))
	})
}

// AuthorURLHasSuffix applies the HasSuffix predicate on the "author_url" field.
func AuthorURLHasSuffix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAuthorURL), v))
	})
}

// AuthorURLIsNil applies the IsNil predicate on the "author_url" field.
func AuthorURLIsNil() predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAuthorURL)))
	})
}

// AuthorURLNotNil applies the NotNil predicate on the "author_url" field.
func AuthorURLNotNil() predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAuthorURL)))
	})
}

// AuthorURLEqualFold applies the EqualFold predicate on the "author_url" field.
func AuthorURLEqualFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAuthorURL), v))
	})
}

// AuthorURLContainsFold applies the ContainsFold predicate on the "author_url" field.
func AuthorURLContainsFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAuthorURL), v))
	})
}

// ImageURLEQ applies the EQ predicate on the "image_url" field.
func ImageURLEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImageURL), v))
	})
}

// ImageURLNEQ applies the NEQ predicate on the "image_url" field.
func ImageURLNEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImageURL), v))
	})
}

// ImageURLIn applies the In predicate on the "image_url" field.
func ImageURLIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldImageURL), v...))
	})
}

// ImageURLNotIn applies the NotIn predicate on the "image_url" field.
func ImageURLNotIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldImageURL), v...))
	})
}

// ImageURLGT applies the GT predicate on the "image_url" field.
func ImageURLGT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImageURL), v))
	})
}

// ImageURLGTE applies the GTE predicate on the "image_url" field.
func ImageURLGTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImageURL), v))
	})
}

// ImageURLLT applies the LT predicate on the "image_url" field.
func ImageURLLT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImageURL), v))
	})
}

// ImageURLLTE applies the LTE predicate on the "image_url" field.
func ImageURLLTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImageURL), v))
	})
}

// ImageURLContains applies the Contains predicate on the "image_url" field.
func ImageURLContains(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImageURL), v))
	})
}

// ImageURLHasPrefix applies the HasPrefix predicate on the "image_url" field.
func ImageURLHasPrefix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImageURL), v))
	})
}

// ImageURLHasSuffix applies the HasSuffix predicate on the "image_url" field.
func ImageURLHasSuffix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImageURL), v))
	})
}

// ImageURLIsNil applies the IsNil predicate on the "image_url" field.
func ImageURLIsNil() predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldImageURL)))
	})
}

// ImageURLNotNil applies the NotNil predicate on the "image_url" field.
func ImageURLNotNil() predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldImageURL)))
	})
}

// ImageURLEqualFold applies the EqualFold predicate on the "image_url" field.
func ImageURLEqualFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImageURL), v))
	})
}

// ImageURLContainsFold applies the ContainsFold predicate on the "image_url" field.
func ImageURLContainsFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImageURL), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// ViewsEQ applies the EQ predicate on the "views" field.
func ViewsEQ(v int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldViews), v))
	})
}

// ViewsNEQ applies the NEQ predicate on the "views" field.
func ViewsNEQ(v int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldViews), v))
	})
}

// ViewsIn applies the In predicate on the "views" field.
func ViewsIn(vs ...int) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldViews), v...))
	})
}

// ViewsNotIn applies the NotIn predicate on the "views" field.
func ViewsNotIn(vs ...int) predicate.Page {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldViews), v...))
	})
}

// ViewsGT applies the GT predicate on the "views" field.
func ViewsGT(v int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldViews), v))
	})
}

// ViewsGTE applies the GTE predicate on the "views" field.
func ViewsGTE(v int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldViews), v))
	})
}

// ViewsLT applies the LT predicate on the "views" field.
func ViewsLT(v int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldViews), v))
	})
}

// ViewsLTE applies the LTE predicate on the "views" field.
func ViewsLTE(v int) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldViews), v))
	})
}

// CanEditEQ applies the EQ predicate on the "can_edit" field.
func CanEditEQ(v bool) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCanEdit), v))
	})
}

// CanEditNEQ applies the NEQ predicate on the "can_edit" field.
func CanEditNEQ(v bool) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCanEdit), v))
	})
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.Account) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Page) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Page) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Page) predicate.Page {
	return predicate.Page(func(s *sql.Selector) {
		p(s.Not())
	})
}
