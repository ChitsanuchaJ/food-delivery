package producer

import (
	"encoding/json"
	"events"

	"github.com/IBM/sarama"
)

type EventProducer interface {
	Produce(events.Event) error
}

type eventProducer struct {
	producer sarama.SyncProducer
}

func NewEventProducer(producer sarama.SyncProducer) EventProducer {
	return eventProducer{producer}
}

func (ep eventProducer) Produce(event events.Event) error {
	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := sarama.ProducerMessage{
		Topic: event.GetName(),
		Value: sarama.ByteEncoder(value),
	}

	_, _, err = ep.producer.SendMessage(&msg)
	if err != nil {
		return err
	}

	return nil
}
