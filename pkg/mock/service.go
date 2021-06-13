package mock

import (
	interfaces "github.com/cryptellation/binance.go/pkg/binance"
)

// MockedService represents the Binance service mocked
type MockedService struct {
	candleSticks []CandleSticks
	nextError    error
}

// New will create a mocked service
func New() *MockedService {
	return &MockedService{}
}

// NewCandleStickService will create a new candlestick service
func (m *MockedService) NewCandleStickService() interfaces.CandleStickServiceInterface {
	candleService := newCandleStickService(m.candleSticks)
	candleService.SetError(m.nextError)
	return candleService
}

// AddCandleSticks will add fake candlesticks to service that can be used in candlestick services
func (m *MockedService) AddCandleSticks(cs []CandleSticks) {
	m.candleSticks = append(m.candleSticks, cs...)
}

// NextError will set an error for the next Do() on any child service
func (m *MockedService) NextError(err error) {
	m.nextError = err
}
