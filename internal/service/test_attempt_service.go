package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"mnp-tests-server/internal/dto"
	"mnp-tests-server/internal/repo"
)

type TestAttemptService struct {
	attemptRepo  *repo.TestAttemptRepo
	answerRepo   *repo.TestAttemptAnswerRepo
	testRepo     *repo.TestRepo
	questionRepo *repo.QuestionRepo
	optionRepo   *repo.OptionRepo
}

func NewTestAttemptService(
	attemptRepo *repo.TestAttemptRepo,
	answerRepo *repo.TestAttemptAnswerRepo,
	testRepo *repo.TestRepo,
	questionRepo *repo.QuestionRepo,
	optionRepo *repo.OptionRepo,
) *TestAttemptService {
	return &TestAttemptService{
		attemptRepo:  attemptRepo,
		answerRepo:   answerRepo,
		testRepo:     testRepo,
		questionRepo: questionRepo,
		optionRepo:   optionRepo,
	}
}

func (s *TestAttemptService) StartAttempt(userID uuid.UUID, testID int, assignmentID int) (int, error) {
	_, err := s.testRepo.GetByID(testID)
	if err != nil {
		return 0, errors.New("test not found")
	}
	attempt := &dto.TestAttempt{
		TestID:       testID,
		AssignmentID: assignmentID,
		UserID:       userID,
		StartedAt:    time.Now(),
		Status:       "in_progress",
	}
	id, err := s.attemptRepo.Create(attempt)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *TestAttemptService) GetQuestionsForAttempt(attemptID int) ([]dto.Question, error) {
	attempt, err := s.attemptRepo.GetByID(attemptID)
	if err != nil {
		return nil, err
	}

	poolIDs, err := s.testRepo.GetPoolsByTestID(attempt.TestID)
	if err != nil {
		return nil, err
	}

	var questions []dto.Question
	for _, poolID := range poolIDs {
		qs, _ := s.questionRepo.GetByPoolID(poolID)
		for j := range qs {
			opts, _ := s.optionRepo.GetByQuestionID(qs[j].ID)
			qs[j].Options = opts
		}
		questions = append(questions, qs...)
	}

	return questions, nil
}

func (s *TestAttemptService) SubmitAnswer(attemptID, questionID int, optionID *int, text *string) error {
	question, err := s.questionRepo.GetByID(questionID)
	if err != nil {
		return errors.New("question not found")
	}

	answer := &dto.TestAttemptAnswer{
		AttemptID:  attemptID,
		QuestionID: questionID,
		OptionID:   optionID,
		Text:       text,
		AnsweredAt: time.Now(),
	}

	if optionID != nil && (question.QuestionType == "single_choice" || question.QuestionType == "multiple_choice") {
		option, err := s.optionRepo.GetByID(*optionID)
		if err != nil {
			return errors.New("option not found")
		}
		answer.IsCorrect = &option.IsCorrect // ✅ указатель на bool
	} else if question.QuestionType == "text" {
		isCorrect := false
		answer.IsCorrect = &isCorrect
	}

	_, err = s.answerRepo.Create(answer)
	return err
}

func (s *TestAttemptService) FinishAttempt(attemptID int) error {
	attempt, err := s.attemptRepo.GetByID(attemptID)
	if err != nil {
		return err
	}
	score, maxScore, _ := s.answerRepo.CalculateScore(attemptID)
	attempt.Score = &score
	attempt.MaxScore = &maxScore
	attempt.Status = "finished"
	now := time.Now()
	attempt.FinishedAt = &now
	return s.attemptRepo.Update(attempt)
}
