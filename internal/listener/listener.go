package listener

import (
	"context"
	"log/slog"

	"github.com/Nadzyoki/wowmq/internal/config"
	"github.com/Nadzyoki/wowmq/internal/models"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

type source interface {
	GetChannelRawMessages() chan models.RawMessage
	StartListen() error
	StopListen(context.Context) error
}

type Listener struct {
	src source
}

func NewListener(lc fx.Lifecycle, cfg *config.Config, logger *slog.Logger) *Listener {
	listener := &Listener{}

	listener.src = buildSource(cfg.ListenerConfig.TypeListen, logger)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := listener.src.StartListen()
			if err != nil {
				return errors.Wrap(err, "listener : start listen in source")
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := listener.src.StopListen(ctx)
			if err != nil {
				return errors.Wrap(err, "listener : stop listen in source")
			}
			return nil
		},
	})

	return listener
}

func (lsnr *Listener) GetChannelRawMessages() <-chan models.RawMessage {
	return lsnr.src.GetChannelRawMessages()
}
