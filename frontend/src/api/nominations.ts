import { Category, NominationDTO } from "./types";

async function fetchNominationsByCategory(
    category: Category,
): Promise<NominationDTO[]> {
    const resp = await fetch(
        `${import.meta.env.VITE_API_BASE_URL}/p/categories/${category.id}/nominations`,
        {
            headers: { "Content-Type": "application/json" },
            credentials: "include",
        },
    );
    if (!resp.ok) {
        throw new Error(
            `Failed to fetch nominations for category ${category.name}`,
        );
    }

    const data = await resp.json();

    return data.nominations;
}

export { fetchNominationsByCategory };
