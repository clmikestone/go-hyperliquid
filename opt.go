package hyperliquid

import (
	"os"

	"github.com/sonirico/vago/lol"
)

type Opt[T any] func(*T)

func (o Opt[T]) Apply(opt *T) {
	o(opt)
}

type (
	ClientOpt   = Opt[Client]
	ExchangeOpt = Opt[Exchange]
	InfoOpt     = Opt[Info]
	WsOpt       = Opt[WebsocketClient]
)

func WsOptDebugMode() WsOpt {
	return func(w *WebsocketClient) {
		w.debug = true
		w.logger = lol.NewZerolog(
			lol.WithLevel(lol.LevelTrace),
			lol.WithWriter(os.Stderr),
			lol.WithEnv(lol.EnvDev),
		)
	}
}

func InfoOptDebugMode() InfoOpt {
	return func(i *Info) {
		i.debug = true
	}
}

func ExchangeOptDebugMode() ExchangeOpt {
	return func(e *Exchange) {
		e.debug = true
	}
}

func ClientOptDebugMode() ClientOpt {
	return func(c *Client) {
		c.debug = true
		c.logger = lol.NewZerolog(
			lol.WithLevel(lol.LevelTrace),
			lol.WithWriter(os.Stderr),
			lol.WithEnv(lol.EnvDev),
		)
	}
}

// ExchangeOptClientOptions allows passing of ClientOpt to Client
func ExchangeOptClientOptions(opts ...ClientOpt) ExchangeOpt {
	return func(e *Exchange) {
		e.clientOpts = append(e.clientOpts, opts...)
	}
}

// ExchangeOptInfoOptions allows passing of InfoOpt to Info
func ExchangeOptInfoOptions(opts ...InfoOpt) ExchangeOpt {
	return func(e *Exchange) {
		e.infoOpts = append(e.infoOpts, opts...)
	}
}

// InfoOptClientOptions allows passing of ClientOpt to Info
func InfoOptClientOptions(opts ...ClientOpt) InfoOpt {
	return func(i *Info) {
		i.clientOpts = append(i.clientOpts, opts...)
	}
}
