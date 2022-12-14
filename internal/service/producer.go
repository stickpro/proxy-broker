package service

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"time"
)

type ProducerConfigOption func(*sarama.Config)

func NewKafkaProducerConfig(version string, opts ...ProducerConfigOption) (*sarama.Config, error) {
	kafkaVersion, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		return nil, err
	}

	config := sarama.NewConfig()
	config.Version = kafkaVersion

	for _, opt := range opts {
		opt(config)
	}

	return config, nil
}

func WithClientID(clientID string) ProducerConfigOption {
	return func(c *sarama.Config) {
		c.ClientID = clientID
	}
}

func WithVerbose() ProducerConfigOption {
	return func(_ *sarama.Config) {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}
}

func WithRetryMax(max int) ProducerConfigOption {
	return func(c *sarama.Config) {
		c.Producer.Retry.Max = max
	}
}

func WithRetryBackoff(duration time.Duration) ProducerConfigOption {
	return func(c *sarama.Config) {
		c.Producer.Retry.Backoff = duration
	}
}

func WithReturnSuccesses(isReturnSuccess bool) ProducerConfigOption {
	return func(c *sarama.Config) {
		c.Producer.Return.Successes = isReturnSuccess
	}
}

func WithReturnErrors(isReturnError bool) ProducerConfigOption {
	return func(c *sarama.Config) {
		c.Producer.Return.Errors = isReturnError
	}
}

func WithRequiredAcks(reqAcks sarama.RequiredAcks) ProducerConfigOption {
	return func(c *sarama.Config) {
		c.Producer.RequiredAcks = reqAcks
	}
}

func WithMaxMessageBytes(max int) ProducerConfigOption {
	return func(c *sarama.Config) {
		c.Producer.MaxMessageBytes = max
	}
}

func WithTLS(withTLS bool) ProducerConfigOption {
	return func(c *sarama.Config) {
		c.Net.TLS.Enable = withTLS
	}
}

func WithSASL(user, pass string) ProducerConfigOption {
	return func(c *sarama.Config) {
		c.Net.SASL.Enable = true
		c.Net.SASL.User = user
		c.Net.SASL.Password = pass
	}
}

func WithoutSASL() ProducerConfigOption {
	return func(c *sarama.Config) {
		c.Net.SASL.Enable = false
		c.Net.SASL.User = ""
		c.Net.SASL.Password = ""
	}
}

func NewKafkaProducer(brokers []string, version string, opts ...ProducerConfigOption) (producer sarama.SyncProducer, err error) {
	opts = append(opts, WithReturnSuccesses(true))
	config, err := NewKafkaProducerConfig(version, opts...)
	if err != nil {
		return
	}

	producer, err = sarama.NewSyncProducer(brokers, config)
	return
}

func NewKafkaAsyncProducer(brokers []string, version string, opts ...ProducerConfigOption) (producer sarama.AsyncProducer, err error) {
	config, err := NewKafkaProducerConfig(version, opts...)
	if err != nil {
		return
	}

	producer, err = sarama.NewAsyncProducer(brokers, config)
	return
}
