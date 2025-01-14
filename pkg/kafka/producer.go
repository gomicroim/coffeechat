package kafka

import (
	"github.com/Shopify/sarama"
)

// NewSyncProducer A nil sarama.config use the default config
func NewSyncProducer(addrs []string, config *sarama.Config) (sarama.SyncProducer, error) {
	if config == nil {
		config = sarama.NewConfig()
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Partitioner = sarama.NewRandomPartitioner
	}

	// refer to:
	// https://github.com/Shopify/sarama/blob/v1.32.0/sync_producer.go#L13-L14
	// https://github.com/Shopify/sarama/blob/v1.32.0/sync_producer.go#L43
	config.Producer.Return.Successes = true

	return sarama.NewSyncProducer(addrs, config)
}

// NewSyncProducerClient A nil sarama.config use the default config
func NewSyncProducerClient(addrs []string, config *sarama.Config) (sarama.Client, sarama.SyncProducer, error) {
	if config == nil {
		config = sarama.NewConfig()
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Partitioner = sarama.NewRandomPartitioner
	}

	client, err := sarama.NewClient(addrs, config)
	if err != nil {
		return nil, nil, err
	}

	config.Producer.Return.Successes = true

	p, err := sarama.NewSyncProducerFromClient(client)
	return client, p, err
}
