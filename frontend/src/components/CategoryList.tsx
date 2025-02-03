import { useCategories } from "../hooks/useCategories";
import { NominationsCard } from "./NominationsCard";

function CategoryList() {
    const { categories, loading, error } = useCategories();

    if (loading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <ul className="w-[50%] flex flex-col gap-4">
            <NominationsCard category={categories[1]} />
            {/*categories.map((category) => (
                <NominationsCard category={category} />
            ))*/}
        </ul>
    );
}

export { CategoryList };
