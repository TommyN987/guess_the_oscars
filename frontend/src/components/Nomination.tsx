import { Nomination } from "../api/types";

type Props = {
    onSubmitGuess: (nominationID: number) => Promise<void>;
    nomination: Nomination;
};

function NominationItem({ nomination, onSubmitGuess }: Props) {
    return (
        <li
            className={`w-[95%] sm:w-[45%] rounded-xl py-3 pl-6 text-white cursor-pointer ${nomination.isGuessed ? "bg-linear-to-br from-sunset to-gold" : "bg-gradient-to-br from-gold from-55% to-sunny"}`}
            onClick={() => onSubmitGuess(nomination.id)}
        >
            <p className="font-bold test-base lg:text-lg">
                {nomination.movie.title}
            </p>
            <ul className="flex flex-col">
                {nomination.people?.map((person) => (
                    <p className="text-base font-semibold" key={person.id}>
                        {person.name}
                    </p>
                ))}
            </ul>
        </li>
    );
}

export { NominationItem };
