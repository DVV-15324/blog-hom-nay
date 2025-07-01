import React from "react";
import { useNavigate } from "react-router-dom";
import { Categories } from "../module/categories";

type Props = {
    categories: Categories;
};

const ItemCategories: React.FC<Props> = ({ categories }) => {
    const navigate = useNavigate();

    const handleClick = () => {
        navigate(`/post/categories/${categories.id}`);
    };

    return (
        <div
            className="rounded-2xl cursor-pointer overflow-hidden shadow-md border border-gray-200 bg-white transition hover:shadow-xl hover:scale-[1.02] duration-300"
            onClick={handleClick}
        >
            {/* Hình ảnh */}
            <img
                src={categories.img}
                alt={categories.name}
                className="w-full h-48 object-cover"
            />

            {/* Nội dung */}
            <div className="p-4 space-y-2">
                <h2 className="text-xl font-semibold text-gray-800">{categories.name}</h2>
                <p className="text-gray-600 text-sm line-clamp-5">{categories.description}</p>
            </div>
        </div>
    );
};

export default ItemCategories;
