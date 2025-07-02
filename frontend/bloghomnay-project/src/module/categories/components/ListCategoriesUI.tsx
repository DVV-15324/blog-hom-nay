import { Categories } from "../module/categories";
import ItemCategories from "./ItemCategoriesUI";

type Props = {
    categories: Categories[];
};

const ListCategories: React.FC<Props> = ({ categories }) => {
    if (!categories.length) {
        return <p className="text-center text-gray-500 mt-10">Không có danh mục nào.</p>;
    }

    return (
        <div className="max-w-6xl mx-auto p-4 grid xl:grid-cols-4 grid-cols-1 md:grid-cols-3 gap-4">
            {categories.map((category) => (
                <ItemCategories key={category.id} categories={category} />
            ))}
        </div>
    );
};

export default ListCategories;
