package clients

import (
	"encoding/json"
	"fmt"

	"github.com/RSOI/gateway/model"
	"github.com/RSOI/gateway/ui"
	"github.com/RSOI/gateway/utils"
)

// UserClient instance
type UserClient struct {
	connection Connection
	host       string
}

// NewUserClient returns instance of user
func NewUserClient() *UserClient {
	utils.LOG("Setting up user client")

	return &UserClient{
		host: "http://localhost:8082",
	}
}

// Signup add new user
func (uc *UserClient) Signup(nickname string) (*model.User, int) {
	utils.LOG("Sign up new user")

	code, body := uc.connection.PUT(uc.host+"/signup", []byte("{\"nickname\":\""+nickname+"\"}"))

	var response ui.Response
	data := &model.User{}
	temp := ui.Response{Data: data}
	json.Unmarshal(body, &temp)

	if code != 201 {
		utils.LOG(fmt.Sprintf("Error occured while signing up: %s", response.Error))
		return nil, code
	}

	return data, code
}

// Rating update user rating
func (uc *UserClient) Rating(id, diff int) (*model.User, int) {
	utils.LOG("Updating user rating")

	bodyReq := fmt.Sprintf("{\"id\":%d,\"rating\":%d}", id, diff)
	code, body := uc.connection.PATCH(uc.host+"/update", []byte(bodyReq))

	var response ui.Response
	data := &model.User{}
	temp := ui.Response{Data: data}
	json.Unmarshal(body, &temp)

	if code != 200 {
		utils.LOG(fmt.Sprintf("Error occured while updating user up: %s", response.Error))
		return nil, code
	}

	return data, code
}

// Get user stat
func (uc *UserClient) Get(id int) (*model.User, int) {
	utils.LOG("Getting user")

	code, body := uc.connection.GET(fmt.Sprintf("%s/user/id%d", uc.host, id))

	var response ui.Response
	data := &model.User{}
	temp := ui.Response{Data: data}
	json.Unmarshal(body, &temp)

	if code != 200 {
		utils.LOG(fmt.Sprintf("Error occured while getting user: %s", response.Error))
		return nil, code
	}

	return data, code
}
