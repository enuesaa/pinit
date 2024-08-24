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
	"github.com/enuesaa/pinit/internal/ent/note"
	"github.com/enuesaa/pinit/internal/ent/predicate"
)

// NoteUpdate is the builder for updating Note entities.
type NoteUpdate struct {
	config
	hooks    []Hook
	mutation *NoteMutation
}

// Where appends a list predicates to the NoteUpdate builder.
func (nu *NoteUpdate) Where(ps ...predicate.Note) *NoteUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetBinderID sets the "binder_id" field.
func (nu *NoteUpdate) SetBinderID(u uint) *NoteUpdate {
	nu.mutation.ResetBinderID()
	nu.mutation.SetBinderID(u)
	return nu
}

// SetNillableBinderID sets the "binder_id" field if the given value is not nil.
func (nu *NoteUpdate) SetNillableBinderID(u *uint) *NoteUpdate {
	if u != nil {
		nu.SetBinderID(*u)
	}
	return nu
}

// AddBinderID adds u to the "binder_id" field.
func (nu *NoteUpdate) AddBinderID(u int) *NoteUpdate {
	nu.mutation.AddBinderID(u)
	return nu
}

// ClearBinderID clears the value of the "binder_id" field.
func (nu *NoteUpdate) ClearBinderID() *NoteUpdate {
	nu.mutation.ClearBinderID()
	return nu
}

// SetPublisher sets the "publisher" field.
func (nu *NoteUpdate) SetPublisher(s string) *NoteUpdate {
	nu.mutation.SetPublisher(s)
	return nu
}

// SetNillablePublisher sets the "publisher" field if the given value is not nil.
func (nu *NoteUpdate) SetNillablePublisher(s *string) *NoteUpdate {
	if s != nil {
		nu.SetPublisher(*s)
	}
	return nu
}

