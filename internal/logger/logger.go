package logger

import (
	"log/slog"
	"os"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
)

func Setup(appName string, debug bool) (*os.File, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		cacheDir = os.TempDir()
	}

	appDir := filepath.Join(cacheDir, appName)
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return nil, err
	}

	logPath := filepath.Join(appDir, appName+".log")

	f, err := tea.LogToFile(logPath, appName+": ")
	if err != nil {
		return nil, err
	}

	logLevel := slog.LevelInfo
	if debug {
		logLevel = slog.LevelDebug
	}

	opt := &slog.HandlerOptions{
		Level: logLevel,
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(f, opt)))

	return f, nil
}
