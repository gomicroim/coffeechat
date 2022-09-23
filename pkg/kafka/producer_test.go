package kafka

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/Shopify/sarama"
)

var (
	kafkaAddr       = []string{"127.0.0.1:9092"}
	consumerGroupId = "msg-gate"
	topic           = "test-assign"
)

func TestSyncProducer_SendMessage(t *testing.T) {
	producer, err := NewSyncProducer(kafkaAddr, nil)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		key := "recvId1"
		if i%2 == 0 {
			key = "recvId2"
		}
		value := "hello" + strconv.Itoa(i)
		msg := sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.ByteEncoder(key),
			Value: sarama.StringEncoder(value),
		}
		p, o, err := producer.SendMessage(&msg)
		if err != nil {
			t.Fatal(err)
		} else {
			log.Printf("success produce, msg=%s, partition=%d, offset=%d\n", value, p, o)
		}
		time.Sleep(time.Second * 3)
	}
}
