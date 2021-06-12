package mock

import (
	interfaces "github.com/cryptellation/binance.go/pkg/binance"
)

// MockedService represents the Binance service mocked
type MockedService struct {
	candleSticks []CandleSticks
}

// New will create a mocked service
func New() *MockedService {
	return &MockedService{}
}

// NewCandleStickService will create a new candlestick service
func (m *MockedService) NewCandleStickService() interfaces.CandleStickServiceInterface {
	return newCandleStickservice(m.candleSticks)
}

// AddCandleSticks will add fake candlesticks to service that can be used in candlestick services
func (m *MockedService) AddCandleSticks(cs []CandleSticks) {
	m.candleSticks = append(m.candleSticks, cs...)
}
