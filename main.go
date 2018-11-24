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
	log.Println("***request***")
	spew.Dump(request)

	if(request.body.session.new) {
		speechtext = "Welcome to Foodie"
	}
	switch request.Body.Intent.Name {
	default:
		speechtext = "You ended up in defaut"
	}
	return alexa.NewSimpleResponse("Dispatch", speechtext)
}

// Handler is the lambda hander
func Handler(request alexa.Request) (alexa.Response, error) {
		return DispatchIntents(request), nil
	}

func main() {
	lambda.Start(Handler)
}