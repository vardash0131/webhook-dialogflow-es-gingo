package intents

// Use map[string]interface{} to pair functions to name
// Could maybe use anonymous functions instead. Might be clean
// in certain cases
var FuncMap = map[string]interface{}{
	"welcome": welcome,
	"getname": GetName,
}

func welcome() string {
	return "hi there"
}
func GetName() string {
	return "Hi from Get name"
}
