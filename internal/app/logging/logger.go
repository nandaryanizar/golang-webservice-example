package logging

import "go.uber.org/zap"

// Logger instance
var Logger *zap.Logger

func init() {
	Logger, _ = zap.NewProduction()
}
