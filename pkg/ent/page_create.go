// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tomato3713/nullwiki/pkg/ent/page"
	"github.com/tomato3713/nullwiki/pkg/ent/user"
)

// PageCreate is the builder for creating a Page entity.
type PageCreate struct {
	config
	mutation *PageMutation
	hooks    []Hook
}

// SetBody sets the "body" field.
func (pc *PageCreate) SetBody(s string) *PageCreate {
	pc.mutation.SetBody(s)
	return pc
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (pc *PageCreate) SetNillableBody(s *string) *PageCreate {
	if s != nil {
		pc.SetBody(*s)
	}
	return pc
}

// SetTextFormat sets the "text_format" field.
func (pc *PageCreate) SetTextFormat(u uint) *PageCreate {
	pc.mutation.SetTextFormat(u)
	return pc
}

// SetNillableTextFormat sets the "text_format" field if the given value is not nil.
func (pc *PageCreate) SetNillableTextFormat(u *uint) *PageCreate {
	if u != nil {
		pc.SetTextFormat(*u)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PageCreate) SetCreatedAt(t time.Time) *PageCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PageCreate) SetNillableCreatedAt(t *time.Time) *PageCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PageCreate) SetUpdatedAt(t time.Time) *PageCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PageCreate) SetNillableUpdatedAt(t *time.Time) *PageCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pc *PageCreate) SetOwnerID(id int) *PageCreate {
	pc.mutation.SetOwnerID(id)
	return pc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (pc *PageCreate) SetNillableOwnerID(id *int) *PageCreate {
	if id != nil {
		pc = pc.SetOwnerID(*id)
	}
	return pc
}

// SetOwner sets the "owner" edge to the User entity.
func (pc *PageCreate) SetOwner(u *User) *PageCreate {
	return pc.SetOwnerID(u.ID)
}

// Mutation returns the PageMutation object of the builder.
func (pc *PageCreate) Mutation() *PageMutation {
	return pc.mutation
}

// Save creates the Page in the database.
func (pc *PageCreate) Save(ctx context.Context) (*Page, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PageCreate) SaveX(ctx context.Context) *Page {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PageCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PageCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PageCreate) defaults() {
	if _, ok := pc.mutation.Body(); !ok {
		v := page.DefaultBody
		pc.mutation.SetBody(v)
	}
	if _, ok := pc.mutation.TextFormat(); !ok {
		v := page.DefaultTextFormat
		pc.mutation.SetTextFormat(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := page.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := page.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PageCreate) check() error {
	if _, ok := pc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "Page.body"`)}
	}
	if _, ok := pc.mutation.TextFormat(); !ok {
		return &ValidationError{Name: "text_format", err: errors.New(`ent: missing required field "Page.text_format"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Page.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Page.updated_at"`)}
	}
	return nil
}

func (pc *PageCreate) sqlSave(ctx context.Context) (*Page, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PageCreate) createSpec() (*Page, *sqlgraph.CreateSpec) {
	var (
		_node = &Page{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(page.Table, sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Body(); ok {
		_spec.SetField(page.FieldBody, field.TypeString, value)
		_node.Body = value
	}
	if value, ok := pc.mutation.TextFormat(); ok {
		_spec.SetField(page.FieldTextFormat, field.TypeUint, value)
		_node.TextFormat = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(page.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(page.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   page.OwnerTable,
			Columns: []string{page.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PageCreateBulk is the builder for creating many Page entities in bulk.
type PageCreateBulk struct {
	config
	builders []*PageCreate
}

// Save creates the Page entities in the database.
func (pcb *PageCreateBulk) Save(ctx context.Context) ([]*Page, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Page, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PageCreateBulk) SaveX(ctx context.Context) []*Page {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PageCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PageCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
