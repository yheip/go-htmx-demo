package httpserver

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

type Server struct {
	httpServer *http.Server
}

type ServerOption func(*http.Server)

func WithTimeout(read, write time.Duration) ServerOption {
	return func(s *http.Server) {
		s.ReadTimeout = read
		s.WriteTimeout = write
	}
}

func WithPort(port string) ServerOption {
	return func(s *http.Server) {
		s.Addr = ":" + port
	}
}

func WithHandler(handler http.Handler) ServerOption {
	return func(s *http.Server) {
		s.Handler = handler
	}
}

func New(ctx context.Context, opts ...ServerOption) *Server {
	const (
		readTimeout  = 10 * time.Second
		writeTimeout = 60 * time.Second
		addr         = ":8080"
	)

	defaultHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK")) // nolint:errcheck
	}
	s := &http.Server{
		Addr:         addr,
		Handler:      http.HandlerFunc(defaultHandler),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		BaseContext:  func(l net.Listener) context.Context { return ctx },
	}

	for _, opt := range opts {
		opt(s)
	}

	return &Server{httpServer: s}
}

func (r *Server) Start(ctx context.Context) error {
	log := zerolog.Ctx(ctx)

	listener, err := net.Listen("tcp", r.httpServer.Addr)
	if err != nil {
		return err
	}

	log.Info().Msgf("http server started on %s", r.httpServer.Addr)

	done := make(chan error)
	go func() {
		defer close(done)
		err = r.httpServer.Serve(listener)
		if err != nil {
			if err == http.ErrServerClosed {
				log.Info().Msg("http server closed")
				done <- err
				return
			}
			log.Error().Err(err).Msg("server error")
			done <- err
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case e := <-done:
		return e
	}
}

func (r *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	r.httpServer.Shutdown(ctx)
}
