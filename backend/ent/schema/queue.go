package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Queue holds the schema definition for the Queue entity.
type Queue struct {
	ent.Schema
}

// Fields of the Queue.
func (Queue) Fields() []ent.Field {
	return []ent.Field{
		field.String("dental"),
		field.Time("queue_time"),
	}
}

// Edges of the Queue.
func (Queue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dentist", Dentist.Type).Ref("queue").Unique(),
		edge.From("employee", Employee.Type).Ref("queue").Unique(),
		edge.From("patient", Patient.Type).Ref("queue").Unique(),
	}
}
