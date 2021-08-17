import { CardQueue } from "./cardQueue";
import { Card, Faces, Suits } from "./card";

export class Deck{
    totalSize: number;
    currentSize: number;
    cards: CardQueue;

    constructor(){
        this.totalSize = 0;
        this.currentSize = 0;
        this.cards = new CardQueue();

        //creates a standard deck of 52 cards
        for (let [fk, fv] of Faces) {
            for (let [sk, sv] of Suits) {
                let card = new Card(sk, sv, fk, fv);
                this.cards.Enqueue(card);
                this.totalSize++;
                this.currentSize++;
            }
        }
        this.Shuffle();
    }

    // shuffles the deck of cards
    public Shuffle(){
            var currentIndex = this.cards.queue.length,  randomIndex;
          
            // While there remain elements to shuffle...
            while (currentIndex != 0) {
          
              // Pick a remaining element...
              randomIndex = Math.floor(Math.random() * currentIndex);
              currentIndex--;
          
              // And swap it with the current element.
              [this.cards.queue[currentIndex], this.cards.queue[randomIndex]] = [
                this.cards.queue[randomIndex], this.cards.queue[currentIndex]];
            }
    }

    // deals out size amount of cards
    public DealHand(size: number): Card[]{
        let respCards: Card[] = [];

        for(let i = 0; i < size; i++ ){
            let card: Card = this.cards.Dequeue();
            this.currentSize--;
            respCards.push(card);
        }

        return respCards;
    }

    // deals out size amount of cards
    public ReturnCards(cards: Card[]){

        for(let card of cards){
            this.cards.Enqueue(card);
            this.currentSize++;
        }
    }

    // deals a 4 of a kind of Aces from the deck
    public LuckyHand(): Card[]{
        let aces: Card[] = this.cards.queue.filter(card => card.faceString === "Ace");
        aces.forEach(ace => this.cards.queue.splice(this.cards.queue.findIndex(card => card.face === ace.face),1));

        let hand: Card[] = aces;
        let randCard = this.cards.Dequeue();
        hand.push(randCard);

        this.currentSize -= 5;
        return hand;
    }
}