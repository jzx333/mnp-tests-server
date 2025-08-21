package service

import (
	"time"

	"mnp-tests-server/internal/dto"
	"mnp-tests-server/internal/repo"
)

type TestAttemptAnswerService struct {
	answerRepo *repo.TestAttemptAnswerRepo
}

func NewTestAttemptAnswerService(answerRepo *repo.TestAttemptAnswerRepo) *TestAttemptAnswerService {
	return &TestAttemptAnswerService{
		answerRepo: answerRepo,
	}
}

func (s *TestAttemptAnswerService) CreateAnswer(answer *dto.TestAttemptAnswer) (int, error) {
	if answer.AnsweredAt.IsZero() {
		answer.AnsweredAt = time.Now()
	}
	return s.answerRepo.Create(answer)
}

func (s *TestAttemptAnswerService) UpdateAnswer(answer *dto.TestAttemptAnswer) error {
	answer.AnsweredAt = time.Now()
	return s.answerRepo.Update(answer)
}

func (s *TestAttemptAnswerService) GetAnswersByAttempt(attemptID int) ([]*dto.TestAttemptAnswer, error) {
	return s.answerRepo.GetByAttempt(attemptID)
}

func (s *TestAttemptAnswerService) DeleteAnswer(id int) error {
	return s.answerRepo.Delete(id)
}
