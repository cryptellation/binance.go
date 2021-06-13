package mock

import (
	"context"
	"time"

	interfaces "github.com/cryptellation/binance.go/pkg/binance"
	"github.com/cryptellation/models.go"
)

// DefaultCandleStickServiceLimit is the limit for CandleStick service if none is specified
var DefaultCandleStickServiceLimit = 1000

// TestCandleSticks are candle sticks that can be used for test
var TestCandleSticks = []CandleSticks{
	{
		Symbol: "BTC-USDC", Period: models.M1, CandleSticks: []models.CandleStick{
			{Time: time.Time{}, Open: 10, High: 10, Low: 10, Close: 10},
			{Time: time.Time{}, Open: 15, High: 15, Low: 15, Close: 15}},
	},
	{
		Symbol: "ETH-USDC", Period: models.M5, CandleSticks: []models.CandleStick{
			{Time: time.Time{}, Open: 20, High: 20, Low: 20, Close: 20},
			{Time: time.Time{}, Open: 25, High: 25, Low: 25, Close: 25}},
	},
	{
		Symbol: "IOTA-USDC", Period: models.M15, CandleSticks: []models.CandleStick{
			{Time: time.Unix(1257894000, 0), Open: 30, High: 30, Low: 30, Close: 30},
			{Time: time.Unix(1257894900, 0), Open: 35, High: 35, Low: 35, Close: 35}},
	},
	{
		Symbol: "BTC-USDC", Period: models.M5, CandleSticks: []models.CandleStick{
			{Time: time.Unix(1257894000, 0), Open: 30, High: 30, Low: 30, Close: 30},
			{Time: time.Unix(1257894300, 0), Open: 35, High: 35, Low: 35, Close: 35}},
	},
}

// TestCandleSticksCount will return the count of candles for TestCandleSticks
func TestCandleSticksCount() (count int) {
	for _, c := range TestCandleSticks {
		count += len(c.CandleSticks)
	}
	return count
}

// CandleSticks are candlesticks that can be used in MockCandleStickService
type CandleSticks struct {
	Symbol       string
	Period       int64
	CandleSticks []models.CandleStick
}

// CandleStickService is the mocked service for candlesticks
type CandleStickService struct {
	candleSticks []CandleSticks

	symbol  string
	period  int64
	endTime time.Time
	limit   int
	err     error
}

func newCandleStickService(cs []CandleSticks) *CandleStickService {
	return &CandleStickService{
		candleSticks: cs,
		limit:        DefaultCandleStickServiceLimit,
	}
}

// Do will execute a request for candlesticks
func (m *CandleStickService) Do(ctx context.Context) ([]models.CandleStick, error) {
	cs := make([]models.CandleStick, 0)
	count := 0

	if m.err != nil {
		return cs, m.err
	}

	for _, t := range m.candleSticks {
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
func (m *CandleStickService) Symbol(symbol string) interfaces.CandleStickServiceInterface {
	m.symbol = symbol
	return m
}

// Period will specify a period for next candlesticks request
func (m *CandleStickService) Period(period int64) interfaces.CandleStickServiceInterface {
	m.period = period
	return m
}

// EndTime will specify the time where the list ends (earliest time) for
// next candlesticks request
func (m *CandleStickService) EndTime(endTime time.Time) interfaces.CandleStickServiceInterface {
	m.endTime = endTime
	return m
}

// Limit will specify the number of candlesticks the list should have at its maximum
// If the limit is higher than the default limit, it will be limited to this one
func (m *CandleStickService) Limit(limit int) interfaces.CandleStickServiceInterface {
	if limit < DefaultCandleStickServiceLimit {
		m.limit = limit
	} else {
		m.limit = DefaultCandleStickServiceLimit
	}
	return m
}

// SetError will set an error that will be raised each time a Do() is executed
// You can set it at nil if you want to deactivate it
func (m *CandleStickService) SetError(err error) {
	m.err = err
}
