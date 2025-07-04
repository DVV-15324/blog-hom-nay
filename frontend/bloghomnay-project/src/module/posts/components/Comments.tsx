import { useEffect, useState } from "react";
import { enqueueSnackbar } from "notistack";
import { CommentBase, CreateCommentPayload } from "../models/post";
import { ApiCreateComment } from "../services/api";
import { Response } from "../../common/model";
import { useHookAuth } from "../../auth/hooks/authHooks";
import { useNavigate } from "react-router-dom";
interface CommentBoxProps {
    postId: string;
    initialComments?: CommentBase[];
}

export const CommentBox = ({ postId, initialComments = [] }: CommentBoxProps) => {
    const [content, setContent] = useState("");
    const [loading, setLoading] = useState(false);
    const [commentList, setCommentList] = useState<CommentBase[]>(initialComments ?? []);
    const { profile } = useHookAuth();
    const navigate = useNavigate();

    useEffect(() => {
        setCommentList(initialComments ?? []);
    }, [postId, initialComments]);

    const handleSubmit = async () => {
        if (!content.trim()) {
            enqueueSnackbar("Nội dung bình luận không được để trống", { variant: "warning" });
            return;
        }
        setLoading(true);
        try {
            if (profile != null) {
                const payload: CreateCommentPayload = { post_id: postId, content };
                const newComment = await ApiCreateComment<Response<CommentBase>>(payload);
                setCommentList((prev) => [...prev, newComment.data]);
                setContent("");
            } else {
                navigate("/login");
            }
        } catch (err) {
            enqueueSnackbar("Không thể gửi bình luận", { variant: "error" });
        } finally {
            setLoading(false);
        }
    };

    const handleProfileOthers = (e: React.MouseEvent, user_id: string) => {
        e.stopPropagation();
        navigate(`/user/${user_id}`);
    };
    return (
        <div className="bg-white rounded-xl p-4 shadow-sm">
            {/* Textarea */}
            <textarea
                value={content}
                onChange={(e) => setContent(e.target.value)}
                className="w-full p-3 border border-gray-300 rounded-xl shadow-inner focus:outline-none focus:ring-2 focus:ring-blue-500 transition resize-none text-sm text-black"
                placeholder="Nhập bình luận..."
                rows={3}
                disabled={loading}
            />

            {/* Submit Button */}
            <div className="flex justify-end mt-2">
                <button
                    onClick={handleSubmit}
                    disabled={loading}
                    className={`px-4 py-2 rounded-lg text-sm font-medium transition ${loading
                        ? "bg-blue-400 cursor-not-allowed"
                        : "bg-blue-600 hover:bg-blue-700"
                        } text-white`}
                >
                    {loading ? "Đang gửi..." : "Gửi bình luận"}
                </button>
            </div>

            {/* Danh sách bình luận */}
            <div className="mt-6 space-y-4">
                {commentList.map((c) => (
                    <div key={c.id} className="flex items-start gap-3 border-b pb-4">
                        <img
                            src={c.user?.avatar?.String || "/av.png"}
                            alt={`${c.user?.first_name || ""} ${c.user?.last_name || ""}`}
                            className="w-10 h-10 rounded-full object-cover cursor-pointer"
                            onClick={(e) => handleProfileOthers(e, c.user_id)}

                        />
                        <div className="flex-1">
                            <div className="flex items-center gap-2">
                                <span className="font-semibold text-sm text-gray-800">
                                    {c.user?.first_name || "Ẩn danh"} {c.user?.last_name || ""}
                                </span>
                                <span className="text-xs text-gray-400">
                                    {new Date(c.created_at).toLocaleString()}
                                </span>
                            </div>
                            <p className="mt-1 text-sm text-gray-700">{c.content}</p>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};
