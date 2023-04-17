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

// FieldLabel applies equality check predicate on the "field_label" field. It's identical to FieldLabelEQ.
func FieldLabel(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldFieldLabel, v))
}

// FiledName applies equality check predicate on the "filed_name" field. It's identical to FiledNameEQ.
func FiledName(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldFiledName, v))
}

// IsRequired applies equality check predicate on the "is_required" field. It's identical to IsRequiredEQ.
func IsRequired(v int) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldIsRequired, v))
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

// Options applies equality check predicate on the "options" field. It's identical to OptionsEQ.
func Options(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldOptions, v))
}

// Placeholder applies equality check predicate on the "placeholder" field. It's identical to PlaceholderEQ.
func Placeholder(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldPlaceholder, v))
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

// FieldLabelEQ applies the EQ predicate on the "field_label" field.
func FieldLabelEQ(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldFieldLabel, v))
}

// FieldLabelNEQ applies the NEQ predicate on the "field_label" field.
func FieldLabelNEQ(v string) predicate.FormField {
	return predicate.FormField(sql.FieldNEQ(FieldFieldLabel, v))
}

// FieldLabelIn applies the In predicate on the "field_label" field.
func FieldLabelIn(vs ...string) predicate.FormField {
	return predicate.FormField(sql.FieldIn(FieldFieldLabel, vs...))
}

// FieldLabelNotIn applies the NotIn predicate on the "field_label" field.
func FieldLabelNotIn(vs ...string) predicate.FormField {
	return predicate.FormField(sql.FieldNotIn(FieldFieldLabel, vs...))
}

// FieldLabelGT applies the GT predicate on the "field_label" field.
func FieldLabelGT(v string) predicate.FormField {
	return predicate.FormField(sql.FieldGT(FieldFieldLabel, v))
}

// FieldLabelGTE applies the GTE predicate on the "field_label" field.
func FieldLabelGTE(v string) predicate.FormField {
	return predicate.FormField(sql.FieldGTE(FieldFieldLabel, v))
}

// FieldLabelLT applies the LT predicate on the "field_label" field.
func FieldLabelLT(v string) predicate.FormField {
	return predicate.FormField(sql.FieldLT(FieldFieldLabel, v))
}

// FieldLabelLTE applies the LTE predicate on the "field_label" field.
func FieldLabelLTE(v string) predicate.FormField {
	return predicate.FormField(sql.FieldLTE(FieldFieldLabel, v))
}

// FieldLabelContains applies the Contains predicate on the "field_label" field.
func FieldLabelContains(v string) predicate.FormField {
	return predicate.FormField(sql.FieldContains(FieldFieldLabel, v))
}

// FieldLabelHasPrefix applies the HasPrefix predicate on the "field_label" field.
func FieldLabelHasPrefix(v string) predicate.FormField {
	return predicate.FormField(sql.FieldHasPrefix(FieldFieldLabel, v))
}

// FieldLabelHasSuffix applies the HasSuffix predicate on the "field_label" field.
func FieldLabelHasSuffix(v string) predicate.FormField {
	return predicate.FormField(sql.FieldHasSuffix(FieldFieldLabel, v))
}

// FieldLabelEqualFold applies the EqualFold predicate on the "field_label" field.
func FieldLabelEqualFold(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEqualFold(FieldFieldLabel, v))
}

// FieldLabelContainsFold applies the ContainsFold predicate on the "field_label" field.
func FieldLabelContainsFold(v string) predicate.FormField {
	return predicate.FormField(sql.FieldContainsFold(FieldFieldLabel, v))
}

// FiledNameEQ applies the EQ predicate on the "filed_name" field.
func FiledNameEQ(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldFiledName, v))
}

// FiledNameNEQ applies the NEQ predicate on the "filed_name" field.
func FiledNameNEQ(v string) predicate.FormField {
	return predicate.FormField(sql.FieldNEQ(FieldFiledName, v))
}

// FiledNameIn applies the In predicate on the "filed_name" field.
func FiledNameIn(vs ...string) predicate.FormField {
	return predicate.FormField(sql.FieldIn(FieldFiledName, vs...))
}

// FiledNameNotIn applies the NotIn predicate on the "filed_name" field.
func FiledNameNotIn(vs ...string) predicate.FormField {
	return predicate.FormField(sql.FieldNotIn(FieldFiledName, vs...))
}

// FiledNameGT applies the GT predicate on the "filed_name" field.
func FiledNameGT(v string) predicate.FormField {
	return predicate.FormField(sql.FieldGT(FieldFiledName, v))
}

// FiledNameGTE applies the GTE predicate on the "filed_name" field.
func FiledNameGTE(v string) predicate.FormField {
	return predicate.FormField(sql.FieldGTE(FieldFiledName, v))
}

