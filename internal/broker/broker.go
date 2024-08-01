package broker

import (
	"errors"
	"log/slog"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type OnDeliveryFunc func(err error)

type Broker struct {
	producer *kafka.Producer
	topic    string
}

func New(p *kafka.Producer, topic string) *Broker {
	return &Broker{
		producer: p,
		topic:    topic,
	}
}

func (k *Broker) Produce(data []byte) error {
	return k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &k.topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)
}

func (k *Broker) ProduceWithDelivery(data []byte, onDelivery OnDeliveryFunc) error {
	deliveryChan := make(chan kafka.Event)

	err := k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &k.topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, deliveryChan)

	if err != nil {
		close(deliveryChan)
		return err
	}

	go func() {
		defer close(deliveryChan)
		event := <-deliveryChan
		slog.Debug("kafka event", "event", event.String())
		m, ok := event.(*kafka.Message)
		if !ok {
			onDelivery(errors.New("unexpected event type"))
			return
		}
		onDelivery(m.TopicPartition.Error)
	}()

	return nil
}
