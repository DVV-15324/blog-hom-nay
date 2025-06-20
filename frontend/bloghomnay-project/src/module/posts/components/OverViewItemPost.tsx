import { useNavigate } from "react-router-dom";
import { PostResponse } from "../../posts/models/post";


type Props = {
    post: PostResponse;
};

const OverViewItemPost: React.FC<Props> = ({ post }) => {
    const navigate = useNavigate();
    const handleClick = () => {
        navigate(`/post/${post.id}`);
    };
    return (
        <div className="rounded-xl shadow p-5 bg-white mb-6 border border-gray-200 transition hover:shadow-xl hover:scale-[1.005] duration-300 cursor-pointer" onClick={() => {
            handleClick()
        }}>

            <h2 className="text-xl font-semibold mb-2  text-black">{post.title}</h2>
            <p className="text-sm mb-2 text-black">{post.description}</p>

            <div className="flex flex-wrap gap-2 mt-2">
                {post.tags.map((tag) => (
                    <span
                        key={tag.id}
                        className="text-sm bg-blue-100 text-blue-700 px-3 py-1 rounded-full"
                    >
                        #{tag.name}
                    </span>
                ))}
            </div>

            <div className="text-right text-xs text-gray-400 mt-4">
                Cập nhật lúc: {new Date(post.updated_at).toLocaleString()}
            </div>
        </div >
    );
};

export default OverViewItemPost;