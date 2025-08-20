package repo

import (
	"github.com/jmoiron/sqlx"
	"mnp-tests-server/internal/dto"
)

type UserGroupMemberRepo struct {
	db *sqlx.DB
}

func NewUserGroupMemberRepo(db *sqlx.DB) *UserGroupMemberRepo {
	return &UserGroupMemberRepo{db: db}
}

func (r *UserGroupMemberRepo) Add(member *dto.UserGroupMember) error {
	_, err := r.db.Exec(
		"INSERT INTO user_group_members (group_id, user_id) VALUES ($1, $2)",
		member.GroupID, member.UserID,
	)
	return err
}

func (r *UserGroupMemberRepo) Remove(member *dto.UserGroupMember) error {
	_, err := r.db.Exec(
		"DELETE FROM user_group_members WHERE group_id=$1 AND user_id=$2",
		member.GroupID, member.UserID,
	)
	return err
}

func (r *UserGroupMemberRepo) GetByGroup(groupID int) ([]*dto.UserGroupMember, error) {
	var members []*dto.UserGroupMember
	err := r.db.Select(&members,
		"SELECT group_id, user_id FROM user_group_members WHERE group_id=$1", groupID)
	return members, err
}

func (r *UserGroupMemberRepo) GetByUser(userID string) ([]*dto.UserGroupMember, error) {
	var members []*dto.UserGroupMember
	err := r.db.Select(&members,
		"SELECT group_id, user_id FROM user_group_members WHERE user_id=$1", userID)
	return members, err
}
