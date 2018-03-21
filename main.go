package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
	"iamproto/polices"
	"net/http"
)

// Binding from JSON
type Input struct {
	Description string `form:"name" json:"name"`
	Name        string `form:"name" json:"name" binding:"required"`
	Age         int    `form:"age" json:"age" binding:"required"`
	City        string `form:"city" json:"city" binding:"required"`
}

type RequestInput struct {
	// Resource is the resource that access is requested to.
	Resource string `json:"Resource" binding:"required"`

	// Action is the action that is requested on the resource.
	Action string `json:"Action" binding:"required"`

	// Subejct is the subject that is requesting access.
	Subject string `json:"Principal" binding:"required"`
}

func main() {
	router := gin.Default()

	router.GET("/hi", getting)
	router.POST("/request", request)

	router.Run()
}

func getting(c *gin.Context) {
	c.String(http.StatusOK, "Hello Lucas")
}

func request(c *gin.Context) {
	json := &RequestInput{}
	warden := &ladon.Ladon{
		Manager: memory.NewMemoryManager(),
	}

	if err := c.ShouldBindJSON(json); err == nil {
		request := &ladon.Request{
			Subject:  json.Subject,
			Action:   json.Action,
			Resource: json.Resource,
		}
		for _, pol := range polices.Polices {
			warden.Manager.Create(pol)
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
			c.JSON(http.StatusForbidden, gin.H{"status": err.Error()})
			//default:
			//	// unknown error
			//	c.JSON(http.StatusForbidden, gin.H{"status": et})
			//}

		} else {
			c.JSON(http.StatusOK, gin.H{"status": "Allow"})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
