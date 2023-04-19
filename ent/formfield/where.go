// Code generated by ent, DO NOT EDIT.

package formfield

import (
	"magnum/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.FormField {
	return predicate.FormField(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.FormField {
	return predicate.FormField(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.FormField {
	return predicate.FormField(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.FormField {
	return predicate.FormField(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.FormField {
	return predicate.FormField(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.FormField {
	return predicate.FormField(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.FormField {
	return predicate.FormField(sql.FieldLTE(FieldID, id))
}

// FormID applies equality check predicate on the "form_id" field. It's identical to FormIDEQ.
func FormID(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldFormID, v))
}

// FieldType applies equality check predicate on the "field_type" field. It's identical to FieldTypeEQ.
func FieldType(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldFieldType, v))
}

// FieldName applies equality check predicate on the "field_name" field. It's identical to FieldNameEQ.
func FieldName(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldFieldName, v))
}

// OrderIndex applies equality check predicate on the "order_index" field. It's identical to OrderIndexEQ.
func OrderIndex(v int) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldOrderIndex, v))
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldCreateAt, v))
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldUpdateAt, v))
}

// Disabled applies equality check predicate on the "disabled" field. It's identical to DisabledEQ.
func Disabled(v int) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldDisabled, v))
}

// FormIDEQ applies the EQ predicate on the "form_id" field.
func FormIDEQ(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldFormID, v))
}

// FormIDNEQ applies the NEQ predicate on the "form_id" field.
func FormIDNEQ(v string) predicate.FormField {
	return predicate.FormField(sql.FieldNEQ(FieldFormID, v))
}

// FormIDIn applies the In predicate on the "form_id" field.
func FormIDIn(vs ...string) predicate.FormField {
	return predicate.FormField(sql.FieldIn(FieldFormID, vs...))
}

// FormIDNotIn applies the NotIn predicate on the "form_id" field.
func FormIDNotIn(vs ...string) predicate.FormField {
	return predicate.FormField(sql.FieldNotIn(FieldFormID, vs...))
}

// FormIDGT applies the GT predicate on the "form_id" field.
func FormIDGT(v string) predicate.FormField {
	return predicate.FormField(sql.FieldGT(FieldFormID, v))
}

// FormIDGTE applies the GTE predicate on the "form_id" field.
func FormIDGTE(v string) predicate.FormField {
	return predicate.FormField(sql.FieldGTE(FieldFormID, v))
}

// FormIDLT applies the LT predicate on the "form_id" field.
func FormIDLT(v string) predicate.FormField {
	return predicate.FormField(sql.FieldLT(FieldFormID, v))
}

// FormIDLTE applies the LTE predicate on the "form_id" field.
func FormIDLTE(v string) predicate.FormField {
	return predicate.FormField(sql.FieldLTE(FieldFormID, v))
}

// FormIDContains applies the Contains predicate on the "form_id" field.
func FormIDContains(v string) predicate.FormField {
	return predicate.FormField(sql.FieldContains(FieldFormID, v))
}

// FormIDHasPrefix applies the HasPrefix predicate on the "form_id" field.
func FormIDHasPrefix(v string) predicate.FormField {
	return predicate.FormField(sql.FieldHasPrefix(FieldFormID, v))
}

// FormIDHasSuffix applies the HasSuffix predicate on the "form_id" field.
func FormIDHasSuffix(v string) predicate.FormField {
	return predicate.FormField(sql.FieldHasSuffix(FieldFormID, v))
}

// FormIDEqualFold applies the EqualFold predicate on the "form_id" field.
func FormIDEqualFold(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEqualFold(FieldFormID, v))
}

// FormIDContainsFold applies the ContainsFold predicate on the "form_id" field.
func FormIDContainsFold(v string) predicate.FormField {
	return predicate.FormField(sql.FieldContainsFold(FieldFormID, v))
}

// FieldTypeEQ applies the EQ predicate on the "field_type" field.
func FieldTypeEQ(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldFieldType, v))
}

// FieldTypeNEQ applies the NEQ predicate on the "field_type" field.
func FieldTypeNEQ(v string) predicate.FormField {
	return predicate.FormField(sql.FieldNEQ(FieldFieldType, v))
}

// FieldTypeIn applies the In predicate on the "field_type" field.
func FieldTypeIn(vs ...string) predicate.FormField {
	return predicate.FormField(sql.FieldIn(FieldFieldType, vs...))
}

// FieldTypeNotIn applies the NotIn predicate on the "field_type" field.
func FieldTypeNotIn(vs ...string) predicate.FormField {
	return predicate.FormField(sql.FieldNotIn(FieldFieldType, vs...))
}

// FieldTypeGT applies the GT predicate on the "field_type" field.
func FieldTypeGT(v string) predicate.FormField {
	return predicate.FormField(sql.FieldGT(FieldFieldType, v))
}

