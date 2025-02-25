package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable().
			Comment("Unique identifier for the role"),

		field.String("org_id").
			NotEmpty().
			Comment("Organization this role belongs to"),

		field.String("name").
			NotEmpty().
			MaxLen(255).
			Comment("Name of the role"),

		field.String("description").
			Optional().
			MaxLen(2048).
			Comment("Description of the role"),

		field.JSON("permissions", []string{}).
			Optional().
			Comment("Permission strings for this role"),

		field.Bool("is_system").
			Default(false).
			Comment("Whether this is a system-defined role"),

		field.Bool("is_default").
			Default(false).
			Comment("Whether this role is assigned by default to new users"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the role was created"),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the role was last updated"),

		field.String("created_by").
			Optional().
			MaxLen(255).
			Comment("User ID who created this role"),

		field.String("updated_by").
			Optional().
			MaxLen(255).
			Comment("User ID who last updated this role"),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("roles").
			Field("org_id").
			Unique().
			Required().
			Comment("Organization this role belongs to"),

		edge.To("users", User.Type).
			Comment("Users assigned to this role"),

		edge.To("required_routes", Route.Type).
			Comment("Routes that require this role"),
	}
}

// Indexes of the Role.
func (Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("org_id", "name").
			Unique(),
		index.Fields("is_system"),
		index.Fields("is_default"),
	}
}

// Annotations of the Role.
func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "roles",
		},
	}
}
