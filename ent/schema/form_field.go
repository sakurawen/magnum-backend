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
	OrderIndex int `json:"order_index"`
}

func (FormField) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(uuid.NewString).Unique(),
		field.String("form_id").Comment("字段对应的表单id"),
		field.String("field_type").Comment("字段类型"),
		field.String("field_name").Comment("字段名称"),
		field.Int("order_index").Comment("排序字段").StructTag(`json:"order_index"`).Optional(),
		field.Time("create_at"),
		field.Time("update_at").Optional(),
		field.Int("disabled").Comment("是否禁用").Default(0).StructTag(`json:"disabled"`),
	}
}

func (FormField) Edges() []ent.Edge {
	return nil
}
