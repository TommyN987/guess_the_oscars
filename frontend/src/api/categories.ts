import { Category } from "./types";

async function fetchCategories(): Promise<Category[]> {
    const resp = await fetch(
        `${import.meta.env.VITE_API_BASE_URL}/p/categories`,
        {
            headers: { "Content-Type": "application/json" },
            credentials: "include",
        },
    );
    if (!resp.ok) {
        throw new Error("Failed to fetch categories");
    }

    return resp.json();
}

export { fetchCategories };
