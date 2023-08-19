package usecase

import "github.com/demig00d/zakaty-service/pkg/puzzlebot"

type Raiting = string

type Tournament interface {
	// GetRating gets top three by earned sum from table and user
	GetRating(puzzlebot.User) (Raiting, error)
}
