package position

import "github.com/raphaelleveque/IRGlobal/backend/internal/domain"

type PositionHandler struct {
	positionService domain.PositionService
}

func NewPositionHandler(positionService domain.PositionService) *PositionHandler {
	return &PositionHandler{positionService: positionService}
}