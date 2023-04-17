package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FormField
// 表单字段
type FormField struct {
	ent.Schema
}

func (FormField) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(uuid.New().String).Unique(),
		field.String("form_id").Comment("字段对应的表单id"),
		field.String("field_type").Comment("字段类型"),
		field.String("field_label").Comment("字段标签"),
		field.String("filed_name").Comment("字段名称"),
		field.Int("is_required").Comment("是否必填 1是0否").Default(0),
		field.Int("order_index").Comment("排序字段"),
		field.Time("create_at"),
		field.Time("update_at").Optional(),
		field.String("options").Optional().Comment("存放选项"),
		field.String("placeholder").Optional(),
	}
}

func (FormField) Edges() []ent.Edge {
	return nil
}
