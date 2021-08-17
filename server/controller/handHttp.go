package controller

import "github.com/rallinator7/akita-poker/server/card"

type HandCheckRequest struct {
	Hand []card.Card `json:"hand"`
}

type HandCheckResponse struct {
	Name  string      `json:"name"`
	Cards []card.Card `json:"cards"`
}
