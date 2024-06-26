// Code generated by ent, DO NOT EDIT.

package formfieldconfig

const (
	// Label holds the string label denoting the formfieldconfig type in the database.
	Label = "form_field_config"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFieldID holds the string denoting the field_id field in the database.
	FieldFieldID = "field_id"
	// FieldFormID holds the string denoting the form_id field in the database.
	FieldFormID = "form_id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldJSONStringValue holds the string denoting the json_string_value field in the database.
	FieldJSONStringValue = "json_string_value"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldOrderIndex holds the string denoting the order_index field in the database.
	FieldOrderIndex = "order_index"
	// FieldDisabled holds the string denoting the disabled field in the database.
	FieldDisabled = "disabled"
	// Table holds the table name of the formfieldconfig in the database.
	Table = "form_field_configs"
)

// Columns holds all SQL columns for formfieldconfig fields.
var Columns = []string{
	FieldID,
	FieldFieldID,
	FieldFormID,
	FieldKey,
	FieldType,
	FieldValue,
	FieldJSONStringValue,
	FieldText,
	FieldOrderIndex,
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
