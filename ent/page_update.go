// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"telegra/ent/account"
	"telegra/ent/page"
	"telegra/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PageUpdate is the builder for updating Page entities.
type PageUpdate struct {
	config
	hooks    []Hook
	mutation *PageMutation
}

// Where appends a list predicates to the PageUpdate builder.
func (pu *PageUpdate) Where(ps ...predicate.Page) *PageUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetPath sets the "path" field.
func (pu *PageUpdate) SetPath(s string) *PageUpdate {
	pu.mutation.SetPath(s)
	return pu
}

// SetURL sets the "url" field.
func (pu *PageUpdate) SetURL(s string) *PageUpdate {
	pu.mutation.SetURL(s)
	return pu
}

// SetTitle sets the "title" field.
func (pu *PageUpdate) SetTitle(s string) *PageUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetDescription sets the "description" field.
func (pu *PageUpdate) SetDescription(s string) *PageUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PageUpdate) SetNillableDescription(s *string) *PageUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *PageUpdate) ClearDescription() *PageUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetAuthorName sets the "author_name" field.
func (pu *PageUpdate) SetAuthorName(s string) *PageUpdate {
	pu.mutation.SetAuthorName(s)
	return pu
}

// SetNillableAuthorName sets the "author_name" field if the given value is not nil.
func (pu *PageUpdate) SetNillableAuthorName(s *string) *PageUpdate {
	if s != nil {
		pu.SetAuthorName(*s)
	}
	return pu
}

// ClearAuthorName clears the value of the "author_name" field.
func (pu *PageUpdate) ClearAuthorName() *PageUpdate {
	pu.mutation.ClearAuthorName()
	return pu
}

// SetAuthorURL sets the "author_url" field.
func (pu *PageUpdate) SetAuthorURL(s string) *PageUpdate {
	pu.mutation.SetAuthorURL(s)
	return pu
}

// SetNillableAuthorURL sets the "author_url" field if the given value is not nil.
func (pu *PageUpdate) SetNillableAuthorURL(s *string) *PageUpdate {
	if s != nil {
		pu.SetAuthorURL(*s)
	}
	return pu
}

// ClearAuthorURL clears the value of the "author_url" field.
func (pu *PageUpdate) ClearAuthorURL() *PageUpdate {
	pu.mutation.ClearAuthorURL()
	return pu
}

