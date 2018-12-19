// QUESTION SERVICE

package main

import (
	"fmt"
	"os"

	"github.com/RSOI/gateway/clients"
	"github.com/RSOI/gateway/utils"
	"github.com/valyala/fasthttp"
)

var (
	uc *clients.UserClient
	qc *clients.QuestionClient
	ac *clients.AnswerClient
)

// PORT application port
const PORT = 8083

// Init client services
func Init() {
	utils.LOG("Initialize clients...")
	uc = clients.NewUserClient()
	qc = clients.NewQuestionClient()
	ac = clients.NewAnswerClient()
}

func main() {
	if len(os.Args) > 1 {
		utils.DEBUG = os.Args[1] == "debug"
	}
	utils.LOG("Launched in debug mode...")
	utils.LOG(fmt.Sprintf("Gateway service is starting on localhost: %d", PORT))

	Init()
	fasthttp.ListenAndServe(fmt.Sprintf(":%d", PORT), initRoutes().Handler)
}
