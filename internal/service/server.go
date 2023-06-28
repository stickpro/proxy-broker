package service

import (
	"asocks-ws/internal/domain"
	"asocks-ws/internal/repository"
)

type ServerService struct {
	repository repository.Server
}

type ServerServiceInterface interface {
	FindByIP(string) (domain.Server, error)
}

func NewServerService(repository repository.Server) *ServerService {
	return &ServerService{repository: repository}
}

func (s *ServerService) FindByIP(ip string) (domain.Server, error) {
	return s.repository.FindByColumn(ip, "ip")
}
