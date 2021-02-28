package service

import "testing"

func TestNewService(t *testing.T) {
	s := New("apiKey", "secretKey")

	// Check nil
	if s == nil {
		t.Fatal("Service should not be nil")
	}

	// Check type
	if _, ok := s.(*Service); !ok {
		t.Fatal("Service is not the good type")
	}
}

func TestNewCandleStickService(t *testing.T) {
	// Create a service
	m := New("apiKey", "secretKey")

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
