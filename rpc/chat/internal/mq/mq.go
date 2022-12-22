package mq

import "github.com/google/wire"

// ProviderSet is mq providers
var ProviderSet = wire.NewSet(NewMsgProducer, NewChatConsumerGroup)
