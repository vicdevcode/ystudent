package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

type Config struct {
	Port            string        `yaml:"port"             env-required:"true"`
	Host            string        `yaml:"host"             env-required:"true"`
	ReadTimeout     time.Duration `yaml:"read_timeout"     env-required:"true"`
	WriteTimeout    time.Duration `yaml:"write_timeout"    env-required:"true"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env-required:"true"`
}

func New(handler http.Handler, cfg *Config, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: cfg.ShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
