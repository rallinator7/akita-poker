package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rallinator7/akita-poker/server/checker"
)

type HandController struct {
	HandChecker checker.Checker
}

func NewHandController(c checker.Checker) *HandController {
	hc := HandController{
		HandChecker: c,
	}

	return &hc
}

func (hc *HandController) PostHandCheck(c *gin.Context) {
	hcreq := HandCheckRequest{}

	err := c.ShouldBindJSON(&hcreq)
	if err != nil {
		hc.returnBadRequest("could not bind to json", c)
		return
	}

	name, cards, err := hc.HandChecker.CheckHand(hcreq.Hand)
	if err != nil {
		hc.returnServerError(err, c)
		return
	}

	hcresp := HandCheckResponse{
		Name:  name,
		Cards: cards,
	}

	c.JSON(http.StatusOK, gin.H{
		"handCheckResponse": hcresp,
	})
}

// returns a server error back to the client
func (hc *HandController) returnServerError(err error, c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

// returns a bad request message back to the client
func (hc *HandController) returnBadRequest(message string, c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": message,
	})
}
