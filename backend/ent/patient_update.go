// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/B5916177/app/ent/patient"
	"github.com/B5916177/app/ent/predicate"
	"github.com/B5916177/app/ent/queue"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PatientUpdate is the builder for updating Patient entities.
type PatientUpdate struct {
	config
	hooks      []Hook
	mutation   *PatientMutation
	predicates []predicate.Patient
}

// Where adds a new predicate for the builder.
func (pu *PatientUpdate) Where(ps ...predicate.Patient) *PatientUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetPatientName sets the patient_name field.
func (pu *PatientUpdate) SetPatientName(s string) *PatientUpdate {
	pu.mutation.SetPatientName(s)
	return pu
}

// SetPatientGender sets the patient_gender field.
func (pu *PatientUpdate) SetPatientGender(s string) *PatientUpdate {
	pu.mutation.SetPatientGender(s)
	return pu
}

// SetPatientAge sets the patient_age field.
func (pu *PatientUpdate) SetPatientAge(i int) *PatientUpdate {
	pu.mutation.ResetPatientAge()
	pu.mutation.SetPatientAge(i)
	return pu
}

// AddPatientAge adds i to patient_age.
func (pu *PatientUpdate) AddPatientAge(i int) *PatientUpdate {
	pu.mutation.AddPatientAge(i)
	return pu
}

// SetPatientPhone sets the patient_phone field.
func (pu *PatientUpdate) SetPatientPhone(i int) *PatientUpdate {
	pu.mutation.ResetPatientPhone()
	pu.mutation.SetPatientPhone(i)
	return pu
}

// AddPatientPhone adds i to patient_phone.
func (pu *PatientUpdate) AddPatientPhone(i int) *PatientUpdate {
	pu.mutation.AddPatientPhone(i)
	return pu
}

// AddQueueIDs adds the queue edge to Queue by ids.
func (pu *PatientUpdate) AddQueueIDs(ids ...int) *PatientUpdate {
	pu.mutation.AddQueueIDs(ids...)
	return pu
}

// AddQueue adds the queue edges to Queue.
func (pu *PatientUpdate) AddQueue(q ...*Queue) *PatientUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return pu.AddQueueIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (pu *PatientUpdate) Mutation() *PatientMutation {
	return pu.mutation
}

// RemoveQueueIDs removes the queue edge to Queue by ids.
func (pu *PatientUpdate) RemoveQueueIDs(ids ...int) *PatientUpdate {
	pu.mutation.RemoveQueueIDs(ids...)
	return pu
}

