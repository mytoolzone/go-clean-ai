// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/glog"

	"github.com/gin-gonic/gin"

	"go-clean/config"
	v1 "go-clean/internal/controller/http/v1"
	"go-clean/internal/usecase"
	"go-clean/internal/usecase/repo"
	"go-clean/pkg/httpserver"
	"go-clean/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		glog.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	translationUseCase := usecase.New(
		repo.New(pg),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, translationUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		glog.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		glog.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		glog.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
