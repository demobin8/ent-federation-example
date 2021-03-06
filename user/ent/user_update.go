// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"user/ent/predicate"
	"user/ent/user"
	"user/ent/userlog"
	"user/ent/useroainfo"
	"user/ent/useropinfo"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where adds a new predicate for the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.predicates = append(uu.mutation.predicates, ps...)
	return uu
}

// SetAge sets the "age" field.
func (uu *UserUpdate) SetAge(i int) *UserUpdate {
	uu.mutation.ResetAge()
	uu.mutation.SetAge(i)
	return uu
}

// AddAge adds i to the "age" field.
func (uu *UserUpdate) AddAge(i int) *UserUpdate {
	uu.mutation.AddAge(i)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// AddUserLogIDs adds the "userLogs" edge to the UserLog entity by IDs.
func (uu *UserUpdate) AddUserLogIDs(ids ...int) *UserUpdate {
	uu.mutation.AddUserLogIDs(ids...)
	return uu
}

// AddUserLogs adds the "userLogs" edges to the UserLog entity.
func (uu *UserUpdate) AddUserLogs(u ...*UserLog) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddUserLogIDs(ids...)
}

// AddUserOaInfoIDs adds the "userOaInfo" edge to the UserOaInfo entity by IDs.
func (uu *UserUpdate) AddUserOaInfoIDs(ids ...int) *UserUpdate {
	uu.mutation.AddUserOaInfoIDs(ids...)
	return uu
}

// AddUserOaInfo adds the "userOaInfo" edges to the UserOaInfo entity.
func (uu *UserUpdate) AddUserOaInfo(u ...*UserOaInfo) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddUserOaInfoIDs(ids...)
}

// AddUserOpInfoIDs adds the "userOpInfo" edge to the UserOpInfo entity by IDs.
func (uu *UserUpdate) AddUserOpInfoIDs(ids ...int) *UserUpdate {
	uu.mutation.AddUserOpInfoIDs(ids...)
	return uu
}

// AddUserOpInfo adds the "userOpInfo" edges to the UserOpInfo entity.
func (uu *UserUpdate) AddUserOpInfo(u ...*UserOpInfo) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddUserOpInfoIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearUserLogs clears all "userLogs" edges to the UserLog entity.
func (uu *UserUpdate) ClearUserLogs() *UserUpdate {
	uu.mutation.ClearUserLogs()
	return uu
}

// RemoveUserLogIDs removes the "userLogs" edge to UserLog entities by IDs.
func (uu *UserUpdate) RemoveUserLogIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveUserLogIDs(ids...)
	return uu
}

// RemoveUserLogs removes "userLogs" edges to UserLog entities.
func (uu *UserUpdate) RemoveUserLogs(u ...*UserLog) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveUserLogIDs(ids...)
}

// ClearUserOaInfo clears all "userOaInfo" edges to the UserOaInfo entity.
func (uu *UserUpdate) ClearUserOaInfo() *UserUpdate {
	uu.mutation.ClearUserOaInfo()
	return uu
}

// RemoveUserOaInfoIDs removes the "userOaInfo" edge to UserOaInfo entities by IDs.
func (uu *UserUpdate) RemoveUserOaInfoIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveUserOaInfoIDs(ids...)
	return uu
}

// RemoveUserOaInfo removes "userOaInfo" edges to UserOaInfo entities.
func (uu *UserUpdate) RemoveUserOaInfo(u ...*UserOaInfo) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveUserOaInfoIDs(ids...)
}

// ClearUserOpInfo clears all "userOpInfo" edges to the UserOpInfo entity.
func (uu *UserUpdate) ClearUserOpInfo() *UserUpdate {
	uu.mutation.ClearUserOpInfo()
	return uu
}

// RemoveUserOpInfoIDs removes the "userOpInfo" edge to UserOpInfo entities by IDs.
func (uu *UserUpdate) RemoveUserOpInfoIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveUserOpInfoIDs(ids...)
	return uu
}

