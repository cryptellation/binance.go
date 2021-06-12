package mock

import (
	"context"
	"testing"
)

func TestNewService(t *testing.T) {
	m := New()

	// Check nil
	if m == nil {
		t.Fatal("Mock service should not be nil")
	}
}

func TestNewCandleStickservice(t *testing.T) {
	// Create a mock service
	m := New()

	// Check nil on new candlestick service
	s := m.NewCandleStickService()
	if s == nil {
		t.Fatal("New candlestick service should not be nil")
	}

	// Check type
	if _, ok := s.(*CandleStickservice); !ok {
		t.Fatal("Service is not the good type")
	}
}

func TestAddCandleSticks(t *testing.T) {
	// Create a new mock service
	m := New()

	// Add candlesticks
	m.AddCandleSticks(TestCandleSticks)

	// Get candlestick service and get candles
	css := m.NewCandleStickService()
	cs, _ := css.Do(context.TODO())

	// Check candlesticks count
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
