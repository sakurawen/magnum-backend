// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"magnum/ent/form"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Form is the model entity for the Form schema.
type Form struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// 用户的id
	UserID string `json:"user_id,omitempty"`
	// 标题
	Title string `json:"title,omitempty"`
	// 描述
	Description string `json:"description,omitempty"`
	// 创建日期
	CreateAt time.Time `json:"create_at,omitempty"`
	// 更新日期
	UpdateAt time.Time `json:"update_at,omitempty"`
	// 是否发布，1是0否
	IsRelease int `json:"is_release"`
	// Disabled holds the value of the "disabled" field.
	Disabled int `json:"disabled"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Form) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case form.FieldIsRelease, form.FieldDisabled:
			values[i] = new(sql.NullInt64)
		case form.FieldID, form.FieldUserID, form.FieldTitle, form.FieldDescription:
			values[i] = new(sql.NullString)
		case form.FieldCreateAt, form.FieldUpdateAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Form", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Form fields.
func (f *Form) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case form.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				f.ID = value.String
			}
		case form.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				f.UserID = value.String
			}
		case form.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				f.Title = value.String
			}
		case form.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				f.Description = value.String
			}
		case form.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				f.CreateAt = value.Time
			}
		case form.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				f.UpdateAt = value.Time
			}
		case form.FieldIsRelease:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_release", values[i])
			} else if value.Valid {
				f.IsRelease = int(value.Int64)
			}
		case form.FieldDisabled:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				f.Disabled = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Form.
// Note that you need to call Form.Unwrap() before calling this method if this Form
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Form) Update() *FormUpdateOne {
	return NewFormClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Form entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Form) Unwrap() *Form {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Form is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Form) String() string {
	var builder strings.Builder
	builder.WriteString("Form(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("user_id=")
	builder.WriteString(f.UserID)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(f.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(f.Description)
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(f.CreateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(f.UpdateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_release=")
	builder.WriteString(fmt.Sprintf("%v", f.IsRelease))
	builder.WriteString(", ")
	builder.WriteString("disabled=")
	builder.WriteString(fmt.Sprintf("%v", f.Disabled))
	builder.WriteByte(')')
	return builder.String()
}

// Forms is a parsable slice of Form.
type Forms []*Form
