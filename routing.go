package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/RSOI/gateway/model"
	"github.com/RSOI/gateway/ui"
	"github.com/RSOI/gateway/utils"
	"github.com/RSOI/gateway/view"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func sendResponse(ctx *fasthttp.RequestCtx, r ui.Response) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(r.Status)
	utils.LOG(fmt.Sprintf("Sending response. Status: %d", r.Status))

	content, _ := json.Marshal(r)
	ctx.Write(content)
}

func signup(ctx *fasthttp.RequestCtx) {
	utils.LOG(fmt.Sprintf("Request: Signup (%s)", ctx.Path()))
	var r ui.Response

	var NewUser model.User
	var code int
	err := json.Unmarshal(ctx.PostBody(), &NewUser)
	if err != nil {
		utils.LOG(fmt.Sprintf("Broken body. Error: %s", err.Error()))
	} else {
		r.Data, code = uc.Signup(NewUser.Nickname)
	}

	r.Error = ui.CodeToError(code)
	r.Status = code
	sendResponse(ctx, r)
}

func ask(ctx *fasthttp.RequestCtx) {
	utils.LOG(fmt.Sprintf("Request: Ask (%s)", ctx.Path()))
	var r ui.Response

	var NewQuestion model.Question
	var QuestionAuthor *model.User
	var code int

	err := json.Unmarshal(ctx.PostBody(), &NewQuestion)
	if err != nil {
		utils.LOG(fmt.Sprintf("Broken body. Error: %s", err.Error()))
		r.Error = ui.CodeToError(code)
		r.Status = code
		sendResponse(ctx, r)
		return
	}

	err = view.ValidateNewQuestion(NewQuestion)
	if err != nil {
		utils.LOG(fmt.Sprintf("Validation error: %s", err.Error()))
		r.Status, r.Error = ui.ErrToResponse(err)
		sendResponse(ctx, r)
		return
	}

	QuestionAuthor, code = uc.Get(NewQuestion.AuthorID)
	if code != 200 {
		utils.LOG("Can't check author. Aborting.")
		r.Error = ui.CodeToError(code)
		r.Status = code
		sendResponse(ctx, r)
		return
	}

	r.Data, code = qc.Ask(NewQuestion.Title, NewQuestion.Content, QuestionAuthor.ID, QuestionAuthor.Nickname)
	r.Error = ui.CodeToError(code)
	r.Status = code
	sendResponse(ctx, r)
}

func answer(ctx *fasthttp.RequestCtx) {
	utils.LOG(fmt.Sprintf("Request: Answer (%s)", ctx.Path()))
	var r ui.Response

	var NewAnswer model.Answer
	var Question *model.Question
	var AnswerAuthor *model.User
	var code int

	err := json.Unmarshal(ctx.PostBody(), &NewAnswer)
	if err != nil {
		utils.LOG(fmt.Sprintf("Broken body. Error: %s", err.Error()))
		r.Error = ui.CodeToError(code)
		r.Status = code
		sendResponse(ctx, r)
		return
	}

	err = view.ValidateNewAnswer(NewAnswer)
	if err != nil {
		utils.LOG(fmt.Sprintf("Validation error: %s", err.Error()))
		r.Status, r.Error = ui.ErrToResponse(err)
		sendResponse(ctx, r)
		return
	}

	AnswerAuthor, code = uc.Get(NewAnswer.AuthorID)
	if code != 200 {
		utils.LOG("Can't check author. Aborting.")
		r.Error = ui.CodeToError(code)
		r.Status = code
		sendResponse(ctx, r)
		return
	}

	Question, code = qc.Get(NewAnswer.QuestionID)
	if code != 200 {
		utils.LOG("Can't check question. Aborting.")
		r.Error = ui.CodeToError(code)
		r.Status = code
		sendResponse(ctx, r)
		return
	}

	r.Data, code = ac.Answer(NewAnswer.Content, AnswerAuthor.ID, AnswerAuthor.Nickname, Question.ID)
	r.Error = ui.CodeToError(code)
	r.Status = code
	sendResponse(ctx, r)
}

