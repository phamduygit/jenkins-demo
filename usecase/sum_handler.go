package usecase

import (
	"errors"

	"github.com/phamduygit/jenkins-demo/domain"
)

type SumHandler struct {
	SumService domain.SumService
}

func (h *SumHandler) Sum(req *domain.SumNumberRequest) (*domain.SumNumberResponse, error) {
	sum, err := h.SumService.Sum(req.Number1, req.Number2)
	if err != nil {
        return nil, errors.New("failed: " + err.Error())
    }

	return &domain.SumNumberResponse{Result: sum}, nil
}

func NewSumHandler(sumService domain.SumService) *SumHandler {
    return &SumHandler{SumService: sumService}
}