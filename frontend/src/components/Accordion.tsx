import { useState } from "react";
import { Category, Nomination } from "../api/types";
import { useNominations } from "../hooks/useNominations";
import { NominationItem } from "./Nomination";

type Props = {
    category: Category;
};

function Accordion({ category }: Props) {
    const { nominationsForCategory, loading, error } = useNominations(category);
    const [isOpen, setIsOpen] = useState(false);

    if (loading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <div className="py-2">
            <button
                onClick={() => setIsOpen((prev) => !prev)}
                className="flex justify-between w-full"
            >
                <span className="text-2xl font-serif font-bold mb-4">
                    {category.name}
                </span>
                <svg
                    className="fill-indigo-500 shrink-0 ml-8"
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
                <ul className="flex flex-wrap gap-2 overflow-hidden">
                    {nominationsForCategory.map((nomination: Nomination) => (
                        <NominationItem
                            key={nomination.id}
                            nomination={nomination}
                        />
                    ))}
                </ul>
            </div>
        </div>
    );
}

export { Accordion };
