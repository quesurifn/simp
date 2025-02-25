package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserCreatedOrgs represents an edge between users and organizations they created.
type UserCreatedOrgs struct {
	ent.Schema
}

// Fields of the UserCreatedOrgs.
func (UserCreatedOrgs) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			Comment("User who created the organization"),
		field.String("organization_id").
			Comment("Organization that was created"),
	}
}

// Edges of the UserCreatedOrgs.
func (UserCreatedOrgs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("user_id").
			Unique().
			Required(),
		edge.To("organization", Organization.Type).
			Field("organization_id").
			Unique().
			Required(),
	}
}
