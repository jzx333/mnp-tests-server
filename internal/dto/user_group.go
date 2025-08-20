package dto

type UserGroup struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
