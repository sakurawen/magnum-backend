// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"magnum/ent/formfieldconfig"
	"magnum/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FormFieldConfigUpdate is the builder for updating FormFieldConfig entities.
type FormFieldConfigUpdate struct {
	config
	hooks    []Hook
	mutation *FormFieldConfigMutation
}

// Where appends a list predicates to the FormFieldConfigUpdate builder.
func (ffcu *FormFieldConfigUpdate) Where(ps ...predicate.FormFieldConfig) *FormFieldConfigUpdate {
	ffcu.mutation.Where(ps...)
	return ffcu
}

// SetFieldID sets the "field_id" field.
func (ffcu *FormFieldConfigUpdate) SetFieldID(s string) *FormFieldConfigUpdate {
	ffcu.mutation.SetFieldID(s)
	return ffcu
}

// SetFormID sets the "form_id" field.
func (ffcu *FormFieldConfigUpdate) SetFormID(s string) *FormFieldConfigUpdate {
	ffcu.mutation.SetFormID(s)
	return ffcu
}

// SetKey sets the "key" field.
func (ffcu *FormFieldConfigUpdate) SetKey(s string) *FormFieldConfigUpdate {
	ffcu.mutation.SetKey(s)
	return ffcu
}

// SetType sets the "type" field.
func (ffcu *FormFieldConfigUpdate) SetType(s string) *FormFieldConfigUpdate {
	ffcu.mutation.SetType(s)
	return ffcu
}

// SetValue sets the "value" field.
func (ffcu *FormFieldConfigUpdate) SetValue(s string) *FormFieldConfigUpdate {
	ffcu.mutation.SetValue(s)
	return ffcu
}

// SetJSONStringValue sets the "json_string_value" field.
func (ffcu *FormFieldConfigUpdate) SetJSONStringValue(s string) *FormFieldConfigUpdate {
	ffcu.mutation.SetJSONStringValue(s)
	return ffcu
}

// SetNillableJSONStringValue sets the "json_string_value" field if the given value is not nil.
func (ffcu *FormFieldConfigUpdate) SetNillableJSONStringValue(s *string) *FormFieldConfigUpdate {
	if s != nil {
		ffcu.SetJSONStringValue(*s)
	}
	return ffcu
}

// ClearJSONStringValue clears the value of the "json_string_value" field.
func (ffcu *FormFieldConfigUpdate) ClearJSONStringValue() *FormFieldConfigUpdate {
	ffcu.mutation.ClearJSONStringValue()
	return ffcu
}

// SetText sets the "text" field.
func (ffcu *FormFieldConfigUpdate) SetText(s string) *FormFieldConfigUpdate {
	ffcu.mutation.SetText(s)
	return ffcu
}

// SetOrderIndex sets the "order_index" field.
func (ffcu *FormFieldConfigUpdate) SetOrderIndex(i int) *FormFieldConfigUpdate {
	ffcu.mutation.ResetOrderIndex()
	ffcu.mutation.SetOrderIndex(i)
	return ffcu
}

// AddOrderIndex adds i to the "order_index" field.
func (ffcu *FormFieldConfigUpdate) AddOrderIndex(i int) *FormFieldConfigUpdate {
	ffcu.mutation.AddOrderIndex(i)
	return ffcu
}

