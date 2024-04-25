package api

import "sync"

type ApiClient struct {
	mu   sync.RWMutex
	data *string
}

func NewApiClient() *ApiClient {
	return &ApiClient{}
}

func (ac *ApiClient) GetInfo() error {
	return nil
}

func (ac *ApiClient) UpdateInfo() error {
	return nil
}

func (ac *ApiClient) GetInfoByName() error {
	return nil
}