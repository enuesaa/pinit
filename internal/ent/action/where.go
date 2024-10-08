// Code generated by ent, DO NOT EDIT.

package action

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/enuesaa/pinit/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint) predicate.Action {
	return predicate.Action(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint) predicate.Action {
	return predicate.Action(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint) predicate.Action {
	return predicate.Action(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint) predicate.Action {
	return predicate.Action(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint) predicate.Action {
	return predicate.Action(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint) predicate.Action {
	return predicate.Action(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldName, v))
}

// Template applies equality check predicate on the "template" field. It's identical to TemplateEQ.
func Template(v string) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldTemplate, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Action {
	return predicate.Action(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Action {
	return predicate.Action(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Action {
	return predicate.Action(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Action {
	return predicate.Action(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Action {
	return predicate.Action(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Action {
	return predicate.Action(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Action {
	return predicate.Action(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Action {
	return predicate.Action(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Action {
	return predicate.Action(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Action {
	return predicate.Action(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Action {
	return predicate.Action(sql.FieldContainsFold(FieldName, v))
}

// TemplateEQ applies the EQ predicate on the "template" field.
func TemplateEQ(v string) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldTemplate, v))
}

// TemplateNEQ applies the NEQ predicate on the "template" field.
func TemplateNEQ(v string) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldTemplate, v))
}

// TemplateIn applies the In predicate on the "template" field.
func TemplateIn(vs ...string) predicate.Action {
	return predicate.Action(sql.FieldIn(FieldTemplate, vs...))
}

// TemplateNotIn applies the NotIn predicate on the "template" field.
func TemplateNotIn(vs ...string) predicate.Action {
	return predicate.Action(sql.FieldNotIn(FieldTemplate, vs...))
}

// TemplateGT applies the GT predicate on the "template" field.
func TemplateGT(v string) predicate.Action {
	return predicate.Action(sql.FieldGT(FieldTemplate, v))
}

// TemplateGTE applies the GTE predicate on the "template" field.
func TemplateGTE(v string) predicate.Action {
	return predicate.Action(sql.FieldGTE(FieldTemplate, v))
}

// TemplateLT applies the LT predicate on the "template" field.
func TemplateLT(v string) predicate.Action {
	return predicate.Action(sql.FieldLT(FieldTemplate, v))
}

// TemplateLTE applies the LTE predicate on the "template" field.
func TemplateLTE(v string) predicate.Action {
	return predicate.Action(sql.FieldLTE(FieldTemplate, v))
}

// TemplateContains applies the Contains predicate on the "template" field.
func TemplateContains(v string) predicate.Action {
	return predicate.Action(sql.FieldContains(FieldTemplate, v))
}

// TemplateHasPrefix applies the HasPrefix predicate on the "template" field.
func TemplateHasPrefix(v string) predicate.Action {
	return predicate.Action(sql.FieldHasPrefix(FieldTemplate, v))
}

// TemplateHasSuffix applies the HasSuffix predicate on the "template" field.
func TemplateHasSuffix(v string) predicate.Action {
	return predicate.Action(sql.FieldHasSuffix(FieldTemplate, v))
}

// TemplateEqualFold applies the EqualFold predicate on the "template" field.
func TemplateEqualFold(v string) predicate.Action {
	return predicate.Action(sql.FieldEqualFold(FieldTemplate, v))
}

// TemplateContainsFold applies the ContainsFold predicate on the "template" field.
func TemplateContainsFold(v string) predicate.Action {
	return predicate.Action(sql.FieldContainsFold(FieldTemplate, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Action {
	return predicate.Action(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Action {
	return predicate.Action(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Action {
	return predicate.Action(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Action {
	return predicate.Action(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Action {
	return predicate.Action(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Action) predicate.Action {
	return predicate.Action(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Action) predicate.Action {
	return predicate.Action(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Action) predicate.Action {
	return predicate.Action(sql.NotPredicates(p))
}
