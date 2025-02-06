import { useEffect, useState } from "react";
import { Category, Guess, Nomination } from "../api/types";
import { useNominations } from "../hooks/useNominations";
import { NominationItem } from "./Nomination";
import { useGuess } from "../hooks/useGuess";

type Props = {
    category: Category;
};

function NominationsCard({ category }: Props) {
    const { nominationsForCategory, loading, error } = useNominations(category);
    const { submitGuess } = useGuess();

    const [nominations, setNominations] = useState<Nomination[]>([]);

    useEffect(() => {
        setNominations(nominationsForCategory);
    }, [nominationsForCategory]);

    const handleSubmitGuess = async (nominationID: number) => {
        setNominations((prev) =>
            prev.map((nom) =>
                nom.id === nominationID ? { ...nom, isGuessed: true } : nom,
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
        <ul className="flex flex-wrap gap-2">
            {nominations.map((nomination: Nomination) => (
                <NominationItem
                    key={nomination.id}
                    nomination={nomination}
                    onSubmitGuess={handleSubmitGuess}
                />
            ))}
        </ul>
    );
}

export { NominationsCard };
