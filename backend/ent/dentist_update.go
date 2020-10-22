// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/B5916177/app/ent/dentist"
	"github.com/B5916177/app/ent/predicate"
	"github.com/B5916177/app/ent/queue"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DentistUpdate is the builder for updating Dentist entities.
type DentistUpdate struct {
	config
	hooks      []Hook
	mutation   *DentistMutation
	predicates []predicate.Dentist
}

// Where adds a new predicate for the builder.
func (du *DentistUpdate) Where(ps ...predicate.Dentist) *DentistUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetDentistName sets the dentist_name field.
func (du *DentistUpdate) SetDentistName(s string) *DentistUpdate {
	du.mutation.SetDentistName(s)
	return du
}

// SetDentistEmail sets the dentist_email field.
func (du *DentistUpdate) SetDentistEmail(s string) *DentistUpdate {
	du.mutation.SetDentistEmail(s)
	return du
}

// SetDentistPhone sets the dentist_phone field.
func (du *DentistUpdate) SetDentistPhone(i int) *DentistUpdate {
	du.mutation.ResetDentistPhone()
	du.mutation.SetDentistPhone(i)
	return du
}

// AddDentistPhone adds i to dentist_phone.
func (du *DentistUpdate) AddDentistPhone(i int) *DentistUpdate {
	du.mutation.AddDentistPhone(i)
	return du
}

// AddQueueIDs adds the queue edge to Queue by ids.
func (du *DentistUpdate) AddQueueIDs(ids ...int) *DentistUpdate {
	du.mutation.AddQueueIDs(ids...)
	return du
}

// AddQueue adds the queue edges to Queue.
func (du *DentistUpdate) AddQueue(q ...*Queue) *DentistUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return du.AddQueueIDs(ids...)
}

// Mutation returns the DentistMutation object of the builder.
func (du *DentistUpdate) Mutation() *DentistMutation {
	return du.mutation
}

// RemoveQueueIDs removes the queue edge to Queue by ids.
func (du *DentistUpdate) RemoveQueueIDs(ids ...int) *DentistUpdate {
	du.mutation.RemoveQueueIDs(ids...)
	return du
}

// RemoveQueue removes queue edges to Queue.
func (du *DentistUpdate) RemoveQueue(q ...*Queue) *DentistUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return du.RemoveQueueIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DentistUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.DentistName(); ok {
		if err := dentist.DentistNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "dentist_name", err: fmt.Errorf("ent: validator failed for field \"dentist_name\": %w", err)}
		}
	}
	if v, ok := du.mutation.DentistEmail(); ok {
		if err := dentist.DentistEmailValidator(v); err != nil {
			return 0, &ValidationError{Name: "dentist_email", err: fmt.Errorf("ent: validator failed for field \"dentist_email\": %w", err)}
		}
	}
	if v, ok := du.mutation.DentistPhone(); ok {
		if err := dentist.DentistPhoneValidator(v); err != nil {
			return 0, &ValidationError{Name: "dentist_phone", err: fmt.Errorf("ent: validator failed for field \"dentist_phone\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DentistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DentistUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DentistUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DentistUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DentistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dentist.Table,
			Columns: dentist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dentist.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.DentistName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentist.FieldDentistName,
		})
	}
	if value, ok := du.mutation.DentistEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentist.FieldDentistEmail,
		})
	}
	if value, ok := du.mutation.DentistPhone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dentist.FieldDentistPhone,
		})
	}
	if value, ok := du.mutation.AddedDentistPhone(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dentist.FieldDentistPhone,
		})
	}
	if nodes := du.mutation.RemovedQueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentist.QueueTable,
			Columns: []string{dentist.QueueColumn},
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
	if nodes := du.mutation.QueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentist.QueueTable,
			Columns: []string{dentist.QueueColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dentist.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DentistUpdateOne is the builder for updating a single Dentist entity.
type DentistUpdateOne struct {
	config
	hooks    []Hook
	mutation *DentistMutation
}

// SetDentistName sets the dentist_name field.
func (duo *DentistUpdateOne) SetDentistName(s string) *DentistUpdateOne {
	duo.mutation.SetDentistName(s)
	return duo
}

// SetDentistEmail sets the dentist_email field.
func (duo *DentistUpdateOne) SetDentistEmail(s string) *DentistUpdateOne {
	duo.mutation.SetDentistEmail(s)
	return duo
}

// SetDentistPhone sets the dentist_phone field.
func (duo *DentistUpdateOne) SetDentistPhone(i int) *DentistUpdateOne {
	duo.mutation.ResetDentistPhone()
	duo.mutation.SetDentistPhone(i)
	return duo
}

// AddDentistPhone adds i to dentist_phone.
func (duo *DentistUpdateOne) AddDentistPhone(i int) *DentistUpdateOne {
	duo.mutation.AddDentistPhone(i)
	return duo
}

// AddQueueIDs adds the queue edge to Queue by ids.
func (duo *DentistUpdateOne) AddQueueIDs(ids ...int) *DentistUpdateOne {
	duo.mutation.AddQueueIDs(ids...)
	return duo
}

// AddQueue adds the queue edges to Queue.
func (duo *DentistUpdateOne) AddQueue(q ...*Queue) *DentistUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return duo.AddQueueIDs(ids...)
}

// Mutation returns the DentistMutation object of the builder.
func (duo *DentistUpdateOne) Mutation() *DentistMutation {
	return duo.mutation
}

// RemoveQueueIDs removes the queue edge to Queue by ids.
func (duo *DentistUpdateOne) RemoveQueueIDs(ids ...int) *DentistUpdateOne {
	duo.mutation.RemoveQueueIDs(ids...)
	return duo
}

// RemoveQueue removes queue edges to Queue.
func (duo *DentistUpdateOne) RemoveQueue(q ...*Queue) *DentistUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return duo.RemoveQueueIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DentistUpdateOne) Save(ctx context.Context) (*Dentist, error) {
	if v, ok := duo.mutation.DentistName(); ok {
		if err := dentist.DentistNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "dentist_name", err: fmt.Errorf("ent: validator failed for field \"dentist_name\": %w", err)}
		}
	}
	if v, ok := duo.mutation.DentistEmail(); ok {
		if err := dentist.DentistEmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "dentist_email", err: fmt.Errorf("ent: validator failed for field \"dentist_email\": %w", err)}
		}
	}
	if v, ok := duo.mutation.DentistPhone(); ok {
		if err := dentist.DentistPhoneValidator(v); err != nil {
			return nil, &ValidationError{Name: "dentist_phone", err: fmt.Errorf("ent: validator failed for field \"dentist_phone\": %w", err)}
		}
	}

	var (
		err  error
		node *Dentist
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DentistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DentistUpdateOne) SaveX(ctx context.Context) *Dentist {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DentistUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DentistUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DentistUpdateOne) sqlSave(ctx context.Context) (d *Dentist, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dentist.Table,
			Columns: dentist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dentist.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Dentist.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.DentistName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentist.FieldDentistName,
		})
	}
	if value, ok := duo.mutation.DentistEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentist.FieldDentistEmail,
		})
	}
	if value, ok := duo.mutation.DentistPhone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dentist.FieldDentistPhone,
		})
	}
	if value, ok := duo.mutation.AddedDentistPhone(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dentist.FieldDentistPhone,
		})
	}
	if nodes := duo.mutation.RemovedQueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentist.QueueTable,
			Columns: []string{dentist.QueueColumn},
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
	if nodes := duo.mutation.QueueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentist.QueueTable,
			Columns: []string{dentist.QueueColumn},
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
	d = &Dentist{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dentist.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