// SetImageURL sets the "image_url" field.
func (pu *PageUpdate) SetImageURL(s string) *PageUpdate {
	pu.mutation.SetImageURL(s)
	return pu
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (pu *PageUpdate) SetNillableImageURL(s *string) *PageUpdate {
	if s != nil {
		pu.SetImageURL(*s)
	}
	return pu
}

// ClearImageURL clears the value of the "image_url" field.
func (pu *PageUpdate) ClearImageURL() *PageUpdate {
	pu.mutation.ClearImageURL()
	return pu
}

// SetContent sets the "content" field.
func (pu *PageUpdate) SetContent(s string) *PageUpdate {
	pu.mutation.SetContent(s)
	return pu
}

// SetViews sets the "views" field.
func (pu *PageUpdate) SetViews(i int) *PageUpdate {
	pu.mutation.ResetViews()
	pu.mutation.SetViews(i)
	return pu
}

// SetNillableViews sets the "views" field if the given value is not nil.
func (pu *PageUpdate) SetNillableViews(i *int) *PageUpdate {
	if i != nil {
		pu.SetViews(*i)
	}
	return pu
}

// AddViews adds i to the "views" field.
func (pu *PageUpdate) AddViews(i int) *PageUpdate {
	pu.mutation.AddViews(i)
	return pu
}

// SetCanEdit sets the "can_edit" field.
func (pu *PageUpdate) SetCanEdit(b bool) *PageUpdate {
	pu.mutation.SetCanEdit(b)
	return pu
}

// SetNillableCanEdit sets the "can_edit" field if the given value is not nil.
func (pu *PageUpdate) SetNillableCanEdit(b *bool) *PageUpdate {
	if b != nil {
		pu.SetCanEdit(*b)
	}
	return pu
}

// SetAuthorID sets the "author" edge to the Account entity by ID.
func (pu *PageUpdate) SetAuthorID(id int) *PageUpdate {
	pu.mutation.SetAuthorID(id)
	return pu
}

// SetNillableAuthorID sets the "author" edge to the Account entity by ID if the given value is not nil.
func (pu *PageUpdate) SetNillableAuthorID(id *int) *PageUpdate {
	if id != nil {
		pu = pu.SetAuthorID(*id)
	}
	return pu
}

// SetAuthor sets the "author" edge to the Account entity.
func (pu *PageUpdate) SetAuthor(a *Account) *PageUpdate {
	return pu.SetAuthorID(a.ID)
}

// Mutation returns the PageMutation object of the builder.
func (pu *PageUpdate) Mutation() *PageMutation {
	return pu.mutation
}

// ClearAuthor clears the "author" edge to the Account entity.
func (pu *PageUpdate) ClearAuthor() *PageUpdate {
	pu.mutation.ClearAuthor()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PageUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PageUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PageUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PageUpdate) check() error {
	if v, ok := pu.mutation.Path(); ok {
		if err := page.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Page.path": %w`, err)}
		}
	}
	if v, ok := pu.mutation.URL(); ok {
		if err := page.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Page.url": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Title(); ok {
		if err := page.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Page.title": %w`, err)}
		}
	}
	if v, ok := pu.mutation.AuthorName(); ok {
		if err := page.AuthorNameValidator(v); err != nil {
			return &ValidationError{Name: "author_name", err: fmt.Errorf(`ent: validator failed for field "Page.author_name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.AuthorURL(); ok {
		if err := page.AuthorURLValidator(v); err != nil {
			return &ValidationError{Name: "author_url", err: fmt.Errorf(`ent: validator failed for field "Page.author_url": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Content(); ok {
		if err := page.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Page.content": %w`, err)}
		}
	}
	return nil
}

func (pu *PageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   page.Table,
			Columns: page.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: page.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Path(); ok {
		_spec.SetField(page.FieldPath, field.TypeString, value)
	}
	if value, ok := pu.mutation.URL(); ok {
		_spec.SetField(page.FieldURL, field.TypeString, value)
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(page.FieldTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(page.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(page.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.AuthorName(); ok {
		_spec.SetField(page.FieldAuthorName, field.TypeString, value)
	}
	if pu.mutation.AuthorNameCleared() {
		_spec.ClearField(page.FieldAuthorName, field.TypeString)
	}
	if value, ok := pu.mutation.AuthorURL(); ok {
		_spec.SetField(page.FieldAuthorURL, field.TypeString, value)
	}
	if pu.mutation.AuthorURLCleared() {
		_spec.ClearField(page.FieldAuthorURL, field.TypeString)
	}
	if value, ok := pu.mutation.ImageURL(); ok {
		_spec.SetField(page.FieldImageURL, field.TypeString, value)
	}
	if pu.mutation.ImageURLCleared() {
		_spec.ClearField(page.FieldImageURL, field.TypeString)
	}
	if value, ok := pu.mutation.Content(); ok {
		_spec.SetField(page.FieldContent, field.TypeString, value)
	}
	if value, ok := pu.mutation.Views(); ok {
		_spec.SetField(page.FieldViews, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedViews(); ok {
		_spec.AddField(page.FieldViews, field.TypeInt, value)
	}
	if value, ok := pu.mutation.CanEdit(); ok {
		_spec.SetField(page.FieldCanEdit, field.TypeBool, value)
	}
	if pu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   page.AuthorTable,
			Columns: []string{page.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   page.AuthorTable,
			Columns: []string{page.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{page.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PageUpdateOne is the builder for updating a single Page entity.
type PageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PageMutation
}

// SetPath sets the "path" field.
func (puo *PageUpdateOne) SetPath(s string) *PageUpdateOne {
	puo.mutation.SetPath(s)
	return puo
}

// SetURL sets the "url" field.
func (puo *PageUpdateOne) SetURL(s string) *PageUpdateOne {
	puo.mutation.SetURL(s)
	return puo
}

// SetTitle sets the "title" field.
func (puo *PageUpdateOne) SetTitle(s string) *PageUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetDescription sets the "description" field.
func (puo *PageUpdateOne) SetDescription(s string) *PageUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableDescription(s *string) *PageUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *PageUpdateOne) ClearDescription() *PageUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetAuthorName sets the "author_name" field.
func (puo *PageUpdateOne) SetAuthorName(s string) *PageUpdateOne {
	puo.mutation.SetAuthorName(s)
	return puo
}

// SetNillableAuthorName sets the "author_name" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableAuthorName(s *string) *PageUpdateOne {
	if s != nil {
		puo.SetAuthorName(*s)
	}
	return puo
}

// ClearAuthorName clears the value of the "author_name" field.
func (puo *PageUpdateOne) ClearAuthorName() *PageUpdateOne {
	puo.mutation.ClearAuthorName()
	return puo
}

// SetAuthorURL sets the "author_url" field.
func (puo *PageUpdateOne) SetAuthorURL(s string) *PageUpdateOne {
	puo.mutation.SetAuthorURL(s)
	return puo
}

// SetNillableAuthorURL sets the "author_url" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableAuthorURL(s *string) *PageUpdateOne {
	if s != nil {
		puo.SetAuthorURL(*s)
	}
	return puo
}

// ClearAuthorURL clears the value of the "author_url" field.
func (puo *PageUpdateOne) ClearAuthorURL() *PageUpdateOne {
	puo.mutation.ClearAuthorURL()
	return puo
}

// SetImageURL sets the "image_url" field.
func (puo *PageUpdateOne) SetImageURL(s string) *PageUpdateOne {
	puo.mutation.SetImageURL(s)
	return puo
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableImageURL(s *string) *PageUpdateOne {
	if s != nil {
		puo.SetImageURL(*s)
	}
	return puo
}

// ClearImageURL clears the value of the "image_url" field.
func (puo *PageUpdateOne) ClearImageURL() *PageUpdateOne {
	puo.mutation.ClearImageURL()
	return puo
}

// SetContent sets the "content" field.
func (puo *PageUpdateOne) SetContent(s string) *PageUpdateOne {
	puo.mutation.SetContent(s)
	return puo
}

// SetViews sets the "views" field.
func (puo *PageUpdateOne) SetViews(i int) *PageUpdateOne {
	puo.mutation.ResetViews()
	puo.mutation.SetViews(i)
	return puo
}

// SetNillableViews sets the "views" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableViews(i *int) *PageUpdateOne {
	if i != nil {
		puo.SetViews(*i)
	}
	return puo
}

// AddViews adds i to the "views" field.
func (puo *PageUpdateOne) AddViews(i int) *PageUpdateOne {
	puo.mutation.AddViews(i)
	return puo
}

// SetCanEdit sets the "can_edit" field.
func (puo *PageUpdateOne) SetCanEdit(b bool) *PageUpdateOne {
	puo.mutation.SetCanEdit(b)
	return puo
}

// SetNillableCanEdit sets the "can_edit" field if the given value is not nil.
func (puo *PageUpdateOne) SetNillableCanEdit(b *bool) *PageUpdateOne {
	if b != nil {
		puo.SetCanEdit(*b)
	}
	return puo
}

// SetAuthorID sets the "author" edge to the Account entity by ID.
func (puo *PageUpdateOne) SetAuthorID(id int) *PageUpdateOne {
	puo.mutation.SetAuthorID(id)
	return puo
}

// SetNillableAuthorID sets the "author" edge to the Account entity by ID if the given value is not nil.
func (puo *PageUpdateOne) SetNillableAuthorID(id *int) *PageUpdateOne {
	if id != nil {
		puo = puo.SetAuthorID(*id)
	}
	return puo
}

// SetAuthor sets the "author" edge to the Account entity.
func (puo *PageUpdateOne) SetAuthor(a *Account) *PageUpdateOne {
	return puo.SetAuthorID(a.ID)
}

// Mutation returns the PageMutation object of the builder.
func (puo *PageUpdateOne) Mutation() *PageMutation {
	return puo.mutation
}

// ClearAuthor clears the "author" edge to the Account entity.
func (puo *PageUpdateOne) ClearAuthor() *PageUpdateOne {
	puo.mutation.ClearAuthor()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PageUpdateOne) Select(field string, fields ...string) *PageUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Page entity.
func (puo *PageUpdateOne) Save(ctx context.Context) (*Page, error) {
	var (
		err  error
		node *Page
	)
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Page)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PageMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PageUpdateOne) SaveX(ctx context.Context) *Page {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PageUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PageUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PageUpdateOne) check() error {
	if v, ok := puo.mutation.Path(); ok {
		if err := page.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Page.path": %w`, err)}
		}
	}
	if v, ok := puo.mutation.URL(); ok {
		if err := page.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Page.url": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Title(); ok {
		if err := page.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Page.title": %w`, err)}
		}
	}
	if v, ok := puo.mutation.AuthorName(); ok {
		if err := page.AuthorNameValidator(v); err != nil {
			return &ValidationError{Name: "author_name", err: fmt.Errorf(`ent: validator failed for field "Page.author_name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.AuthorURL(); ok {
		if err := page.AuthorURLValidator(v); err != nil {
			return &ValidationError{Name: "author_url", err: fmt.Errorf(`ent: validator failed for field "Page.author_url": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Content(); ok {
		if err := page.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Page.content": %w`, err)}
		}
	}
	return nil
}

func (puo *PageUpdateOne) sqlSave(ctx context.Context) (_node *Page, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   page.Table,
			Columns: page.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: page.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Page.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, page.FieldID)
		for _, f := range fields {
			if !page.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != page.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Path(); ok {
		_spec.SetField(page.FieldPath, field.TypeString, value)
	}
	if value, ok := puo.mutation.URL(); ok {
		_spec.SetField(page.FieldURL, field.TypeString, value)
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(page.FieldTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(page.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(page.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.AuthorName(); ok {
		_spec.SetField(page.FieldAuthorName, field.TypeString, value)
	}
	if puo.mutation.AuthorNameCleared() {
		_spec.ClearField(page.FieldAuthorName, field.TypeString)
	}
	if value, ok := puo.mutation.AuthorURL(); ok {
		_spec.SetField(page.FieldAuthorURL, field.TypeString, value)
	}
	if puo.mutation.AuthorURLCleared() {
		_spec.ClearField(page.FieldAuthorURL, field.TypeString)
	}
	if value, ok := puo.mutation.ImageURL(); ok {
		_spec.SetField(page.FieldImageURL, field.TypeString, value)
	}
	if puo.mutation.ImageURLCleared() {
		_spec.ClearField(page.FieldImageURL, field.TypeString)
	}
	if value, ok := puo.mutation.Content(); ok {
		_spec.SetField(page.FieldContent, field.TypeString, value)
	}
	if value, ok := puo.mutation.Views(); ok {
		_spec.SetField(page.FieldViews, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedViews(); ok {
		_spec.AddField(page.FieldViews, field.TypeInt, value)
	}
	if value, ok := puo.mutation.CanEdit(); ok {
		_spec.SetField(page.FieldCanEdit, field.TypeBool, value)
	}
	if puo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   page.AuthorTable,
			Columns: []string{page.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   page.AuthorTable,
			Columns: []string{page.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Page{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{page.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
