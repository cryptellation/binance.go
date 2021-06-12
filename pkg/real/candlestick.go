package real

import (
	"context"
	"time"

	"github.com/cryptellation/binance.go/internal/adapters"

	binance "github.com/adshao/go-binance/v2"
	"github.com/cryptellation/models.go"

	"github.com/cryptellation/binance.go/pkg/interfaces"
)

// CandleStickService is the real service for candlesticks
type CandleStickService struct {
	service *binance.KlinesService
}

// Do will execute a request for candlesticks
func (s *CandleStickService) Do(ctx context.Context) ([]models.CandleStick, error) {
	// Get KLines
	kl, err := s.service.Do(ctx)
	if err != nil {
		return nil, err
	}

	// Change them to right format
	return adapters.KLinesToCandleSticks(kl)
}

// Symbol will specify a symbol for next candlesticks request
func (s *CandleStickService) Symbol(symbol string) interfaces.CandleStickServiceInterface {
	s.service.Symbol(symbol)
	return s
}

// Period will specify a period for next candlesticks request
func (s *CandleStickService) Period(period int64) interfaces.CandleStickServiceInterface {
	interval, err := adapters.PeriodToInterval(period)
	if err != nil {
		interval = "unknown"
	}

	s.service.Interval(interval)
	return s
}

// EndTime will specify the time where the list ends (earliest time) for
// next candlesticks request
func (s *CandleStickService) EndTime(endTime time.Time) interfaces.CandleStickServiceInterface {
	binanceTime := adapters.TimeCandleStickToKLine(endTime)
	s.service.EndTime(binanceTime)
	return s
}

// Limit will specify the number of candlesticks the list should have at its maximum
// If the limit is higher than the default limit, it will be limited to this one
func (s *CandleStickService) Limit(limit int) interfaces.CandleStickServiceInterface {
	s.service.Limit(limit)
	return s
}
