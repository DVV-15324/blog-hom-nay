<<<<<<< HEAD
import { useState, useEffect } from "react";
import { PostResponse } from "../models/post";
import MyItemPost from "./MyItemPost";

=======
import { PostResponse } from "../models/post";
import MyItemPost from "./MyItemPost";


>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
type Props = {
    posts?: PostResponse[] | null;
};

const MyListPost: React.FC<Props> = ({ posts }) => {
<<<<<<< HEAD
    const [listPost, setListPost] = useState<PostResponse[]>([]);

    useEffect(() => {
        if (posts) {
            setListPost(posts);
        }
    }, [posts]);

    const handleDeletePost = (id: string) => {
        setListPost((prev) => prev.filter((post) => post.id !== id));
    };

    if (listPost.length < 1) {
=======
    const safePosts = posts || [];

    console.log("Posts length:", safePosts.length);

    if (safePosts.length < 1) {
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
        return <p className="text-center text-gray-500 mt-10">Không có bài viết nào.</p>;
    }

    return (
        <div className="p-4">
            <div className="p-4 h-full">
<<<<<<< HEAD
                {listPost.map((post) => (
                    <MyItemPost key={post.id} post={post} onDelete={handleDeletePost} />
                ))}
            </div>
        </div>
    );
};

export default MyListPost;
=======

                {safePosts.map((post) => (
                    <MyItemPost key={post.id} post={post} />
                ))}
            </div>
        </div>

    );
};

export default MyListPost;
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