// SetComment sets the "comment" field.
func (nu *NoteUpdate) SetComment(s string) *NoteUpdate {
	nu.mutation.SetComment(s)
	return nu
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (nu *NoteUpdate) SetNillableComment(s *string) *NoteUpdate {
	if s != nil {
		nu.SetComment(*s)
	}
	return nu
}

// ClearComment clears the value of the "comment" field.
func (nu *NoteUpdate) ClearComment() *NoteUpdate {
	nu.mutation.ClearComment()
	return nu
}

// SetContent sets the "content" field.
func (nu *NoteUpdate) SetContent(s string) *NoteUpdate {
	nu.mutation.SetContent(s)
	return nu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (nu *NoteUpdate) SetNillableContent(s *string) *NoteUpdate {
	if s != nil {
		nu.SetContent(*s)
	}
	return nu
}

// ClearContent clears the value of the "content" field.
func (nu *NoteUpdate) ClearContent() *NoteUpdate {
	nu.mutation.ClearContent()
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NoteUpdate) SetUpdatedAt(t time.Time) *NoteUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// Mutation returns the NoteMutation object of the builder.
func (nu *NoteUpdate) Mutation() *NoteMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NoteUpdate) Save(ctx context.Context) (int, error) {
	nu.defaults()
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NoteUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NoteUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NoteUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NoteUpdate) defaults() {
	if _, ok := nu.mutation.UpdatedAt(); !ok {
		v := note.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nu *NoteUpdate) check() error {
	if v, ok := nu.mutation.Publisher(); ok {
		if err := note.PublisherValidator(v); err != nil {
			return &ValidationError{Name: "publisher", err: fmt.Errorf(`ent: validator failed for field "Note.publisher": %w`, err)}
		}
	}
	return nil
}

func (nu *NoteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(note.Table, note.Columns, sqlgraph.NewFieldSpec(note.FieldID, field.TypeUint))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.BinderID(); ok {
		_spec.SetField(note.FieldBinderID, field.TypeUint, value)
	}
	if value, ok := nu.mutation.AddedBinderID(); ok {
		_spec.AddField(note.FieldBinderID, field.TypeUint, value)
	}
	if nu.mutation.BinderIDCleared() {
		_spec.ClearField(note.FieldBinderID, field.TypeUint)
	}
	if value, ok := nu.mutation.Publisher(); ok {
		_spec.SetField(note.FieldPublisher, field.TypeString, value)
	}
	if value, ok := nu.mutation.Comment(); ok {
		_spec.SetField(note.FieldComment, field.TypeString, value)
	}
	if nu.mutation.CommentCleared() {
		_spec.ClearField(note.FieldComment, field.TypeString)
	}
	if value, ok := nu.mutation.Content(); ok {
		_spec.SetField(note.FieldContent, field.TypeString, value)
	}
	if nu.mutation.ContentCleared() {
		_spec.ClearField(note.FieldContent, field.TypeString)
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.SetField(note.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{note.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NoteUpdateOne is the builder for updating a single Note entity.
type NoteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NoteMutation
}

// SetBinderID sets the "binder_id" field.
func (nuo *NoteUpdateOne) SetBinderID(u uint) *NoteUpdateOne {
	nuo.mutation.ResetBinderID()
	nuo.mutation.SetBinderID(u)
	return nuo
}

// SetNillableBinderID sets the "binder_id" field if the given value is not nil.
func (nuo *NoteUpdateOne) SetNillableBinderID(u *uint) *NoteUpdateOne {
	if u != nil {
		nuo.SetBinderID(*u)
	}
	return nuo
}

// AddBinderID adds u to the "binder_id" field.
func (nuo *NoteUpdateOne) AddBinderID(u int) *NoteUpdateOne {
	nuo.mutation.AddBinderID(u)
	return nuo
}

// ClearBinderID clears the value of the "binder_id" field.
func (nuo *NoteUpdateOne) ClearBinderID() *NoteUpdateOne {
	nuo.mutation.ClearBinderID()
	return nuo
}

// SetPublisher sets the "publisher" field.
func (nuo *NoteUpdateOne) SetPublisher(s string) *NoteUpdateOne {
	nuo.mutation.SetPublisher(s)
	return nuo
}

// SetNillablePublisher sets the "publisher" field if the given value is not nil.
func (nuo *NoteUpdateOne) SetNillablePublisher(s *string) *NoteUpdateOne {
	if s != nil {
		nuo.SetPublisher(*s)
	}
	return nuo
}

// SetComment sets the "comment" field.
func (nuo *NoteUpdateOne) SetComment(s string) *NoteUpdateOne {
	nuo.mutation.SetComment(s)
	return nuo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (nuo *NoteUpdateOne) SetNillableComment(s *string) *NoteUpdateOne {
	if s != nil {
		nuo.SetComment(*s)
	}
	return nuo
}

// ClearComment clears the value of the "comment" field.
func (nuo *NoteUpdateOne) ClearComment() *NoteUpdateOne {
	nuo.mutation.ClearComment()
	return nuo
}

// SetContent sets the "content" field.
func (nuo *NoteUpdateOne) SetContent(s string) *NoteUpdateOne {
	nuo.mutation.SetContent(s)
	return nuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (nuo *NoteUpdateOne) SetNillableContent(s *string) *NoteUpdateOne {
	if s != nil {
		nuo.SetContent(*s)
	}
	return nuo
}

// ClearContent clears the value of the "content" field.
func (nuo *NoteUpdateOne) ClearContent() *NoteUpdateOne {
	nuo.mutation.ClearContent()
	return nuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NoteUpdateOne) SetUpdatedAt(t time.Time) *NoteUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// Mutation returns the NoteMutation object of the builder.
func (nuo *NoteUpdateOne) Mutation() *NoteMutation {
	return nuo.mutation
}

// Where appends a list predicates to the NoteUpdate builder.
func (nuo *NoteUpdateOne) Where(ps ...predicate.Note) *NoteUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NoteUpdateOne) Select(field string, fields ...string) *NoteUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Note entity.
func (nuo *NoteUpdateOne) Save(ctx context.Context) (*Note, error) {
	nuo.defaults()
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NoteUpdateOne) SaveX(ctx context.Context) *Note {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NoteUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NoteUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NoteUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdatedAt(); !ok {
		v := note.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nuo *NoteUpdateOne) check() error {
	if v, ok := nuo.mutation.Publisher(); ok {
		if err := note.PublisherValidator(v); err != nil {
			return &ValidationError{Name: "publisher", err: fmt.Errorf(`ent: validator failed for field "Note.publisher": %w`, err)}
		}
	}
	return nil
}

func (nuo *NoteUpdateOne) sqlSave(ctx context.Context) (_node *Note, err error) {
	if err := nuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(note.Table, note.Columns, sqlgraph.NewFieldSpec(note.FieldID, field.TypeUint))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Note.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, note.FieldID)
		for _, f := range fields {
			if !note.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != note.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.BinderID(); ok {
		_spec.SetField(note.FieldBinderID, field.TypeUint, value)
	}
	if value, ok := nuo.mutation.AddedBinderID(); ok {
		_spec.AddField(note.FieldBinderID, field.TypeUint, value)
	}
	if nuo.mutation.BinderIDCleared() {
		_spec.ClearField(note.FieldBinderID, field.TypeUint)
	}
	if value, ok := nuo.mutation.Publisher(); ok {
		_spec.SetField(note.FieldPublisher, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Comment(); ok {
		_spec.SetField(note.FieldComment, field.TypeString, value)
	}
	if nuo.mutation.CommentCleared() {
		_spec.ClearField(note.FieldComment, field.TypeString)
	}
	if value, ok := nuo.mutation.Content(); ok {
		_spec.SetField(note.FieldContent, field.TypeString, value)
	}
	if nuo.mutation.ContentCleared() {
		_spec.ClearField(note.FieldContent, field.TypeString)
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.SetField(note.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Note{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{note.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}