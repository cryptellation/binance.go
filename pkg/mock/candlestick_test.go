package mock

import (
	"context"
	"testing"
	"time"

	"github.com/cryptellation/models.go"
)

func TestMockedDo(t *testing.T) {
	// Get new service
	s := newCandleStickservice(TestCandleSticks)

	// Do the service
	cs, _ := s.Do(context.TODO())
	if len(cs) != TestCandleSticksCount() {
		t.Fatal("There should be", TestCandleSticksCount(), "candlesticks, but there is", len(cs))
	}

	// Test first case
	for i, c := range TestCandleSticks[0].CandleSticks {
		if c != cs[i] {
			t.Error("Candlesticks", i, "don't correspond")
		}
	}

	// Test second case
	for i, c := range TestCandleSticks[1].CandleSticks {
		if c != cs[i+2] {
			t.Error("Candlesticks", i, "don't correspond")
		}
	}
}

func TestMockedDo_NoData(t *testing.T) {
	// Get new service
	s := newCandleStickservice(nil)

	// Do the service
	cs, _ := s.Do(context.TODO())
	if len(cs) != 0 {
		t.Fatal("There should be 0 candlesticks, but there is", len(cs))
	}
}

func TestMockedDo_DefaultLimit(t *testing.T) {
	localTest := []CandleSticks{{"BTC-USDC", models.M1, []models.CandleStick{}}}
	for i := 0; i < DefaultCandleStickServiceLimit+100; i++ {
		localTest[0].CandleSticks = append(localTest[0].CandleSticks, models.CandleStick{})
	}

	// Get new service
	s := newCandleStickservice(localTest)

	// Do the service
	cs, _ := s.Do(context.TODO())
	if len(cs) != DefaultCandleStickServiceLimit {
		t.Error("There should be", DefaultCandleStickServiceLimit, "candlesticks but there is", len(cs))
	}
}

func TestMockedSymbolDo(t *testing.T) {
	// Get new service
	s := newCandleStickservice(TestCandleSticks)

	// Do the service with symbol
	cs, _ := s.Symbol("BTC-USDC").Do(context.TODO())
	if len(cs) != 4 {
		t.Error("There should be 4 candlesticks but there is", len(cs))
	}

	// Test corresponding case
	for i, c := range TestCandleSticks[0].CandleSticks {
		if c != cs[i] {
			t.Error("Candlesticks", i, "don't correspond")
		}
	}
}

func TestMockedIntervalDo(t *testing.T) {
	// Get new service
	s := newCandleStickservice(TestCandleSticks)

	// Do the service with period
	cs, _ := s.Period(models.M5).Do(context.TODO())
	if len(cs) != 4 {
		t.Error("There should be 4 candlesticks but there is", len(cs))
	}

	// Test corresponding case
	for i, c := range TestCandleSticks[1].CandleSticks {
		if c != cs[i] {
			t.Error("Candlesticks", i, "don't correspond")
		}
	}
}

func TestMockedEndTimeDo(t *testing.T) {
	// Get new service
	s := newCandleStickservice(TestCandleSticks)

	// Do the service with endtime
	cs, _ := s.EndTime(time.Unix(1257893900, 0)).Do(context.TODO())
	if len(cs) != 4 {
		t.Error("There should be 4 candlesticks but there is", len(cs))
	}

	// Test corresponding case
	for i, c := range TestCandleSticks[2].CandleSticks {
		if c != cs[i] {
			t.Error("Candlesticks", i, "don't correspond: should be", c, "but is", cs[i])
		}
	}
}

func TestMockedLimitDo(t *testing.T) {
	// Get new service
	s := newCandleStickservice(TestCandleSticks)

	// Do the service with limit
	cs, _ := s.Limit(4).Do(context.TODO())
	if len(cs) != 4 {
		t.Error("There should be 4 candlesticks but there is", len(cs))
	}
}

func TestMockedLimitDo_DefaultLimitTrespassed(t *testing.T) {
	localTest := []CandleSticks{{"BTC-USDC", models.M1, []models.CandleStick{}}}
	for i := 0; i < DefaultCandleStickServiceLimit+100; i++ {
		localTest[0].CandleSticks = append(localTest[0].CandleSticks, models.CandleStick{})
	}

	// Get new service
	s := newCandleStickservice(localTest)

	// Do the service with limit
	cs, _ := s.Limit(2000).Do(context.TODO())
	if len(cs) != DefaultCandleStickServiceLimit {
		t.Error("There should be", DefaultCandleStickServiceLimit, "candlesticks but there is", len(cs))
	}
}

func TestMockedAllDo(t *testing.T) {
	// Get new service
	s := newCandleStickservice(TestCandleSticks)

	// Do the service with limit
	tm := time.Unix(1257894200, 0)
	cs, _ := s.Symbol("BTC-USDC").Period(models.M5).EndTime(tm).Limit(1).Do(context.TODO())
	if len(cs) != 1 {
		t.Error("There should be 1 candlesticks but there is", len(cs))
	}

	c := TestCandleSticks[3].CandleSticks[1]
	if c != cs[0] {
		t.Error("Candlestick don't correspond: should be", c, "but is", cs[0])
	}
}
