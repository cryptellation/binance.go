package mock

import (
	"context"
	"time"

	binance "github.com/cryptellation/binance.go"
	"github.com/cryptellation/models.go"
)

// DefaultCandleStickServiceLimit is the limit for CandleStick service if none is specified
var DefaultCandleStickServiceLimit = 1000

// MockedCandleSticks are candlesticks that can be used in MockCandleStickService
type MockedCandleSticks struct {
	Symbol       string
	Period       int64
	CandleSticks []models.CandleStick
}

// MockedCandleStickService is the mocked service for candlesticks
type MockedCandleStickService struct {
	MockedCandleSticks []MockedCandleSticks

	// Next request specifications
	symbol  string
	period  int64
	endTime time.Time
	limit   int
}

func newMockedCandleStickService(cs []MockedCandleSticks) *MockedCandleStickService {
	return &MockedCandleStickService{
		MockedCandleSticks: cs,
		limit:              DefaultCandleStickServiceLimit,
	}
}

// Do will execute a request for candlesticks
func (m *MockedCandleStickService) Do(ctx context.Context) ([]models.CandleStick, error) {
	cs := make([]models.CandleStick, 0)
	count := 0
	for _, t := range m.MockedCandleSticks {
		// Check if symbol is set and correspond
		if m.symbol != "" && t.Symbol != m.symbol {
			continue
		}

		// Check if period is set and correspond
		// TODO check if period is valid or throw an error
		if m.period != 0 && t.Period != m.period {
			continue
		}

		// Check each candle
		for _, c := range t.CandleSticks {
			// Check if count as trespassed limit
			if count >= m.limit {
				break
			}

			// Check if endtime is send and correspond
			if !m.endTime.IsZero() && m.endTime.After(c.Time) {
				continue
			}

			// Add it if it passed tests
			cs = append(cs, c)
			count++
		}
	}
	return cs, nil
}

// Symbol will specify a symbol for next candlesticks request
func (m *MockedCandleStickService) Symbol(symbol string) binance.CandleStickServiceInterface {
	m.symbol = symbol
	return m
}

// Period will specify a period for next candlesticks request
func (m *MockedCandleStickService) Period(period int64) binance.CandleStickServiceInterface {
	m.period = period
	return m
}

// EndTime will specify the time where the list ends (earliest time) for
// next candlesticks request
func (m *MockedCandleStickService) EndTime(endTime time.Time) binance.CandleStickServiceInterface {
	m.endTime = endTime
	return m
}

// Limit will specify the number of candlesticks the list should have at its maximum
// If the limit is higher than the default limit, it will be limited to this one
func (m *MockedCandleStickService) Limit(limit int) binance.CandleStickServiceInterface {
	if limit < DefaultCandleStickServiceLimit {
		m.limit = limit
	} else {
		m.limit = DefaultCandleStickServiceLimit
	}
	return m
}
