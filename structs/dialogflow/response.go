package dialogflow

// Structures

type Response struct {
	FulfillmentText     string          `json:"fulfillmentText"`
	FulfillmentMessages []interface{}   `json:"fulfillmentMessages"`
	OutputContexts      []OutputContext `json:"outputContexts"`
}

type OutputContext struct {
	Name          string         `json:"name"`
	LifespanCount int32          `json:"lifespanCount" binding:"exists"`
	Parameters    map[string]any `json:"parameters" bson:",omitempty"`
}

type Text struct {
	Text TextKV `json:"text"`
}

type TextKV struct {
	Text []string `json:"text"`
}

type Button struct {
	Text     string `json:"text"`
	Postback string `json:"postback"`
}

type Card struct {
	Card CardKV `json:"card"`
}
type CardKV struct {
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Description string   `json:"description"`
	ImageURI    string   `json:"imageUri"`
	Buttons     []Button `json:"buttons"`
}

type QuickReplies struct {
	QuickReplies QuickRepliesKV `json:"quickReplies"`
}

type QuickRepliesKV struct {
	QuickReplies []string `json:"quickReplies"`
}

// Functions

func (r Response) AddText(content string, res *Response) {
	res.FulfillmentMessages = append(res.FulfillmentMessages, Text{
		Text: TextKV{
			Text: []string{
				content,
			},
		},
	})
}

func (r Response) AddCard(content CardKV, res *Response) {
	res.FulfillmentMessages = append(res.FulfillmentMessages, Card{
		Card: content,
	})
}

func (r Response) AddQuickReplies(content []string, res *Response) {
	res.FulfillmentMessages = append(res.FulfillmentMessages, QuickReplies{
		QuickReplies: QuickRepliesKV{
			QuickReplies: content,
		},
	})
}

func (r Response) AddOutputContext(outputContext OutputContext, res *Response) {
	res.OutputContexts = append(res.OutputContexts, outputContext)
}
