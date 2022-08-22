package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

// middleware - функция принимающая и возвращающая эндпоинт
// Тип Middleware предоставляется Go kit
// Поскольку StringService определен как интерфейс, необходимо добавить новый тип,
// оборачивающий существующий StringService и делающие дополнительные обязанности
type loggingMiddleware struct {
	logger log.Logger
	next   StringService
}

func (mw loggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Uppercase(s)
	return
}

func (mw loggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.next.Count(s)
	return
}