// FiledNameLT applies the LT predicate on the "filed_name" field.
func FiledNameLT(v string) predicate.FormField {
	return predicate.FormField(sql.FieldLT(FieldFiledName, v))
}

// FiledNameLTE applies the LTE predicate on the "filed_name" field.
func FiledNameLTE(v string) predicate.FormField {
	return predicate.FormField(sql.FieldLTE(FieldFiledName, v))
}

// FiledNameContains applies the Contains predicate on the "filed_name" field.
func FiledNameContains(v string) predicate.FormField {
	return predicate.FormField(sql.FieldContains(FieldFiledName, v))
}

// FiledNameHasPrefix applies the HasPrefix predicate on the "filed_name" field.
func FiledNameHasPrefix(v string) predicate.FormField {
	return predicate.FormField(sql.FieldHasPrefix(FieldFiledName, v))
}

// FiledNameHasSuffix applies the HasSuffix predicate on the "filed_name" field.
func FiledNameHasSuffix(v string) predicate.FormField {
	return predicate.FormField(sql.FieldHasSuffix(FieldFiledName, v))
}

// FiledNameEqualFold applies the EqualFold predicate on the "filed_name" field.
func FiledNameEqualFold(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEqualFold(FieldFiledName, v))
}

// FiledNameContainsFold applies the ContainsFold predicate on the "filed_name" field.
func FiledNameContainsFold(v string) predicate.FormField {
	return predicate.FormField(sql.FieldContainsFold(FieldFiledName, v))
}

// IsRequiredEQ applies the EQ predicate on the "is_required" field.
func IsRequiredEQ(v int) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldIsRequired, v))
}

// IsRequiredNEQ applies the NEQ predicate on the "is_required" field.
func IsRequiredNEQ(v int) predicate.FormField {
	return predicate.FormField(sql.FieldNEQ(FieldIsRequired, v))
}

// IsRequiredIn applies the In predicate on the "is_required" field.
func IsRequiredIn(vs ...int) predicate.FormField {
	return predicate.FormField(sql.FieldIn(FieldIsRequired, vs...))
}

// IsRequiredNotIn applies the NotIn predicate on the "is_required" field.
func IsRequiredNotIn(vs ...int) predicate.FormField {
	return predicate.FormField(sql.FieldNotIn(FieldIsRequired, vs...))
}

// IsRequiredGT applies the GT predicate on the "is_required" field.
func IsRequiredGT(v int) predicate.FormField {
	return predicate.FormField(sql.FieldGT(FieldIsRequired, v))
}

// IsRequiredGTE applies the GTE predicate on the "is_required" field.
func IsRequiredGTE(v int) predicate.FormField {
	return predicate.FormField(sql.FieldGTE(FieldIsRequired, v))
}

// IsRequiredLT applies the LT predicate on the "is_required" field.
func IsRequiredLT(v int) predicate.FormField {
	return predicate.FormField(sql.FieldLT(FieldIsRequired, v))
}

// IsRequiredLTE applies the LTE predicate on the "is_required" field.
func IsRequiredLTE(v int) predicate.FormField {
	return predicate.FormField(sql.FieldLTE(FieldIsRequired, v))
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

// OptionsEQ applies the EQ predicate on the "options" field.
func OptionsEQ(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldOptions, v))
}

// OptionsNEQ applies the NEQ predicate on the "options" field.
func OptionsNEQ(v string) predicate.FormField {
	return predicate.FormField(sql.FieldNEQ(FieldOptions, v))
}

// OptionsIn applies the In predicate on the "options" field.
func OptionsIn(vs ...string) predicate.FormField {
	return predicate.FormField(sql.FieldIn(FieldOptions, vs...))
}

// OptionsNotIn applies the NotIn predicate on the "options" field.
func OptionsNotIn(vs ...string) predicate.FormField {
	return predicate.FormField(sql.FieldNotIn(FieldOptions, vs...))
}

// OptionsGT applies the GT predicate on the "options" field.
func OptionsGT(v string) predicate.FormField {
	return predicate.FormField(sql.FieldGT(FieldOptions, v))
}

// OptionsGTE applies the GTE predicate on the "options" field.
func OptionsGTE(v string) predicate.FormField {
	return predicate.FormField(sql.FieldGTE(FieldOptions, v))
}

// OptionsLT applies the LT predicate on the "options" field.
func OptionsLT(v string) predicate.FormField {
	return predicate.FormField(sql.FieldLT(FieldOptions, v))
}

// OptionsLTE applies the LTE predicate on the "options" field.
func OptionsLTE(v string) predicate.FormField {
	return predicate.FormField(sql.FieldLTE(FieldOptions, v))
}

