package config

type AppConfig struct {
	Kafka KafkaConfig `yaml:"kafka"`
	MQTT  MQTTConfig  `yaml:"mqtt"`
}
