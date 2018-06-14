package log

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
}

func Logger() *zap.Logger {
	return logger
}

func SubscribeToErrorChan(errors <-chan error) {
	go func() {
		for {
			e := <-errors
			logger.Error(e.Error())
		}
	}()
}