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
};

type NominationsForCategory = {
    category: Category;
    nominations: Nomination[];
};

export {
    type User,
    type Movie,
    type Person,
    type Category,
    type Nomination,
    type NominationsForCategory,
};
