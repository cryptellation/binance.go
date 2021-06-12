package adapters

import (
	"testing"
	"time"

	binance "github.com/adshao/go-binance/v2"

	"github.com/cryptellation/models.go"
)

var testCasesKLineToCandleStick = []struct {
	KLine       binance.Kline
	CandleStick models.CandleStick
}{
	{
		KLine:       binance.Kline{OpenTime: 0, Open: "1.0", High: "2.0", Low: "0.5", Close: "1.5"},
		CandleStick: models.CandleStick{Time: time.Unix(0, 0), Open: 1, High: 2, Low: 0.5, Close: 1.5},
	},
	{
		KLine:       binance.Kline{OpenTime: 0, Open: "2.0", High: "4.0", Low: "1", Close: "3"},
		CandleStick: models.CandleStick{Time: time.Unix(0, 0), Open: 2, High: 4, Low: 1, Close: 3},
	},
}

func TestKLineToCandleStick(t *testing.T) {
	for i, test := range testCasesKLineToCandleStick {
		cs, err := KLineToCandleStick(test.KLine)
		if err != nil {
			t.Error("There should be no error on CandleStick", i, ":", err)
		} else if test.CandleStick != cs {
			t.Error("CandleStick", i, "is not transformed correctly:", test.CandleStick, cs)
		}
	}
}

func TestKLineToCandleStick_IncorrectOpen(t *testing.T) {
	c := binance.Kline{OpenTime: 0, Open: "error", High: "2.0", Low: "0.5", Close: "1.5"}
	if _, err := KLineToCandleStick(c); err == nil {
		t.Error("There should be an error on open")
	}
}

func TestKLineToCandleStick_IncorrectHigh(t *testing.T) {
	c := binance.Kline{OpenTime: 0, Open: "1.0", High: "error", Low: "0.5", Close: "1.5"}
	if _, err := KLineToCandleStick(c); err == nil {
		t.Error("There should be an error on high")
	}
}

func TestKLineToCandleStick_IncorrectLow(t *testing.T) {
	c := binance.Kline{OpenTime: 0, Open: "1.0", High: "2.0", Low: "error", Close: "1.5"}
	if _, err := KLineToCandleStick(c); err == nil {
		t.Error("There should be an error on low")
	}
}

func TestKLineToCandleStick_IncorrectClose(t *testing.T) {
	c := binance.Kline{OpenTime: 0, Open: "1.0", High: "2.0", Low: "0.5", Close: "error"}
	if _, err := KLineToCandleStick(c); err == nil {
		t.Error("There should be an error on close")
	}
}

func TestKLinesToCandleSticks(t *testing.T) {
	// Only get klines
	kl := make([]*binance.Kline, len(testCasesKLineToCandleStick))
	for i := range testCasesKLineToCandleStick {
		kl[i] = &testCasesKLineToCandleStick[i].KLine
	}

	// Test function
	cs, err := KLinesToCandleSticks(kl)
	if err != nil {
		t.Error("There should be no error:", err)
	}

	for i, test := range testCasesKLineToCandleStick {
		if test.CandleStick != cs[i] {
			t.Error("CandleStick", i, "is not transformed correctly:", test.CandleStick, cs[i])
		}
	}
}

func TestKLinesToCandleSticks_IncorrectOpen(t *testing.T) {
	c := []*binance.Kline{{OpenTime: 0, Open: "error", High: "2.0", Low: "0.5", Close: "1.5"}}
	if _, err := KLinesToCandleSticks(c); err == nil {
		t.Error("There should be an error on open")
	}
}

func TestKLinesToCandleSticks_IncorrectHigh(t *testing.T) {
	c := []*binance.Kline{{OpenTime: 0, Open: "1.0", High: "error", Low: "0.5", Close: "1.5"}}
	if _, err := KLinesToCandleSticks(c); err == nil {
		t.Error("There should be an error on high")
	}
}

func TestKLinesToCandleSticks_IncorrectLow(t *testing.T) {
	c := []*binance.Kline{{OpenTime: 0, Open: "1.0", High: "2.0", Low: "error", Close: "1.5"}}
	if _, err := KLinesToCandleSticks(c); err == nil {
		t.Error("There should be an error on low")
	}
}

func TestKLinesToCandleSticks_IncorrectClose(t *testing.T) {
	c := []*binance.Kline{{OpenTime: 0, Open: "1.0", High: "2.0", Low: "0.5", Close: "error"}}
	if _, err := KLinesToCandleSticks(c); err == nil {
		t.Error("There should be an error on close")
	}
}

var timeKLineToCandleStickTests = []struct {
	BinanceTimestamp int64
	Time             time.Time
}{
	{BinanceTimestamp: 1257894000000, Time: time.Unix(1257894000, 0)},
}

func TestTimeKLineToCandleStick(t *testing.T) {
	for i, c := range timeKLineToCandleStickTests {
		r := TimeKLineToCandleStick(c.BinanceTimestamp)
		if r != c.Time {
			t.Error("Times don't match on test", i, ":", c.Time, r)
		}
	}
}

func TestTimeCandleStickToKLine(t *testing.T) {
	for i, c := range timeKLineToCandleStickTests {
		r := TimeCandleStickToKLine(c.Time)
		if r != c.BinanceTimestamp {
			t.Error("Times don't match on test", i, ":", c.BinanceTimestamp, r)
		}
	}
}

var possibleIntervals = map[int64]string{
	models.M1:  "1m",
	models.M3:  "3m",
	models.M5:  "5m",
	models.M15: "15m",
	models.M30: "30m",
	models.H1:  "1h",
	models.H2:  "2h",
	models.H4:  "4h",
	models.H6:  "6h",
	models.H8:  "8h",
	models.H12: "12h",
	models.D1:  "1d",
	models.D3:  "3d",
	models.W1:  "1w",
}

func TestPeriodToInterval(t *testing.T) {
	for k, v := range possibleIntervals {
		if e, err := PeriodToInterval(k); err != nil {
			t.Error("Interval for Period", k, "should not throw an error:", err)
		} else if e != v {
			t.Error("Interval for Period", k, "does not correspond : should be", v, "but is", e)
		}
	}
}

func TestPeriodToInterval_InexistantPeriod(t *testing.T) {
	if _, err := PeriodToInterval(0); err == nil {
		t.Error("Period 0 should throw an error")
	}
}
