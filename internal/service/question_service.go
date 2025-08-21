package service

import (
	"errors"
	"github.com/google/uuid"
	"mnp-tests-server/internal/dto"
	"mnp-tests-server/internal/repo"
)

var allowedQuestionTypes = map[string]bool{
	"single_choice":   true,
	"multiple_choice": true,
	"text":            true,
	"matching":        true,
}

type QuestionService struct {
	questionRepo *repo.QuestionRepo
	poolRepo     *repo.QuestionPoolRepo
}

func NewQuestionService(questionRepo *repo.QuestionRepo, poolRepo *repo.QuestionPoolRepo) *QuestionService {
	return &QuestionService{
		questionRepo: questionRepo,
		poolRepo:     poolRepo,
	}
}

func (s *QuestionService) CreateQuestion(userRole string, userID uuid.UUID, q *dto.Question) (int, error) {
	if q.PoolID == nil {
		return 0, errors.New("pool_id is required")
	}

	if q.Text == "" {
		return 0, errors.New("question text is required")
	}

	if q.Score <= 0 {
		q.Score = 1
	}

	// проверка и установка типа вопроса
	if q.QuestionType == "" {
		q.QuestionType = "single_choice"
	} else if !allowedQuestionTypes[q.QuestionType] {
		return 0, errors.New("invalid question_type")
	}

	pool, err := s.poolRepo.GetByID(*q.PoolID)
	if err != nil {
		return 0, errors.New("pool not found")
	}

	switch userRole {
	case "hr":
	case "curator":
		if pool.OwnerID == nil || *pool.OwnerID != userID {
			return 0, errors.New("curator can only add questions to own pool")
		}
	default:
		return 0, errors.New("unauthorized role")
	}

	if q.Position <= 0 {
		lastPos, _ := s.questionRepo.GetMaxPosition(*q.PoolID)
		q.Position = lastPos + 1
	}

	id, err := s.questionRepo.Create(q)
	if err != nil {
		return 0, err
	}

	return id, nil
}
