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

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable().
			Comment("Unique identifier for the user"),

		field.String("org_id").
			NotEmpty().
			Comment("Organization this user belongs to"),

		field.String("external_id").
			Optional().
			Comment("External system identifier"),

		field.String("email").
			Unique().
			NotEmpty().
			MaxLen(255).
			Comment("User's email address"),

		field.String("first_name").
			MaxLen(100).
			Optional().
			Comment("User's first name"),

		field.String("last_name").
			MaxLen(100).
			Optional().
			Comment("User's last name"),

		field.String("phone").
			MaxLen(50).
			Optional().
			Comment("User's contact phone number"),

		field.String("password_hash").
			Sensitive().
			MaxLen(255).
			Optional().
			Comment("Hashed password"),

		field.String("auth_provider").
			MaxLen(50).
			Optional().
			Comment("Authentication provider (e.g., 'local', 'google', 'github')"),

		field.String("auth_provider_id").
			MaxLen(255).
			Optional().
			Comment("User ID in the authentication provider"),

		field.Enum("status").
			Values("STATUS_UNSPECIFIED", "STATUS_ACTIVE", "STATUS_INACTIVE", "STATUS_SUSPENDED", "STATUS_DELETED").
			Default("STATUS_ACTIVE").
			Comment("Current status of the user"),

		field.Bool("email_verified").
			Default(false).
			Comment("Whether the email has been verified"),

		field.Bool("phone_verified").
			Default(false).
			Comment("Whether the phone has been verified"),

		field.Bool("mfa_enabled").
			Default(false).
			Comment("Whether multi-factor authentication is enabled"),

		field.String("mfa_type").
			MaxLen(50).
			Optional().
			Comment("Type of MFA used (e.g., 'totp', 'sms')"),

		field.String("mfa_secret").
			Sensitive().
			MaxLen(255).
			Optional().
			Comment("Secret for MFA"),

		field.JSON("role_ids", []string{}).
			Optional().
			Comment("IDs of roles assigned to user"),

		field.JSON("attributes", map[string]string{}).
			Optional().
			Comment("User attributes stored as key-value pairs"),

		field.JSON("preferences", map[string]string{}).
			Optional().
			Comment("User preferences stored as key-value pairs"),

		field.Time("last_login").
			Optional().
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the user last logged in"),

		field.Int("login_count").
			Default(0).
			Comment("Number of times user has logged in"),

		field.String("last_ip").
			MaxLen(45).
			Optional().
			Comment("Last IP address from which the user logged in"),

		field.Time("password_changed_at").
			Optional().
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the password was last changed"),

		field.Time("email_verified_at").
			Optional().
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the email was verified"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the user was created"),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone",
			}).
			Comment("When the user was last updated"),

		field.String("created_by").
			Optional().
			MaxLen(255).
			Comment("User ID who created this user"),

		field.String("updated_by").
			Optional().
			MaxLen(255).
			Comment("User ID who last updated this user"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("users").
			Field("org_id").
			Unique().
			Required().
			Comment("Organization this user belongs to"),

		edge.From("roles", Role.Type).
			Ref("users").
			Comment("Roles assigned to this user"),

		edge.To("created_organizations", Organization.Type).
			Comment("Organizations created by this user").
			Through("user_created_orgs", UserCreatedOrgs.Type),

		edge.To("updated_organizations", Organization.Type).
			Comment("Organizations updated by this user").
			Through("user_updated_orgs", UserUpdatedOrgs.Type),
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").
			Unique(),
		index.Fields("org_id"),
		index.Fields("status"),
		index.Fields("email_verified"),
		index.Fields("auth_provider", "auth_provider_id"),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "users",
		},
	}
}
