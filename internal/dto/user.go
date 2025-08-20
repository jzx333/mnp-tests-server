package dto

import "github.com/google/uuid"

type User struct {
	ID            uuid.UUID  `db:"id" json:"id"`
	FullName      string     `db:"full_name" json:"full_name"`
	RoleID        int        `db:"role_id" json:"role_id"`
	CuratorID     *uuid.UUID `db:"curator_id" json:"curator_id"`
	UnitID        *int       `db:"unit_id" json:"unit_id"`
	DepartmentID  *int       `db:"department_id" json:"department_id"`
	PracticeStart *string    `db:"practice_start" json:"practice_start"`
	PracticeEnd   *string    `db:"practice_end" json:"practice_end"`
}