// RemoveUserOpInfo removes "userOpInfo" edges to UserOpInfo entities.
func (uu *UserUpdate) RemoveUserOpInfo(u ...*UserOpInfo) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveUserOpInfoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
	}
	if value, ok := uu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
	}
	if uu.mutation.UserLogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserLogsTable,
			Columns: []string{user.UserLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userlog.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedUserLogsIDs(); len(nodes) > 0 && !uu.mutation.UserLogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserLogsTable,
			Columns: []string{user.UserLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserLogsTable,
			Columns: []string{user.UserLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.UserOaInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOaInfoTable,
			Columns: []string{user.UserOaInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useroainfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedUserOaInfoIDs(); len(nodes) > 0 && !uu.mutation.UserOaInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOaInfoTable,
			Columns: []string{user.UserOaInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useroainfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserOaInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOaInfoTable,
			Columns: []string{user.UserOaInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useroainfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.UserOpInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOpInfoTable,
			Columns: []string{user.UserOpInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useropinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedUserOpInfoIDs(); len(nodes) > 0 && !uu.mutation.UserOpInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOpInfoTable,
			Columns: []string{user.UserOpInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useropinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserOpInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOpInfoTable,
			Columns: []string{user.UserOpInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useropinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetAge sets the "age" field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.mutation.ResetAge()
	uuo.mutation.SetAge(i)
	return uuo
}

// AddAge adds i to the "age" field.
func (uuo *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	uuo.mutation.AddAge(i)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// AddUserLogIDs adds the "userLogs" edge to the UserLog entity by IDs.
func (uuo *UserUpdateOne) AddUserLogIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddUserLogIDs(ids...)
	return uuo
}

// AddUserLogs adds the "userLogs" edges to the UserLog entity.
func (uuo *UserUpdateOne) AddUserLogs(u ...*UserLog) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddUserLogIDs(ids...)
}

// AddUserOaInfoIDs adds the "userOaInfo" edge to the UserOaInfo entity by IDs.
func (uuo *UserUpdateOne) AddUserOaInfoIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddUserOaInfoIDs(ids...)
	return uuo
}

// AddUserOaInfo adds the "userOaInfo" edges to the UserOaInfo entity.
func (uuo *UserUpdateOne) AddUserOaInfo(u ...*UserOaInfo) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddUserOaInfoIDs(ids...)
}

// AddUserOpInfoIDs adds the "userOpInfo" edge to the UserOpInfo entity by IDs.
func (uuo *UserUpdateOne) AddUserOpInfoIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddUserOpInfoIDs(ids...)
	return uuo
}

// AddUserOpInfo adds the "userOpInfo" edges to the UserOpInfo entity.
func (uuo *UserUpdateOne) AddUserOpInfo(u ...*UserOpInfo) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddUserOpInfoIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearUserLogs clears all "userLogs" edges to the UserLog entity.
func (uuo *UserUpdateOne) ClearUserLogs() *UserUpdateOne {
	uuo.mutation.ClearUserLogs()
	return uuo
}

// RemoveUserLogIDs removes the "userLogs" edge to UserLog entities by IDs.
func (uuo *UserUpdateOne) RemoveUserLogIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveUserLogIDs(ids...)
	return uuo
}

// RemoveUserLogs removes "userLogs" edges to UserLog entities.
func (uuo *UserUpdateOne) RemoveUserLogs(u ...*UserLog) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveUserLogIDs(ids...)
}

// ClearUserOaInfo clears all "userOaInfo" edges to the UserOaInfo entity.
func (uuo *UserUpdateOne) ClearUserOaInfo() *UserUpdateOne {
	uuo.mutation.ClearUserOaInfo()
	return uuo
}

// RemoveUserOaInfoIDs removes the "userOaInfo" edge to UserOaInfo entities by IDs.
func (uuo *UserUpdateOne) RemoveUserOaInfoIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveUserOaInfoIDs(ids...)
	return uuo
}

// RemoveUserOaInfo removes "userOaInfo" edges to UserOaInfo entities.
func (uuo *UserUpdateOne) RemoveUserOaInfo(u ...*UserOaInfo) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveUserOaInfoIDs(ids...)
}

// ClearUserOpInfo clears all "userOpInfo" edges to the UserOpInfo entity.
func (uuo *UserUpdateOne) ClearUserOpInfo() *UserUpdateOne {
	uuo.mutation.ClearUserOpInfo()
	return uuo
}

// RemoveUserOpInfoIDs removes the "userOpInfo" edge to UserOpInfo entities by IDs.
func (uuo *UserUpdateOne) RemoveUserOpInfoIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveUserOpInfoIDs(ids...)
	return uuo
}

// RemoveUserOpInfo removes "userOpInfo" edges to UserOpInfo entities.
func (uuo *UserUpdateOne) RemoveUserOpInfo(u ...*UserOpInfo) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveUserOpInfoIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
	}
	if value, ok := uuo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
	}
	if uuo.mutation.UserLogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserLogsTable,
			Columns: []string{user.UserLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userlog.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedUserLogsIDs(); len(nodes) > 0 && !uuo.mutation.UserLogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserLogsTable,
			Columns: []string{user.UserLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserLogsTable,
			Columns: []string{user.UserLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.UserOaInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOaInfoTable,
			Columns: []string{user.UserOaInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useroainfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedUserOaInfoIDs(); len(nodes) > 0 && !uuo.mutation.UserOaInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOaInfoTable,
			Columns: []string{user.UserOaInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useroainfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserOaInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOaInfoTable,
			Columns: []string{user.UserOaInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useroainfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.UserOpInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOpInfoTable,
			Columns: []string{user.UserOpInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useropinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedUserOpInfoIDs(); len(nodes) > 0 && !uuo.mutation.UserOpInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOpInfoTable,
			Columns: []string{user.UserOpInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useropinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserOpInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOpInfoTable,
			Columns: []string{user.UserOpInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: useropinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
