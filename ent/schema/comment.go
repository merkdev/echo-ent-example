package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(20).
			Default("名無し"),
		field.String("text").
			NotEmpty().
			MaxLen(200).
			Default(""),
		field.Time("created").
			Default(func() time.Time {
				return time.Now()
			}),
		field.Time("updated").
			Default(func() time.Time {
				return time.Now()
			}).
			UpdateDefault(func() time.Time {
				return time.Now()
			}),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}