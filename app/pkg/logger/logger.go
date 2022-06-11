package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	zapcore "ozon-url-shortener/app/pkg/logger/zap"
)

func New(path string) (*zap.Logger, func() error, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("create/open log file (%s): %w", path, err)
	}

	info, err := file.Stat()
	if err != nil {
		return nil, nil, fmt.Errorf("get file stats: %w", err)
	}

	if info.Size() > 0 {
		_, err = file.Write([]byte("\n\n"))
		if err != nil {
			return nil, nil, fmt.Errorf("write to file: %w", err)
		}
	}

	return zapcore.NewLogger(os.Stdout, file), file.Close, nil
}
