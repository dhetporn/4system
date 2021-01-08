// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/dhetporn/team08/ent/diagnose"
	"github.com/dhetporn/team08/ent/doctor"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DoctorCreate is the builder for creating a Doctor entity.
type DoctorCreate struct {
	config
	mutation *DoctorMutation
	hooks    []Hook
}

// SetDoctorName sets the Doctor_Name field.
func (dc *DoctorCreate) SetDoctorName(s string) *DoctorCreate {
	dc.mutation.SetDoctorName(s)
	return dc
}

// SetDoctorPassword sets the Doctor_Password field.
func (dc *DoctorCreate) SetDoctorPassword(s string) *DoctorCreate {
	dc.mutation.SetDoctorPassword(s)
	return dc
}

// SetDoctorEmail sets the Doctor_Email field.
func (dc *DoctorCreate) SetDoctorEmail(s string) *DoctorCreate {
	dc.mutation.SetDoctorEmail(s)
	return dc
}

// SetDoctorTel sets the Doctor_tel field.
func (dc *DoctorCreate) SetDoctorTel(s string) *DoctorCreate {
	dc.mutation.SetDoctorTel(s)
	return dc
}

// AddDoctorDiagnoseIDs adds the doctor_diagnose edge to Diagnose by ids.
func (dc *DoctorCreate) AddDoctorDiagnoseIDs(ids ...int) *DoctorCreate {
	dc.mutation.AddDoctorDiagnoseIDs(ids...)
	return dc
}

// AddDoctorDiagnose adds the doctor_diagnose edges to Diagnose.
func (dc *DoctorCreate) AddDoctorDiagnose(d ...*Diagnose) *DoctorCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDoctorDiagnoseIDs(ids...)
}

// Mutation returns the DoctorMutation object of the builder.
func (dc *DoctorCreate) Mutation() *DoctorMutation {
	return dc.mutation
}

// Save creates the Doctor in the database.
func (dc *DoctorCreate) Save(ctx context.Context) (*Doctor, error) {
	if _, ok := dc.mutation.DoctorName(); !ok {
		return nil, &ValidationError{Name: "Doctor_Name", err: errors.New("ent: missing required field \"Doctor_Name\"")}
	}
	if v, ok := dc.mutation.DoctorName(); ok {
		if err := doctor.DoctorNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Doctor_Name", err: fmt.Errorf("ent: validator failed for field \"Doctor_Name\": %w", err)}
		}
	}
	if _, ok := dc.mutation.DoctorPassword(); !ok {
		return nil, &ValidationError{Name: "Doctor_Password", err: errors.New("ent: missing required field \"Doctor_Password\"")}
	}
	if v, ok := dc.mutation.DoctorPassword(); ok {
		if err := doctor.DoctorPasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "Doctor_Password", err: fmt.Errorf("ent: validator failed for field \"Doctor_Password\": %w", err)}
		}
	}
	if _, ok := dc.mutation.DoctorEmail(); !ok {
		return nil, &ValidationError{Name: "Doctor_Email", err: errors.New("ent: missing required field \"Doctor_Email\"")}
	}
	if v, ok := dc.mutation.DoctorEmail(); ok {
		if err := doctor.DoctorEmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "Doctor_Email", err: fmt.Errorf("ent: validator failed for field \"Doctor_Email\": %w", err)}
		}
	}
	if _, ok := dc.mutation.DoctorTel(); !ok {
		return nil, &ValidationError{Name: "Doctor_tel", err: errors.New("ent: missing required field \"Doctor_tel\"")}
	}
	if v, ok := dc.mutation.DoctorTel(); ok {
		if err := doctor.DoctorTelValidator(v); err != nil {
			return nil, &ValidationError{Name: "Doctor_tel", err: fmt.Errorf("ent: validator failed for field \"Doctor_tel\": %w", err)}
		}
	}
	var (
		err  error
		node *Doctor
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoctorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DoctorCreate) SaveX(ctx context.Context) *Doctor {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DoctorCreate) sqlSave(ctx context.Context) (*Doctor, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DoctorCreate) createSpec() (*Doctor, *sqlgraph.CreateSpec) {
	var (
		d     = &Doctor{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: doctor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctor.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.DoctorName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorName,
		})
		d.DoctorName = value
	}
	if value, ok := dc.mutation.DoctorPassword(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorPassword,
		})
		d.DoctorPassword = value
	}
	if value, ok := dc.mutation.DoctorEmail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorEmail,
		})
		d.DoctorEmail = value
	}
	if value, ok := dc.mutation.DoctorTel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldDoctorTel,
		})
		d.DoctorTel = value
	}
	if nodes := dc.mutation.DoctorDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorDiagnoseTable,
			Columns: []string{doctor.DoctorDiagnoseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return d, _spec
}