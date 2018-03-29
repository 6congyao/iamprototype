package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
	"net/http"
	"os"
)

// Binding from JSON
type AuthRequestInput struct {
	// Resource is the resource that access is requested to.
	Resource string `json:"resource" binding:"required"`

	// Action is the action that is requested on the resource.
	Action string `json:"action" binding:"required"`

	// Subejct is the subject that is requesting access.
	Subject string `json:"principal" binding:"required"`
}

type PolicyRequestInput struct {
	Description string `json:"description"`

	Effect string `json:"effect" binding:"required"`

	// Resource is the resource that access is requested to.
	Resources string `json:"resource" binding:"required"`

	// Action is the action that is requested on the resource.
	Actions string `json:"action" binding:"required"`

	// Subejct is the subject that is requesting access.
	Subjects string `json:"principal" binding:"required"`
}

var hostname string
var warden *ladon.Ladon

func main() {

	iamInit()
	router := gin.Default()

	router.GET("/hi", greeting)
	router.POST("/evaluation", auth)
	router.POST("/policy", createPolicy)
	router.GET("/policy", getPolicy)

	router.Run()
}

func iamInit() {
	hostname, _ = os.Hostname()
	warden = &ladon.Ladon{
		Manager: memory.NewMemoryManager(),
	}
	for _, pol := range Polices {
		warden.Manager.Create(pol)
	}
}

func greeting(c *gin.Context) {
	c.String(http.StatusOK, "Greetings! This is from %s \n", hostname)
}

func auth(c *gin.Context) {
	json := &AuthRequestInput{}


	if err := c.ShouldBindJSON(json); err == nil {
		request := &ladon.Request{
			Subject:  json.Subject,
			Action:   json.Action,
			Resource: json.Resource,
		}

		err := warden.IsAllowed(request)
		if err != nil {
			//ret := errors.Cause(err)
			//ret2 := ladon.ErrRequestDenied
			//if ret == ret2 {
			//	fmt.Printf("Type: %T\n", errors.Cause(err))
			//}
			//switch et := errors.Cause(err).(type) {
			//case *ladon.errorWithContext:
			//	// handle specifically
			c.JSON(http.StatusForbidden, gin.H{"status": err.Error(), "from": hostname})
			//default:
			//	// unknown error
			//	c.JSON(http.StatusForbidden, gin.H{"status": et})
			//}

		} else {
			c.JSON(http.StatusOK, gin.H{"status": "Allow", "from": hostname})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func getPolicy(c *gin.Context) {
	polices, err := warden.Manager.GetAll(100, 0)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"polices": polices})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func createPolicy(c *gin.Context) {
	json := &PolicyRequestInput{}
	if err := c.ShouldBindJSON(json); err == nil {

		c.JSON(http.StatusOK, gin.H{"status": "create successfully", "from": hostname})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}