import { useNavigate } from "react-router-dom";
import { PostResponse } from "../../posts/models/post";
import { ConvertProfileToString } from "../../common/profile.slug.ts";

type Props = {
    post: PostResponse;
};

const AddDetailsPost: React.FC<Props> = ({ post }) => {
    const navigate = useNavigate();
    const handleTagClick = (e: React.MouseEvent, tagName: string) => {
        e.stopPropagation();
        navigate(`/search?q=${encodeURIComponent(`[${tagName}]`)}`);
    };
    const handleClick = () => {
        navigate(`/post/${post.slug}`);
    };

    const handleProfileOthers = (e: React.MouseEvent) => {
        e.stopPropagation();
        const url = ConvertProfileToString({ first_name: post.user.first_name, last_name: post.user.last_name, id: post.user.user_id })
        navigate(`/user/${url}`);
    };

    return (
        <div
            className="rounded-2xl shadow-md p-4 bg-white border border-gray-200 hover:shadow-lg transition flex flex-col justify-between h-full"
            onClick={handleClick}
        >
            <div className="flex items-center gap-4 mb-4" >
                <img
                    src={post.user.avatar.String || "/av.png"}
                    alt={post.user.first_name}
                    className="w-10 h-10 rounded-full object-contain cursor-pointer"
                    onClick={handleProfileOthers}
                />
                <div>
                    <div className="font-medium text-black">
                        {post.user.first_name} {post.user.last_name}
                    </div>
                    <div className="text-xs text-gray-500">
                        {new Date(post.created_at).toLocaleDateString()}
                    </div>
                </div>
            </div>

            <h2 className="text-xl font-semibold mb-2 text-black">{post.title}</h2>
            <p className="text-sm mb-2 text-black line-clamp-3 break-words">{post.description}</p>

            <div className="flex flex-wrap gap-2 mt-2">
                {post.tags.map((tag) => (
                    <span
                        key={tag.id}
                        className="text-sm bg-blue-100 text-blue-700 px-3 py-1 rounded-full cursor-pointer"
                        onClick={(e) => handleTagClick(e, tag.name)}
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
export default AddDetailsPost;