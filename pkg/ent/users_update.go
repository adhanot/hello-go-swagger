// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hello/pkg/ent/predicate"
	"hello/pkg/ent/users"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UsersUpdate is the builder for updating Users entities.
type UsersUpdate struct {
	config
	hooks    []Hook
	mutation *UsersMutation
}

// Where adds a new predicate for the UsersUpdate builder.
func (uu *UsersUpdate) Where(ps ...predicate.Users) *UsersUpdate {
	uu.mutation.predicates = append(uu.mutation.predicates, ps...)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UsersUpdate) SetPassword(s string) *UsersUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetFname sets the "fname" field.
func (uu *UsersUpdate) SetFname(s string) *UsersUpdate {
	uu.mutation.SetFname(s)
	return uu
}

// SetLname sets the "lname" field.
func (uu *UsersUpdate) SetLname(s string) *UsersUpdate {
	uu.mutation.SetLname(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UsersUpdate) SetEmail(s string) *UsersUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UsersUpdate) SetCreatedAt(t time.Time) *UsersUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableCreatedAt(t *time.Time) *UsersUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// Mutation returns the UsersMutation object of the builder.
func (uu *UsersUpdate) Mutation() *UsersMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UsersUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UsersMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UsersUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UsersUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UsersUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UsersUpdate) check() error {
	if v, ok := uu.mutation.Password(); ok {
		if err := users.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Fname(); ok {
		if err := users.FnameValidator(v); err != nil {
			return &ValidationError{Name: "fname", err: fmt.Errorf("ent: validator failed for field \"fname\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Lname(); ok {
		if err := users.LnameValidator(v); err != nil {
			return &ValidationError{Name: "lname", err: fmt.Errorf("ent: validator failed for field \"lname\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := users.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	return nil
}

func (uu *UsersUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   users.Table,
			Columns: users.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: users.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldPassword,
		})
	}
	if value, ok := uu.mutation.Fname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldFname,
		})
	}
	if value, ok := uu.mutation.Lname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldLname,
		})
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldEmail,
		})
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: users.FieldCreatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{users.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UsersUpdateOne is the builder for updating a single Users entity.
type UsersUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UsersMutation
}

// SetPassword sets the "password" field.
func (uuo *UsersUpdateOne) SetPassword(s string) *UsersUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetFname sets the "fname" field.
func (uuo *UsersUpdateOne) SetFname(s string) *UsersUpdateOne {
	uuo.mutation.SetFname(s)
	return uuo
}

// SetLname sets the "lname" field.
func (uuo *UsersUpdateOne) SetLname(s string) *UsersUpdateOne {
	uuo.mutation.SetLname(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UsersUpdateOne) SetEmail(s string) *UsersUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UsersUpdateOne) SetCreatedAt(t time.Time) *UsersUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableCreatedAt(t *time.Time) *UsersUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// Mutation returns the UsersMutation object of the builder.
func (uuo *UsersUpdateOne) Mutation() *UsersMutation {
	return uuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UsersUpdateOne) Select(field string, fields ...string) *UsersUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Users entity.
func (uuo *UsersUpdateOne) Save(ctx context.Context) (*Users, error) {
	var (
		err  error
		node *Users
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UsersMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UsersUpdateOne) SaveX(ctx context.Context) *Users {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UsersUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UsersUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UsersUpdateOne) check() error {
	if v, ok := uuo.mutation.Password(); ok {
		if err := users.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Fname(); ok {
		if err := users.FnameValidator(v); err != nil {
			return &ValidationError{Name: "fname", err: fmt.Errorf("ent: validator failed for field \"fname\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Lname(); ok {
		if err := users.LnameValidator(v); err != nil {
			return &ValidationError{Name: "lname", err: fmt.Errorf("ent: validator failed for field \"lname\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := users.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	return nil
}

func (uuo *UsersUpdateOne) sqlSave(ctx context.Context) (_node *Users, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   users.Table,
			Columns: users.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: users.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Users.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, users.FieldID)
		for _, f := range fields {
			if !users.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != users.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldPassword,
		})
	}
	if value, ok := uuo.mutation.Fname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldFname,
		})
	}
	if value, ok := uuo.mutation.Lname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldLname,
		})
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: users.FieldCreatedAt,
		})
	}
	_node = &Users{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{users.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
