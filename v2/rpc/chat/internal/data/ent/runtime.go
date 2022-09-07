// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat/internal/data/ent/message"
	"chat/internal/data/ent/session"
	"chat/internal/data/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	messageMixin := schema.Message{}.Mixin()
	messageMixinFields0 := messageMixin[0].Fields()
	_ = messageMixinFields0
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCreated is the schema descriptor for created field.
	messageDescCreated := messageMixinFields0[0].Descriptor()
	// message.DefaultCreated holds the default value on creation for the created field.
	message.DefaultCreated = messageDescCreated.Default.(func() time.Time)
	// messageDescUpdated is the schema descriptor for updated field.
	messageDescUpdated := messageMixinFields0[1].Descriptor()
	// message.DefaultUpdated holds the default value on creation for the updated field.
	message.DefaultUpdated = messageDescUpdated.Default.(func() time.Time)
	// message.UpdateDefaultUpdated holds the default value on update for the updated field.
	message.UpdateDefaultUpdated = messageDescUpdated.UpdateDefault.(func() time.Time)
	// messageDescSessionKey is the schema descriptor for sessionKey field.
	messageDescSessionKey := messageFields[1].Descriptor()
	// message.SessionKeyValidator is a validator for the "sessionKey" field. It is called by the builders before save.
	message.SessionKeyValidator = messageDescSessionKey.Validators[0].(func(string) error)
	// messageDescTo is the schema descriptor for to field.
	messageDescTo := messageFields[3].Descriptor()
	// message.ToValidator is a validator for the "to" field. It is called by the builders before save.
	message.ToValidator = messageDescTo.Validators[0].(func(string) error)
	// messageDescClientMsgID is the schema descriptor for client_msg_id field.
	messageDescClientMsgID := messageFields[5].Descriptor()
	// message.ClientMsgIDValidator is a validator for the "client_msg_id" field. It is called by the builders before save.
	message.ClientMsgIDValidator = messageDescClientMsgID.Validators[0].(func(string) error)
	// messageDescMsgData is the schema descriptor for msg_data field.
	messageDescMsgData := messageFields[8].Descriptor()
	// message.MsgDataValidator is a validator for the "msg_data" field. It is called by the builders before save.
	message.MsgDataValidator = messageDescMsgData.Validators[0].(func(string) error)
	sessionMixin := schema.Session{}.Mixin()
	sessionMixinFields0 := sessionMixin[0].Fields()
	_ = sessionMixinFields0
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescCreated is the schema descriptor for created field.
	sessionDescCreated := sessionMixinFields0[0].Descriptor()
	// session.DefaultCreated holds the default value on creation for the created field.
	session.DefaultCreated = sessionDescCreated.Default.(func() time.Time)
	// sessionDescUpdated is the schema descriptor for updated field.
	sessionDescUpdated := sessionMixinFields0[1].Descriptor()
	// session.DefaultUpdated holds the default value on creation for the updated field.
	session.DefaultUpdated = sessionDescUpdated.Default.(func() time.Time)
	// session.UpdateDefaultUpdated holds the default value on update for the updated field.
	session.UpdateDefaultUpdated = sessionDescUpdated.UpdateDefault.(func() time.Time)
	// sessionDescUserID is the schema descriptor for user_id field.
	sessionDescUserID := sessionFields[1].Descriptor()
	// session.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	session.UserIDValidator = sessionDescUserID.Validators[0].(func(string) error)
	// sessionDescPeerID is the schema descriptor for peer_id field.
	sessionDescPeerID := sessionFields[2].Descriptor()
	// session.PeerIDValidator is a validator for the "peer_id" field. It is called by the builders before save.
	session.PeerIDValidator = sessionDescPeerID.Validators[0].(func(string) error)
}
