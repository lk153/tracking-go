package server

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"
	"golang.org/x/sync/errgroup"

	"factory/exam/utils/logger"
	"factory/exam/utils/shutdown"
	"factory/exam/utils/tracer"
)

//Manager ...
type Manager struct {
	httpServer  *HTTPServer
	tracerFlush func()
}

//NewServerManager ...
func NewServerManager(
	httpServer *HTTPServer,
) *Manager {
	return &Manager{
		httpServer: httpServer,
	}
}

//StartAll ...
func (m *Manager) StartAll(parentCtx context.Context) error {
	logger.InitLogger()
	m.tracerFlush = tracer.InitTracer()
	eg, ctx := errgroup.WithContext(parentCtx)
	eg.Go(func() error {
		return shutdown.BlockListen(ctx, m.httpServer)
	})

	return eg.Wait()
}

//CloseAll ...
func (m *Manager) CloseAll() error {
	sentry.Flush(2 * time.Second)
	m.tracerFlush()
	return nil
}
