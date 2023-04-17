// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FormsColumns holds the columns for the "forms" table.
	FormsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "user_id", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime, Nullable: true},
	}
	// FormsTable holds the schema information for the "forms" table.
	FormsTable = &schema.Table{
		Name:       "forms",
		Columns:    FormsColumns,
		PrimaryKey: []*schema.Column{FormsColumns[0]},
	}
	// FormFieldsColumns holds the columns for the "form_fields" table.
	FormFieldsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "form_id", Type: field.TypeString},
		{Name: "field_type", Type: field.TypeString},
		{Name: "field_label", Type: field.TypeString},
		{Name: "filed_name", Type: field.TypeString},
		{Name: "is_required", Type: field.TypeInt, Default: 0},
		{Name: "order_index", Type: field.TypeInt},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime, Nullable: true},
		{Name: "options", Type: field.TypeString, Nullable: true},
		{Name: "placeholder", Type: field.TypeString, Nullable: true},
	}
	// FormFieldsTable holds the schema information for the "form_fields" table.
	FormFieldsTable = &schema.Table{
		Name:       "form_fields",
		Columns:    FormFieldsColumns,
		PrimaryKey: []*schema.Column{FormFieldsColumns[0]},
	}
	// FormSubmissionsColumns holds the columns for the "form_submissions" table.
	FormSubmissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "form_id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
		{Name: "create_id", Type: field.TypeTime},
	}
	// FormSubmissionsTable holds the schema information for the "form_submissions" table.
	FormSubmissionsTable = &schema.Table{
		Name:       "form_submissions",
		Columns:    FormSubmissionsColumns,
		PrimaryKey: []*schema.Column{FormSubmissionsColumns[0]},
	}
	// FormSubmissionDataColumns holds the columns for the "form_submission_data" table.
	FormSubmissionDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "submission_id", Type: field.TypeString},
		{Name: "field_id", Type: field.TypeString},
		{Name: "field_value", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeTime},
	}
	// FormSubmissionDataTable holds the schema information for the "form_submission_data" table.
	FormSubmissionDataTable = &schema.Table{
		Name:       "form_submission_data",
		Columns:    FormSubmissionDataColumns,
		PrimaryKey: []*schema.Column{FormSubmissionDataColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "phone", Type: field.TypeString},
		{Name: "role", Type: field.TypeString, Default: "user"},
		{Name: "account", Type: field.TypeString, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FormsTable,
		FormFieldsTable,
		FormSubmissionsTable,
		FormSubmissionDataTable,
		UsersTable,
	}
)

func init() {
}
