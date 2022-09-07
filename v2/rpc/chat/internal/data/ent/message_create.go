// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat/internal/data/ent/message"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
}

// SetCreated sets the "created" field.
func (mc *MessageCreate) SetCreated(t time.Time) *MessageCreate {
	mc.mutation.SetCreated(t)
	return mc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (mc *MessageCreate) SetNillableCreated(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetCreated(*t)
	}
	return mc
}

// SetUpdated sets the "updated" field.
func (mc *MessageCreate) SetUpdated(t time.Time) *MessageCreate {
	mc.mutation.SetUpdated(t)
	return mc
}

// SetNillableUpdated sets the "updated" field if the given value is not nil.
func (mc *MessageCreate) SetNillableUpdated(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetUpdated(*t)
	}
	return mc
}

// SetSessionKey sets the "sessionKey" field.
func (mc *MessageCreate) SetSessionKey(s string) *MessageCreate {
	mc.mutation.SetSessionKey(s)
	return mc
}

// SetFrom sets the "from" field.
func (mc *MessageCreate) SetFrom(i int64) *MessageCreate {
	mc.mutation.SetFrom(i)
	return mc
}

// SetTo sets the "to" field.
func (mc *MessageCreate) SetTo(s string) *MessageCreate {
	mc.mutation.SetTo(s)
	return mc
}

// SetSessionType sets the "session_type" field.
func (mc *MessageCreate) SetSessionType(i int8) *MessageCreate {
	mc.mutation.SetSessionType(i)
	return mc
}

// SetClientMsgID sets the "client_msg_id" field.
func (mc *MessageCreate) SetClientMsgID(s string) *MessageCreate {
	mc.mutation.SetClientMsgID(s)
	return mc
}

// SetServerMsgSeq sets the "server_msg_seq" field.
func (mc *MessageCreate) SetServerMsgSeq(i int64) *MessageCreate {
	mc.mutation.SetServerMsgSeq(i)
	return mc
}

// SetMsgType sets the "msg_type" field.
func (mc *MessageCreate) SetMsgType(i int8) *MessageCreate {
	mc.mutation.SetMsgType(i)
	return mc
}

// SetMsgData sets the "msg_data" field.
func (mc *MessageCreate) SetMsgData(s string) *MessageCreate {
	mc.mutation.SetMsgData(s)
	return mc
}

// SetMsgResCode sets the "msg_res_code" field.
func (mc *MessageCreate) SetMsgResCode(i int8) *MessageCreate {
	mc.mutation.SetMsgResCode(i)
	return mc
}

// SetMsgFeature sets the "msg_feature" field.
func (mc *MessageCreate) SetMsgFeature(i int8) *MessageCreate {
	mc.mutation.SetMsgFeature(i)
	return mc
}

// SetMsgStatus sets the "msg_status" field.
func (mc *MessageCreate) SetMsgStatus(i int8) *MessageCreate {
	mc.mutation.SetMsgStatus(i)
	return mc
}

