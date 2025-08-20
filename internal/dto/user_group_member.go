package dto

type UserGroupMember struct {
	GroupID int    `db:"group_id" json:"group_id"`
	UserID  string `db:"user_id" json:"user_id"`
}
