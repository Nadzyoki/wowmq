package sorter

import (
	"context"
	"log/slog"

	"github.com/Nadzyoki/wowmq/internal/listener"
	"github.com/Nadzyoki/wowmq/internal/models"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)



type Sorter struct {
	logger *slog.Logger
}

func NewSorter(lc fx.Lifecycle, listenerRaw *listener.Listener, logger *slog.Logger) *Sorter {
	srt := &Sorter{
		logger: logger,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := srt.StartSort(listenerRaw.GetChannelRawMessages())
			if err != nil {
				return errors.Wrap(err, "start sorter")
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := srt.StopSort()
			if err != nil {
				return errors.Wrap(err, "stop sorter")
			}
			return nil
		},
	})

	return srt
}

func (srt *Sorter) StartSort(chRaw <-chan models.RawMessage) error {
	srt.logger.Debug("start sorter")
	return nil
}

func (srt *Sorter) StopSort() error {
	srt.logger.Debug("stop sorter")
	return nil
}
