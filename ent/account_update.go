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

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetShortName sets the "short_name" field.
func (au *AccountUpdate) SetShortName(s string) *AccountUpdate {
	au.mutation.SetShortName(s)
	return au
}

// SetAuthorName sets the "author_name" field.
func (au *AccountUpdate) SetAuthorName(s string) *AccountUpdate {
	au.mutation.SetAuthorName(s)
	return au
}

// SetAuthorURL sets the "author_url" field.
func (au *AccountUpdate) SetAuthorURL(s string) *AccountUpdate {
	au.mutation.SetAuthorURL(s)
	return au
}

// SetAccessToken sets the "access_token" field.
func (au *AccountUpdate) SetAccessToken(s string) *AccountUpdate {
	au.mutation.SetAccessToken(s)
	return au
}

// SetAuthURL sets the "auth_url" field.
func (au *AccountUpdate) SetAuthURL(s string) *AccountUpdate {
	au.mutation.SetAuthURL(s)
	return au
}

// SetPageCount sets the "page_count" field.
func (au *AccountUpdate) SetPageCount(i int) *AccountUpdate {
	au.mutation.ResetPageCount()
	au.mutation.SetPageCount(i)
	return au
}

// SetNillablePageCount sets the "page_count" field if the given value is not nil.
func (au *AccountUpdate) SetNillablePageCount(i *int) *AccountUpdate {
	if i != nil {
		au.SetPageCount(*i)
	}
	return au
}

// AddPageCount adds i to the "page_count" field.
func (au *AccountUpdate) AddPageCount(i int) *AccountUpdate {
	au.mutation.AddPageCount(i)
	return au
}

// AddPageIDs adds the "pages" edge to the Page entity by IDs.
func (au *AccountUpdate) AddPageIDs(ids ...int) *AccountUpdate {
	au.mutation.AddPageIDs(ids...)
	return au
}

// AddPages adds the "pages" edges to the Page entity.
func (au *AccountUpdate) AddPages(p ...*Page) *AccountUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return au.AddPageIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearPages clears all "pages" edges to the Page entity.
func (au *AccountUpdate) ClearPages() *AccountUpdate {
	au.mutation.ClearPages()
	return au
}

// RemovePageIDs removes the "pages" edge to Page entities by IDs.
func (au *AccountUpdate) RemovePageIDs(ids ...int) *AccountUpdate {
	au.mutation.RemovePageIDs(ids...)
	return au
}

