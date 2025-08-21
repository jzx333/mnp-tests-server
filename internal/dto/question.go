package dto

type Question struct {
	ID           int      `db:"id" json:"id"`
	PoolID       *int     `db:"pool_id" json:"pool_id,omitempty"`
	Text         string   `db:"text" json:"text"`
	QuestionType string   `db:"question_type" json:"question_type"`
	Score        int      `db:"score" json:"score"`
	Position     int      `db:"position" json:"position"`
	MediaURL     *string  `db:"media_url" json:"media_url,omitempty"`
	Options      []Option `json:"options,omitempty"`
}
