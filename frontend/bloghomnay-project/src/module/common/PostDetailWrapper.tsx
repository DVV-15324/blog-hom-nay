import { useParams } from "react-router-dom";
import PostsDetail from "../posts/components/PostDetailUI";


export const PostDetailWrapper = () => {
    const { id } = useParams();

    // ép remount PostDetail mỗi lần id thay đổi
    return <PostsDetail key={id} />;
};
