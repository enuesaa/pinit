// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/enuesaa/pinit/internal/ent/action"
	"github.com/enuesaa/pinit/internal/ent/predicate"
)

// ActionUpdate is the builder for updating Action entities.
type ActionUpdate struct {
	config
	hooks    []Hook
	mutation *ActionMutation
}

// Where appends a list predicates to the ActionUpdate builder.
func (au *ActionUpdate) Where(ps ...predicate.Action) *ActionUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *ActionUpdate) SetName(s string) *ActionUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *ActionUpdate) SetNillableName(s *string) *ActionUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetTemplate sets the "template" field.
func (au *ActionUpdate) SetTemplate(s string) *ActionUpdate {
	au.mutation.SetTemplate(s)
	return au
}

// SetNillableTemplate sets the "template" field if the given value is not nil.
func (au *ActionUpdate) SetNillableTemplate(s *string) *ActionUpdate {
	if s != nil {
		au.SetTemplate(*s)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *ActionUpdate) SetUpdatedAt(t time.Time) *ActionUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// Mutation returns the ActionMutation object of the builder.
func (au *ActionUpdate) Mutation() *ActionMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ActionUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ActionUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ActionUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ActionUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *ActionUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := action.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *ActionUpdate) check() error {
	if v, ok := au.mutation.Name(); ok {
		if err := action.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Action.name": %w`, err)}
		}
	}
	return nil
}

func (au *ActionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(action.Table, action.Columns, sqlgraph.NewFieldSpec(action.FieldID, field.TypeUint))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(action.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.Template(); ok {
		_spec.SetField(action.FieldTemplate, field.TypeString, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(action.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{action.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ActionUpdateOne is the builder for updating a single Action entity.
type ActionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ActionMutation
}

// SetName sets the "name" field.
func (auo *ActionUpdateOne) SetName(s string) *ActionUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *ActionUpdateOne) SetNillableName(s *string) *ActionUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetTemplate sets the "template" field.
func (auo *ActionUpdateOne) SetTemplate(s string) *ActionUpdateOne {
	auo.mutation.SetTemplate(s)
	return auo
}

// SetNillableTemplate sets the "template" field if the given value is not nil.
func (auo *ActionUpdateOne) SetNillableTemplate(s *string) *ActionUpdateOne {
	if s != nil {
		auo.SetTemplate(*s)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *ActionUpdateOne) SetUpdatedAt(t time.Time) *ActionUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// Mutation returns the ActionMutation object of the builder.
func (auo *ActionUpdateOne) Mutation() *ActionMutation {
	return auo.mutation
}

// Where appends a list predicates to the ActionUpdate builder.
func (auo *ActionUpdateOne) Where(ps ...predicate.Action) *ActionUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ActionUpdateOne) Select(field string, fields ...string) *ActionUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Action entity.
func (auo *ActionUpdateOne) Save(ctx context.Context) (*Action, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ActionUpdateOne) SaveX(ctx context.Context) *Action {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ActionUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ActionUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *ActionUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := action.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *ActionUpdateOne) check() error {
	if v, ok := auo.mutation.Name(); ok {
		if err := action.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Action.name": %w`, err)}
		}
	}
	return nil
}

func (auo *ActionUpdateOne) sqlSave(ctx context.Context) (_node *Action, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(action.Table, action.Columns, sqlgraph.NewFieldSpec(action.FieldID, field.TypeUint))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Action.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, action.FieldID)
		for _, f := range fields {
			if !action.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != action.FieldID {
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
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(action.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.Template(); ok {
		_spec.SetField(action.FieldTemplate, field.TypeString, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(action.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Action{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{action.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