// RemoveQueue removes queue edges to Queue.
func (pu *PatientUpdate) RemoveQueue(q ...*Queue) *PatientUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return pu.RemoveQueueIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PatientUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.PatientName(); ok {
		if err := patient.PatientNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "patient_name", err: fmt.Errorf("ent: validator failed for field \"patient_name\": %w", err)}
		}
	}
	if v, ok := pu.mutation.PatientGender(); ok {
		if err := patient.PatientGenderValidator(v); err != nil {
			return 0, &ValidationError{Name: "patient_gender", err: fmt.Errorf("ent: validator failed for field \"patient_gender\": %w", err)}
		}
	}
	if v, ok := pu.mutation.PatientPhone(); ok {
		if err := patient.PatientPhoneValidator(v); err != nil {
			return 0, &ValidationError{Name: "patient_phone", err: fmt.Errorf("ent: validator failed for field \"patient_phone\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PatientUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PatientUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PatientUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PatientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.PatientName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientName,
		})
	}
	if value, ok := pu.mutation.PatientGender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientGender,
		})
	}
	if value, ok := pu.mutation.PatientAge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldPatientAge,
		})
	}
	if value, ok := pu.mutation.AddedPatientAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldPatientAge,
		})
	}
	if value, ok := pu.mutation.PatientPhone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldPatientPhone,
		})
	}
	if value, ok := pu.mutation.AddedPatientPhone(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldPatientPhone,
		})
	}
	if nodes := pu.mutation.RemovedQueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.QueueTable,
			Columns: []string{patient.QueueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: queue.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.QueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.QueueTable,
			Columns: []string{patient.QueueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: queue.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PatientUpdateOne is the builder for updating a single Patient entity.
type PatientUpdateOne struct {
	config
	hooks    []Hook
	mutation *PatientMutation
}

// SetPatientName sets the patient_name field.
func (puo *PatientUpdateOne) SetPatientName(s string) *PatientUpdateOne {
	puo.mutation.SetPatientName(s)
	return puo
}

// SetPatientGender sets the patient_gender field.
func (puo *PatientUpdateOne) SetPatientGender(s string) *PatientUpdateOne {
	puo.mutation.SetPatientGender(s)
	return puo
}

// SetPatientAge sets the patient_age field.
func (puo *PatientUpdateOne) SetPatientAge(i int) *PatientUpdateOne {
	puo.mutation.ResetPatientAge()
	puo.mutation.SetPatientAge(i)
	return puo
}

// AddPatientAge adds i to patient_age.
func (puo *PatientUpdateOne) AddPatientAge(i int) *PatientUpdateOne {
	puo.mutation.AddPatientAge(i)
	return puo
}

// SetPatientPhone sets the patient_phone field.
func (puo *PatientUpdateOne) SetPatientPhone(i int) *PatientUpdateOne {
	puo.mutation.ResetPatientPhone()
	puo.mutation.SetPatientPhone(i)
	return puo
}

// AddPatientPhone adds i to patient_phone.
func (puo *PatientUpdateOne) AddPatientPhone(i int) *PatientUpdateOne {
	puo.mutation.AddPatientPhone(i)
	return puo
}

// AddQueueIDs adds the queue edge to Queue by ids.
func (puo *PatientUpdateOne) AddQueueIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.AddQueueIDs(ids...)
	return puo
}

// AddQueue adds the queue edges to Queue.
func (puo *PatientUpdateOne) AddQueue(q ...*Queue) *PatientUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return puo.AddQueueIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (puo *PatientUpdateOne) Mutation() *PatientMutation {
	return puo.mutation
}

// RemoveQueueIDs removes the queue edge to Queue by ids.
func (puo *PatientUpdateOne) RemoveQueueIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.RemoveQueueIDs(ids...)
	return puo
}

// RemoveQueue removes queue edges to Queue.
func (puo *PatientUpdateOne) RemoveQueue(q ...*Queue) *PatientUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return puo.RemoveQueueIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PatientUpdateOne) Save(ctx context.Context) (*Patient, error) {
	if v, ok := puo.mutation.PatientName(); ok {
		if err := patient.PatientNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "patient_name", err: fmt.Errorf("ent: validator failed for field \"patient_name\": %w", err)}
		}
	}
	if v, ok := puo.mutation.PatientGender(); ok {
		if err := patient.PatientGenderValidator(v); err != nil {
			return nil, &ValidationError{Name: "patient_gender", err: fmt.Errorf("ent: validator failed for field \"patient_gender\": %w", err)}
		}
	}
	if v, ok := puo.mutation.PatientPhone(); ok {
		if err := patient.PatientPhoneValidator(v); err != nil {
			return nil, &ValidationError{Name: "patient_phone", err: fmt.Errorf("ent: validator failed for field \"patient_phone\": %w", err)}
		}
	}

	var (
		err  error
		node *Patient
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PatientUpdateOne) SaveX(ctx context.Context) *Patient {
	pa, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// Exec executes the query on the entity.
func (puo *PatientUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PatientUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PatientUpdateOne) sqlSave(ctx context.Context) (pa *Patient, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Patient.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.PatientName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientName,
		})
	}
	if value, ok := puo.mutation.PatientGender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientGender,
		})
	}
	if value, ok := puo.mutation.PatientAge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldPatientAge,
		})
	}
	if value, ok := puo.mutation.AddedPatientAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldPatientAge,
		})
	}
	if value, ok := puo.mutation.PatientPhone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldPatientPhone,
		})
	}
	if value, ok := puo.mutation.AddedPatientPhone(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldPatientPhone,
		})
	}
	if nodes := puo.mutation.RemovedQueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.QueueTable,
			Columns: []string{patient.QueueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: queue.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.QueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.QueueTable,
			Columns: []string{patient.QueueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: queue.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pa = &Patient{config: puo.config}
	_spec.Assign = pa.assignValues
	_spec.ScanValues = pa.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pa, nil
}
