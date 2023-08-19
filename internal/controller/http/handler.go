package http

import (
	"github.com/demig00d/zakaty-service/internal/usecase"
	"github.com/demig00d/zakaty-service/pkg/logger"
	"github.com/demig00d/zakaty-service/pkg/puzzlebot"
)

type tournamentRouter struct {
	uc usecase.Tournament
	pb puzzlebot.Puzzlebot
	l  logger.Interface
}
