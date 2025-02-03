import { Nomination } from "../api/types";

type Props = {
    nomination: Nomination;
};

function NominationItem({ nomination }: Props) {
    return (
        <li className="w-[30%] rounded-xl p-3 bg-radial from-zinc-600 to-zinc-900 to-75%">
            <p className="font-bold">{nomination.movie.title}</p>
            <ul className="flex gap-1">
                {nomination.people?.map((person) => (
                    <p>{`${person.name} &`}</p>
                ))}
            </ul>
        </li>
    );
}

export { NominationItem };
