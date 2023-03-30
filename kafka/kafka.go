package kafka

import (
	"context"
	"errors"
	"log"
	"time"

	"git.brilliantmonkey.net/frigate/frigate-clips/types"
	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	config       *types.KafkaConfigSection
	writerConfig kafka.WriterConfig
	readerConfig kafka.ReaderConfig
}

func NewKafkaClient(config *types.KafkaConfigSection) *KafkaClient {
	var writerConfig kafka.WriterConfig
	if config.Producer == nil {
		log.Println("Kafka producer config not set.")
	} else {
		writerConfig = kafka.WriterConfig{
			Brokers: config.Brokers,
			Topic:   config.Producer.Topic,
		}
	}

	var readerConfig kafka.ReaderConfig
	if config.Consumer == nil {
		log.Println("Kafka consumer config not set.")
	} else {
		readerConfig = kafka.ReaderConfig{
			Brokers: config.Brokers,
			Topic:   config.Consumer.Topic,
			GroupID: config.GroupId,
		}
	}

	return &KafkaClient{
		config,
		writerConfig,
		readerConfig,
	}
}

func (client KafkaClient) Publish(message []byte) (err error) {
	writer := kafka.NewWriter(client.writerConfig)
	writer.AllowAutoTopicCreation = true

	defer writer.Close()

	const retries = 3
	for i := 0; i < retries; i++ {
		err = writer.WriteMessages(context.Background(), kafka.Message{
			Value: message,
		})
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			log.Println("Failed to publish message. Leader not available yet. Retrying...")
			time.Sleep(time.Millisecond * 250)
			continue
		}
	}
	return
}

func (client KafkaClient) Consume(callback func(message []byte) error) {
	r := kafka.NewReader(client.readerConfig)
	defer func() {
		if err := r.Close(); err != nil {
			log.Printf("Failed to close Kafka reader: %s", err)
		}
	}()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		if err = callback(m.Value); err != nil {
			log.Print(err)
		}
	}
}
