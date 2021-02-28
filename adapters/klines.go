package adapters

import (
	"fmt"
	"strconv"
	"time"

	binance "github.com/adshao/go-binance/v2"

	"github.com/cryptellation/models.go"
)

// Intervals represents every intervals supported by Binance API
func Intervals() []int64 {
	return []int64{
		models.M1,
		models.M3,
		models.M5,
		models.M15,
		models.M30,
		models.H1,
		models.H2,
		models.H4,
		models.H6,
		models.H8,
		models.H12,
		models.D1,
		models.D3,
		models.W1,
	}
}

// PeriodToInterval converts an interval to its corresponding epoch
func PeriodToInterval(period int64) (e string, err error) {
	switch period {
	case models.M1:
		return "1m", nil
	case models.M3:
		return "3m", nil
	case models.M5:
		return "5m", nil
	case models.M15:
		return "15m", nil
	case models.M30:
		return "30m", nil
	case models.H1:
		return "1h", nil
	case models.H2:
		return "2h", nil
	case models.H4:
		return "4h", nil
	case models.H6:
		return "6h", nil
	case models.H8:
		return "8h", nil
	case models.H12:
		return "12h", nil
	case models.D1:
		return "1d", nil
	case models.D3:
		return "3d", nil
	case models.W1:
		return "1w", nil
	default:
		return e, fmt.Errorf("interval error: unknown period")
	}
}

// TimeCandleStickToKLine will take the time from a candle and will convert it to Kline time
func TimeCandleStickToKLine(t time.Time) int64 {
	return t.Unix() * 1000
}

// TimeKLineToCandleStick will take the time from a kline and will convert it to candle time
func TimeKLineToCandleStick(t int64) time.Time {
	return time.Unix(t/1000, 0)
}

// KLineToCandleStick will convert KLine binance format for CandleStick
func KLineToCandleStick(k binance.Kline) (models.CandleStick, error) {
	var c models.CandleStick

	// Convert Open
	open, err := strconv.ParseFloat(k.Open, 64)
	if err != nil {
		return c, err
	}

	// Convert High
	high, err := strconv.ParseFloat(k.High, 64)
	if err != nil {
		return c, err
	}

	// Convert Low
	low, err := strconv.ParseFloat(k.Low, 64)
	if err != nil {
		return c, err
	}

	// Convert Close
	close, err := strconv.ParseFloat(k.Close, 64)
	if err != nil {
		return c, err
	}

	// Instanciate Candle
	c = models.CandleStick{
		Time:  TimeKLineToCandleStick(k.OpenTime),
		Open:  open,
		High:  high,
		Low:   low,
		Close: close,
	}

	return c, nil
}

// KLinesToCandleSticks will transform a slice of binance format for CandleStick
func KLinesToCandleSticks(kl []*binance.Kline) ([]models.CandleStick, error) {
	var err error

	cs := make([]models.CandleStick, len(kl))
	for i, k := range kl {
		if cs[i], err = KLineToCandleStick(*k); err != nil {
			return nil, err
		}
	}

	return cs, nil
}
