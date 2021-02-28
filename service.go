package service

import (
	binance "github.com/adshao/go-binance/v2"
)

// Interface is an interface for service
type Interface interface {
	NewCandleStickService() CandleStickServiceInterface
}

// Service represents the Binance service
type Service struct {
	client *binance.Client
}

// New will create a new binance service
func New(apiKey, secretKey string) Interface {
	return &Service{
		client: binance.NewClient(apiKey, secretKey),
	}
}

// NewCandleStickService will create a new candlestick service
func (s *Service) NewCandleStickService() CandleStickServiceInterface {
	return &CandleStickService{
		service: s.client.NewKlinesService(),
	}
}
