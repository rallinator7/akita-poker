
export const Faces = new Map([
    ["Two", 0],
    ["Three", 1],
    ["Four", 2],
    ["Five", 3],
    ["Six", 4],
    ["Seven", 5],
    ["Eight", 6],
    ["Nine", 7],
    ["Ten", 8],
    ["Jack", 9],
    ["Queen", 10],
    ["King", 11],
    ["Ace", 12],
]); 

export const Suits = new Map([
    ["Heart", 0],
    ["Diamond", 1],
    ["Spade", 2],
    ["Club", 3],
]); 

export class Card{
    faceString: string;
    face: number;
    suitString: string;
    suit: number;


    constructor(suitString: string, suit: number, faceString: string, face: number){
        this.suit = suit,
        this.suitString = suitString;
        this.face = face;
        this.faceString = faceString;
    }

    public String(): string {
        return `${this.faceString} of ${this.suitString}s`
    }
}

