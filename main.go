package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"Foodie/alexa-go"
	"github.com/davecgh/go-spew/spew"
	"log"
)

// DispatchIntents dispatches each intent to the right handler
func DispatchIntents(request alexa.Request) alexa.Response {
	//var response alexa.Response
	var speechtext string
	var ssmltext string

	log.Println("***request***")
	spew.Dump(request)

	// Need new NewSimpleResponse for the SSML responce **** 
	if(request.Body.Type == "LaunchRequest") {
		ssmltext = "<speak><prosody pitch=\"high\">Launch Request Yaaa!</prosody></speak>"
		speechtext = ""
	}
	if(request.Session.New) {
		speechtext = speechtext + " Welcome to Foodie"
	}
	//switch request.Body.Intent.Name {
	//default:
	//	speechtext = "You ended up in defaut"
	//}
	return alexa.NewSimpleResponse("Dispatch", speechtext, ssmltext)
}

// Handler is the lambda hander
func Handler(request alexa.Request) (alexa.Response, error) {
		return DispatchIntents(request), nil
	}

func main() {
	lambda.Start(Handler)
}