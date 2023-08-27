package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.Text("body").Default(""),
		field.Uint("text_format").Default(0), // 0: Markdown
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("pages").Unique(),
	}
}
