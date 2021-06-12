package binance

import (
	"github.com/adshao/go-binance/v2"
)

// Service represents the real Binance service
type Service struct {
	client *binance.Client
}

// New will create a new real binance service
func New(apiKey, secretKey string) ServiceInterface {
	return &Service{
		client: binance.NewClient(apiKey, secretKey),
	}
}

// NewCandleStickService will create a new real candlestick service
func (s *Service) NewCandleStickService() CandleStickServiceInterface {
	return &CandleStickService{
		service: s.client.NewKlinesService(),
	}
}
