import { Category, NominationsForCategory } from "./types";

async function fetchNominationsByCategory(
    category: Category,
): Promise<NominationsForCategory> {
    const resp = await fetch(
        `${import.meta.env.VITE_API_BASE_URL}/categories/${category.id}/nominations`,
    );
    if (!resp.ok) {
        throw new Error(
            `Failed to fetch nominations for category ${category.name}`,
        );
    }

    return resp.json();
}

export { fetchNominationsByCategory };
