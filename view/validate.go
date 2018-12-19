package view

import (
	"github.com/RSOI/gateway/model"
	"github.com/RSOI/gateway/ui"
)

// ValidateNewQuestion returns nil if all the required form values are passed
func ValidateNewQuestion(data model.Question) error {
	if data.Title == "" ||
		data.Content == "" ||
		data.AuthorID == 0 {
		return ui.ErrFieldsRequired
	}
	return nil
}

// ValidateNewAnswer returns nil if all the required form values are passed
func ValidateNewAnswer(data model.Answer) error {
	if data.Content == "" ||
		data.AuthorID == 0 ||
		data.QuestionID == 0 {
		return ui.ErrFieldsRequired
	}
	return nil
}
