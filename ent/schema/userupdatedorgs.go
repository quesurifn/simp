package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserUpdatedOrgs represents an edge between users and organizations they updated.
type UserUpdatedOrgs struct {
	ent.Schema
}

// Fields of the UserUpdatedOrgs.
func (UserUpdatedOrgs) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			Comment("User who updated the organization"),
		field.String("organization_id").
			Comment("Organization that was updated"),
	}
}

// Edges of the UserUpdatedOrgs.
func (UserUpdatedOrgs) Edges() []ent.Edge {
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
