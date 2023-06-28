package service

import (
	"encoding/json"
	"fmt"
	"time"

	"asocks-ws/internal/config"
	"github.com/Shopify/sarama"
)

type Consumer struct {
	services  *Services
	kafkaHost string
}

type InitMessage struct {
	Command string `json:"command"`
	Data    string `json:"data"`
}

func NewConsumer(cfg config.KafkaConfig, services *Services) *Consumer {
	return &Consumer{
		services:  services,
		kafkaHost: cfg.Host + ":" + cfg.Port,
	}
}

func (c *Consumer) InitKafkaConsumer() {
	fmt.Printf("Consumer")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Interval = time.Millisecond * 50
	// consumer
	consumer, err := sarama.NewConsumer([]string{c.kafkaHost}, config)
	if err != nil {
		fmt.Printf("Create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("lpm-commands", 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var messageInitLpm InitMessage
			err := json.Unmarshal(msg.Value, &messageInitLpm)
			if err != nil {
				return
			}

			if messageInitLpm.Command == "init" {
				fmt.Println("[data]", messageInitLpm.Data)

				server, err := c.services.Server.FindByIP(messageInitLpm.Data)
				if err != nil {
					return
				}

				// TODO change from request
				userProxies, err := c.services.UserProxy.FindByServerIP(server.Ip)
				if err != nil {
					return
				}
	

				go func() {
					c.services.UserProxy.SendKafkaInitMessage(userProxies, messageInitLpm.Data)
				}()
			}

		case err := <-partitionConsumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}
}
