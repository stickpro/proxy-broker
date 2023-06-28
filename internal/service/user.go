package service

import (
	"encoding/json"
	"sync"

	"asocks-ws/internal/config"
	"asocks-ws/internal/domain"
	"asocks-ws/internal/repository"
	"asocks-ws/pkg/logger"
	"github.com/Shopify/sarama"
)

type UserServiceInterface interface {
	GetAllUsers() ([]domain.User, error)
	GetUserByID(ID int) (domain.User, error)
	SendKafkaInitMessage(users []domain.User, topic string)
	SendUserToTSTopics(user domain.User) error
}

type UserService struct {
	repository   repository.User
	kafkaHost    string
	suffixTopics map[string]struct{}

	mu sync.RWMutex
}

func (u *UserService) GetUserByID(ID int) (domain.User, error) {
	return u.repository.GetByID(ID)
}

func (u *UserService) GetAllUsers() ([]domain.User, error) {
	return u.repository.GetAll()
}

func (u *UserService) SendKafkaInitMessage(users []domain.User, topic string) {
	u.mu.Lock()
	u.suffixTopics[topic] = struct{}{}
	u.mu.Unlock()

	producer, err := NewKafkaAsyncProducer([]string{u.kafkaHost}, "2.5.0", WithRetryMax(5))
	if err != nil {
		logger.Error("Panic", err)
	}

	for _, user := range users {
		bUser, err := json.Marshal(user.ToKafkaMsg())
		if err != nil {
			logger.Error("[Error formatting json]", err)
		}

		message := &sarama.ProducerMessage{Topic: "lpm-" + topic, Value: sarama.StringEncoder(bUser)}
		producer.Input() <- message
	}

	go func() {
		for {
			select {
			case <-producer.Successes():
			case err, ok := <-producer.Errors():
				if ok {
					logger.Errorf("SendAsyncMessage, err=%v", err)
				}
			}
		}
	}()
}

func (u *UserService) SendUserToTSTopics(user domain.User) error {
	producer, err := NewKafkaProducer([]string{u.kafkaHost}, "2.5.0", WithRetryMax(5))
	if err != nil {
		return err
	}

	bUser, err := json.Marshal(user.ToKafkaMsg())
	if err != nil {
		return err
	}

	u.mu.RLock()
	defer u.mu.RUnlock()

	for suffixTopic := range u.suffixTopics {
		message := &sarama.ProducerMessage{Topic: "traffic-server-" + suffixTopic, Value: sarama.StringEncoder(bUser)}
		if _, _, err := producer.SendMessage(message); err != nil {
			logger.Error("[Error send message kafka]", err)
			return err
		}
	}

	return nil
}

func NewUserService(repository repository.User, kafkaConfig config.KafkaConfig) *UserService {
	return &UserService{
		repository:   repository,
		kafkaHost:    kafkaConfig.Host + ":" + kafkaConfig.Port,
		suffixTopics: make(map[string]struct{}),
	}
}
