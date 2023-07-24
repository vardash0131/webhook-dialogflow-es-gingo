package intents

import (
	"dialogflow-webhook/structs/dialogflow"
)

// "dialogflow-webhook/structs/dialogflow"

// Use map[string]interface{} to pair functions to name
// Could maybe use anonymous functions instead. Might be clean
// in certain cases
var FuncMap = map[string]interface{}{
	"welcome": welcome,
	"getname": GetName,
}

var response = dialogflow.Response{
	FulfillmentText:     "",
	FulfillmentMessages: []interface{}{},
}

func welcome() dialogflow.Response {

	//Examples to test welcome

	response.AddText("This is a text", &response)

	response.AddCard(dialogflow.CardKV{
		Title:       "Hello world!",
		Subtitle:    "This is a subtitle",
		ImageURI:    "https://this.is.an.image.com/image.png",
		Description: "This is a huge description...",
		Buttons: []dialogflow.Button{
			{
				Text:     "This is a button",
				Postback: "Action after clicking",
			},
		},
	}, &response)

	response.AddText("This is another text", &response)

	response.AddQuickReplies([]string{
		"Option 1",
		"Option 2",
	}, &response)

	response.AddOutputContext(dialogflow.OutputContext{
		Name:          "projects/cdmx-poc/agent/sessions/40efe566-29e3-5059-4561-91456b751d96/contexts/verificacion_tramite_requisitos",
		LifespanCount: 1,
		Parameters: map[string]any{
			"fallbacks": 0,
			"query":     "",
		},
	}, &response)

	return response
}

func GetName() dialogflow.Response {
	response.AddText("Hi from Get name", &response)
	return response
}
