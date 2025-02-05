type User = {
    id: number;
    name: string;
    email: string;
};
type Movie = {
    id: number;
    title: string;
    additionalInfo?: string;
};

type Person = {
    id: number;
    name: string;
};

type Category = {
    id: number;
    name: string;
};

type Nomination = {
    id: number;
    movie: Movie;
    people: Person[];
    isGuessed: boolean;
};

type NominationDTO = {
    id: number;
    movie: Movie;
    people: Person[];
    is_guessed: boolean;
};

type NominationsForCategory = {
    nominations: Nomination[];
};

type Guess = {
    userID: number;
    nominationID: number;
};

export {
    type User,
    type Movie,
    type Person,
    type Category,
    type Nomination,
    type NominationDTO,
    type NominationsForCategory,
    type Guess,
};
