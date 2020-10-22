package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.String("patient_name").NotEmpty(),
		field.String("patient_gender").NotEmpty(),
		field.Int("patient_age"),
		field.Int("patient_phone").Positive(),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("queue", Queue.Type).StorageKey(edge.Column("patient_id")),
	}
}
