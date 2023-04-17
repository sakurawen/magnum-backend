package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FormSubmission
// 表单提交记录
type FormSubmission struct {
	ent.Schema
}

func (FormSubmission) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(uuid.New().String).Unique(),
		field.String("form_id").Comment("表单id"),
		field.String("user_id").Comment("用户id"),
		field.Time("create_id").Comment("创建时间"),
	}
}
