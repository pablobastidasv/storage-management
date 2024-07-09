package bootstrap

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
)

func withLog(action func() error) error {
	// https://betterstack.com/community/guides/logging/go/zap/
	logger := zap.Must(zap.NewProduction())
	if env := os.Getenv(EnvVar); env == "" {
		logger = zap.Must(zap.NewDevelopment())
	}
	defer logger.Sync()

	// https://betterstack.com/community/guides/logging/logging-in-go/
	opts := &zapslog.HandlerOptions{
		AddSource: true,
	}
	newLogger := slog.New(zapslog.NewHandler(logger.Core(), opts))
	slog.SetDefault(newLogger)

	return action()
}