// SetID sets the "id" field.
func (mc *MessageCreate) SetID(i int64) *MessageCreate {
	mc.mutation.SetID(i)
	return mc
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Message)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MessageMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessageCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessageCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MessageCreate) defaults() {
	if _, ok := mc.mutation.Created(); !ok {
		v := message.DefaultCreated()
		mc.mutation.SetCreated(v)
	}
	if _, ok := mc.mutation.Updated(); !ok {
		v := message.DefaultUpdated()
		mc.mutation.SetUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessageCreate) check() error {
	if _, ok := mc.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New(`ent: missing required field "Message.created"`)}
	}
	if _, ok := mc.mutation.Updated(); !ok {
		return &ValidationError{Name: "updated", err: errors.New(`ent: missing required field "Message.updated"`)}
	}
	if _, ok := mc.mutation.SessionKey(); !ok {
		return &ValidationError{Name: "sessionKey", err: errors.New(`ent: missing required field "Message.sessionKey"`)}
	}
	if v, ok := mc.mutation.SessionKey(); ok {
		if err := message.SessionKeyValidator(v); err != nil {
			return &ValidationError{Name: "sessionKey", err: fmt.Errorf(`ent: validator failed for field "Message.sessionKey": %w`, err)}
		}
	}
	if _, ok := mc.mutation.From(); !ok {
		return &ValidationError{Name: "from", err: errors.New(`ent: missing required field "Message.from"`)}
	}
	if _, ok := mc.mutation.To(); !ok {
		return &ValidationError{Name: "to", err: errors.New(`ent: missing required field "Message.to"`)}
	}
	if v, ok := mc.mutation.To(); ok {
		if err := message.ToValidator(v); err != nil {
			return &ValidationError{Name: "to", err: fmt.Errorf(`ent: validator failed for field "Message.to": %w`, err)}
		}
	}
	if _, ok := mc.mutation.SessionType(); !ok {
		return &ValidationError{Name: "session_type", err: errors.New(`ent: missing required field "Message.session_type"`)}
	}
	if _, ok := mc.mutation.ClientMsgID(); !ok {
		return &ValidationError{Name: "client_msg_id", err: errors.New(`ent: missing required field "Message.client_msg_id"`)}
	}
	if v, ok := mc.mutation.ClientMsgID(); ok {
		if err := message.ClientMsgIDValidator(v); err != nil {
			return &ValidationError{Name: "client_msg_id", err: fmt.Errorf(`ent: validator failed for field "Message.client_msg_id": %w`, err)}
		}
	}
	if _, ok := mc.mutation.ServerMsgSeq(); !ok {
		return &ValidationError{Name: "server_msg_seq", err: errors.New(`ent: missing required field "Message.server_msg_seq"`)}
	}
	if _, ok := mc.mutation.MsgType(); !ok {
		return &ValidationError{Name: "msg_type", err: errors.New(`ent: missing required field "Message.msg_type"`)}
	}
	if _, ok := mc.mutation.MsgData(); !ok {
		return &ValidationError{Name: "msg_data", err: errors.New(`ent: missing required field "Message.msg_data"`)}
	}
	if v, ok := mc.mutation.MsgData(); ok {
		if err := message.MsgDataValidator(v); err != nil {
			return &ValidationError{Name: "msg_data", err: fmt.Errorf(`ent: validator failed for field "Message.msg_data": %w`, err)}
		}
	}
	if _, ok := mc.mutation.MsgResCode(); !ok {
		return &ValidationError{Name: "msg_res_code", err: errors.New(`ent: missing required field "Message.msg_res_code"`)}
	}
	if _, ok := mc.mutation.MsgFeature(); !ok {
		return &ValidationError{Name: "msg_feature", err: errors.New(`ent: missing required field "Message.msg_feature"`)}
	}
	if _, ok := mc.mutation.MsgStatus(); !ok {
		return &ValidationError{Name: "msg_status", err: errors.New(`ent: missing required field "Message.msg_status"`)}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		_node = &Message{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: message.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: message.FieldID,
			},
		}
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.Created(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldCreated,
		})
		_node.Created = value
	}
	if value, ok := mc.mutation.Updated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: message.FieldUpdated,
		})
		_node.Updated = value
	}
	if value, ok := mc.mutation.SessionKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldSessionKey,
		})
		_node.SessionKey = value
	}
	if value, ok := mc.mutation.From(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: message.FieldFrom,
		})
		_node.From = value
	}
	if value, ok := mc.mutation.To(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldTo,
		})
		_node.To = value
	}
	if value, ok := mc.mutation.SessionType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: message.FieldSessionType,
		})
		_node.SessionType = value
	}
	if value, ok := mc.mutation.ClientMsgID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldClientMsgID,
		})
		_node.ClientMsgID = value
	}
	if value, ok := mc.mutation.ServerMsgSeq(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: message.FieldServerMsgSeq,
		})
		_node.ServerMsgSeq = value
	}
	if value, ok := mc.mutation.MsgType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: message.FieldMsgType,
		})
		_node.MsgType = value
	}
	if value, ok := mc.mutation.MsgData(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldMsgData,
		})
		_node.MsgData = value
	}
	if value, ok := mc.mutation.MsgResCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: message.FieldMsgResCode,
		})
		_node.MsgResCode = value
	}
	if value, ok := mc.mutation.MsgFeature(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: message.FieldMsgFeature,
		})
		_node.MsgFeature = value
	}
	if value, ok := mc.mutation.MsgStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: message.FieldMsgStatus,
		})
		_node.MsgStatus = value
	}
	return _node, _spec
}

// MessageCreateBulk is the builder for creating many Message entities in bulk.
type MessageCreateBulk struct {
	config
	builders []*MessageCreate
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessageCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessageCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
