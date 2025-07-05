import { useNavigate } from "react-router-dom";
import { PostResponse } from "../../posts/models/post";


type Props = {
    post: PostResponse;
};

const OverViewItemPost: React.FC<Props> = ({ post }) => {
    const navigate = useNavigate();
    const handleClick = () => {
        navigate(`/post/${post.slug}`);
    };
    const handleTagClick = (e: React.MouseEvent, tagName: string) => {
        e.stopPropagation();
        navigate(`/search?q=${encodeURIComponent(`[${tagName}]`)}`);
    };
    return (
        <div className="container mx-auto xl:px-4 flex justify-center items-center">
            <div
                className="relative w-4xl  m-2 rounded-2xl shadow-md p-4 bg-white border border-gray-200 hover:shadow-lg transition cursor-pointer group flex flex-col justify-between"
                onClick={handleClick}
            >

                <h2 className="text-xl font-semibold mb-2  text-black">{post.title}</h2>
                <p className="text-sm mb-2 text-black line-clamp-3 mb-2 break-words">{post.description}</p>

                <div className="flex flex-wrap gap-2 mt-2">
                    {post.tags.map((tag) => (
                        <span
                            key={tag.id}
                            className="text-sm bg-blue-100 text-blue-700 px-3 py-1 rounded-full"
                            onClick={(e) => handleTagClick(e, tag.name)}
                        >
                            #{tag.name}
                        </span>
                    ))}
                </div>

                <div className="text-right text-xs text-gray-400 mt-4">
                    Cập nhật lúc: {new Date(post.updated_at).toLocaleString()}
                </div>
            </div >
        </div>
    );
};

export default OverViewItemPost;