// FieldTypeGTE applies the GTE predicate on the "field_type" field.
func FieldTypeGTE(v string) predicate.FormField {
	return predicate.FormField(sql.FieldGTE(FieldFieldType, v))
}

// FieldTypeLT applies the LT predicate on the "field_type" field.
func FieldTypeLT(v string) predicate.FormField {
	return predicate.FormField(sql.FieldLT(FieldFieldType, v))
}

// FieldTypeLTE applies the LTE predicate on the "field_type" field.
func FieldTypeLTE(v string) predicate.FormField {
	return predicate.FormField(sql.FieldLTE(FieldFieldType, v))
}

// FieldTypeContains applies the Contains predicate on the "field_type" field.
func FieldTypeContains(v string) predicate.FormField {
	return predicate.FormField(sql.FieldContains(FieldFieldType, v))
}

// FieldTypeHasPrefix applies the HasPrefix predicate on the "field_type" field.
func FieldTypeHasPrefix(v string) predicate.FormField {
	return predicate.FormField(sql.FieldHasPrefix(FieldFieldType, v))
}

// FieldTypeHasSuffix applies the HasSuffix predicate on the "field_type" field.
func FieldTypeHasSuffix(v string) predicate.FormField {
	return predicate.FormField(sql.FieldHasSuffix(FieldFieldType, v))
}

// FieldTypeEqualFold applies the EqualFold predicate on the "field_type" field.
func FieldTypeEqualFold(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEqualFold(FieldFieldType, v))
}

// FieldTypeContainsFold applies the ContainsFold predicate on the "field_type" field.
func FieldTypeContainsFold(v string) predicate.FormField {
	return predicate.FormField(sql.FieldContainsFold(FieldFieldType, v))
}

// FieldNameEQ applies the EQ predicate on the "field_name" field.
func FieldNameEQ(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldFieldName, v))
}

// FieldNameNEQ applies the NEQ predicate on the "field_name" field.
func FieldNameNEQ(v string) predicate.FormField {
	return predicate.FormField(sql.FieldNEQ(FieldFieldName, v))
}

// FieldNameIn applies the In predicate on the "field_name" field.
func FieldNameIn(vs ...string) predicate.FormField {
	return predicate.FormField(sql.FieldIn(FieldFieldName, vs...))
}

// FieldNameNotIn applies the NotIn predicate on the "field_name" field.
func FieldNameNotIn(vs ...string) predicate.FormField {
	return predicate.FormField(sql.FieldNotIn(FieldFieldName, vs...))
}

// FieldNameGT applies the GT predicate on the "field_name" field.
func FieldNameGT(v string) predicate.FormField {
	return predicate.FormField(sql.FieldGT(FieldFieldName, v))
}

// FieldNameGTE applies the GTE predicate on the "field_name" field.
func FieldNameGTE(v string) predicate.FormField {
	return predicate.FormField(sql.FieldGTE(FieldFieldName, v))
}

// FieldNameLT applies the LT predicate on the "field_name" field.
func FieldNameLT(v string) predicate.FormField {
	return predicate.FormField(sql.FieldLT(FieldFieldName, v))
}

// FieldNameLTE applies the LTE predicate on the "field_name" field.
func FieldNameLTE(v string) predicate.FormField {
	return predicate.FormField(sql.FieldLTE(FieldFieldName, v))
}

// FieldNameContains applies the Contains predicate on the "field_name" field.
func FieldNameContains(v string) predicate.FormField {
	return predicate.FormField(sql.FieldContains(FieldFieldName, v))
}

// FieldNameHasPrefix applies the HasPrefix predicate on the "field_name" field.
func FieldNameHasPrefix(v string) predicate.FormField {
	return predicate.FormField(sql.FieldHasPrefix(FieldFieldName, v))
}

// FieldNameHasSuffix applies the HasSuffix predicate on the "field_name" field.
func FieldNameHasSuffix(v string) predicate.FormField {
	return predicate.FormField(sql.FieldHasSuffix(FieldFieldName, v))
}

// FieldNameEqualFold applies the EqualFold predicate on the "field_name" field.
func FieldNameEqualFold(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEqualFold(FieldFieldName, v))
}

// FieldNameContainsFold applies the ContainsFold predicate on the "field_name" field.
func FieldNameContainsFold(v string) predicate.FormField {
	return predicate.FormField(sql.FieldContainsFold(FieldFieldName, v))
}

// OrderIndexEQ applies the EQ predicate on the "order_index" field.
func OrderIndexEQ(v int) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldOrderIndex, v))
}

// OrderIndexNEQ applies the NEQ predicate on the "order_index" field.
func OrderIndexNEQ(v int) predicate.FormField {
	return predicate.FormField(sql.FieldNEQ(FieldOrderIndex, v))
}

