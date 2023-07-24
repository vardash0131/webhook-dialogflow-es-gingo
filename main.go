package main

import (
	"dialogflow-webhook/structs"
	"dialogflow-webhook/structs/agent"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", detectIntent)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func detectIntent(c *gin.Context) {

	body := structs.Request{}

	// using BindJson method to serialize body with struct
	err := c.BindJSON(&body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	response := agent.CreateResponse(&body)

	c.SecureJSON(http.StatusOK, response)

}
