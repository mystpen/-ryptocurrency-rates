package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"

	"github.com/mystpen/cryptocurrency-rates/config"
	"github.com/mystpen/cryptocurrency-rates/internal/model"
)

var ErrNoResponce = errors.New("no responce from API")
var ErrNoData = errors.New("no data")

type ApiClient struct {
	mu   sync.RWMutex
	data map[string]model.CryptoCoin
}

func NewApiClient() *ApiClient {
	return &ApiClient{
		data: make(map[string]model.CryptoCoin),
	}
}

func (ac *ApiClient) GetInfo() (*[]model.CryptoCoin, error) {
	resp, err := http.Get(config.Config.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotModified {
		return nil, ErrNoResponce
	}

	var CryptoCoins []model.CryptoCoin

	if err := json.NewDecoder(resp.Body).Decode(&CryptoCoins); err != nil {
		return nil, err
	}

	return &CryptoCoins, nil
}

func (ac *ApiClient) UpdateInfo(CryptoCoins *[]model.CryptoCoin) error {
	ac.mu.Lock()
	defer ac.mu.Unlock()

	for _, coinData := range *CryptoCoins {
		ac.data[coinData.Name] = coinData
	}

	return nil
}

func (ac *ApiClient) GetInfoByName(name string) (*model.CryptoCoin, error) {
	ac.mu.RLock()
	defer ac.mu.RUnlock()

	var info model.CryptoCoin
	info, ok := ac.data[name]
	if !ok {
		return nil, ErrNoData
	}

	return &info, nil
}
