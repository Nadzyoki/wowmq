package http

import (
	"context"
	"log/slog"

	"github.com/Nadzyoki/wowmq/internal/models"
)

type HTTPListener struct {
	ch     chan models.RawMessage
	logger *slog.Logger
}

func NewHTTPListener(logger *slog.Logger) *HTTPListener {
	httpListener := &HTTPListener{
		ch:     make(chan models.RawMessage),
		logger: logger,
	}

	return httpListener
}

func (hl *HTTPListener) StartListen() error {
	hl.logger.Debug("start http listener")
	return nil
}
func (hl *HTTPListener) StopListen(context.Context) error {
	close(hl.ch)
	hl.logger.Debug("stop http listener")
	return nil
}
