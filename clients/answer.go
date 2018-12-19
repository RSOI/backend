package clients

import (
	"encoding/json"
	"fmt"

	"github.com/RSOI/gateway/model"
	"github.com/RSOI/gateway/ui"
	"github.com/RSOI/gateway/utils"
)

// AnswerClient instance
type AnswerClient struct {
	connection Connection
	host       string
}

// NewAnswerClient returns instance of answer
func NewAnswerClient() *AnswerClient {
	utils.LOG("Setting up answer client")

	return &AnswerClient{
		host: "http://localhost:8081",
	}
}

// Answer add new answer
func (ac *AnswerClient) Answer(content string, authorID int, authorNickname string, questionID int) (*model.Answer, int) {
	utils.LOG("Adding new answer")

	bodyReq := fmt.Sprintf("{\"content\":\"%s\",\"author_id\":%d,\"author_nickname\":\"%s\",\"question_id\":%d}", content, authorID, authorNickname, questionID)
	code, body := ac.connection.PUT(ac.host+"/answer", []byte(bodyReq))
	var response ui.Response
	data := &model.Answer{}
	temp := ui.Response{Data: data}
	json.Unmarshal(body, &temp)

	if code != 201 {
		utils.LOG(fmt.Sprintf("Error occured while asking new answer: %s", response.Error))
		return nil, code
	}

	return data, code
}

// GetByQuestionID add new answer
func (ac *AnswerClient) GetByQuestionID(questionID int, page int, countPerPage int) (*[]model.Answer, int) {
	utils.LOG("Geting answers by question id")

	uri := fmt.Sprintf("%s/answers/question%d", ac.host, questionID)
	if page > 0 && countPerPage > 0 {
		uri = fmt.Sprintf("%s?offset=%d&limit=%d", uri, page, countPerPage)
	}
	code, body := ac.connection.GET(uri)
	var response ui.Response
	data := &[]model.Answer{}
	temp := ui.Response{Data: data}
	json.Unmarshal(body, &temp)

	if code != 200 {
		utils.LOG(fmt.Sprintf("Error occured while asking new answer: %s", response.Error))
		return nil, code
	}

	return data, code
}

// DeleteByQuestionID delete answers by question
func (ac *AnswerClient) DeleteByQuestionID(questionID int) int {
	utils.LOG("Deleting answers by question id")

	bodyReq := fmt.Sprintf("{\"question_id\":%d}", questionID)
	code, body := ac.connection.DELETE(ac.host+"/delete", []byte(bodyReq))
	var response ui.Response
	json.Unmarshal(body, &response)

	if code != 200 {
		utils.LOG(fmt.Sprintf("Error occured while deleting answer: %s", response.Error))
	}

	return code
}

// MarkAnswerBest marks answer as best
func (ac *AnswerClient) MarkAnswerBest(id int) (*model.Answer, int) {
	utils.LOG("Mark answer as best")

	bodyReq := fmt.Sprintf("{\"id\":%d}", id)
	code, body := ac.connection.PATCH(ac.host+"/best", []byte(bodyReq))
	var response ui.Response
	data := &model.Answer{}
	temp := ui.Response{Data: data}
	json.Unmarshal(body, &temp)

	if code != 200 {
		utils.LOG(fmt.Sprintf("Error occured while updating answer: %s", response.Error))
	}

	return data, code
}
