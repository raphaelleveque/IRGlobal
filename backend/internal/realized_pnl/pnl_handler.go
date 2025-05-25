package realizedpnl

import "github.com/raphaelleveque/IRGlobal/backend/internal/domain"

type RealizedPNLHandler struct {
	service domain.RealizedPNLService
}

func NewRealizedPNLHandler(service domain.RealizedPNLService) *RealizedPNLHandler {
	return &RealizedPNLHandler{service: service}
}