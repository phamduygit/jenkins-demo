package service

import (
	"errors"

	"github.com/phamduygit/jenkins-demo/domain"
)

type SumService struct {
}

// Sum implements domain.SumService.
func (s *SumService) Sum(num1 int, num2 int) (int, error) {
	if num1 < 0 || num2 < 0 {
		return 0, errors.New("numbers must be positive")
	}
	return num1 + num2, nil
}

func NewSumService() domain.SumService {
	return &SumService{}
}
