import { useState, useEffect } from "react";
import { PostResponse } from "../models/post";
import MyItemPost from "./MyItemPost";

type Props = {
    posts?: PostResponse[] | null;
};

const MyListPost: React.FC<Props> = ({ posts }) => {
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
        return <p className="text-center text-gray-500 mt-10">Không có bài viết nào.</p>;
    }

    return (

        <div className="xl:p-4 h-full">
            {listPost.map((post) => (
                <MyItemPost key={post.id} post={post} onDelete={handleDeletePost} />
            ))}
        </div>

    );
};

export default MyListPost;
