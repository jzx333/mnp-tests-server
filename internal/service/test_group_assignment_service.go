package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"mnp-tests-server/internal/dto"
	"mnp-tests-server/internal/repo"
)

type TestGroupAssignmentService struct {
	assignmentRepo *repo.TestGroupAssignmentRepo
	userRepo       *repo.UserRepo
	testRepo       *repo.TestRepo
}

func NewTestGroupAssignmentService(
	assignmentRepo *repo.TestGroupAssignmentRepo,
	userRepo *repo.UserRepo,
	testRepo *repo.TestRepo,
) *TestGroupAssignmentService {
	return &TestGroupAssignmentService{
		assignmentRepo: assignmentRepo,
		userRepo:       userRepo,
		testRepo:       testRepo,
	}
}

func (s *TestGroupAssignmentService) AssignTestToGroup(curatorID uuid.UUID, testID, groupID int, deadline *time.Time) (int, error) {
	if _, err := s.testRepo.GetByID(testID); err != nil {
		return 0, errors.New("test not found")
	}

	users, err := s.userRepo.GetByCuratorID(curatorID)
	if err != nil {
		return 0, errors.New("cannot fetch users")
	}
	if len(users) == 0 {
		return 0, errors.New("curator has no users")
	}

	id, err := s.assignmentRepo.Create(&dto.TestGroupAssignment{
		TestID:     testID,
		GroupID:    groupID,
		AssignedBy: curatorID,
		AssignedAt: time.Now(),
		Deadline:   deadline,
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *TestGroupAssignmentService) GetAssignmentsByCurator(curatorID uuid.UUID) ([]dto.TestGroupAssignment, error) {
	assignments, err := s.assignmentRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var filtered []dto.TestGroupAssignment
	for _, a := range assignments {
		users, _ := s.userRepo.GetByGroupID(a.GroupID)
		for _, u := range users {
			if u.CuratorID != nil && *u.CuratorID == curatorID {
				filtered = append(filtered, a)
				break
			}
		}
	}

	return filtered, nil
}

func (s *TestGroupAssignmentService) UpdateDeadline(assignmentID int, deadline *time.Time) error {
	assignment, err := s.assignmentRepo.GetByID(assignmentID)
	if err != nil {
		return err
	}
	assignment.Deadline = deadline
	return s.assignmentRepo.Update(assignment)
}
