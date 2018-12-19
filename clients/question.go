package clients

import (
	"encoding/json"
	"fmt"

	"github.com/RSOI/gateway/model"
	"github.com/RSOI/gateway/ui"
	"github.com/RSOI/gateway/utils"
)

// QuestionClient instance
type QuestionClient struct {
	connection Connection
	host       string
}

// NewQuestionClient returns instance of question
func NewQuestionClient() *QuestionClient {
	utils.LOG("Setting up question client")

	return &QuestionClient{
		host: "http://localhost:8080",
	}
}

// Ask add new question
func (qc *QuestionClient) Ask(title string, content string, aid int, anickname string) (*model.Question, int) {
	utils.LOG("Asking new question")

	bodyReq := fmt.Sprintf("{\"title\":\"%s\",\"content\":\"%s\",\"author_id\":%d,\"author_nickname\":\"%s\"}", title, content, aid, anickname)
	code, body := qc.connection.PUT(qc.host+"/ask", []byte(bodyReq))
	var response ui.Response
	data := &model.Question{}
	temp := ui.Response{Data: data}
	json.Unmarshal(body, &temp)

	if code != 201 {
		utils.LOG(fmt.Sprintf("Error occured while asking new question: %s", response.Error))
		return nil, code
	}

	return data, code
}

// Get Getting question
func (qc *QuestionClient) Get(id int) (*model.Question, int) {
	utils.LOG("Getting question")

	code, body := qc.connection.GET(fmt.Sprintf("%s/question/id%d", qc.host, id))
	var response ui.Response
	data := &model.Question{}
	temp := ui.Response{Data: data}
	json.Unmarshal(body, &temp)

	if code != 200 {
		utils.LOG(fmt.Sprintf("Error occured while getting question: %s", response.Error))
		return nil, code
	}

	return data, code
}

// Delete question
func (qc *QuestionClient) Delete(id int) int {
	utils.LOG("Deleting question")

	bodyReq := fmt.Sprintf("{\"id\":%d}", id)
	code, body := qc.connection.DELETE(qc.host+"/delete", []byte(bodyReq))
	var response ui.Response
	json.Unmarshal(body, &response)

	if code != 200 {
		utils.LOG(fmt.Sprintf("Error occured while deleting question: %s", response.Error))
	}

	return code
}

// UpdateBest question
func (qc *QuestionClient) UpdateBest(id int) (*model.Question, int) {
	utils.LOG("Updating question as contained best answer")

	bodyReq := fmt.Sprintf("{\"id\":%d, \"has_best\": true}", id)
	code, body := qc.connection.PATCH(qc.host+"/update", []byte(bodyReq))
	var response ui.Response
	data := &model.Question{}
	temp := ui.Response{Data: data}
	json.Unmarshal(body, &temp)

	if code != 200 {
		utils.LOG(fmt.Sprintf("Error occured while updating question: %s", response.Error))
	}

	return data, code
}
