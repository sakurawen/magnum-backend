package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Form
// 表单
type Form struct {
	ent.Schema
}

func (Form) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(uuid.New().String).Unique(),
		field.String("user_id").Comment("用户的id"),
		field.String("title").Comment("标题"),
		field.String("description").Comment("描述"),
		field.Time("create_at").Comment("创建日期"),
		field.Time("update_at").Comment("更新日期").Optional(),
	}
}

func (Form) Edges() []ent.Edge {
	return nil
}
