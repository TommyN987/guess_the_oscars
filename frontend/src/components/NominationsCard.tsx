import { Category, Nomination } from "../api/types";
import { useNominations } from "../hooks/useNominations";
import { NominationItem } from "./Nomination";

type Props = {
    category: Category;
};

function NominationsCard({ category }: Props) {
    const { nominationsForCategory, loading, error } = useNominations(category);

    if (loading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <ul className="flex flex-wrap gap-2">
            {nominationsForCategory.map((nomination: Nomination) => (
                <NominationItem key={nomination.id} nomination={nomination} />
            ))}
        </ul>
    );
}

export { NominationsCard };
