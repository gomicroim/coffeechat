package schema

import (
	"CoffeeChat/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the User entity.
type Message struct {
	ent.Schema
}

// Fields of the Message
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Comment("自增ID"),
		field.String("sessionKey").MaxLen(32).Comment("session唯一标志"),
		field.Int64("from").Comment("发送人Id"),
		field.String("to").MaxLen(32).Comment("接收目标ID"),
		field.Int8("session_type").Comment("会话类型，0:未知 1:单聊 2:群聊"),

		field.String("client_msg_id").MaxLen(32).Comment("客户端生成的消息UUID(去重)"),
		field.Int64("server_msg_seq").Comment("服务端生成的消息序号(乱序处理)"),

		field.Int8("msg_type").Comment("消息类型"),
		field.String("msg_data").MaxLen(4096).Comment("消息内容"),

		field.Int8("msg_res_code").Comment("消息错误码"),
		field.Int8("msg_feature").Comment("消息属性"),
		field.Int8("msg_status").Comment("消息状态"),
	}
}

func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Edges of the Message
func (Message) Edges() []ent.Edge {
	return nil
}
