import { useNavigate } from "react-router-dom";
import { PostResponse } from "../../posts/models/post";
import { Pencil, Trash2 } from "lucide-react";

type Props = {
    post: PostResponse;
};

const MyItemPost: React.FC<Props> = ({ post }) => {
    const navigate = useNavigate();

    const handleClick = () => {
        navigate(`/post/${post.id}`);
    };

    const handleUpdatePost = (e: React.MouseEvent) => {
        e.stopPropagation();
        navigate(`/my_post/update_post/${post.id}`);
    };

    const handleDeletePost = (e: React.MouseEvent) => {
        e.stopPropagation();
        // TODO: call delete API or open confirm modal
        alert(`Xóa post với id: ${post.id}`);
    };

    return (
        <div
            className="relative rounded-2xl shadow-md p-6 bg-white mb-6 border border-gray-200 hover:shadow-lg transition cursor-pointer group"
            onClick={handleClick}
        >
            {/* Buttons: Edit & Delete */}
            <div className="absolute top-3 right-3 flex space-x-2 ">
                <button
                    className="text-gray-500 hover:text-blue-600"
                    onClick={handleUpdatePost}
                    title="Chỉnh sửa"
                >
                    <Pencil size={20} />
                </button>
                <button
                    className="text-gray-500 hover:text-red-600"
                    onClick={handleDeletePost}
                    title="Xoá bài viết"
                >
                    <Trash2 size={20} />
                </button>
            </div>

            <h2 className="text-xl font-bold text-gray-800 mb-2 break-words">{post.title}</h2>
            <p className="text-sm text-gray-700 mb-3 break-words">{post.description}</p>

            <div className="flex flex-wrap gap-2 mt-2">
                {post.tags.map((tag) => (
                    <span
                        key={tag.id}
                        className="text-xs bg-blue-100 text-blue-700 px-3 py-1 rounded-full"
                    >
                        #{tag.name}
                    </span>
                ))}
            </div>

            <div className="text-right text-xs text-gray-400 mt-4">
                Cập nhật lúc: {new Date(post.updated_at).toLocaleString()}
            </div>
        </div>
    );
};

export default MyItemPost;