// OptionsContains applies the Contains predicate on the "options" field.
func OptionsContains(v string) predicate.FormField {
	return predicate.FormField(sql.FieldContains(FieldOptions, v))
}

// OptionsHasPrefix applies the HasPrefix predicate on the "options" field.
func OptionsHasPrefix(v string) predicate.FormField {
	return predicate.FormField(sql.FieldHasPrefix(FieldOptions, v))
}

// OptionsHasSuffix applies the HasSuffix predicate on the "options" field.
func OptionsHasSuffix(v string) predicate.FormField {
	return predicate.FormField(sql.FieldHasSuffix(FieldOptions, v))
}

// OptionsIsNil applies the IsNil predicate on the "options" field.
func OptionsIsNil() predicate.FormField {
	return predicate.FormField(sql.FieldIsNull(FieldOptions))
}

// OptionsNotNil applies the NotNil predicate on the "options" field.
func OptionsNotNil() predicate.FormField {
	return predicate.FormField(sql.FieldNotNull(FieldOptions))
}

// OptionsEqualFold applies the EqualFold predicate on the "options" field.
func OptionsEqualFold(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEqualFold(FieldOptions, v))
}

// OptionsContainsFold applies the ContainsFold predicate on the "options" field.
func OptionsContainsFold(v string) predicate.FormField {
	return predicate.FormField(sql.FieldContainsFold(FieldOptions, v))
}

// PlaceholderEQ applies the EQ predicate on the "placeholder" field.
func PlaceholderEQ(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEQ(FieldPlaceholder, v))
}

// PlaceholderNEQ applies the NEQ predicate on the "placeholder" field.
func PlaceholderNEQ(v string) predicate.FormField {
	return predicate.FormField(sql.FieldNEQ(FieldPlaceholder, v))
}

// PlaceholderIn applies the In predicate on the "placeholder" field.
func PlaceholderIn(vs ...string) predicate.FormField {
	return predicate.FormField(sql.FieldIn(FieldPlaceholder, vs...))
}

// PlaceholderNotIn applies the NotIn predicate on the "placeholder" field.
func PlaceholderNotIn(vs ...string) predicate.FormField {
	return predicate.FormField(sql.FieldNotIn(FieldPlaceholder, vs...))
}

// PlaceholderGT applies the GT predicate on the "placeholder" field.
func PlaceholderGT(v string) predicate.FormField {
	return predicate.FormField(sql.FieldGT(FieldPlaceholder, v))
}

// PlaceholderGTE applies the GTE predicate on the "placeholder" field.
func PlaceholderGTE(v string) predicate.FormField {
	return predicate.FormField(sql.FieldGTE(FieldPlaceholder, v))
}

// PlaceholderLT applies the LT predicate on the "placeholder" field.
func PlaceholderLT(v string) predicate.FormField {
	return predicate.FormField(sql.FieldLT(FieldPlaceholder, v))
}

// PlaceholderLTE applies the LTE predicate on the "placeholder" field.
func PlaceholderLTE(v string) predicate.FormField {
	return predicate.FormField(sql.FieldLTE(FieldPlaceholder, v))
}

// PlaceholderContains applies the Contains predicate on the "placeholder" field.
func PlaceholderContains(v string) predicate.FormField {
	return predicate.FormField(sql.FieldContains(FieldPlaceholder, v))
}

// PlaceholderHasPrefix applies the HasPrefix predicate on the "placeholder" field.
func PlaceholderHasPrefix(v string) predicate.FormField {
	return predicate.FormField(sql.FieldHasPrefix(FieldPlaceholder, v))
}

// PlaceholderHasSuffix applies the HasSuffix predicate on the "placeholder" field.
func PlaceholderHasSuffix(v string) predicate.FormField {
	return predicate.FormField(sql.FieldHasSuffix(FieldPlaceholder, v))
}

// PlaceholderIsNil applies the IsNil predicate on the "placeholder" field.
func PlaceholderIsNil() predicate.FormField {
	return predicate.FormField(sql.FieldIsNull(FieldPlaceholder))
}

// PlaceholderNotNil applies the NotNil predicate on the "placeholder" field.
func PlaceholderNotNil() predicate.FormField {
	return predicate.FormField(sql.FieldNotNull(FieldPlaceholder))
}

// PlaceholderEqualFold applies the EqualFold predicate on the "placeholder" field.
func PlaceholderEqualFold(v string) predicate.FormField {
	return predicate.FormField(sql.FieldEqualFold(FieldPlaceholder, v))
}

// PlaceholderContainsFold applies the ContainsFold predicate on the "placeholder" field.
func PlaceholderContainsFold(v string) predicate.FormField {
	return predicate.FormField(sql.FieldContainsFold(FieldPlaceholder, v))
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