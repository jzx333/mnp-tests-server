package repo

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type UserGroupMemberRepo struct {
	db *sqlx.DB
}

func NewUserGroupMemberRepo(db *sqlx.DB) *UserGroupMemberRepo {
	return &UserGroupMemberRepo{db: db}
}

// Add member
func (r *UserGroupMemberRepo) Add(member *dto.UserGroupMember) error {
	_, err := r.db.Exec(
		`INSERT INTO user_group_members (group_id, user_id) VALUES ($1, $2)`,
		member.GroupID, member.UserID,
	)
	return err
}

// Remove member
func (r *UserGroupMemberRepo) Remove(groupID int, userID uuid.UUID) error {
	_, err := r.db.Exec(
		`DELETE FROM user_group_members WHERE group_id=$1 AND user_id=$2`,
		groupID, userID,
	)
	return err
}

// List members by group
func (r *UserGroupMemberRepo) ListByGroup(groupID int) ([]dto.UserGroupMember, error) {
	var members []dto.UserGroupMember
	err := r.db.Select(&members, `SELECT * FROM user_group_members WHERE group_id=$1`, groupID)
	return members, err
}
