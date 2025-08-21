package service

import (
	"errors"

	"github.com/google/uuid"
	"mnp-tests-server/internal/dto"
	"mnp-tests-server/internal/repo"
)

type UserGroupService struct {
	groupRepo  *repo.UserGroupRepo
	memberRepo *repo.UserGroupMemberRepo
	userRepo   *repo.UserRepo
}

func NewUserGroupService(
	groupRepo *repo.UserGroupRepo,
	memberRepo *repo.UserGroupMemberRepo,
	userRepo *repo.UserRepo,
) *UserGroupService {
	return &UserGroupService{
		groupRepo:  groupRepo,
		memberRepo: memberRepo,
		userRepo:   userRepo,
	}
}

func (s *UserGroupService) CreateGroup(curatorID uuid.UUID, name string, userIDs []uuid.UUID) (int, error) {
	students, err := s.userRepo.GetByCuratorID(curatorID)
	if err != nil {
		return 0, errors.New("cannot fetch users")
	}

	validUserIDs := map[uuid.UUID]struct{}{}
	for _, u := range students {
		validUserIDs[u.ID] = struct{}{}
	}

	for _, id := range userIDs {
		if _, ok := validUserIDs[id]; !ok {
			return 0, errors.New("some users are not under this curator")
		}
	}

	group := &dto.UserGroup{Name: name}
	groupID, err := s.groupRepo.Create(group)
	if err != nil {
		return 0, err
	}

	for _, id := range userIDs {
		_ = s.memberRepo.Add(&dto.UserGroupMember{
			GroupID: groupID,
			UserID:  id,
		})
	}

	return groupID, nil
}

func (s *UserGroupService) ListGroupsByCurator(curatorID uuid.UUID) ([]dto.UserGroup, error) {
	groups, err := s.groupRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var filtered []dto.UserGroup
	for _, g := range groups {
		members, _ := s.memberRepo.ListByGroup(g.ID)
		for _, m := range members {
			user, _ := s.userRepo.GetByID(m.UserID)
			if user.CuratorID != nil && *user.CuratorID == curatorID {
				filtered = append(filtered, g)
				break
			}
		}
	}

	return filtered, nil
}
