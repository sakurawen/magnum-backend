package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(uuid.New().String).Unique(),
		field.String("password"),
		field.String("username").Unique(),
		field.String("phone"),
		field.String("role").Default("user"),
		field.String("account").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
