package examples

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/sonirico/go-hyperliquid"
)

func TestCandlesSnapshot(t *testing.T) {
	godotenv.Overload()
	info := hyperliquid.NewInfo(hyperliquid.MainnetAPIURL, true, nil, nil)

	now := time.Now()
	startTime := now.Add(-1 * time.Hour).UnixMilli()
	endTime := now.UnixMilli()

	tests := []struct {
		name     string
		coin     string
		interval string
	}{
		{name: "BTC 1m", coin: "BTC", interval: "1m"},
		{name: "ETH 5m", coin: "ETH", interval: "5m"},
		{name: "BTC 15m", coin: "BTC", interval: "15m"},
		{name: "ETH 1h", coin: "ETH", interval: "1h"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("Fetching candles for %s with interval %s", tt.coin, tt.interval)
			candles, err := info.CandlesSnapshot(tt.coin, tt.interval, startTime, endTime)
			if err != nil {
				t.Fatalf("Failed to fetch candles: %v", err)
			}

			if len(candles) == 0 {
				t.Error("Expected non-empty candles response")
			}

			// Print first candle for inspection
			first := candles[0]
			t.Logf("First candle: %+v", first)
		})
	}
}

func TestCandleWebSocket(t *testing.T) {
	ws := hyperliquid.NewWebsocketClient("")

	if err := ws.Connect(context.Background()); err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer ws.Close()

	done := make(chan bool)

	sub, err := ws.Candles(
		hyperliquid.CandlesSubscriptionParams{
			Coin:     "BTC",
			Interval: "1m",
		},
		func(candles []hyperliquid.Candle, err error) {
			if err != nil {
				t.Errorf("Error in candle callback: %v", err)
				return
			}

			for _, candle := range candles {
				t.Logf("Received candle: %+v", candle)
			}

			done <- true
		},
	)

	if err != nil {
		t.Fatalf("Failed to subscribe: %v", err)
	}

	defer sub.Close()

	select {
	case <-done:
		// Test passed
	case <-time.After(10 * time.Second):
		t.Error("Timeout waiting for candle update")
	}
}
