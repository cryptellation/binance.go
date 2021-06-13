package mock

import (
	"context"
	"errors"
	"testing"
)

func TestNewService(t *testing.T) {
	m := New()

	// Check nil
	if m == nil {
		t.Fatal("Mock service should not be nil")
	}
}

func TestNewCandleStickService(t *testing.T) {
	// Create a mock service
	m := New()

	// Check nil on new candlestick service
	s := m.NewCandleStickService()
	if s == nil {
		t.Fatal("New candlestick service should not be nil")
	}

	// Check type
	if _, ok := s.(*CandleStickService); !ok {
		t.Fatal("Service is not the good type")
	}
}

func TestAddCandleSticks(t *testing.T) {
	m := New()
	m.AddCandleSticks(TestCandleSticks)

	css := m.NewCandleStickService()
	cs, _ := css.Do(context.TODO())

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

func TestNextError(t *testing.T) {
	m := New()
	m.NextError(errors.New("Some error"))
	if _, err := m.NewCandleStickService().Do(context.TODO()); err == nil {
		t.Error("There should be an error on candlestick service")
	}
}
