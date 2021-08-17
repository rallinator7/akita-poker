package hand

import (
	"context"
	"fmt"

	"github.com/rallinator7/akita-poker/server/card"
	"github.com/rallinator7/akita-poker/server/checker"
	"go.uber.org/zap"
)

type Server struct {
	HandChecker checker.Checker
	UnimplementedHandServerServer
}

func NewServer(c checker.Checker) *Server {
	s := Server{
		HandChecker: c,
	}

	return &s
}

func (s *Server) CheckHand(ctx context.Context, req *CheckHandRequest) (*CheckHandResponse, error) {
	zap.S().Infof("received hand: %s", req.Hand.String())

	hand := req.GetHand().GetCards()

	cards := []card.Card{}

	for _, c := range hand {
		cards = append(cards, *card.NewCard(card.Suit(c.GetSuit()), card.Face(c.GetFace())))
	}

	name, cards, err := s.HandChecker.CheckHand(cards)
	if err != nil {
		zap.S().Errorf("could not check hand: %s", err)
		return nil, fmt.Errorf("could not check hand: %s", err)
	}

	respCards := []*Card{}

	for _, c := range cards {
		newCard := &Card{
			Face: int64(c.Face),
			Suit: int64(c.Suit),
		}
		respCards = append(respCards, newCard)
	}

	resp := CheckHandResponse{
		Name: name,
		Hand: &Hand{
			Cards: respCards,
		},
	}

	return &resp, nil
}
