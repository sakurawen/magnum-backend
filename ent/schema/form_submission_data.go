package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type FormSubmissionData struct {
	ent.Schema
}

func (FormSubmissionData) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(uuid.New().String).Unique(),
		field.String("submission_id").Comment("提交记录id"),
		field.String("field_id").Comment("表单字段id"),
		field.String("field_value").Comment("表单字段value"),
		field.Time("create_at").Comment("提交时间"),
	}
}
