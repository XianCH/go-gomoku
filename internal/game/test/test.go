package main

import xlog "go-gomoku/internal/log"

func main() {
	logger, err := xlog.NewLogger()
	if err != nil {
		panic(err)
	}

	logger.INFO("hello my log")
	logger.DEBUG("FUCK U")
	logger.WARN("FUCK U")
	logger.ERROR("ERROR")
}
