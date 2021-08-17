# akita-poker

## What is this?

Akita-poker is solution to evaluate poker hands.  There are two pieces that make up this repository, a client written in Typescript and a server written in Go.

## Typescript Client

The client creates a deck of cards (basic queue), shuffles the deck, pulls five cards from the deck, and then sends it to the server to be evaluated.  Before sending to the server, the client logs what the cards are so we know what the hand looks like before it is evaluated.  Because randomness doesn't provide for very fun poker hands, every fifth call the client searches the deck for 4 Aces and sends it to the server for a pretty good poker hand.  When the client recieves the evaluation, it then puts the cards back into the deck and shuffles it.  This process repeats itself until the client is stopped.

## Go Server

This is where most of the work is taking place.  The server has a single endpoint that it listens for hand check requests.  When it receives a hand. it concurrently evaluates the hand for each type of possible outcome, and pushes any eligible poker hand into a priority queue.  When each evaluation is finished, it signals back to the main process it has completed.  When all the evaluations are complete, the main process then pops the priority queue for the best possible poker hand and returns it to the client.

## Running the Services

There is a simple GitHub workflow that tests the server and builds containers for both the client and sever.  In this repository, there is a docker-compose file at the root of the directory.  If you have docker installed on your machine, you can simply run:

```
docker-compose up
```

and you will see the client and server interacting with each other.