// SetDisabled sets the "disabled" field.
func (ffcu *FormFieldConfigUpdate) SetDisabled(i int) *FormFieldConfigUpdate {
	ffcu.mutation.ResetDisabled()
	ffcu.mutation.SetDisabled(i)
	return ffcu
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (ffcu *FormFieldConfigUpdate) SetNillableDisabled(i *int) *FormFieldConfigUpdate {
	if i != nil {
		ffcu.SetDisabled(*i)
	}
	return ffcu
}

// AddDisabled adds i to the "disabled" field.
func (ffcu *FormFieldConfigUpdate) AddDisabled(i int) *FormFieldConfigUpdate {
	ffcu.mutation.AddDisabled(i)
	return ffcu
}

// Mutation returns the FormFieldConfigMutation object of the builder.
func (ffcu *FormFieldConfigUpdate) Mutation() *FormFieldConfigMutation {
	return ffcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ffcu *FormFieldConfigUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, FormFieldConfigMutation](ctx, ffcu.sqlSave, ffcu.mutation, ffcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ffcu *FormFieldConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := ffcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ffcu *FormFieldConfigUpdate) Exec(ctx context.Context) error {
	_, err := ffcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ffcu *FormFieldConfigUpdate) ExecX(ctx context.Context) {
	if err := ffcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ffcu *FormFieldConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(formfieldconfig.Table, formfieldconfig.Columns, sqlgraph.NewFieldSpec(formfieldconfig.FieldID, field.TypeString))
	if ps := ffcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ffcu.mutation.FieldID(); ok {
		_spec.SetField(formfieldconfig.FieldFieldID, field.TypeString, value)
	}
	if value, ok := ffcu.mutation.FormID(); ok {
		_spec.SetField(formfieldconfig.FieldFormID, field.TypeString, value)
	}
	if value, ok := ffcu.mutation.Key(); ok {
		_spec.SetField(formfieldconfig.FieldKey, field.TypeString, value)
	}
	if value, ok := ffcu.mutation.GetType(); ok {
		_spec.SetField(formfieldconfig.FieldType, field.TypeString, value)
	}
	if value, ok := ffcu.mutation.Value(); ok {
		_spec.SetField(formfieldconfig.FieldValue, field.TypeString, value)
	}
	if value, ok := ffcu.mutation.JSONStringValue(); ok {
		_spec.SetField(formfieldconfig.FieldJSONStringValue, field.TypeString, value)
	}
	if ffcu.mutation.JSONStringValueCleared() {
		_spec.ClearField(formfieldconfig.FieldJSONStringValue, field.TypeString)
	}
	if value, ok := ffcu.mutation.Text(); ok {
		_spec.SetField(formfieldconfig.FieldText, field.TypeString, value)
	}
	if value, ok := ffcu.mutation.OrderIndex(); ok {
		_spec.SetField(formfieldconfig.FieldOrderIndex, field.TypeInt, value)
	}
	if value, ok := ffcu.mutation.AddedOrderIndex(); ok {
		_spec.AddField(formfieldconfig.FieldOrderIndex, field.TypeInt, value)
	}
	if value, ok := ffcu.mutation.Disabled(); ok {
		_spec.SetField(formfieldconfig.FieldDisabled, field.TypeInt, value)
	}
	if value, ok := ffcu.mutation.AddedDisabled(); ok {
		_spec.AddField(formfieldconfig.FieldDisabled, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ffcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{formfieldconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ffcu.mutation.done = true
	return n, nil
}

// FormFieldConfigUpdateOne is the builder for updating a single FormFieldConfig entity.
type FormFieldConfigUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FormFieldConfigMutation
}

// SetFieldID sets the "field_id" field.
func (ffcuo *FormFieldConfigUpdateOne) SetFieldID(s string) *FormFieldConfigUpdateOne {
	ffcuo.mutation.SetFieldID(s)
	return ffcuo
}

// SetFormID sets the "form_id" field.
func (ffcuo *FormFieldConfigUpdateOne) SetFormID(s string) *FormFieldConfigUpdateOne {
	ffcuo.mutation.SetFormID(s)
	return ffcuo
}

// SetKey sets the "key" field.
func (ffcuo *FormFieldConfigUpdateOne) SetKey(s string) *FormFieldConfigUpdateOne {
	ffcuo.mutation.SetKey(s)
	return ffcuo
}

// SetType sets the "type" field.
func (ffcuo *FormFieldConfigUpdateOne) SetType(s string) *FormFieldConfigUpdateOne {
	ffcuo.mutation.SetType(s)
	return ffcuo
}

// SetValue sets the "value" field.
func (ffcuo *FormFieldConfigUpdateOne) SetValue(s string) *FormFieldConfigUpdateOne {
	ffcuo.mutation.SetValue(s)
	return ffcuo
}

// SetJSONStringValue sets the "json_string_value" field.
func (ffcuo *FormFieldConfigUpdateOne) SetJSONStringValue(s string) *FormFieldConfigUpdateOne {
	ffcuo.mutation.SetJSONStringValue(s)
	return ffcuo
}

// SetNillableJSONStringValue sets the "json_string_value" field if the given value is not nil.
func (ffcuo *FormFieldConfigUpdateOne) SetNillableJSONStringValue(s *string) *FormFieldConfigUpdateOne {
	if s != nil {
		ffcuo.SetJSONStringValue(*s)
	}
	return ffcuo
}

// ClearJSONStringValue clears the value of the "json_string_value" field.
func (ffcuo *FormFieldConfigUpdateOne) ClearJSONStringValue() *FormFieldConfigUpdateOne {
	ffcuo.mutation.ClearJSONStringValue()
	return ffcuo
}

// SetText sets the "text" field.
func (ffcuo *FormFieldConfigUpdateOne) SetText(s string) *FormFieldConfigUpdateOne {
	ffcuo.mutation.SetText(s)
	return ffcuo
}

// SetOrderIndex sets the "order_index" field.
func (ffcuo *FormFieldConfigUpdateOne) SetOrderIndex(i int) *FormFieldConfigUpdateOne {
	ffcuo.mutation.ResetOrderIndex()
	ffcuo.mutation.SetOrderIndex(i)
	return ffcuo
}

// AddOrderIndex adds i to the "order_index" field.
func (ffcuo *FormFieldConfigUpdateOne) AddOrderIndex(i int) *FormFieldConfigUpdateOne {
	ffcuo.mutation.AddOrderIndex(i)
	return ffcuo
}

// SetDisabled sets the "disabled" field.
func (ffcuo *FormFieldConfigUpdateOne) SetDisabled(i int) *FormFieldConfigUpdateOne {
	ffcuo.mutation.ResetDisabled()
	ffcuo.mutation.SetDisabled(i)
	return ffcuo
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (ffcuo *FormFieldConfigUpdateOne) SetNillableDisabled(i *int) *FormFieldConfigUpdateOne {
	if i != nil {
		ffcuo.SetDisabled(*i)
	}
	return ffcuo
}

// AddDisabled adds i to the "disabled" field.
func (ffcuo *FormFieldConfigUpdateOne) AddDisabled(i int) *FormFieldConfigUpdateOne {
	ffcuo.mutation.AddDisabled(i)
	return ffcuo
}

// Mutation returns the FormFieldConfigMutation object of the builder.
func (ffcuo *FormFieldConfigUpdateOne) Mutation() *FormFieldConfigMutation {
	return ffcuo.mutation
}

// Where appends a list predicates to the FormFieldConfigUpdate builder.
func (ffcuo *FormFieldConfigUpdateOne) Where(ps ...predicate.FormFieldConfig) *FormFieldConfigUpdateOne {
	ffcuo.mutation.Where(ps...)
	return ffcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ffcuo *FormFieldConfigUpdateOne) Select(field string, fields ...string) *FormFieldConfigUpdateOne {
	ffcuo.fields = append([]string{field}, fields...)
	return ffcuo
}

// Save executes the query and returns the updated FormFieldConfig entity.
func (ffcuo *FormFieldConfigUpdateOne) Save(ctx context.Context) (*FormFieldConfig, error) {
	return withHooks[*FormFieldConfig, FormFieldConfigMutation](ctx, ffcuo.sqlSave, ffcuo.mutation, ffcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ffcuo *FormFieldConfigUpdateOne) SaveX(ctx context.Context) *FormFieldConfig {
	node, err := ffcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ffcuo *FormFieldConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := ffcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ffcuo *FormFieldConfigUpdateOne) ExecX(ctx context.Context) {
	if err := ffcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ffcuo *FormFieldConfigUpdateOne) sqlSave(ctx context.Context) (_node *FormFieldConfig, err error) {
	_spec := sqlgraph.NewUpdateSpec(formfieldconfig.Table, formfieldconfig.Columns, sqlgraph.NewFieldSpec(formfieldconfig.FieldID, field.TypeString))
	id, ok := ffcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FormFieldConfig.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ffcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, formfieldconfig.FieldID)
		for _, f := range fields {
			if !formfieldconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != formfieldconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ffcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ffcuo.mutation.FieldID(); ok {
		_spec.SetField(formfieldconfig.FieldFieldID, field.TypeString, value)
	}
	if value, ok := ffcuo.mutation.FormID(); ok {
		_spec.SetField(formfieldconfig.FieldFormID, field.TypeString, value)
	}
	if value, ok := ffcuo.mutation.Key(); ok {
		_spec.SetField(formfieldconfig.FieldKey, field.TypeString, value)
	}
	if value, ok := ffcuo.mutation.GetType(); ok {
		_spec.SetField(formfieldconfig.FieldType, field.TypeString, value)
	}
	if value, ok := ffcuo.mutation.Value(); ok {
		_spec.SetField(formfieldconfig.FieldValue, field.TypeString, value)
	}
	if value, ok := ffcuo.mutation.JSONStringValue(); ok {
		_spec.SetField(formfieldconfig.FieldJSONStringValue, field.TypeString, value)
	}
	if ffcuo.mutation.JSONStringValueCleared() {
		_spec.ClearField(formfieldconfig.FieldJSONStringValue, field.TypeString)
	}
	if value, ok := ffcuo.mutation.Text(); ok {
		_spec.SetField(formfieldconfig.FieldText, field.TypeString, value)
	}
	if value, ok := ffcuo.mutation.OrderIndex(); ok {
		_spec.SetField(formfieldconfig.FieldOrderIndex, field.TypeInt, value)
	}
	if value, ok := ffcuo.mutation.AddedOrderIndex(); ok {
		_spec.AddField(formfieldconfig.FieldOrderIndex, field.TypeInt, value)
	}
	if value, ok := ffcuo.mutation.Disabled(); ok {
		_spec.SetField(formfieldconfig.FieldDisabled, field.TypeInt, value)
	}
	if value, ok := ffcuo.mutation.AddedDisabled(); ok {
		_spec.AddField(formfieldconfig.FieldDisabled, field.TypeInt, value)
	}
	_node = &FormFieldConfig{config: ffcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ffcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{formfieldconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ffcuo.mutation.done = true
	return _node, nil
}