// OrderIndexIn applies the In predicate on the "order_index" field.
func OrderIndexIn(vs ...int) predicate.FormField {
	return predicate.FormField(sql.FieldIn(FieldOrderIndex, vs...))
}

// OrderIndexNotIn applies the NotIn predicate on the "order_index" field.
func OrderIndexNotIn(vs ...int) predicate.FormField {
	return predicate.FormField(sql.FieldNotIn(FieldOrderIndex, vs...))
}

// OrderIndexGT applies the GT predicate on the "order_index" field.
func OrderIndexGT(v int) predicate.FormField {
	return predicate.FormField(sql.FieldGT(FieldOrderIndex, v))
}

// OrderIndexGTE applies the GTE predicate on the "order_index" field.
func OrderIndexGTE(v int) predicate.FormField {
	return predicate.FormField(sql.FieldGTE(FieldOrderIndex, v))
}

// OrderIndexLT applies the LT predicate on the "order_index" field.
func OrderIndexLT(v int) predicate.FormField {
	return predicate.FormField(sql.FieldLT(FieldOrderIndex, v))
}

// OrderIndexLTE applies the LTE predicate on the "order_index" field.
func OrderIndexLTE(v int) predicate.FormField {
	return predicate.FormField(sql.FieldLTE(FieldOrderIndex, v))
}

// OrderIndexIsNil applies the IsNil predicate on the "order_index" field.
func OrderIndexIsNil() predicate.FormField {
	return predicate.FormField(sql.FieldIsNull(FieldOrderIndex))
}

// OrderIndexNotNil applies the NotNil predicate on the "order_index" field.
func OrderIndexNotNil() predicate.FormField {
	return predicate.FormField(sql.FieldNotNull(FieldOrderIndex))
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldCreateAt, v))
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldNEQ(FieldCreateAt, v))
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldIn(FieldCreateAt, vs...))
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldNotIn(FieldCreateAt, vs...))
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldGT(FieldCreateAt, v))
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldGTE(FieldCreateAt, v))
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldLT(FieldCreateAt, v))
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldLTE(FieldCreateAt, v))
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldUpdateAt, v))
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldNEQ(FieldUpdateAt, v))
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldIn(FieldUpdateAt, vs...))
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldNotIn(FieldUpdateAt, vs...))
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldGT(FieldUpdateAt, v))
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldGTE(FieldUpdateAt, v))
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldLT(FieldUpdateAt, v))
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v time.Time) predicate.FormField {
	return predicate.FormField(sql.FieldLTE(FieldUpdateAt, v))
}

// UpdateAtIsNil applies the IsNil predicate on the "update_at" field.
func UpdateAtIsNil() predicate.FormField {
	return predicate.FormField(sql.FieldIsNull(FieldUpdateAt))
}

// UpdateAtNotNil applies the NotNil predicate on the "update_at" field.
func UpdateAtNotNil() predicate.FormField {
	return predicate.FormField(sql.FieldNotNull(FieldUpdateAt))
}

// DisabledEQ applies the EQ predicate on the "disabled" field.
func DisabledEQ(v int) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldDisabled, v))
}

// DisabledNEQ applies the NEQ predicate on the "disabled" field.
func DisabledNEQ(v int) predicate.FormField {
	return predicate.FormField(sql.FieldNEQ(FieldDisabled, v))
}

// DisabledIn applies the In predicate on the "disabled" field.
func DisabledIn(vs ...int) predicate.FormField {
	return predicate.FormField(sql.FieldIn(FieldDisabled, vs...))
}

// DisabledNotIn applies the NotIn predicate on the "disabled" field.
func DisabledNotIn(vs ...int) predicate.FormField {
	return predicate.FormField(sql.FieldNotIn(FieldDisabled, vs...))
}

// DisabledGT applies the GT predicate on the "disabled" field.
func DisabledGT(v int) predicate.FormField {
	return predicate.FormField(sql.FieldGT(FieldDisabled, v))
}

// DisabledGTE applies the GTE predicate on the "disabled" field.
func DisabledGTE(v int) predicate.FormField {
	return predicate.FormField(sql.FieldGTE(FieldDisabled, v))
}

// DisabledLT applies the LT predicate on the "disabled" field.
func DisabledLT(v int) predicate.FormField {
	return predicate.FormField(sql.FieldLT(FieldDisabled, v))
}

// DisabledLTE applies the LTE predicate on the "disabled" field.
func DisabledLTE(v int) predicate.FormField {
	return predicate.FormField(sql.FieldLTE(FieldDisabled, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FormField) predicate.FormField {
	return predicate.FormField(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FormField) predicate.FormField {
	return predicate.FormField(func(s *sql.Selector) {
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
func Not(p predicate.FormField) predicate.FormField {
	return predicate.FormField(func(s *sql.Selector) {
		p(s.Not())
	})
}
