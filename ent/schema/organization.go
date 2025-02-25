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

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable().
			Comment("Unique identifier for the organization"),

		field.String("name").
			NotEmpty().
			MaxLen(255).
			Unique().
			Comment("Internal name of the organization, used for references"),

		field.String("display_name").
			NotEmpty().
			MaxLen(255).
			Comment("Human-readable name of the organization"),

		field.String("logo_url").
			Optional().
			MaxLen(2083).
			Comment("URL to the organization's logo"),

		field.String("primary_email").
			MaxLen(255).
			Optional().
			Comment("Primary email address for the organization"),

		field.String("billing_email").
			MaxLen(255).
			Optional().
			Comment("Email address for billing communications"),

		field.String("phone").
			MaxLen(50).
			Optional().
			Comment("Contact phone number"),

		field.String("stripe_customer_id").
			MaxLen(255).
			Optional().
			Comment("Stripe customer ID for billing"),

		field.String("subscription_plan").
			MaxLen(100).
			Optional().
			Comment("Current subscription plan"),

		field.Enum("status").
			Values("STATUS_UNSPECIFIED", "STATUS_ACTIVE", "STATUS_INACTIVE", "STATUS_SUSPENDED", "STATUS_DELETED").
			Default("STATUS_ACTIVE").
			Comment("Current status of the organization"),

		field.JSON("settings", map[string]string{}).
			Optional().
			Comment("Organization-specific settings stored as key-value pairs"),

		field.Bool("is_active").
			Default(true).
			Comment("Whether the organization is currently active"),

		field.Time("subscription_ends_at").
			Optional().
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the current subscription expires"),

		field.Time("trial_ends_at").
			Optional().
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the trial period expires"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the organization was created"),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the organization was last updated"),

		field.String("created_by").
			Optional().
			MaxLen(255).
			Comment("User ID who created the organization"),

		field.String("updated_by").
			Optional().
			MaxLen(255).
			Comment("User ID who last updated the organization"),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		// You can define relationships here, such as:
		edge.To("users", User.Type).
			Comment("Users belonging to this organization"),

		edge.To("routes", Route.Type).
			Comment("Routes associated with this organization"),

		edge.To("roles", Role.Type).
			Comment("Roles defined in this organization"),
	}
}

// Indexes of the Organization.
func (Organization) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique(),
		index.Fields("status"),
		index.Fields("is_active"),
	}
}

// Annotations of the Organization.
func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "organizations",
		},
	}
}
