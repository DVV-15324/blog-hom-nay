import { PostResponse } from "../models/post";
import ItemPost from "./ItemPost";

type Props = {
    posts?: PostResponse[] | null;
};

const ListPost: React.FC<Props> = ({ posts }) => {
    const safePosts = posts || [];


    if (safePosts.length < 1) {
        return <p className="text-center text-gray-500 mt-10">Không có bài viết nào.</p>;
    }

    return (
        <div className="xl:p-4">
            {safePosts.map((post) => (
                <ItemPost key={post.id} post={post} />
            ))}
        </div>
    );
};


export default ListPost;