// Code generated by ent, DO NOT EDIT.

package session

import (
	"time"
)

const (
	// Label holds the string label denoting the session type in the database.
	Label = "session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldUpdated holds the string denoting the updated field in the database.
	FieldUpdated = "updated"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldPeerID holds the string denoting the peer_id field in the database.
	FieldPeerID = "peer_id"
	// FieldSessionType holds the string denoting the session_type field in the database.
	FieldSessionType = "session_type"
	// FieldSessionStatus holds the string denoting the session_status field in the database.
	FieldSessionStatus = "session_status"
	// Table holds the table name of the session in the database.
	Table = "sessions"
)

// Columns holds all SQL columns for session fields.
var Columns = []string{
	FieldID,
	FieldCreated,
	FieldUpdated,
	FieldUserID,
	FieldPeerID,
	FieldSessionType,
	FieldSessionStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreated holds the default value on creation for the "created" field.
	DefaultCreated func() time.Time
	// DefaultUpdated holds the default value on creation for the "updated" field.
	DefaultUpdated func() time.Time
	// UpdateDefaultUpdated holds the default value on update for the "updated" field.
	UpdateDefaultUpdated func() time.Time
	// UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	UserIDValidator func(string) error
	// PeerIDValidator is a validator for the "peer_id" field. It is called by the builders before save.
	PeerIDValidator func(string) error
)
