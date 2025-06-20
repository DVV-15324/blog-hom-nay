import { PostResponse } from "../models/post";
import OverViewItemPost from "./OverViewItemPost";


type Props = {
    posts?: PostResponse[] | null;
};

const OverViewListPost: React.FC<Props> = ({ posts }) => {
    const safePosts = posts || [];

    console.log("Posts length:", safePosts.length);

    if (safePosts.length < 1) {
        return <p className="text-center text-gray-500 mt-10">Không có bài viết nào.</p>;
    }

    return (
        <div className="p-4 h-full">

            {safePosts.map((post) => (
                <OverViewItemPost key={post.id} post={post} />
            ))}
        </div>
    );
};

export default OverViewListPost;