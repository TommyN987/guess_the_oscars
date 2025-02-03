import { useEffect, useState } from "react";
import { Category, NominationsForCategory } from "../api/types";
import { fetchNominationsByCategory } from "../api/nominations";

function useNominations(category: Category) {
    const [nominationsForCategory, setNominationsForCategory] =
        useState<NominationsForCategory | null>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async (category: Category) => {
            try {
                const data = await fetchNominationsByCategory(category);
                data.nominations.sort((a, b) =>
                    a.movie.title.localeCompare(b.movie.title),
                );
                setNominationsForCategory(data);
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