func getUserQuestion(ctx *fasthttp.RequestCtx) {
	utils.LOG(fmt.Sprintf("Request: getUserQuestion (%s)", ctx.Path()))
	var r ui.Response

	var DataResponse model.UserQuestion
	var Question *model.Question
	var Answers *[]model.Answer
	var QuestionAuthor *model.User
	var code int

	qidv := ctx.UserValue("id").(string)
	qid, _ := strconv.Atoi(qidv)
	page := ctx.QueryArgs().Peek("page")
	countOnPage := ctx.QueryArgs().Peek("conp")
	p := -1
	c := -1
	if len(page) > 0 && len(countOnPage) > 0 {
		p, _ = strconv.Atoi(string(page))
		c, _ = strconv.Atoi(string(countOnPage))
	}

	Question, code = qc.Get(qid)
	if code != 200 {
		utils.LOG("Can't get question. Aborting.")
		r.Error = ui.CodeToError(code)
		r.Status = code
		sendResponse(ctx, r)
		return
	}

	QuestionAuthor, code = uc.Get(Question.AuthorID)
	if code != 200 {
		utils.LOG("Can't check author. Aborting.")
		r.Error = ui.CodeToError(code)
		r.Status = code
		sendResponse(ctx, r)
		return
	}

	Answers, code = ac.GetByQuestionID(qid, p, c)
	DataResponse.Question = Question
	DataResponse.Author = QuestionAuthor
	DataResponse.Answers = Answers

	r.Data = DataResponse
	r.Error = ui.CodeToError(code)
	r.Status = code
	sendResponse(ctx, r)
}

func deleteQuestion(ctx *fasthttp.RequestCtx) {
	utils.LOG(fmt.Sprintf("Request: Delete question (%s)", ctx.Path()))
	var r ui.Response
	var code int

	var QuestionToDelete model.ID

	err := json.Unmarshal(ctx.PostBody(), &QuestionToDelete)
	if err != nil {
		utils.LOG(fmt.Sprintf("Broken body. Error: %s", err.Error()))
		r.Error = ui.CodeToError(code)
		r.Status = code
		sendResponse(ctx, r)
		return
	}

	code = qc.Delete(QuestionToDelete.ID)
	if code != 200 {
		utils.LOG("Can't delete question. Aborting.")
		r.Error = ui.CodeToError(code)
		r.Status = code
		sendResponse(ctx, r)
		return
	}

	ac.DeleteByQuestionID(QuestionToDelete.ID)
	if code != 200 {
		utils.LOG("Can't delete answer. But question deleted")
	}

	r.Error = ui.CodeToError(code)
	r.Status = code
	sendResponse(ctx, r)
}

func markAnswerBest(ctx *fasthttp.RequestCtx) {
	utils.LOG(fmt.Sprintf("Request: Mark answer best (%s)", ctx.Path()))
	var r ui.Response

	var AnswerToUpdate model.ID
	var UpdatedAnswer *model.Answer
	var code int

	err := json.Unmarshal(ctx.PostBody(), &AnswerToUpdate)
	if err != nil {
		utils.LOG(fmt.Sprintf("Broken body. Error: %s", err.Error()))
		r.Error = ui.CodeToError(code)
		r.Status = code
		sendResponse(ctx, r)
		return
	}

	UpdatedAnswer, code = ac.MarkAnswerBest(AnswerToUpdate.ID)
	if code != 200 {
		utils.LOG("Can't update answer. Aborting.")
		r.Error = ui.CodeToError(code)
		r.Status = code
		sendResponse(ctx, r)
		return
	}

	_, code = qc.UpdateBest(UpdatedAnswer.QuestionID)
	if code != 200 {
		utils.LOG("Can't mark question. Answer was updated.")
		r.Error = ui.CodeToError(code)
		r.Status = code
		sendResponse(ctx, r)
		return
	}

	r.Data = UpdatedAnswer
	r.Error = ui.CodeToError(code)
	r.Status = code
	sendResponse(ctx, r)
}

func initRoutes() *fasthttprouter.Router {
	utils.LOG("Setup router...")
	router := fasthttprouter.New()
	router.PUT("/signup", signup)                     // store 1 user in user ms
	router.PUT("/ask", ask)                           // store 1 question in question ms (check for user exising)
	router.PUT("/answer", answer)                     // store 1 answer in in answer ms (check for user and question existing)
	router.GET("/question/id:id", getUserQuestion)    // get question, author and answers
	router.DELETE("/question/delete", deleteQuestion) // delete question and all child answers
	router.PATCH("/answer/best", markAnswerBest)      // update answer, question

	return router
}
