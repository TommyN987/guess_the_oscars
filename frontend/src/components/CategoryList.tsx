import { useCategories } from "../hooks/useCategories";
import { Accordion } from "./Accordion";

function CategoryList() {
    const { categories, loading, error } = useCategories();

    if (loading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <ul className="w-[50%] flex flex-col gap-4 z-50">
            {categories.map((category) => (
                <Accordion key={category.id} category={category} />
            ))}
        </ul>
    );
}

export { CategoryList };
