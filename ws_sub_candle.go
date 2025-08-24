package hyperliquid

import (
	"fmt"
)

type CandlesSubscriptionParams struct {
	Coin     string
	Interval string
}

func (w *WebsocketClient) Candles(
	params CandlesSubscriptionParams,
	callback func([]Candle, error),
) (*Subscription, error) {
	payload := remoteCandlesSubscriptionPayload{
		Type:     ChannelCandle,
		Coin:     params.Coin,
		Interval: params.Interval,
	}

	return w.subscribe(payload, func(msg any) {
		candles, ok := msg.(Candles)
		if !ok {
			callback(nil, fmt.Errorf("invalid message type"))
			return
		}

		callback(candles, nil)
	})
}
