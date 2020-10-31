package model

import (
	"time"
)

const (
	PRODUCT_SORT_NEWEST  = "newest"
	PRODUCT_SORT_POPULAR = "popular"
)

type Question struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Text      string    `json:"text"`
	Answer    string    `json:"answer"`
	Upvotes   int       `json:"upvotes"`
}

// Validators
type QuestionValidator struct {
	Text string `json:"text" binding:"required"`
}

type QuestionUpvoteValidator struct {
	QuestionID int64  `json:"question_id" binding:"required"`
}

type AnswerValidator struct {
	QuestionID int64  `json:"question_id" binding:"required"`
	Answer     string `json:"answer" binding:"required"`
}
