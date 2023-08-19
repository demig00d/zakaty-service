package app

import (
	"embed"
	"fmt"
	"github.com/demig00d/zakaty-service/config"
	"github.com/demig00d/zakaty-service/pkg/puzzlebot"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	v1 "github.com/demig00d/zakaty-service/internal/controller/http"
	"github.com/demig00d/zakaty-service/internal/usecase/impl"
	"github.com/demig00d/zakaty-service/pkg/httpserver"
	"github.com/demig00d/zakaty-service/pkg/logger"
	"github.com/demig00d/zakaty-service/pkg/sheets"
	"github.com/gin-gonic/gin"
)

//go:embed credentials.json
var credentials embed.FS

func Run(cfg config.Config) {
	l := logger.New(cfg.Level)

	spreadsheet, err := sheets.NewSpreadsheet(credentials, cfg.Sheet.SpreadsheetId)
	if err != nil {
		l.Fatal(err)
	}

	client := &http.Client{}
	pb := puzzlebot.NewPuzzleBot(client, cfg.Token)

	tournamentUseCase := impl.NewTournamentImpl(spreadsheet, cfg.Sheet.Columns)

	// HTTP Server
	handler := gin.New()

	v1.NewRouter(handler, l, tournamentUseCase, pb)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
