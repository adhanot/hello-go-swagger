package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Users schema.
type Users struct {
	ent.Schema
}

// Fields of the user.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.String("pseudo").
			NotEmpty().
			Match(regexp.MustCompile("[a-zA-Z_]+$")).
			Unique().
			Immutable(),
		field.String("password").
			NotEmpty(),
		field.String("fname").
			NotEmpty(),
		field.String("lname").
			NotEmpty(),
		field.String("email").
			Unique().
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
	}
}
