package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rallinator7/akita-poker/server/card"
	"github.com/rallinator7/akita-poker/server/checker"
)

//tests http end point for hand checker
func TestHandController_PostHandCheck(t *testing.T) {
	cards := []card.Card{
		{Face: card.Ace, Suit: card.Club},
		{Face: card.Two, Suit: card.Club},
		{Face: card.Three, Suit: card.Club},
		{Face: card.Four, Suit: card.Club},
		{Face: card.Five, Suit: card.Spade},
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	hcreq := HandCheckRequest{
		Hand: cards,
	}

	b, err := json.Marshal(hcreq)
	if err != nil {
		t.Fatalf("could not marshal json: %s", err)
	}

	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(b))
	handChecker := checker.NewHandChecker()
	handController := NewHandController(handChecker)

	handController.PostHandCheck(c)

	var got map[string]map[string]interface{}

	err = json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	resp := got["handCheckResponse"]
	name := resp["name"]

	if name != checker.HighCard.String() {
		t.Fatalf("expected high card got: %s", name)
	}

}
