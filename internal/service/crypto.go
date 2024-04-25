package service

type ApiClient interface {
	GetInfo() error
	UpdateInfo() error
	GetInfoByName() error
}

type Service struct {
	apiClient  ApiClient
}

func NewService(apiClient ApiClient) *Service{
	return &Service{
		apiClient: apiClient,
	}
}
