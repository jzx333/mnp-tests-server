package dto

import "github.com/google/uuid"

type UserGroupMember struct {
	GroupID int       `db:"group_id" json:"group_id"`
	UserID  uuid.UUID `db:"user_id" json:"user_id"`
}
