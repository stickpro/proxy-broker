package service

import (
	"asocks-ws/internal/config"
	"asocks-ws/internal/domain"
	"asocks-ws/internal/enums"
	"asocks-ws/internal/repository"
	"asocks-ws/pkg/logger"
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

type UserProxyService struct {
	repository repository.UserProxy
	kafkaHost  string
}

type UserProxyServiceInterface interface {
	FindById(uint64) (domain.UserProxy, error)
	UpdateExtIp(message domain.MessageUserProxy, input domain.UpdateUserProxy) (domain.UserProxy, error)
	FindByServerId(uint64) ([]domain.UserProxy, error)
	FindByServerIP(string) ([]domain.UserProxy, error)
	SendKafkaInitMessage(userProxies []domain.UserProxy, topic string)
	SendKafkaMessage(userProxies domain.UserProxy, topic string) (int32, error)
}

func NewUserProxyService(repository repository.UserProxy, kafkaConfig config.KafkaConfig) *UserProxyService {
	return &UserProxyService{
		repository: repository,
		kafkaHost:  kafkaConfig.Host + ":" + kafkaConfig.Port,
	}
}

func (u *UserProxyService) FindById(id uint64) (domain.UserProxy, error) {
	return u.repository.LoadById(id)
}

func (u *UserProxyService) FindByServerId(id uint64) ([]domain.UserProxy, error) {
	return u.repository.FindByColumn(id, "server_id")
}

func (u *UserProxyService) FindByServerIP(ip string) ([]domain.UserProxy, error) {
	return u.repository.FindByIp(ip)
}

func (u *UserProxyService) LoadAll() ([]domain.UserProxy, error) {
	return u.repository.GetAll()
}

func (u *UserProxyService) UpdateExtIp(message domain.MessageUserProxy, input domain.UpdateUserProxy) (domain.UserProxy, error) {
	return u.repository.UpdateOne(message, input)
}

func (u *UserProxyService) SendKafkaInitMessage(userProxies []domain.UserProxy, topic string) {
	producer, err := NewKafkaAsyncProducer([]string{u.kafkaHost}, "2.5.0", WithRetryMax(5))
	if err != nil {
		logger.Error("Panic", err)
	}
	//var messages []*sarama.ProducerMessage

	for _, userProxy := range userProxies {
		j, err := json.Marshal(formatMessageForBroker(userProxy))
		if err != nil {
			logger.Error("[Error formatting json]", err)
		}
		message := &sarama.ProducerMessage{Topic: "lpm-" + topic, Value: sarama.StringEncoder(j)}
		producer.Input() <- message
	}

	var errCount = 1
	go func() {
		c := 0
		for {
			select {
			case <-producer.Successes():
			case err, ok := <-producer.Errors():
				if ok {
					log.Printf("SendAsyncMessage, err=%v", err)
					c++
					if c > errCount {
						time.Sleep(time.Hour)
					}
				}
			}
		}
	}()
}

func (u *UserProxyService) SendKafkaMessage(userProxy domain.UserProxy, topic string) (int32, error) {
	producer, err := NewKafkaProducer([]string{u.kafkaHost}, "2.5.0", WithRetryMax(5))
	if err != nil {
		panic(err)
	}
	j, err := json.Marshal(formatMessageForBroker(userProxy))

	if err != nil {
		logger.Error("[Error formatting json]", err)
	}

	message := &sarama.ProducerMessage{Topic: "lpm-" + topic, Value: sarama.StringEncoder(j)}
	sendMessage, _, err := producer.SendMessage(message)
	if err != nil {
		logger.Error("[Error send message kafka]", err)
		return 0, err
	}

	err = producer.Close()
	if err != nil {
		logger.Error("[Can't close connections", err)
		return 0, err
	}
	return sendMessage, err
}

func ipsList(ips []domain.UserAccessIps) []string {
	var list []string
	for _, ip := range ips {
		list = append(list, ip.Ip)
	}
	return list
}

func formatMessageForBroker(userProxy domain.UserProxy) domain.UserProxyForLpm {
	var proxyType enums.ProxyType
	var proxyAuthType enums.ProxyAuthType
	var session uint64
	var timeout int
	var level int32
	var hold string

	if userProxy.TypeId == enums.RotateConnection {
		switch userProxy.MethodRotate {
		case enums.MethodRotateManually:
			session = userProxy.Session
			hold = "session"
		case enums.MethodRotateEverRequest:
			hold = "query"
		case enums.MethodRotateTimeout:
			hold = "timeout"
			timeout = userProxy.Timeout
		}
	} else {
		session = userProxy.Id

		switch userProxy.TypeId {
		case enums.KeepProxy:
			hold = "hardsession"
		case enums.KeepConnection:
			hold = "session"
			level = 0
		case enums.KeepConnectionLowTrust:
			hold = "session"
			level = 0
		}
	}
	userProxyForLpm := domain.UserProxyForLpm{
		Id:     userProxy.Id,
		UserId: userProxy.UserId,
		Status: userProxy.Status,
		LoginData: domain.LoginData{
			Id:      userProxy.Id,
			Zone:    proxyType.GetProxyType(userProxy.ProxyTypeId),
			Country: userProxy.CountryCode,
			State:   userProxy.StateId,
			City:    userProxy.CityId,
			Asn:     userProxy.Asn,
			LastIp:  userProxy.ExtIp,
			Session: session,
			Hold:    hold,
			Timeout: timeout,
			Level:   level,
		},
		ServerPort: userProxy.Port,
		Auth: domain.AuthProxy{
			Type:     proxyAuthType.GetProxyAuthType(userProxy.AuthTypeId),
			Login:    userProxy.Login,
			Password: userProxy.Password,
			Ips:      ipsList(userProxy.Ips),
		},
		Updated: userProxy.UpdatedAt.Unix(),
	}
	return userProxyForLpm
}
