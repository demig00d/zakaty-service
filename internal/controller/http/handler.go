package http

import (
	"github.com/demig00d/zakaty-service/internal/usecase"
	"github.com/demig00d/zakaty-service/pkg/logger"
)

type tournamentRouter struct {
	uc usecase.Tournament
	l  logger.Interface
}

//
// order, err := r.uc.GetById(c.Request.Context(), orderId)
// if err != nil {
// 	if errors.Unwrap(err) == pgx.ErrNoRows {
// 		r.l.Info("http - v1 - order not found")
// 		c.HTML(http.StatusNotFound, "notfound_order.html", nil)
// 		return
// 	}
// 	r.l.Error(err, "http - v1 - order")
// 	c.HTML(http.StatusInternalServerError, "invalid_data.html", nil)
//
// 	return
// }
//
// if err != nil {
// 	return
// }
