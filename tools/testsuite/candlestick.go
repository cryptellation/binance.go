package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cryptellation/binance.go/pkg/binance"
	"github.com/cryptellation/models.go"
)

func candlestickTest1(bService binance.ServiceInterface) error {
	testPrefix := "[Candlestick][1]"

	t, err := time.Parse("02/01/2006", "10/06/2021")
	if err != nil {
		return fmt.Errorf("%s%s", testPrefix, err)
	}

	cs, err := bService.NewCandleStickService().
		Symbol("ETHUSDT").
		Period(models.D1).
		Limit(1).
		EndTime(t).
		Do(context.TODO())
	if err != nil {
		fmt.Println(testPrefix, err)
		return fmt.Errorf("%s%s", testPrefix, err)
	}

	expectedCS := models.CandleStick{
		Time:  t,
		Open:  2610.2,
		High:  2624.46,
		Low:   2425.11,
		Close: 2471.09,
	}

	if len(cs) == 1 && cs[0].Equal(&expectedCS) {
		return nil
	} else if len(cs) != 1 {
		return fmt.Errorf("%s%s", testPrefix, "Too much data")
	} else {
		return fmt.Errorf(fmt.Sprint(testPrefix, "Expected", expectedCS, "and got", cs[0]))
	}
}

func runCandlestickTests(key, secret string) int {
	fmt.Println("Starting Candlestick tests...")

	bService := binance.New(key, secret)

	count := 0
	count += errToCount(candlestickTest1(bService))

	time.Sleep(time.Second)

	return count
}

func errToCount(err error) int {
	if err == nil {
		return 0
	}

	fmt.Println(err)
	return 1
}
