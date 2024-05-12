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
	"github.com/enuesaa/pinit/pkg/ent/appconf"
	"github.com/enuesaa/pinit/pkg/ent/predicate"
)

// AppconfUpdate is the builder for updating Appconf entities.
type AppconfUpdate struct {
	config
	hooks    []Hook
	mutation *AppconfMutation
}

// Where appends a list predicates to the AppconfUpdate builder.
func (au *AppconfUpdate) Where(ps ...predicate.Appconf) *AppconfUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetKey sets the "key" field.
func (au *AppconfUpdate) SetKey(s string) *AppconfUpdate {
	au.mutation.SetKey(s)
	return au
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (au *AppconfUpdate) SetNillableKey(s *string) *AppconfUpdate {
	if s != nil {
		au.SetKey(*s)
	}
	return au
}

// SetValue sets the "value" field.
func (au *AppconfUpdate) SetValue(s string) *AppconfUpdate {
	au.mutation.SetValue(s)
	return au
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (au *AppconfUpdate) SetNillableValue(s *string) *AppconfUpdate {
	if s != nil {
		au.SetValue(*s)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AppconfUpdate) SetUpdatedAt(t time.Time) *AppconfUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// Mutation returns the AppconfMutation object of the builder.
func (au *AppconfUpdate) Mutation() *AppconfMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AppconfUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AppconfUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AppconfUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AppconfUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AppconfUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := appconf.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

func (au *AppconfUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(appconf.Table, appconf.Columns, sqlgraph.NewFieldSpec(appconf.FieldID, field.TypeUint))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Key(); ok {
		_spec.SetField(appconf.FieldKey, field.TypeString, value)
	}
	if value, ok := au.mutation.Value(); ok {
		_spec.SetField(appconf.FieldValue, field.TypeString, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(appconf.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appconf.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AppconfUpdateOne is the builder for updating a single Appconf entity.
type AppconfUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppconfMutation
}

// SetKey sets the "key" field.
func (auo *AppconfUpdateOne) SetKey(s string) *AppconfUpdateOne {
	auo.mutation.SetKey(s)
	return auo
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (auo *AppconfUpdateOne) SetNillableKey(s *string) *AppconfUpdateOne {
	if s != nil {
		auo.SetKey(*s)
	}
	return auo
}

// SetValue sets the "value" field.
func (auo *AppconfUpdateOne) SetValue(s string) *AppconfUpdateOne {
	auo.mutation.SetValue(s)
	return auo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (auo *AppconfUpdateOne) SetNillableValue(s *string) *AppconfUpdateOne {
	if s != nil {
		auo.SetValue(*s)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AppconfUpdateOne) SetUpdatedAt(t time.Time) *AppconfUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// Mutation returns the AppconfMutation object of the builder.
func (auo *AppconfUpdateOne) Mutation() *AppconfMutation {
	return auo.mutation
}

// Where appends a list predicates to the AppconfUpdate builder.
func (auo *AppconfUpdateOne) Where(ps ...predicate.Appconf) *AppconfUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AppconfUpdateOne) Select(field string, fields ...string) *AppconfUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Appconf entity.
func (auo *AppconfUpdateOne) Save(ctx context.Context) (*Appconf, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AppconfUpdateOne) SaveX(ctx context.Context) *Appconf {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AppconfUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AppconfUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AppconfUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := appconf.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

func (auo *AppconfUpdateOne) sqlSave(ctx context.Context) (_node *Appconf, err error) {
	_spec := sqlgraph.NewUpdateSpec(appconf.Table, appconf.Columns, sqlgraph.NewFieldSpec(appconf.FieldID, field.TypeUint))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Appconf.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appconf.FieldID)
		for _, f := range fields {
			if !appconf.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appconf.FieldID {
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
	if value, ok := auo.mutation.Key(); ok {
		_spec.SetField(appconf.FieldKey, field.TypeString, value)
	}
	if value, ok := auo.mutation.Value(); ok {
		_spec.SetField(appconf.FieldValue, field.TypeString, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(appconf.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Appconf{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appconf.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}