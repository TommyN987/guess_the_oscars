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
        <div className="bg-linear-to-br from-yellow-900 to-gold to-95% p-6 rounded-lg border-4 shadow-lg hover:shadow-xl transition transform">
            <h3 className="text-2xl font-serif font-bold mb-4">
                {nominationsForCategory?.category.name}
            </h3>
            <ul className="flex flex-wrap gap-2">
                {nominationsForCategory?.nominations.map(
                    (nomination: Nomination) => (
                        <NominationItem
                            key={nomination.id}
                            nomination={nomination}
                        />
                    ),
                )}
            </ul>
        </div>
    );
}

export { NominationsCard };
