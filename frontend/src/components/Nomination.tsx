import { Nomination } from "../api/types";

type Props = {
    onSubmitGuess: (nominationID: number) => Promise<void>;
    nomination: Nomination;
};

function NominationItem({ nomination, onSubmitGuess }: Props) {
    return (
        <li
            className={`w-[95%] sm:w-[45%] rounded-xl py-3 pl-6 text-white cursor-pointer bg-gradient-to-br from-55% ${nomination.isGuessed ? " from-sunset to-gold" : "from-gold to-sunny"}`}
            onClick={() => onSubmitGuess(nomination.id)}
        >
            <p className="font-bold text-lg">{nomination.movie.title}</p>
            <ul className="flex flex-col">
                {nomination.people?.map((person) => (
                    <p className="text-base font-semibold pl-4" key={person.id}>
                        {person.name}
                    </p>
                ))}
            </ul>
        </li>
    );
}

export { NominationItem };
