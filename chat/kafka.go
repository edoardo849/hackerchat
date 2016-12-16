package chat

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

type kafkaClient struct {
	client sarama.Client
}

// NewKafkaClient implements basic client using kafka
func NewKafkaClient(addrs []string, conf *sarama.Config) (BasicClient, error) {
	client, err := sarama.NewClient(addrs, conf)
	if err != nil {
		return nil, err
	}

	kafka := kafkaClient{
		client: client,
	}

	return kafka, nil
}

func (kc kafkaClient) Send(dest string, msg Message) error {
	producer, err := sarama.NewSyncProducerFromClient(kc.client)
	if err != nil {
		return err
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	pMsg := &sarama.ProducerMessage{
		Topic:     dest,
		Partition: 0,
		Value:     sarama.ByteEncoder(msgBytes),
	}

	_, _, err = producer.SendMessage(pMsg)
	if err != nil {
		return err
	}

	return nil
}

func (kc kafkaClient) Receive(src string) (Message, error) {
	var message Message

	consumer, err := sarama.NewConsumerFromClient(kc.client)
	if err != nil {
		return message, err
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	pc, err := consumer.ConsumePartition(src, 0, sarama.OffsetNewest)
	if err != nil {
		return message, nil
	}

	defer func() {
		if err := pc.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	msg := <-pc.Messages()
	err = json.Unmarshal(msg.Value, &message)
	return message, err
}
