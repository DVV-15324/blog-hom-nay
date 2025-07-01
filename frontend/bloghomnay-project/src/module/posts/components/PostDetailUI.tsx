import { useNavigate, useParams } from "react-router-dom";
import { PostResponse } from "../../posts/models/post";
import { useEffect, useState } from "react";
import { enqueueSnackbar } from "notistack";
import axios, { AxiosError } from "axios";
import CircularProgress from "@mui/material/CircularProgress";
import { HeartUI } from "./Heart"
import { CommentBox } from "./Comments"
import { Response } from "../../common/model";
import { ApiGetPostById, ApiGetPostByIdP } from "../services/api";

import { useHookAuth } from "../../auth/hooks/authHooks";
import { Helmet } from 'react-helmet';
import PreviewWithCodeBlock from "./PreviewWithCodeBlock";

// Hàm xử lý lỗi
const ErrorHandle = (error: AxiosError | Error) => {
    if (axios.isAxiosError(error)) {
        return {
            message: error.response?.data.error || "Lỗi hệ thống",
            error: error.response?.data.message,
        };
    }
    return { message: error.message || "UnKnown Error" };
};

// Loading component
export const DefaultLoading = () => (
    <div className="flex justify-center items-center h-40">
        <CircularProgress />
    </div>
);

const PostsDetail = () => {
    const { profile } = useHookAuth();
    const { id } = useParams();
    const [posts, setPosts] = useState<PostResponse>();
    const [loading, setLoading] = useState<boolean>(true);
    const [likeCount, setLikeCount] = useState<number>(0);
    const navigate = useNavigate();
    const [isLike, setIsLike] = useState<boolean>(false);
    useEffect(() => {
        if (posts) {
            setLikeCount(posts.like);
            setIsLike(posts.islike);
        }
    }, [posts]);

    const getLastString = (str: string): string => {
        const parts = str.split("-");
        return parts[parts.length - 1];
    };

    const lastString = id ? getLastString(id) : null;

    useEffect(() => {
        if (posts) {
            setLikeCount(posts.like);
        }
    }, [posts]);

    const handleProfileOthers = (e: React.MouseEvent) => {
        e.stopPropagation();
        navigate(`/user/${posts?.user_id}`);
    };

    const handleGetPostById = async () => {
        if (!lastString) return;
        try {
            if (!id) return;
            if (profile != null) {
                const res = await ApiGetPostById<Response<PostResponse>>(lastString);
                setPosts(res.data);
            }
            const res = await ApiGetPostByIdP<Response<PostResponse>>(lastString);
            setPosts(res.data);

        } catch (error) {
            const err = ErrorHandle(error as AxiosError);
            enqueueSnackbar(err.message, { variant: "error" });
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        handleGetPostById();
    }, [id]);




    if (!posts) {
        return
    }

    const {
        user,
        created_at,
        updated_at,
        //title,
        content,
        tags,
    } = posts;

    return (
        <>
            <Helmet>
                <title>{posts.title} | BlogHomNay</title>
                <meta name="description" content={posts.description} />
            </Helmet>

            {loading ? (
                <div className="flex justify-center items-center w-full min-h-screen">
                    <DefaultLoading />
                </div>

            ) : posts ? (
                <div className="grid xl:grid-cols-5 h-full w-full gap-4 px-4 py-10">
                    {/* Sidebar trái cố định */}
                    <aside className="xl:col-span-1 hidden xl:block">
                        <div className="fixed flex flex-col top-24 left-0 w-1/5 flex flex-col items-end gap-4 xl:mt-20" >
                            <div className="flex justify-center items-center flex-col">
                                <img
                                    src={user?.avatar.String || "/av.png"}
                                    alt={user?.first_name}
                                    className="flex w-20 h-20 rounded-full shadow justify-center items-center cursor-pointer object-contain mb-2" onClick={handleProfileOthers}
                                />

                                <h2 className="text-xl font-semibold text-gray-800 cursor-pointer" onClick={handleProfileOthers}>
                                    {user?.first_name} {user?.last_name}
                                </h2>
                                <div className="my-2 ">
                                    <div className="flex justify-center items-center">
                                        <HeartUI
                                            isLike={isLike}
                                            postId={posts.id}
                                            onLikeChange={(liked) => {
                                                setIsLike(liked);
                                                setLikeCount((prev) => liked ? prev + 1 : prev - 1);
                                            }}
                                        />

                                    </div>

                                    <div className="text-sm text-gray-500 mt-1">{likeCount} lượt thích</div>
                                </div>
                                {created_at && (
                                    <div className="text-xs text-gray-500">
                                        Ngày tạo: {new Date(created_at).toLocaleDateString()}
                                    </div>
                                )
                                }
                            </div>


                        </div>
                    </aside>
                    <div className="xl:col-span-3 bg-stone-50 rounded text-white items-center justify-center">
                        <div className="rounded-xl shadow p-5 bg-white mb-6 border border-gray-200">
                            {/*<h2 className="text-xl font-semibold mb-2 text-black">{title}</h2>*/}
                            <div className="text-gray-700 mb-3 ">
                                <PreviewWithCodeBlock content={content} />
                            </div>


                            <div className="flex flex-wrap gap-2 mt-2">
                                {tags.map((tag) => (
                                    <span
                                        key={tag.id}
                                        className="text-sm bg-blue-100 text-blue-700 px-3 py-1 rounded-full"
                                    >
                                        #{tag.name}
                                    </span>
                                ))}
                            </div>

                            <div className="text-right text-xs text-gray-400 mt-4">
                                {updated_at && new Date(updated_at).getFullYear() > 2000
                                    ? `Cập nhật lúc: ${new Date(updated_at).toLocaleString()}`
                                    : "Chưa cập nhật"}
                            </div>

                            <CommentBox initialComments={posts.comments} postId={posts.id} />
                        </div>

                    </div>
                    <div className="hidden xl:col-span-1"></div>
                    {/* Sidebar dưới cho mobile (cố định ở dưới cùng màn hình) */}
                    <div className="fixed bottom-0 left-0 w-full xl:hidden bg-white border-t p-4 shadow z-50">
                        <div className="flex flex-raw  items-center justify-center gap-2">
                            <img
                                src={user?.avatar.String || "/av.png"}
                                alt={user?.first_name}
                                className="w-10 h-10 rounded-full object-contain shadow"
                            />
                            <div className="text-center">
                                <h2 className="text-base font-semibold text-gray-800">
                                    {user?.first_name} {user?.last_name}
                                </h2>
                            </div>
                            <div className="text-center">
                                <div className="flex justify-center items-center">
                                    <HeartUI width={30} height={30} isLike={posts.islike} postId={posts.id} onLikeChange={(liked) => {
                                        setLikeCount((prev) => liked ? prev + 1 : prev - 1);
                                    }} />
                                </div>
                            </div>
                            <div className="text-sm text-gray-500 mt-1">{likeCount} lượt thích</div>
                            {created_at && (
                                <div className="text-xs text-gray-500 mt-1">
                                    Ngày tạo: {new Date(created_at).toLocaleDateString()}
                                </div>
                            )}

                        </div>

                    </div>



                </div>) : <div className="text-center text-red-500">Không tìm thấy bài viết.</div>}

        </>
    );
};

export default PostsDetail;
