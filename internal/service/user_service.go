package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"mnp-tests-server/internal/dto"
	"mnp-tests-server/internal/repo"
)

type UserService struct {
	userRepo       *repo.UserRepo
	roleRepo       *repo.RoleRepo
	unitRepo       *repo.UnitRepo
	departmentRepo *repo.DepartmentRepo
}

func NewUserService(
	userRepo *repo.UserRepo,
	roleRepo *repo.RoleRepo,
	unitRepo *repo.UnitRepo,
	departmentRepo *repo.DepartmentRepo,
) *UserService {
	return &UserService{
		userRepo:       userRepo,
		roleRepo:       roleRepo,
		unitRepo:       unitRepo,
		departmentRepo: departmentRepo,
	}
}

func (s *UserService) CreateUser(currentUserRole string, u *dto.User) (uuid.UUID, error) {
	if currentUserRole != "hr" {
		return uuid.Nil, errors.New("only HR can create users")
	}

	if u.FullName == "" {
		return uuid.Nil, errors.New("full_name is required")
	}
	if u.RoleID == 0 {
		return uuid.Nil, errors.New("role_id is required")
	}

	var practiceStart, practiceEnd time.Time
	var err error

	if u.PracticeStart != nil {
		practiceStart, err = time.Parse("2006-01-02", *u.PracticeStart)
		if err != nil {
			return uuid.Nil, errors.New("invalid practice_start format, expected YYYY-MM-DD")
		}
	}

	if u.PracticeEnd != nil {
		practiceEnd, err = time.Parse("2006-01-02", *u.PracticeEnd)
		if err != nil {
			return uuid.Nil, errors.New("invalid practice_end format, expected YYYY-MM-DD")
		}
	}

	if !practiceStart.IsZero() && !practiceEnd.IsZero() && practiceStart.After(practiceEnd) {
		return uuid.Nil, errors.New("practice_start must be before practice_end")
	}

	if _, err := s.roleRepo.GetByID(u.RoleID); err != nil {
		return uuid.Nil, errors.New("invalid role_id")
	}

	if u.UnitID != nil {
		if _, err := s.unitRepo.GetByID(*u.UnitID); err != nil {
			return uuid.Nil, errors.New("invalid unit_id")
		}
	}

	if u.DepartmentID != nil {
		if _, err := s.departmentRepo.GetByID(*u.DepartmentID); err != nil {
			return uuid.Nil, errors.New("invalid department_id")
		}
	}

	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	id, err := s.userRepo.Create(u)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
