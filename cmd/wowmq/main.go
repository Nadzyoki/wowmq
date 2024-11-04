package main

import (
	"fmt"
	"log/slog"

	"github.com/Nadzyoki/wowmq/internal/config"
	"github.com/Nadzyoki/wowmq/internal/listener"
	"github.com/Nadzyoki/wowmq/internal/logger"
	"github.com/Nadzyoki/wowmq/internal/sorter"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

const(
	configPathEnv = "CONFIG_PATH"
)

func main() {
	cfg, err := config.NewConfig(configPathEnv)
	if err != nil {
		fmt.Printf("failed init config : %s", err)
		return
	}
	sloglogger := logger.NewLogger(cfg)

	fx.New(
		fx.WithLogger(func() fxevent.Logger {
			return &fxevent.SlogLogger{Logger: sloglogger}
		}),
		fx.Provide(func() *slog.Logger {
			return sloglogger
		}),
		fx.Provide(func() *config.Config {
			return cfg
		}),
		fx.Provide(listener.NewListener),
		fx.Provide(sorter.NewSorter),

		fx.Invoke(func(*sorter.Sorter) {}),
	).Run()
}
