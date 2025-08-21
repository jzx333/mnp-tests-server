package service

import (
	"errors"
	"mnp-tests-server/internal/dto"
	"mnp-tests-server/internal/repo"
)

type OptionService struct {
	optionRepo   *repo.OptionRepo
	questionRepo *repo.QuestionRepo
}

func NewOptionService(optionRepo *repo.OptionRepo, questionRepo *repo.QuestionRepo) *OptionService {
	return &OptionService{
		optionRepo:   optionRepo,
		questionRepo: questionRepo,
	}
}

func (s *OptionService) CreateOption(qID int, opt *dto.Option) (int, error) {
	question, err := s.questionRepo.GetByID(qID)
	if err != nil {
		return 0, errors.New("question not found")
	}

	if question.QuestionType == "single_choice" && opt.IsCorrect {
		countCorrect, _ := s.optionRepo.CountCorrectOptions(qID)
		if countCorrect >= 1 {
			return 0, errors.New("single_choice question can have only one correct option")
		}
	}

	opt.QuestionID = qID
	id, err := s.optionRepo.Create(opt)
	if err != nil {
		return 0, err
	}

	return id, nil
}
