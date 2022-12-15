package service

import (
	"asocks-ws/internal/domain"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

type Consumer struct {
	services *Services
}

func NewConsumer(services *Services) *Consumer {
	return &Consumer{
		services: services,
	}
}

func (c *Consumer) InitKafkaConsumer() {
	fmt.Printf("Consumer")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.AutoCommit.Interval = time.Millisecond * 50

	// consumer
	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, config)
	if err != nil {
		fmt.Printf("Create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("asocks-update-ips", 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var messageUserProxy domain.MessageUserProxy
			var input domain.UpdateUserProxy
			err := json.Unmarshal([]byte(msg.Value), &messageUserProxy)
			if err != nil {
				return
			}
			input.ExtIp = messageUserProxy.IP
			ip, err := c.services.UserProxy.UpdateExtIp(messageUserProxy, input)
			if err != nil {
				return
			}
			fmt.Println(messageUserProxy)
			fmt.Println("[UpdatedProxyIp]", ip)
			fmt.Println("msg offset: ", msg.Offset, " partition: ", msg.Partition, " timestrap: ", msg.Timestamp.Format("2006-Jan-02 15:04"), " value: ", string(msg.Value))

		case err := <-partitionConsumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}
}
