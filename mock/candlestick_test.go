package mock

import (
	"context"
	"testing"
	"time"

	"github.com/cryptellation/models.go"
)

var tCS = []MockedCandleSticks{
	{
		"BTC-USDC", models.M1, []models.CandleStick{
			{Time: time.Time{}, Open: 10, High: 10, Low: 10, Close: 10},
			{Time: time.Time{}, Open: 15, High: 15, Low: 15, Close: 15}},
	},
	{
		"ETH-USDC", models.M5, []models.CandleStick{
			{Time: time.Time{}, Open: 20, High: 20, Low: 20, Close: 20},
			{Time: time.Time{}, Open: 25, High: 25, Low: 25, Close: 25}},
	},
	{
		"IOTA-USDC", models.M15, []models.CandleStick{
			{Time: time.Unix(1257894000, 0), Open: 30, High: 30, Low: 30, Close: 30},
			{Time: time.Unix(1257894900, 0), Open: 35, High: 35, Low: 35, Close: 35}},
	},
	{
		"BTC-USDC", models.M5, []models.CandleStick{
			{Time: time.Unix(1257894000, 0), Open: 30, High: 30, Low: 30, Close: 30},
			{Time: time.Unix(1257894300, 0), Open: 35, High: 35, Low: 35, Close: 35}},
	},
}

func TestMockedDo(t *testing.T) {
	// Get new service
	s := newMockedCandleStickService(tCS)

	// Count candlesticks
	var count int
	for _, c := range tCS {
		count += len(c.CandleSticks)
	}

	// Do the service
	cs, _ := s.Do(context.Background())
	if len(cs) != count {
		t.Error("There should be", count, "candlesticks")
	}

	// Test first case
	for i, c := range tCS[0].CandleSticks {
		if c != cs[i] {
			t.Error("Candlesticks", i, "don't correspond")
		}
	}

	// Test second case
	for i, c := range tCS[1].CandleSticks {
		if c != cs[i+2] {
			t.Error("Candlesticks", i, "don't correspond")
		}
	}
}

func TestMockedDo_DefaultLimit(t *testing.T) {
	localTest := []MockedCandleSticks{{"BTC-USDC", models.M1, []models.CandleStick{}}}
	for i := 0; i < DefaultCandleStickServiceLimit+100; i++ {
		localTest[0].CandleSticks = append(localTest[0].CandleSticks, models.CandleStick{})
	}

	// Get new service
	s := newMockedCandleStickService(localTest)

	// Do the service
	cs, _ := s.Do(context.Background())
	if len(cs) != DefaultCandleStickServiceLimit {
		t.Error("There should be", DefaultCandleStickServiceLimit, "candlesticks but there is", len(cs))
	}
}

func TestMockedSymbolDo(t *testing.T) {
	// Get new service
	s := newMockedCandleStickService(tCS)

	// Do the service with symbol
	cs, _ := s.Symbol("BTC-USDC").Do(context.Background())
	if len(cs) != 4 {
		t.Error("There should be 4 candlesticks but there is", len(cs))
	}

	// Test corresponding case
	for i, c := range tCS[0].CandleSticks {
		if c != cs[i] {
			t.Error("Candlesticks", i, "don't correspond")
		}
	}
}

func TestMockedIntervalDo(t *testing.T) {
	// Get new service
	s := newMockedCandleStickService(tCS)

	// Do the service with period
	cs, _ := s.Period(models.M5).Do(context.Background())
	if len(cs) != 4 {
		t.Error("There should be 4 candlesticks but there is", len(cs))
	}

	// Test corresponding case
	for i, c := range tCS[1].CandleSticks {
		if c != cs[i] {
			t.Error("Candlesticks", i, "don't correspond")
		}
	}
}

func TestMockedEndTimeDo(t *testing.T) {
	// Get new service
	s := newMockedCandleStickService(tCS)

	// Do the service with endtime
	cs, _ := s.EndTime(time.Unix(1257893900, 0)).Do(context.Background())
	if len(cs) != 4 {
		t.Error("There should be 4 candlesticks but there is", len(cs))
	}

	// Test corresponding case
	for i, c := range tCS[2].CandleSticks {
		if c != cs[i] {
			t.Error("Candlesticks", i, "don't correspond: should be", c, "but is", cs[i])
		}
	}
}

func TestMockedLimitDo(t *testing.T) {
	// Get new service
	s := newMockedCandleStickService(tCS)

	// Do the service with limit
	cs, _ := s.Limit(4).Do(context.Background())
	if len(cs) != 4 {
		t.Error("There should be 4 candlesticks but there is", len(cs))
	}
}

func TestMockedLimitDo_DefaultLimitTrespassed(t *testing.T) {
	localTest := []MockedCandleSticks{{"BTC-USDC", models.M1, []models.CandleStick{}}}
	for i := 0; i < DefaultCandleStickServiceLimit+100; i++ {
		localTest[0].CandleSticks = append(localTest[0].CandleSticks, models.CandleStick{})
	}

	// Get new service
	s := newMockedCandleStickService(localTest)

	// Do the service with limit
	cs, _ := s.Limit(2000).Do(context.Background())
	if len(cs) != DefaultCandleStickServiceLimit {
		t.Error("There should be", DefaultCandleStickServiceLimit, "candlesticks but there is", len(cs))
	}
}

func TestMockedAllDo(t *testing.T) {
	// Get new service
	s := newMockedCandleStickService(tCS)

	// Do the service with limit
	tm := time.Unix(1257894200, 0)
	cs, _ := s.Symbol("BTC-USDC").Period(models.M5).EndTime(tm).Limit(1).Do(context.Background())
	if len(cs) != 1 {
		t.Error("There should be 1 candlesticks but there is", len(cs))
	}

	c := tCS[3].CandleSticks[1]
	if c != cs[0] {
		t.Error("Candlestick don't correspond: should be", c, "but is", cs[0])
	}
}
