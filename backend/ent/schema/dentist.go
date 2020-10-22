package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Dentist holds the schema definition for the Dentist entity.
type Dentist struct {
	ent.Schema
}

// Fields of the Dentist.
func (Dentist) Fields() []ent.Field {
	return []ent.Field{
		field.String("dentist_name").NotEmpty(),
		field.String("dentist_email").NotEmpty(),
		field.Int("dentist_phone").Positive(),
	}
}

// Edges of the Dentist.
func (Dentist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("queue", Queue.Type).StorageKey(edge.Column("dentist_id")),
	}
}
