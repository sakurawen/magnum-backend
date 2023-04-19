package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
	Disabled int `json:"disabled"`
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(uuid.NewString).Unique(),
		field.String("password"),
		field.String("username").Unique(),
		field.String("phone"),
		field.String("role").Default("user"),
		field.String("account").Unique(),
		field.Int("disabled").Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
