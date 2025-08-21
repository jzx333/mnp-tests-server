package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"mnp-tests-server/internal/dto"
	"mnp-tests-server/internal/repo"
)

type QuestionPoolService struct {
	poolRepo *repo.QuestionPoolRepo
}

func NewQuestionPoolService(poolRepo *repo.QuestionPoolRepo) *QuestionPoolService {
	return &QuestionPoolService{
		poolRepo: poolRepo,
	}
}

func (s *QuestionPoolService) CreatePool(userRole string, userID uuid.UUID, p *dto.QuestionPool) (int, error) {
	if p.Name == "" {
		return 0, errors.New("name is required")
	}

	if p.UnitID == 0 {
		return 0, errors.New("unit_id is required")
	}

	switch userRole {
	case "hr":
		p.OwnerID = nil
	case "curator":
		p.OwnerID = &userID
	default:
		return 0, errors.New("unauthorized role")
	}

	p.CreatedBy = userID

	if p.Description == "" {
		p.Description = "" // оставляем пустым, можно передать nil если поле nullable
	}

	if p.TimeLimitSeconds == 0 {
		p.TimeLimitSeconds = 0 // можно задать любое значение по умолчанию
	}

	p.CreatedAt = time.Now()

	id, err := s.poolRepo.Create(p)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *QuestionPoolService) GetAvailablePoolsForUser(userRole string, userID uuid.UUID, userUnitID int) ([]dto.QuestionPool, error) {
	var pools []dto.QuestionPool
	switch userRole {
	case "hr":
		// HR видит все пулы
		var err error
		pools, err = s.poolRepo.GetAll()
		if err != nil {
			return nil, err
		}
	case "curator":
		// куратор видит свои приватные и общие пулы своего юнита
		var err error
		pools, err = s.poolRepo.GetByCuratorOrUnit(userID, userUnitID)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unauthorized role")
	}
	return pools, nil
}
