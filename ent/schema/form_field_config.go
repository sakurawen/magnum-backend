package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FormFieldConfig 字段配置项
type FormFieldConfig struct {
	ent.Schema
}

func (FormFieldConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(uuid.NewString),
		field.String("field_id").Comment("配置对应字段的id"),
		field.String("form_id").Comment("配置对应的表单id"),
		field.String("key").Comment("配置key"),
		field.String("type").Comment("配置类型"),
		field.String("value").Comment("配置值"),
		field.String("json_string_value").Comment("当值的类型是数组或者对象时，使用序列化的json存储").Optional(),
		field.String("text").Comment("配置名称"),
		field.Int("order_index").StructTag(`json:"order_index"`),
		field.Int("disabled").Default(0).StructTag(`json:"disabled"`),
	}
}
