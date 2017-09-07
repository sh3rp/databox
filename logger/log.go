package logger

import (
	"runtime"

	"github.com/rs/zerolog/log"
)

func E(err error) {
	_, file, line, _ := runtime.Caller(1)
	log.Error().Msgf("[%s][%d] Error: %v", file, line, err)
}

func I(msg string) {
	_, file, line, _ := runtime.Caller(1)
	log.Info().Msgf("[%s][%d] %v", file, line, msg)
}

func OBJ(obj interface{}) {
	_, file, line, _ := runtime.Caller(1)
	log.Info().Msgf("[%s][%d] %v", file, line, obj)
}
