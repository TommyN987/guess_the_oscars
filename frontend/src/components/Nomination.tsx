import { Nomination } from "../api/types";
import { useGuess } from "../hooks/useGuess";

type Props = {
    nomination: Nomination;
};

function NominationItem({ nomination }: Props) {
    const { submitGuess } = useGuess(nomination.id);

    if (nomination.id === 2) {
        console.log(nomination.isGuessed);
    }

    return (
        <li
            className="w-[30%] rounded-xl p-3 bg-radial from-zinc-600 to-zinc-900 to-75%"
            onClick={submitGuess}
        >
            <p className="font-bold">{nomination.movie.title}</p>
            <ul className="flex gap-1">
                {nomination.people?.map((person) => (
                    <p key={person.id}>{`${person.name} &`}</p>
                ))}
            </ul>
            {nomination.isGuessed && <p>Guessed!</p>}
        </li>
    );
}

export { NominationItem };
