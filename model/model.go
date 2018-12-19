package model

import "time"

// User interface
type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Rating   int    `json:"rating"`
}

// Question interface
type Question struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	AuthorID       int       `json:"author_id"`
	AuthorNickname string    `json:"author_nickname"`
	HasBest        bool      `json:"has_best"`
	Created        time.Time `json:"created"`
}

// Answer interface
type Answer struct {
	ID             int       `json:"id"`
	QuestionID     int       `json:"question_id"`
	Content        string    `json:"content"`
	AuthorID       int       `json:"author_id"`
	AuthorNickname string    `json:"author_nickname"`
	IsBest         bool      `json:"is_best"`
	Created        time.Time `json:"created"`
}

// UserQuestion question response
type UserQuestion struct {
	Question *Question `json:"question"`
	Author   *User     `json:"author"`
	Answers  *[]Answer `json:"answers"`
}

// ID (!)
type ID struct {
	ID int `json:"id"`
}
