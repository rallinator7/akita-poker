
import { CheckHandClient } from "./client";
import { Deck } from "./deck";

// creates a deck of cards and runs a forever loop dealing out hands and requesting a hand check
// every 5 rounds a lucky hand is dealt because random poker hands aren't very fun for demonstration
async function main(): Promise<void> {
    //get env variables
    let serverUrl: string = process.env.SERVER_URL!;
    let serverPort: string = process.env.SERVER_PORT!;

    if (serverUrl == null || serverPort == null ){
        throw new Error('env variables SERVER_URL and SERVER_PORT must be set');
    }

    // create deck and client for server
    let deck: Deck = new Deck();
    let client: CheckHandClient = new CheckHandClient(serverUrl, serverPort);
    let count: number =0;

    // loop and evaluate
    while (true) {
        deck.Shuffle();
        let cards = [];
        if(count % 5 == 0){
            cards = deck.LuckyHand();
        }
        else {
            cards = deck.DealHand(5);
        }

        console.log("Your Cards:");
        for (var card of cards){
            console.log(`- ${card.String()}`);
    
        }
    
        let handType: string = await client.CheckHand(cards);
    
        console.log(`\nYour hand type is: ${handType}`);

        deck.ReturnCards(cards);
        await sleep(3000);
        count++
    }   
}

// sleep function so we aren't doing a million requests
function sleep(ms: number) {
    return new Promise(resolve => setTimeout(resolve, ms));
}
  

main()