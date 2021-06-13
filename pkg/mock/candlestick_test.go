package mock

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cryptellation/models.go"
)

func TestMockedDo(t *testing.T) {
	s := newCandleStickService(TestCandleSticks)

	cs, _ := s.Do(context.TODO())
	if len(cs) != TestCandleSticksCount() {
		t.Fatal("There should be", TestCandleSticksCount(), "candlesticks, but there is", len(cs))
	}

	for i, c := range TestCandleSticks[0].CandleSticks {
		if c != cs[i] {
			t.Error("Candlesticks", i, "don't correspond")
		}
	}

	for i, c := range TestCandleSticks[1].CandleSticks {
		if c != cs[i+2] {
			t.Error("Candlesticks", i, "don't correspond")
		}
	}
}

func TestMockedDo_NoData(t *testing.T) {
	s := newCandleStickService(nil)

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

	s := newCandleStickService(localTest)

	cs, _ := s.Do(context.TODO())
	if len(cs) != DefaultCandleStickServiceLimit {
		t.Error("There should be", DefaultCandleStickServiceLimit, "candlesticks but there is", len(cs))
	}
}

func TestMockedSymbolDo(t *testing.T) {
	s := newCandleStickService(TestCandleSticks)

	cs, _ := s.Symbol("BTC-USDC").Do(context.TODO())
	if len(cs) != 4 {
		t.Error("There should be 4 candlesticks but there is", len(cs))
	}

	for i, c := range TestCandleSticks[0].CandleSticks {
		if c != cs[i] {
			t.Error("Candlesticks", i, "don't correspond")
		}
	}
}

func TestMockedIntervalDo(t *testing.T) {
	s := newCandleStickService(TestCandleSticks)

	cs, _ := s.Period(models.M5).Do(context.TODO())
	if len(cs) != 4 {
		t.Error("There should be 4 candlesticks but there is", len(cs))
	}

	for i, c := range TestCandleSticks[1].CandleSticks {
		if c != cs[i] {
			t.Error("Candlesticks", i, "don't correspond")
		}
	}
}

func TestMockedEndTimeDo(t *testing.T) {
	s := newCandleStickService(TestCandleSticks)

	cs, _ := s.EndTime(time.Unix(1257893900, 0)).Do(context.TODO())
	if len(cs) != 4 {
		t.Error("There should be 4 candlesticks but there is", len(cs))
	}

	for i, c := range TestCandleSticks[2].CandleSticks {
		if c != cs[i] {
			t.Error("Candlesticks", i, "don't correspond: should be", c, "but is", cs[i])
		}
	}
}

func TestMockedLimitDo(t *testing.T) {
	s := newCandleStickService(TestCandleSticks)

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

	s := newCandleStickService(localTest)

	cs, _ := s.Limit(2000).Do(context.TODO())
	if len(cs) != DefaultCandleStickServiceLimit {
		t.Error("There should be", DefaultCandleStickServiceLimit, "candlesticks but there is", len(cs))
	}
}

func TestMockedAllDo(t *testing.T) {
	s := newCandleStickService(TestCandleSticks)

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

func TestMockedDo_Error(t *testing.T) {
	s := newCandleStickService(TestCandleSticks)
	s.SetError(errors.New("Some Error"))

	tm := time.Unix(1257894200, 0)
	_, err := s.Symbol("BTC-USDC").Period(models.M5).EndTime(tm).Limit(1).Do(context.TODO())
	if err == nil {
		t.Fatal("There should be an error")
	}
}
