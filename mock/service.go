package mock

import binance "github.com/cryptellation/binance.go"

// MockedService represents the Binance service mocked
type MockedService struct {
}

// NewMock will create a mocked service
func NewMock() binance.Interface {
	return &MockedService{}
}

// NewCandleStickService will create a new candlestick service
func (m *MockedService) NewCandleStickService() binance.CandleStickServiceInterface {
	return &MockedCandleStickService{}
}
