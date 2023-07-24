package agent

import (
	"dialogflow-webhook/intents"
	"dialogflow-webhook/structs"
	"dialogflow-webhook/structs/dialogflow"
)

func CreateResponse(r *structs.Request) dialogflow.Response {
	return intents.FuncMap[r.GetIntentName()].(func() dialogflow.Response)()
}
