// Code generated by ent, DO NOT EDIT.

package formfield

const (
	// Label holds the string label denoting the formfield type in the database.
	Label = "form_field"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFormID holds the string denoting the form_id field in the database.
	FieldFormID = "form_id"
	// FieldFieldType holds the string denoting the field_type field in the database.
	FieldFieldType = "field_type"
	// FieldFieldName holds the string denoting the field_name field in the database.
	FieldFieldName = "field_name"
	// FieldOrderIndex holds the string denoting the order_index field in the database.
	FieldOrderIndex = "order_index"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDisabled holds the string denoting the disabled field in the database.
	FieldDisabled = "disabled"
	// Table holds the table name of the formfield in the database.
	Table = "form_fields"
)

// Columns holds all SQL columns for formfield fields.
var Columns = []string{
	FieldID,
	FieldFormID,
	FieldFieldType,
	FieldFieldName,
	FieldOrderIndex,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDisabled,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDisabled holds the default value on creation for the "disabled" field.
	DefaultDisabled int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)
