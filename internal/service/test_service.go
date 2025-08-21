package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"mnp-tests-server/internal/dto"
	"mnp-tests-server/internal/repo"
)

type TestService struct {
	testRepo     *repo.TestRepo
	poolRepo     *repo.QuestionPoolRepo
	testPoolRepo *repo.TestPoolRepo
}

func NewTestService(
	testRepo *repo.TestRepo,
	poolRepo *repo.QuestionPoolRepo,
	testPoolRepo *repo.TestPoolRepo,
) *TestService {
	return &TestService{
		testRepo:     testRepo,
		poolRepo:     poolRepo,
		testPoolRepo: testPoolRepo,
	}
}

func (s *TestService) CreateTest(title, description string, createdBy uuid.UUID) (int, error) {
	test := &dto.Test{
		Title:       title,
		Description: description,
		CreatedBy:   createdBy,
		CreatedAt:   time.Now(),
	}
	return s.testRepo.Create(test)
}

func (s *TestService) UpdateTest(testID int, title, description string) error {
	test, err := s.testRepo.GetByID(testID)
	if err != nil {
		return errors.New("test not found")
	}
	test.Title = title
	test.Description = description
	return s.testRepo.Update(test)
}

func (s *TestService) DeleteTest(testID int) error {
	_, err := s.testRepo.GetByID(testID)
	if err != nil {
		return errors.New("test not found")
	}
	return s.testRepo.Delete(testID)
}

func (s *TestService) AddPoolsToTest(testID int, poolIDs []int) error {
	for _, poolID := range poolIDs {
		_, err := s.poolRepo.GetByID(poolID)
		if err != nil {
			return errors.New("pool not found")
		}
		err = s.testPoolRepo.Add(&dto.TestPool{
			TestID: testID,
			PoolID: poolID,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *TestService) GetAvailablePoolsForCurator(curatorID uuid.UUID, unitID int) ([]dto.QuestionPool, error) {
	allPools, err := s.poolRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var available []dto.QuestionPool
	for _, p := range allPools {
		if p.UnitID == unitID || (p.OwnerID != nil && *p.OwnerID == curatorID) {
			available = append(available, p)
		}
	}
	return available, nil
}
