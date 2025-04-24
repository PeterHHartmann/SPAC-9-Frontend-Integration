package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// AuditMixin implements the ent.Mixin for sharing
// audit-log capabilities with package schemas.
type AuditMixin struct {
	mixin.Schema
}

// Fields defines the mixin fields.
func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
