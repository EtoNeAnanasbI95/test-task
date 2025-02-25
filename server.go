package test_task

import (
	"context"
	"fmt"
	"github.com/EtoNeAnanasbI95/test-task/internal/config"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler, cfg *config.Config) error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%v", cfg.ApiPort),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    cfg.ApiTimeout,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
