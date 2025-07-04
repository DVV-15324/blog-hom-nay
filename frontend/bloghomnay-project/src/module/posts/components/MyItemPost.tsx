import { useNavigate } from "react-router-dom";
import { PostResponse } from "../../posts/models/post";
import { Pencil, Trash2 } from "lucide-react";
import { useState } from "react";
import {
    Button,
    CircularProgress,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle
} from "@mui/material";
import { ApiDeletePost } from "../services/api";
import { Response } from "../../common/model";
import { enqueueSnackbar } from "notistack";

type Props = {
    post: PostResponse;
    onDelete: (id: string) => void;

};

// Loading component
export const DefaultLoading = () => {
    return (
        <div className="flex justify-center items-center h-40">
            <CircularProgress />
        </div>
    );
};

const MyItemPost: React.FC<Props> = ({ post, onDelete }) => {
    const navigate = useNavigate();
    const [openConfirm, setOpenConfirm] = useState(false);
    const [loading, setLoading] = useState<boolean>(false);

    const handleClick = () => {
        navigate(`/post/${post.slug}`);
    };

    const handleUpdatePost = (e: React.MouseEvent) => {
        e.stopPropagation();
        navigate(`/my_post/update_post/${post.id}`);
    };

    const handleDeleteClick = (e: React.MouseEvent) => {
        e.stopPropagation();
        setOpenConfirm(true);
    };

    const handleConfirmDelete = async () => {
        setOpenConfirm(false);
        setLoading(true);
        try {
            await ApiDeletePost<Response<PostResponse[]>>({ id: post.id });
            enqueueSnackbar("Xoá bài viết thành công", { variant: "success" });
            onDelete(post.id); // Gọi callback để xoá khỏi UI
        } catch (err: any) {
            enqueueSnackbar(err.message || "Đã xảy ra lỗi khi xoá", { variant: "error" });
        } finally {
            setLoading(false);
        }
    };

    const handleTagClick = (e: React.MouseEvent, tagName: string) => {
        e.stopPropagation();
        navigate(`/search?q=${encodeURIComponent(`[${tagName}]`)}`);
    };
    return (
        <>
            {loading && <DefaultLoading />}
            <div className="max-w-5xl container mx-auto px-4 flex justify-center items-center">
                <div
                    className="relative w-4xl  m-2 rounded-2xl shadow-md p-4 bg-white border border-gray-200 hover:shadow-lg transition cursor-pointer group flex flex-col justify-between"
                    onClick={handleClick}
                >
                    <div className="absolute top-3 right-3 flex space-x-2">
                        <button
                            className="text-gray-500 hover:text-blue-600 cursor-pointer"
                            onClick={handleUpdatePost}
                            title="Chỉnh sửa"
                        >
                            <Pencil size={20} />
                        </button>
                        <button
                            className="text-gray-500 hover:text-red-600 cursor-pointer"
                            onClick={handleDeleteClick}
                            title="Xoá bài viết"
                        >
                            <Trash2 size={20} />
                        </button>
                    </div>

                    <h2 className="text-xl font-bold text-gray-800 mb-2 break-words">{post.title}</h2>
                    <p className="text-sm text-gray-700 mb-3 break-words line-clamp-3 mb-2 break-words">{post.description}</p>

                    <div className="flex flex-wrap gap-2 mt-2">
                        {post.tags.map((tag) => (
                            <span
                                key={tag.id}
                                className="text-xs bg-blue-100 text-blue-700 px-3 py-1 rounded-full"
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
            </div>
            <Dialog
                open={openConfirm}
                onClose={() => setOpenConfirm(false)}
            >
                <DialogTitle>Xác nhận xoá</DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        Bạn có chắc chắn muốn xoá bài viết: <strong>{post.title}</strong>?
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={() => setOpenConfirm(false)} color="inherit">
                        Huỷ
                    </Button>
                    <Button onClick={handleConfirmDelete} color="error">
                        Xoá
                    </Button>
                </DialogActions>
            </Dialog>
        </>
    );
};

export default MyItemPost;
