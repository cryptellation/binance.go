package binance

import (
	"context"
	"time"

	"github.com/cryptellation/models.go"
)

// Interface is an interface for service
type ServiceInterface interface {
	NewCandleStickService() CandleStickServiceInterface
}

// CandleStickServiceInterface is the interface for candle stick services
type CandleStickServiceInterface interface {
	Do(ctx context.Context) ([]models.CandleStick, error)
	Symbol(symbol string) CandleStickServiceInterface
	Period(period int64) CandleStickServiceInterface
	EndTime(endTime time.Time) CandleStickServiceInterface
	Limit(limit int) CandleStickServiceInterface
}
