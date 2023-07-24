package structs

import (
	"encoding/json"
)

// Example we need to struct all the  dialogflow fullfilment
type Request struct {
	ResponseId                  string                            `json:"responseId"`
	QueryResult                 *QueryResultstruct                `json:"queryResult" bson:",omitempty"`
	OriginalDetectIntentRequest *OriginalDetectIntentRequesstruct `json:"originalDetectIntentRequest" bson:",omitempty"`
	Session                     string                            `json:"session"`
}
type QueryResultstruct struct {
	QueryText                 string            `json:"queryText" bson:",omitempty"`
	Parameters                interface{}       `json:"parameters" bson:",omitempty"`
	AllRequiredParamsPresent  bool              `json:"allRequiredParamsPresent" bson:",omitempty"`
	OutputContexts            []*OutputContexts `json:"outputContexts" bson:",omitempty"`
	Intent                    Intent            `json:"intent"`
	IntentDetectionConfidence float32           `json:"intentDetectionConfidence" bson:",omitempty"`
	LanguageCode              string            `json:"languageCode" bson:",omitempty"`
	SentimentAnalysisResult   interface{}       `json:"sentimentAnalysisResult" bson:",omitempty"`
}
type Intent struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}
type OriginalDetectIntentRequesstruct struct {
	Source  string      `json:"source"`
	Payload interface{} `json:"payload" bson:",omitempty"`
}
type OutputContexts struct {
	Name          string          `json:"name"`
	LifespanCount *int            `json:"lifespanCount" binding:"exists"`
	Parameters    json.RawMessage `json:"parameters" bson:",omitempty"`
}

func (R *Request) GetIntentName() string {
	return R.QueryResult.Intent.DisplayName
}
