// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mattn/echo-ent-example/ent/comment"
	"github.com/mattn/echo-ent-example/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeComment = "Comment"
)

// CommentMutation represents an operation that mutates the Comment nodes in the graph.
type CommentMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	text          *string
	created       *time.Time
	updated       *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Comment, error)
	predicates    []predicate.Comment
}

var _ ent.Mutation = (*CommentMutation)(nil)

// commentOption allows management of the mutation configuration using functional options.
type commentOption func(*CommentMutation)

// newCommentMutation creates new mutation for the Comment entity.
func newCommentMutation(c config, op Op, opts ...commentOption) *CommentMutation {
	m := &CommentMutation{
		config:        c,
		op:            op,
		typ:           TypeComment,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCommentID sets the ID field of the mutation.
func withCommentID(id int) commentOption {
	return func(m *CommentMutation) {
		var (
			err   error
			once  sync.Once
			value *Comment
		)
		m.oldValue = func(ctx context.Context) (*Comment, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Comment.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withComment sets the old Comment of the mutation.
func withComment(node *Comment) commentOption {
	return func(m *CommentMutation) {
		m.oldValue = func(context.Context) (*Comment, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CommentMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CommentMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *CommentMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *CommentMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *CommentMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Comment entity.
// If the Comment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommentMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *CommentMutation) ResetName() {
	m.name = nil
}

// SetText sets the "text" field.
func (m *CommentMutation) SetText(s string) {
	m.text = &s
}

// Text returns the value of the "text" field in the mutation.
func (m *CommentMutation) Text() (r string, exists bool) {
	v := m.text
	if v == nil {
		return
	}
	return *v, true
}

// OldText returns the old "text" field's value of the Comment entity.
// If the Comment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommentMutation) OldText(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldText is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldText requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldText: %w", err)
	}
	return oldValue.Text, nil
}

// ResetText resets all changes to the "text" field.
func (m *CommentMutation) ResetText() {
	m.text = nil
}

// SetCreated sets the "created" field.
func (m *CommentMutation) SetCreated(t time.Time) {
	m.created = &t
}

// Created returns the value of the "created" field in the mutation.
func (m *CommentMutation) Created() (r time.Time, exists bool) {
	v := m.created
	if v == nil {
		return
	}
	return *v, true
}

// OldCreated returns the old "created" field's value of the Comment entity.
// If the Comment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommentMutation) OldCreated(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreated is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreated requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreated: %w", err)
	}
	return oldValue.Created, nil
}

// ResetCreated resets all changes to the "created" field.
func (m *CommentMutation) ResetCreated() {
	m.created = nil
}

// SetUpdated sets the "updated" field.
func (m *CommentMutation) SetUpdated(t time.Time) {
	m.updated = &t
}

// Updated returns the value of the "updated" field in the mutation.
func (m *CommentMutation) Updated() (r time.Time, exists bool) {
	v := m.updated
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdated returns the old "updated" field's value of the Comment entity.
// If the Comment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommentMutation) OldUpdated(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdated is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdated requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdated: %w", err)
	}
	return oldValue.Updated, nil
}

// ResetUpdated resets all changes to the "updated" field.
func (m *CommentMutation) ResetUpdated() {
	m.updated = nil
}

// Op returns the operation name.
func (m *CommentMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Comment).
func (m *CommentMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CommentMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.name != nil {
		fields = append(fields, comment.FieldName)
	}
	if m.text != nil {
		fields = append(fields, comment.FieldText)
	}
	if m.created != nil {
		fields = append(fields, comment.FieldCreated)
	}
	if m.updated != nil {
		fields = append(fields, comment.FieldUpdated)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CommentMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case comment.FieldName:
		return m.Name()
	case comment.FieldText:
		return m.Text()
	case comment.FieldCreated:
		return m.Created()
	case comment.FieldUpdated:
		return m.Updated()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CommentMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case comment.FieldName:
		return m.OldName(ctx)
	case comment.FieldText:
		return m.OldText(ctx)
	case comment.FieldCreated:
		return m.OldCreated(ctx)
	case comment.FieldUpdated:
		return m.OldUpdated(ctx)
	}
	return nil, fmt.Errorf("unknown Comment field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CommentMutation) SetField(name string, value ent.Value) error {
	switch name {
	case comment.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case comment.FieldText:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetText(v)
		return nil
	case comment.FieldCreated:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreated(v)
		return nil
	case comment.FieldUpdated:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdated(v)
		return nil
	}
	return fmt.Errorf("unknown Comment field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CommentMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CommentMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CommentMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Comment numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CommentMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CommentMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CommentMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Comment nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CommentMutation) ResetField(name string) error {
	switch name {
	case comment.FieldName:
		m.ResetName()
		return nil
	case comment.FieldText:
		m.ResetText()
		return nil
	case comment.FieldCreated:
		m.ResetCreated()
		return nil
	case comment.FieldUpdated:
		m.ResetUpdated()
		return nil
	}
	return fmt.Errorf("unknown Comment field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CommentMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CommentMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CommentMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CommentMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CommentMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CommentMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CommentMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Comment unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CommentMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Comment edge %s", name)
}
