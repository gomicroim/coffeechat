package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/gomicroim/gomicroim/v2/pkg/ent/mixin"
)

type Session struct {
	ent.Schema
}

// Fields of the Session
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").Unique().Comment("自增ID"),
		field.String("user_id").MaxLen(32).Comment("用户ID"),
		field.String("peer_id").MaxLen(32).Comment("对方ID"),
		field.Int8("session_type").Comment("会话类型，1:单聊，2:群聊"),
		field.Int8("session_status").Comment("会话状态，0:未知，1:正常，2:删除"),
	}
}

func (Session) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Edges of the Session
func (Session) Edges() []ent.Edge {
	return nil
}
