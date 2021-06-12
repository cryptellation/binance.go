package real

import (
	"github.com/adshao/go-binance/v2"
	"github.com/cryptellation/binance.go/pkg/interfaces"
)

// Service represents the real Binance service
type Service struct {
	client *binance.Client
}

// New will create a new real binance service
func New(apiKey, secretKey string) interfaces.Interface {
	return &Service{
		client: binance.NewClient(apiKey, secretKey),
	}
}

// NewCandleStickService will create a new real candlestick service
func (s *Service) NewCandleStickService() interfaces.CandleStickServiceInterface {
	return &CandleStickService{
		service: s.client.NewKlinesService(),
	}
}
