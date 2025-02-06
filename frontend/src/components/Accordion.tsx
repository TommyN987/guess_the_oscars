import { useEffect, useState } from "react";
import { Category, Guess, Nomination } from "../api/types";
import { useNominations } from "../hooks/useNominations";
import { NominationItem } from "./Nomination";
import { useGuess } from "../hooks/useGuess";

type Props = {
    category: Category;
};

function Accordion({ category }: Props) {
    const { nominationsForCategory, loading, error } = useNominations(category);
    const [isOpen, setIsOpen] = useState(false);

    const { submitGuess } = useGuess();
    const [nominations, setNominations] = useState<Nomination[]>([]);
    const [guessSubmitted, setGuessSubmitted] = useState<boolean>(false);

    useEffect(() => {
        setNominations(nominationsForCategory);
    }, [nominationsForCategory]);

    useEffect(() => {
        setGuessSubmitted(nominations.some((nom) => nom.isGuessed));
    }, [nominations]);

    const handleSubmitGuess = async (nominationID: number) => {
        setNominations((prev) =>
            prev.map((nom) =>
                nom.id === nominationID
                    ? { ...nom, isGuessed: true }
                    : { ...nom, isGuessed: false },
            ),
        );

        try {
            const guess: Guess = {
                categoryID: category.id,
                nominationID,
            };
            await submitGuess(guess);
        } catch (error) {
            console.error("Error submitting guess:", error);

            setNominations((prev) =>
                prev.map((nom) =>
                    nom.id === nominationID
                        ? { ...nom, isGuessed: false }
                        : nom,
                ),
            );
        }
    };

    if (loading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <div className="px-5 lg:px-10 py-4 bg-radial from-dark to-zinc-800 to-75% border-4 border-gold rounded-xl">
            <button
                onClick={() => setIsOpen((prev) => !prev)}
                className="flex justify-between w-full"
            >
                <span className="text-[16px] md:text-2xl font-serif font-bold mb-4 flex items-center gap-4 cursor-pointer">
                    {category.name}
                    {guessSubmitted && (
                        <svg
                            className="fill-gold"
                            width="16"
                            height="16"
                            version="1.1"
                            id="Capa_1"
                            xmlns="http://www.w3.org/2000/svg"
                            xmlnsXlink="http://www.w3.org/1999/xlink"
                            viewBox="0 0 45.701 45.7"
                            xmlSpace="preserve"
                        >
                            <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
                            <g
                                id="SVGRepo_tracerCarrier"
                                strokeLinecap="round"
                                strokeLinejoin="round"
                            ></g>
                            <g id="SVGRepo_iconCarrier">
                                {" "}
                                <g>
                                    {" "}
                                    <g>
                                        {" "}
                                        <path d="M20.687,38.332c-2.072,2.072-5.434,2.072-7.505,0L1.554,26.704c-2.072-2.071-2.072-5.433,0-7.504 c2.071-2.072,5.433-2.072,7.505,0l6.928,6.927c0.523,0.522,1.372,0.522,1.896,0L36.642,7.368c2.071-2.072,5.433-2.072,7.505,0 c0.995,0.995,1.554,2.345,1.554,3.752c0,1.407-0.559,2.757-1.554,3.752L20.687,38.332z"></path>{" "}
                                    </g>{" "}
                                </g>{" "}
                            </g>
                        </svg>
                    )}
                </span>
                <svg
                    className="fill-gold shrink-0 ml-8 mt-2 cursor-pointer"
                    width="16"
                    height="16"
                    xmlns="http://www.w3.org/2000/svg"
                >
                    <rect
                        y="7"
                        width="16"
                        height="2"
                        rx="1"
                        className={`transform origin-center transition duration-200 ease-out ${
                            isOpen && "!rotate-180"
                        }`}
                    />
                    <rect
                        y="7"
                        width="16"
                        height="2"
                        rx="1"
                        className={`transform origin-center rotate-90 transition duration-200 ease-out ${
                            isOpen && "!rotate-180"
                        }`}
                    />
                </svg>
            </button>
            <div
                className={`grid overflow-hidden transition-all duration-300 ease-in-out text-slate-600 text-sm ${
                    isOpen
                        ? "grid-rows-[1fr] opacity-100"
                        : "grid-rows-[0fr] opacity-0"
                }`}
            >
                <ul className="flex justify-center flex-wrap gap-4 overflow-hidden">
                    {nominations.map((nomination: Nomination) => (
                        <NominationItem
                            key={nomination.id}
                            nomination={nomination}
                            onSubmitGuess={handleSubmitGuess}
                        />
                    ))}
                </ul>
            </div>
        </div>
    );
}

export { Accordion };
