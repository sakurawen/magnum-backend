// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"magnum/ent/formsubmissiondata"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// FormSubmissionData is the model entity for the FormSubmissionData schema.
type FormSubmissionData struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// 提交记录id
	SubmissionID string `json:"submission_id,omitempty"`
	// 表单字段id
	FieldID string `json:"field_id,omitempty"`
	// 表单字段value
	FieldValue string `json:"field_value,omitempty"`
	// 提交时间
	CreateAt time.Time `json:"create_at,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted int `json:"is_deleted"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FormSubmissionData) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case formsubmissiondata.FieldIsDeleted:
			values[i] = new(sql.NullInt64)
		case formsubmissiondata.FieldID, formsubmissiondata.FieldSubmissionID, formsubmissiondata.FieldFieldID, formsubmissiondata.FieldFieldValue:
			values[i] = new(sql.NullString)
		case formsubmissiondata.FieldCreateAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type FormSubmissionData", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FormSubmissionData fields.
func (fsd *FormSubmissionData) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case formsubmissiondata.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				fsd.ID = value.String
			}
		case formsubmissiondata.FieldSubmissionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field submission_id", values[i])
			} else if value.Valid {
				fsd.SubmissionID = value.String
			}
		case formsubmissiondata.FieldFieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field field_id", values[i])
			} else if value.Valid {
				fsd.FieldID = value.String
			}
		case formsubmissiondata.FieldFieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field field_value", values[i])
			} else if value.Valid {
				fsd.FieldValue = value.String
			}
		case formsubmissiondata.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				fsd.CreateAt = value.Time
			}
		case formsubmissiondata.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				fsd.IsDeleted = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this FormSubmissionData.
// Note that you need to call FormSubmissionData.Unwrap() before calling this method if this FormSubmissionData
// was returned from a transaction, and the transaction was committed or rolled back.
func (fsd *FormSubmissionData) Update() *FormSubmissionDataUpdateOne {
	return NewFormSubmissionDataClient(fsd.config).UpdateOne(fsd)
}

// Unwrap unwraps the FormSubmissionData entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fsd *FormSubmissionData) Unwrap() *FormSubmissionData {
	_tx, ok := fsd.config.driver.(*txDriver)
	if !ok {
		panic("ent: FormSubmissionData is not a transactional entity")
	}
	fsd.config.driver = _tx.drv
	return fsd
}

// String implements the fmt.Stringer.
func (fsd *FormSubmissionData) String() string {
	var builder strings.Builder
	builder.WriteString("FormSubmissionData(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fsd.ID))
	builder.WriteString("submission_id=")
	builder.WriteString(fsd.SubmissionID)
	builder.WriteString(", ")
	builder.WriteString("field_id=")
	builder.WriteString(fsd.FieldID)
	builder.WriteString(", ")
	builder.WriteString("field_value=")
	builder.WriteString(fsd.FieldValue)
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(fsd.CreateAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", fsd.IsDeleted))
	builder.WriteByte(')')
	return builder.String()
}

// FormSubmissionDataSlice is a parsable slice of FormSubmissionData.
type FormSubmissionDataSlice []*FormSubmissionData
