package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created").Immutable().Default(time.Now).Comment("创建时间"),
		field.Time("updated").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}
