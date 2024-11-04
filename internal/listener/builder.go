package listener

import (
	"log/slog"

	"github.com/Nadzyoki/wowmq/internal/config"
	"github.com/Nadzyoki/wowmq/internal/listener/http"
)

func buildSource(tp config.TypeListen, logger *slog.Logger) source {
	switch tp {
	case config.HTTP:
		return http.NewHTTPListener(logger)
	default:
		return http.NewHTTPListener(logger)
	}
}
