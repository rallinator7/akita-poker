package hand

import (
	context "context"
	"log"
	"net"
	"testing"

	"github.com/rallinator7/akita-poker/server/card"
	"github.com/rallinator7/akita-poker/server/checker"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

// init function starts before the test call
func init() {
	//implement grpc server over memory is cleaner than assigning a port for test
	lis = bufconn.Listen(bufSize)

	check := checker.NewHandChecker()
	handServer := NewServer(check)
	grpcServer := grpc.NewServer()

	RegisterHandServerServer(grpcServer, handServer)

	// kicks off concurrent grpc server
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

// tests grpc server general connectivity as well as CheckHand grpc
func TestServer_CheckHand(t *testing.T) {
	cards := []card.Card{
		{Face: card.Ace, Suit: card.Spade},
		{Face: card.Queen, Suit: card.Spade},
		{Face: card.King, Suit: card.Spade},
		{Face: card.Jack, Suit: card.Spade},
		{Face: card.Ten, Suit: card.Spade},
	}

	reqCards := []*Card{}
	for _, c := range cards {
		newCard := Card{
			Face: int64(c.Face),
			Suit: int64(c.Suit),
		}

		reqCards = append(reqCards, &newCard)
	}

	// create client
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	reqHand := Hand{
		Cards: reqCards,
	}

	// run CheckHand
	client := NewHandServerClient(conn)
	resp, err := client.CheckHand(ctx, &CheckHandRequest{Hand: &reqHand})
	if err != nil {
		t.Fatalf("CheckHand failed: %v", err)
	}

	if resp.Name != checker.RoyalFlush.String() {
		t.Fatalf("Expected a Royal Flush but got %s", resp.Name)
	}
}
