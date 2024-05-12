// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/enuesaa/pinit/pkg/ent/action"
	"github.com/enuesaa/pinit/pkg/ent/appconf"
	"github.com/enuesaa/pinit/pkg/ent/binder"
	"github.com/enuesaa/pinit/pkg/ent/note"
	"github.com/enuesaa/pinit/pkg/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	actionFields := schema.Action{}.Fields()
	_ = actionFields
	// actionDescName is the schema descriptor for name field.
	actionDescName := actionFields[1].Descriptor()
	// action.NameValidator is a validator for the "name" field. It is called by the builders before save.
	action.NameValidator = actionDescName.Validators[0].(func(string) error)
	// actionDescCreatedAt is the schema descriptor for created_at field.
	actionDescCreatedAt := actionFields[3].Descriptor()
	// action.DefaultCreatedAt holds the default value on creation for the created_at field.
	action.DefaultCreatedAt = actionDescCreatedAt.Default.(func() time.Time)
	// actionDescUpdatedAt is the schema descriptor for updated_at field.
	actionDescUpdatedAt := actionFields[4].Descriptor()
	// action.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	action.DefaultUpdatedAt = actionDescUpdatedAt.Default.(func() time.Time)
	// action.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	action.UpdateDefaultUpdatedAt = actionDescUpdatedAt.UpdateDefault.(func() time.Time)
	appconfFields := schema.Appconf{}.Fields()
	_ = appconfFields
	// appconfDescCreatedAt is the schema descriptor for created_at field.
	appconfDescCreatedAt := appconfFields[3].Descriptor()
	// appconf.DefaultCreatedAt holds the default value on creation for the created_at field.
	appconf.DefaultCreatedAt = appconfDescCreatedAt.Default.(func() time.Time)
	// appconfDescUpdatedAt is the schema descriptor for updated_at field.
	appconfDescUpdatedAt := appconfFields[4].Descriptor()
	// appconf.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	appconf.DefaultUpdatedAt = appconfDescUpdatedAt.Default.(func() time.Time)
	// appconf.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	appconf.UpdateDefaultUpdatedAt = appconfDescUpdatedAt.UpdateDefault.(func() time.Time)
	binderFields := schema.Binder{}.Fields()
	_ = binderFields
	// binderDescName is the schema descriptor for name field.
	binderDescName := binderFields[1].Descriptor()
	// binder.NameValidator is a validator for the "name" field. It is called by the builders before save.
	binder.NameValidator = binderDescName.Validators[0].(func(string) error)
	// binderDescCategory is the schema descriptor for category field.
	binderDescCategory := binderFields[2].Descriptor()
	// binder.CategoryValidator is a validator for the "category" field. It is called by the builders before save.
	binder.CategoryValidator = binderDescCategory.Validators[0].(func(string) error)
	// binderDescCreatedAt is the schema descriptor for created_at field.
	binderDescCreatedAt := binderFields[4].Descriptor()
	// binder.DefaultCreatedAt holds the default value on creation for the created_at field.
	binder.DefaultCreatedAt = binderDescCreatedAt.Default.(func() time.Time)
	// binderDescUpdatedAt is the schema descriptor for updated_at field.
	binderDescUpdatedAt := binderFields[5].Descriptor()
	// binder.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	binder.DefaultUpdatedAt = binderDescUpdatedAt.Default.(func() time.Time)
	// binder.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	binder.UpdateDefaultUpdatedAt = binderDescUpdatedAt.UpdateDefault.(func() time.Time)
	noteFields := schema.Note{}.Fields()
	_ = noteFields
	// noteDescPublisher is the schema descriptor for publisher field.
	noteDescPublisher := noteFields[2].Descriptor()
	// note.PublisherValidator is a validator for the "publisher" field. It is called by the builders before save.
	note.PublisherValidator = noteDescPublisher.Validators[0].(func(string) error)
	// noteDescCreatedAt is the schema descriptor for created_at field.
	noteDescCreatedAt := noteFields[5].Descriptor()
	// note.DefaultCreatedAt holds the default value on creation for the created_at field.
	note.DefaultCreatedAt = noteDescCreatedAt.Default.(func() time.Time)
	// noteDescUpdatedAt is the schema descriptor for updated_at field.
	noteDescUpdatedAt := noteFields[6].Descriptor()
	// note.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	note.DefaultUpdatedAt = noteDescUpdatedAt.Default.(func() time.Time)
	// note.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	note.UpdateDefaultUpdatedAt = noteDescUpdatedAt.UpdateDefault.(func() time.Time)
}
