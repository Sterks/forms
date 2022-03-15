package app

import (
	"forms/internal/config"
	"forms/pkg/logger"
)

func Run(configPath string) {
	_, err := config.Init(configPath)
	if err != nil {
		logger.Error(err)
		return
	}
}
