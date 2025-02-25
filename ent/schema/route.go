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

// Route holds the schema definition for the Route entity.
type Route struct {
	ent.Schema
}

// Fields of the Route.
func (Route) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable().
			Comment("Unique identifier for the route"),

		field.String("org_id").
			NotEmpty().
			Comment("Organization this route belongs to"),

		field.String("path").
			NotEmpty().
			MaxLen(1024).
			Comment("Full API path"),

		field.String("base_path").
			MaxLen(255).
			Comment("Base part of the API path"),

		field.JSON("methods", []string{}).
			Comment("HTTP methods allowed for this route (GET, POST, etc.)"),

		field.String("description").
			Optional().
			MaxLen(2048).
			Comment("Description of the route"),

		field.JSON("metadata", map[string]string{}).
			Optional().
			Comment("Additional metadata as key-value pairs"),

		field.JSON("params", map[string]string{}).
			Optional().
			Comment("Route parameters as key-value pairs"),

		field.Bool("is_protected").
			Default(true).
			Comment("Whether the route requires authentication"),

		field.Bool("is_public").
			Default(false).
			Comment("Whether the route is available to the public"),

		field.Bool("is_deprecated").
			Default(false).
			Comment("Whether the route is deprecated"),

		field.String("version").
			MaxLen(50).
			Default("1.0").
			Comment("API version for this route"),

		field.JSON("required_role_ids", []string{}).
			Optional().
			Comment("IDs of roles required to access this route"),

		field.Time("discovered_at").
			Optional().
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the route was first discovered"),

		field.Time("last_accessed").
			Optional().
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the route was last accessed"),

		field.Int("access_count").
			Default(0).
			Comment("Number of times the route has been accessed"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the route was created"),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the route was last updated"),

		field.String("created_by").
			Optional().
			MaxLen(255).
			Comment("User ID who created this route"),

		field.String("updated_by").
			Optional().
			MaxLen(255).
			Comment("User ID who last updated this route"),
	}
}

// Edges of the Route.
func (Route) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("routes").
			Field("org_id").
			Unique().
			Required().
			Comment("Organization this route belongs to"),

		edge.From("required_roles", Role.Type).
			Ref("required_routes").
			Comment("Roles required to access this route"),
	}
}

// Indexes of the Route.
func (Route) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("org_id", "path").
			Unique(),
		index.Fields("is_protected"),
		index.Fields("is_public"),
		index.Fields("is_deprecated"),
		index.Fields("version"),
	}
}

// Annotations of the Route.
func (Route) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "routes",
		},
	}
}