// RemovePages removes "pages" edges to Page entities.
func (au *AccountUpdate) RemovePages(p ...*Page) *AccountUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return au.RemovePageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.ShortName(); ok {
		if err := account.ShortNameValidator(v); err != nil {
			return &ValidationError{Name: "short_name", err: fmt.Errorf(`ent: validator failed for field "Account.short_name": %w`, err)}
		}
	}
	if v, ok := au.mutation.AuthorName(); ok {
		if err := account.AuthorNameValidator(v); err != nil {
			return &ValidationError{Name: "author_name", err: fmt.Errorf(`ent: validator failed for field "Account.author_name": %w`, err)}
		}
	}
	if v, ok := au.mutation.AuthorURL(); ok {
		if err := account.AuthorURLValidator(v); err != nil {
			return &ValidationError{Name: "author_url", err: fmt.Errorf(`ent: validator failed for field "Account.author_url": %w`, err)}
		}
	}
	if v, ok := au.mutation.AccessToken(); ok {
		if err := account.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "Account.access_token": %w`, err)}
		}
	}
	if v, ok := au.mutation.AuthURL(); ok {
		if err := account.AuthURLValidator(v); err != nil {
			return &ValidationError{Name: "auth_url", err: fmt.Errorf(`ent: validator failed for field "Account.auth_url": %w`, err)}
		}
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.ShortName(); ok {
		_spec.SetField(account.FieldShortName, field.TypeString, value)
	}
	if value, ok := au.mutation.AuthorName(); ok {
		_spec.SetField(account.FieldAuthorName, field.TypeString, value)
	}
	if value, ok := au.mutation.AuthorURL(); ok {
		_spec.SetField(account.FieldAuthorURL, field.TypeString, value)
	}
	if value, ok := au.mutation.AccessToken(); ok {
		_spec.SetField(account.FieldAccessToken, field.TypeString, value)
	}
	if value, ok := au.mutation.AuthURL(); ok {
		_spec.SetField(account.FieldAuthURL, field.TypeString, value)
	}
	if value, ok := au.mutation.PageCount(); ok {
		_spec.SetField(account.FieldPageCount, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedPageCount(); ok {
		_spec.AddField(account.FieldPageCount, field.TypeInt, value)
	}
	if au.mutation.PagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PagesTable,
			Columns: []string{account.PagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: page.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedPagesIDs(); len(nodes) > 0 && !au.mutation.PagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PagesTable,
			Columns: []string{account.PagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: page.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.PagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PagesTable,
			Columns: []string{account.PagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: page.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetShortName sets the "short_name" field.
func (auo *AccountUpdateOne) SetShortName(s string) *AccountUpdateOne {
	auo.mutation.SetShortName(s)
	return auo
}

// SetAuthorName sets the "author_name" field.
func (auo *AccountUpdateOne) SetAuthorName(s string) *AccountUpdateOne {
	auo.mutation.SetAuthorName(s)
	return auo
}

// SetAuthorURL sets the "author_url" field.
func (auo *AccountUpdateOne) SetAuthorURL(s string) *AccountUpdateOne {
	auo.mutation.SetAuthorURL(s)
	return auo
}

// SetAccessToken sets the "access_token" field.
func (auo *AccountUpdateOne) SetAccessToken(s string) *AccountUpdateOne {
	auo.mutation.SetAccessToken(s)
	return auo
}

// SetAuthURL sets the "auth_url" field.
func (auo *AccountUpdateOne) SetAuthURL(s string) *AccountUpdateOne {
	auo.mutation.SetAuthURL(s)
	return auo
}

// SetPageCount sets the "page_count" field.
func (auo *AccountUpdateOne) SetPageCount(i int) *AccountUpdateOne {
	auo.mutation.ResetPageCount()
	auo.mutation.SetPageCount(i)
	return auo
}

// SetNillablePageCount sets the "page_count" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillablePageCount(i *int) *AccountUpdateOne {
	if i != nil {
		auo.SetPageCount(*i)
	}
	return auo
}

// AddPageCount adds i to the "page_count" field.
func (auo *AccountUpdateOne) AddPageCount(i int) *AccountUpdateOne {
	auo.mutation.AddPageCount(i)
	return auo
}

// AddPageIDs adds the "pages" edge to the Page entity by IDs.
func (auo *AccountUpdateOne) AddPageIDs(ids ...int) *AccountUpdateOne {
	auo.mutation.AddPageIDs(ids...)
	return auo
}

// AddPages adds the "pages" edges to the Page entity.
func (auo *AccountUpdateOne) AddPages(p ...*Page) *AccountUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return auo.AddPageIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearPages clears all "pages" edges to the Page entity.
func (auo *AccountUpdateOne) ClearPages() *AccountUpdateOne {
	auo.mutation.ClearPages()
	return auo
}

// RemovePageIDs removes the "pages" edge to Page entities by IDs.
func (auo *AccountUpdateOne) RemovePageIDs(ids ...int) *AccountUpdateOne {
	auo.mutation.RemovePageIDs(ids...)
	return auo
}

// RemovePages removes "pages" edges to Page entities.
func (auo *AccountUpdateOne) RemovePages(p ...*Page) *AccountUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return auo.RemovePageIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Account)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AccountMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.ShortName(); ok {
		if err := account.ShortNameValidator(v); err != nil {
			return &ValidationError{Name: "short_name", err: fmt.Errorf(`ent: validator failed for field "Account.short_name": %w`, err)}
		}
	}
	if v, ok := auo.mutation.AuthorName(); ok {
		if err := account.AuthorNameValidator(v); err != nil {
			return &ValidationError{Name: "author_name", err: fmt.Errorf(`ent: validator failed for field "Account.author_name": %w`, err)}
		}
	}
	if v, ok := auo.mutation.AuthorURL(); ok {
		if err := account.AuthorURLValidator(v); err != nil {
			return &ValidationError{Name: "author_url", err: fmt.Errorf(`ent: validator failed for field "Account.author_url": %w`, err)}
		}
	}
	if v, ok := auo.mutation.AccessToken(); ok {
		if err := account.AccessTokenValidator(v); err != nil {
			return &ValidationError{Name: "access_token", err: fmt.Errorf(`ent: validator failed for field "Account.access_token": %w`, err)}
		}
	}
	if v, ok := auo.mutation.AuthURL(); ok {
		if err := account.AuthURLValidator(v); err != nil {
			return &ValidationError{Name: "auth_url", err: fmt.Errorf(`ent: validator failed for field "Account.auth_url": %w`, err)}
		}
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.ShortName(); ok {
		_spec.SetField(account.FieldShortName, field.TypeString, value)
	}
	if value, ok := auo.mutation.AuthorName(); ok {
		_spec.SetField(account.FieldAuthorName, field.TypeString, value)
	}
	if value, ok := auo.mutation.AuthorURL(); ok {
		_spec.SetField(account.FieldAuthorURL, field.TypeString, value)
	}
	if value, ok := auo.mutation.AccessToken(); ok {
		_spec.SetField(account.FieldAccessToken, field.TypeString, value)
	}
	if value, ok := auo.mutation.AuthURL(); ok {
		_spec.SetField(account.FieldAuthURL, field.TypeString, value)
	}
	if value, ok := auo.mutation.PageCount(); ok {
		_spec.SetField(account.FieldPageCount, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedPageCount(); ok {
		_spec.AddField(account.FieldPageCount, field.TypeInt, value)
	}
	if auo.mutation.PagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PagesTable,
			Columns: []string{account.PagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: page.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedPagesIDs(); len(nodes) > 0 && !auo.mutation.PagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PagesTable,
			Columns: []string{account.PagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: page.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.PagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.PagesTable,
			Columns: []string{account.PagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: page.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
