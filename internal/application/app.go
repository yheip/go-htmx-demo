package application

import (
	"context"
	"io"
	"os"
	"sse-demo/internal/transport/http"
	"sse-demo/pkg/httpserver"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	server *httpserver.Server
}

func New(ctx context.Context) *Application {
	router := chi.NewRouter()

	http.AddRoutes(router)

	return &Application{
		server: httpserver.New(
			ctx,
			httpserver.WithHandler(router),
		),
	}
}

func (a *Application) Start(ctx context.Context) error {
	log := newLogger()

	ctx = log.WithContext(ctx)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return a.server.Start(ctx)
	})

	return g.Wait()
}

func newLogger() *zerolog.Logger {
	var output io.Writer
	if os.Getenv("ENV") == "" {
		output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	} else {
		output = os.Stdout
	}
	log := zerolog.New(output).With().Timestamp().Logger()

	zerolog.DefaultContextLogger = &log

	return &log
}
