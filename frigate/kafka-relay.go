package frigate

import (
	"encoding/json"
	"log"

	"git.brilliantmonkey.net/frigate/frigate-clips/kafka"
	"git.brilliantmonkey.net/frigate/frigate-clips/mqtt"
	"git.brilliantmonkey.net/frigate/frigate-clips/types"
)

type FrigateRelay struct {
	mqtt  *mqtt.MQTTClient
	kafka *kafka.KafkaClient
}

func NewFrigateRelay(config *types.Config) *FrigateRelay {
	mqtt := mqtt.NewMQTTClient(&config.MQTT)
	kafka := kafka.NewKafkaClient(&config.Kafka)
	return &FrigateRelay{
		mqtt:  mqtt,
		kafka: kafka,
	}
}

func (relay *FrigateRelay) mqttMessageCallback(message []byte) {
	var frigateEvent types.FrigateMQTTEvent
	err := json.Unmarshal(message, &frigateEvent)
	if err != nil {
		log.Printf("Failed to deserialize event. %s", err.Error())
		return
	}

	id := frigateEvent.After.Id
	log.Printf("%s: Message received.", id)

	log.Printf("%s: Publishing message to Kafka.", id)
	err = relay.kafka.Publish(message)
	if err != nil {
		log.Printf("%s: Failed to publish message to Kafka. %s", id, err)
		return
	}

	log.Printf("%s: Message published.", id)
}

func (relay *FrigateRelay) Start() (err error) {
	err = relay.mqtt.Connect()
	if err != nil {
		return
	}
	go relay.mqtt.Subscribe(relay.mqttMessageCallback)

	return
}

func (relay *FrigateRelay) Stop() error {
	relay.mqtt.Close()
	return nil
}
