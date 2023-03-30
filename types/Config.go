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

func (c *MQTTConfigSection) GetEndpoint() string {
	return c.Endpoint
}

func (c *MQTTConfigSection) GetClientId() string {
	return c.ClientId
}

func (c *MQTTConfigSection) GetUsername() string {
	return c.Username
}

func (c *MQTTConfigSection) GetPassword() string {
	return c.Password
}

func (c *MQTTConfigSection) GetProducerTopic() *string {
	if c.Producer != nil {
		return &c.Producer.Topic
	}
	return nil
}

func (c *MQTTConfigSection) GetConsumerTopic() *string {
	if c.Consumer != nil {
		return &c.Consumer.Topic
	}
	return nil
}

type Config struct {
	Kafka KafkaConfigSection `yaml:"kafka"`
	MQTT  MQTTConfigSection  `yaml:"mqtt"`
}
