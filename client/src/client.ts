import { Card } from './card'
import axios from 'axios'

interface CheckHandRequest {
    hand: Card[];
}

export class CheckHandClient {
    url: string;
    port: string;

    constructor(url: string, port: string){
        this.url = url;
        this.port = port;

    }

    // calls the check endpoint from the poker server to check poker hands
    // needs some error handling incase the server doesn't return properly but this works for demo
    public async CheckHand(hand: Card[]): Promise<string> {
        try {
            let reqBody: CheckHandRequest = {
                hand: hand,
            };
            const response = await axios.post(`http://${this.url}:${this.port}/check`, reqBody, { timeout: 300 });
            let chresp = response.data;

            return chresp.handCheckResponse.name;
          } catch (error) {
            console.error(error);
            return "";
          }
    }
}