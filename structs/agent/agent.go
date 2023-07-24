package agent

import (
	"dialogflow-webhook/intents"
	"dialogflow-webhook/structs"
)

func CreateResponse(r *structs.Request) string {
	return intents.FuncMap[r.GetIntentName()].(func() string)()
}
