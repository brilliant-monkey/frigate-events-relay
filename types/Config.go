package types

type TopicConfigSection struct {
	Topic string `yaml:"topic"`
}

type KafkaConfigSection struct {
	Brokers  []string            `yaml:"brokers"`
	GroupId  string              `yaml:"group_id,omitempty"`
	Producer *TopicConfigSection `yaml:"producer,omitempty"`
	Consumer *TopicConfigSection `yaml:"consumer,omitempty"`
}

type MQTTConfigSection struct {
	Endpoint string              `yaml:"endpoint"`
	ClientId string              `yaml:"client_id"`
	Username string              `yaml:"username"`
	Password string              `yaml:"password"`
	Producer *TopicConfigSection `yaml:"producer,omitempty"`
	Consumer *TopicConfigSection `yaml:"consumer,omitempty"`
}

type Config struct {
	Kafka KafkaConfigSection `yaml:"kafka"`
	MQTT  MQTTConfigSection  `yaml:"mqtt"`
}
