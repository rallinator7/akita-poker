import { Card } from "./card";

export class CardQueue{
    queue: Card[];

    constructor(){
        this.queue = [];
    }

    // adds cards to the queue
    public Enqueue(card: Card){
        this.queue.push(card)
    }

    //removes the first added card to the queue
    public Dequeue(): Card {
        let card: Card = this.queue.shift()!
        return card
    }
}