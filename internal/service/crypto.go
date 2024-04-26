package service

import (
	"errors"
	"log"
	"time"

	"github.com/mystpen/cryptocurrency-rates/config"
	"github.com/mystpen/cryptocurrency-rates/internal/model"
	"github.com/mystpen/cryptocurrency-rates/internal/repository/api"
)

type ApiClient interface {
	GetInfo() (*[]model.CryptoCoin, error)
	UpdateInfo(*[]model.CryptoCoin) error
	GetInfoByName(name string) (*model.CryptoCoin, error)
}

type Service struct {
	apiClient ApiClient
}

func NewService(apiClient ApiClient) *Service {
	return &Service{
		apiClient: apiClient,
	}
}

func (s *Service) IntervalUpdate() {
	for {
		info, err := s.apiClient.GetInfo()
		if err != nil {
			if !errors.Is(err, api.ErrNoResponce) { //TODO: change err, maybe add chan
				log.Println(err)
			} else {
				log.Println(err)
			}
			
		}

		s.apiClient.UpdateInfo(info)

		time.Sleep(config.Config.Interval)
	}
}

func (s *Service) Init() {
	go s.IntervalUpdate()
}

func (s *Service) GetInfoByName(name string) (*model.CryptoCoin, error){
	return s.apiClient.GetInfoByName(name)
}