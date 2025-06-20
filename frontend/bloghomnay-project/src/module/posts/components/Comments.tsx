import { useState } from "react";
import { enqueueSnackbar } from "notistack";
import { CommentBase, CreateCommentPayload } from "../models/post";
import { ApiCreateComment } from "../services/api";
import { Response } from "../../auth/model/auth";
interface CommentBoxProps {
    postId: string;
    initialComments?: CommentBase[];
}

export const CommentBox = ({ postId, initialComments = [] }: CommentBoxProps) => {
    const [content, setContent] = useState("");
    const [loading, setLoading] = useState(false);
    const [commentList, setCommentList] = useState<CommentBase[]>(initialComments ?? []);


    const handleSubmit = async () => {
        if (!content.trim()) {
            enqueueSnackbar("Nội dung bình luận không được để trống", { variant: "warning" });
            return;
        }
        setLoading(true);
        try {
            const payload: CreateCommentPayload = { post_id: postId, content };
            // Gọi API tạo comment, giả sử API trả về comment mới
            const newComment = await ApiCreateComment<Response<CommentBase>>(payload);
            setCommentList((prev) => [...prev, newComment.data]);
            setContent("");
        } catch (err) {
            enqueueSnackbar("Không thể gửi bình luận", { variant: "error" });
        } finally {
            setLoading(false);
        }
    };

    return (
        <div>
            <textarea
                value={content}
                onChange={(e) => setContent(e.target.value)}
                className="text-black w-full border p-2 rounded mb-2 border-black-2"
                placeholder="Nhập bình luận..."
                disabled={loading}
            />
            <button
                onClick={handleSubmit}
                disabled={loading}
                className="px-4 py-2 bg-blue-600 text-white rounded"
            >
                {loading ? "Đang gửi..." : "Gửi bình luận"}
            </button>

            <div className="mt-4">
                {commentList.map((c) => {
                    return (
                        <div key={c.id} className="border-b py-2">
                            <div className="text-black flex items-center gap-2">
                                <img
                                    src={c.user?.avatar?.String || "/default-avatar.png"}
                                    alt={`${c.user?.first_name || ""} ${c.user?.last_name || ""}`}
                                    className="w-8 h-8 rounded-full object-contain"
                                />
                                <strong>
                                    {c.user?.first_name || "Ẩn danh"} {c.user?.last_name || ""}
                                </strong>
                                <span className="text-xs text-gray-400 ml-2">
                                    {new Date(c.created_at).toLocaleString()}
                                </span>
                            </div>
                            <p className="ml-10 text-gray-800">{c.content}</p>
                        </div>
                    );
                })}

            </div>
        </div >
    );
};
