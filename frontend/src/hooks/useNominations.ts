import { useEffect, useState } from "react";
import { Category, Nomination } from "../api/types";
import { fetchNominationsByCategory } from "../api/nominations";

function useNominations(category: Category) {
    const [nominationsForCategory, setNominationsForCategory] = useState<
        Nomination[]
    >([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async (category: Category) => {
            try {
                const nominationsDTO =
                    await fetchNominationsByCategory(category);
                nominationsDTO.sort((a, b) =>
                    a.movie.title.localeCompare(b.movie.title),
                );

                const nominations = nominationsDTO.map(
                    (nomination): Nomination => {
                        return {
                            id: nomination.id,
                            movie: nomination.movie,
                            people: nomination.people,
                            isGuessed: nomination.is_guessed,
                        };
                    },
                );
                setNominationsForCategory(nominations);
            } catch (err) {
                setError((err as Error).message);
            } finally {
                setLoading(false);
            }
        };

        fetchData(category);
    }, [category]);

    return { nominationsForCategory, loading, error };
}

export { useNominations };
