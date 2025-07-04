import { useNavigate, useParams } from "react-router-dom";
import { PostResponse } from "../../posts/models/post";
import { useEffect, useState } from "react";
import { enqueueSnackbar } from "notistack";
import axios, { AxiosError } from "axios";
import CircularProgress from "@mui/material/CircularProgress";
import { HeartUI } from "./Heart"
import { CommentBox } from "./Comments"
import { Response } from "../../common/model";
import { ApiGetPostById, ApiGetPostByIdP, ApiGetPostByUser, ApiSearchPost } from "../services/api";
import TableOfContents, { TOCItem } from "./TableOfContents"; // đường dẫn điều chỉnh cho đúng

import useMediaQuery from '@mui/material/useMediaQuery';


import { useHookAuth } from "../../auth/hooks/authHooks";
import { Helmet } from 'react-helmet';
import PreviewWithCodeBlock from "./PreviewWithCodeBlock";
import AddDetailsPost from "./AddDetailsPost";
import AddDetailsPostTags from "./AddDetailsPosstTags";
import { useTheme } from "@mui/material/styles";
import { ConvertProfileToString, GetLastString } from "../../common/profile.slug.ts";




function extractTOCFromHTML(html: string): TOCItem[] {
    const parser = new DOMParser();
    const doc = parser.parseFromString(html, "text/html");

    const headings = Array.from(doc.querySelectorAll("h1, h2, h3"));

    let tocIndex = 0;

    return headings
        .filter((heading) =>
            heading.textContent?.trim() &&
            !heading.querySelector("pre")

        )
        .map((heading) => {
            const id = `heading-${tocIndex++}`;
            return {
                id,
                text: heading.textContent!.trim(),
                tag: heading.tagName.toLowerCase() as "h1" | "h2" | "h3",
            };
        });
}


function addIdToHeadings(html: string): string {
    const parser = new DOMParser();
    const doc = parser.parseFromString(html, "text/html");

    const headings = Array.from(doc.querySelectorAll("h1, h2, h3"));
    let validIndex = 0;

    headings.forEach((heading) => {
        const hasPre = heading.querySelector("pre");

        const hasText = heading.textContent?.trim();

        if (!hasPre && hasText) {
            heading.setAttribute("id", `heading-${validIndex}`);
            validIndex++;
        }
    });

    return doc.body.innerHTML;
}

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


interface PostListProps<T> {
    titleContent: string;
    posts: T[];
    RenderItem: React.ComponentType<{ post: T }>;
}



const PostList = <T,>({ posts, RenderItem, titleContent }: PostListProps<T>) => {
    const theme = useTheme();
    const isSmall = useMediaQuery(theme.breakpoints.down("sm"));

    return (
        <div className="w-full mt-10">
            <h2 className="text-black font-bold mb-2">{titleContent}</h2>
            <div
                className={`${isSmall
                    ? "flex flex-col space-y-4" // dọc, cách nhau 16px (space-y-4)
                    : "grid grid-cols-3 gap-4 mb-10"
                    }`}
            >
                {posts.map((post) => (
                    <div
                        key={(post as any).id}
                        className={`${isSmall ? "w-full" : ""}`}
                    >
                        <RenderItem post={post} />
                    </div>
                ))}
            </div>
        </div>
    );
};



const PostsDetail = () => {
    const { profile } = useHookAuth();
    const { id } = useParams();
    const [post, setPost] = useState<PostResponse>();
    const [loading, setLoading] = useState<boolean>(true);
    const [likeCount, setLikeCount] = useState<number>(0);
    const navigate = useNavigate();
    const [isLike, setIsLike] = useState<boolean>(false);
    const [tocList, setTocList] = useState<TOCItem[]>([]);
    const [activeId, setActiveId] = useState<string | null>(null);
    const [isContentReady, setIsContentReady] = useState(false);
    const [posts, setPosts] = useState<PostResponse[]>([]);
    const [postsTag, setPostsTags] = useState<PostResponse[]>([]);


    useEffect(() => {
        if (isContentReady && tocList.length > 0 && !activeId) {
            setActiveId(tocList[0].id);
        }
    }, [isContentReady, tocList]);

    useEffect(() => {
        if (post?.content) {
            const htmlWithIds = addIdToHeadings(post.content);
            const list = extractTOCFromHTML(htmlWithIds);
            setTocList(list);
        }
    }, [post?.content]);
    useEffect(() => {
        if (!post) return;
        handleGetPostByTags();
    }, [post]);
    useEffect(() => {
        if (post) {
            setLikeCount(post.like);
            setIsLike(post.islike);
        }
    }, [post]);

    const lastString = id ? GetLastString(id) : null;

    useEffect(() => {
        if (post) {
            setLikeCount(post.like);
        }
    }, [post]);

    const handleProfileOthers = (e: React.MouseEvent) => {
        e.stopPropagation();
        if (!post) return;
        const url = ConvertProfileToString({ first_name: post.user.first_name, last_name: post.user.last_name, id: post.user.user_id })
        navigate(`/user/${url}`);
    };
    const handleGetPostByTags = async () => {
        try {
            if (post == null) return
            const tagNames = post.tags.map(tag => `[${tag.name}]`).join(" ");
            const res = await ApiSearchPost<Response<PostResponse[]>>(tagNames);
            setPostsTags(res.data);

        } catch (error) {

            const err = ErrorHandle(error as AxiosError);
            enqueueSnackbar(err.message, { variant: "error" });
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        handleGetPostByTags();
    }, []);

    const handleGetPostByUser = async () => {
        try {
            if (!post?.user_id) return
            const res = await ApiGetPostByUser<Response<PostResponse[]>>(post.user_id);

            setPosts(res.data);
        } catch (error) {

            const err = ErrorHandle(error as AxiosError);
            enqueueSnackbar(err.message, { variant: "error" });
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        if (post?.user_id) {
            handleGetPostByUser();
        }
    }, [post]);


    const handleGetPostById = async () => {
        if (!lastString) return;
        try {
            if (!id) return;
            if (profile != null) {
                const res = await ApiGetPostById<Response<PostResponse>>(lastString);
                setPost(res.data);
            }
            const res = await ApiGetPostByIdP<Response<PostResponse>>(lastString);
            setPost(res.data);

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

    useEffect(() => {
        const onScroll = () => {
            const headings = Array.from(document.querySelectorAll("h1, h2, h3")) as HTMLElement[];
            const scrollPos = window.scrollY + 120; // tùy chỉnh offset
            let currentId = "";
            for (let heading of headings) {
                if (heading.offsetTop <= scrollPos) {
                    currentId = heading.id;
                } else break;
            }
            if (currentId !== activeId) setActiveId(currentId);
        };

        window.addEventListener("scroll", onScroll);
        return () => window.removeEventListener("scroll", onScroll);
    }, []);




    if (!post) {
        return
    }

    const {
        user,
        created_at,
        updated_at,
        //title,
        content,
        tags,
    } = post;

    return (
        <>
            <Helmet>
                <title>{post.title} | BlogHomNay</title>
                <meta name="description" content={post.description} />
            </Helmet>

            {loading ? (
                <div className="flex justify-center items-center w-full min-h-screen">
                    <DefaultLoading />
                </div>

            ) : posts ? (
                <div className="container mx-auto px-4 flex justify-center items-center">
                    <div className="grid xl:grid-cols-5 h-full w-full gap-4 py-10">
                        {/* Sidebar trái cố định */}
                        <aside className="xl:col-span-1 hidden xl:block">
                            <div className="sticky top-24 flex flex-col w-full items-end gap-4 xl:mt-20" >

                                <div className="fixed flex justify-center items-center flex-col">
                                    <img
                                        src={user?.avatar.String || "/av.png"}
                                        alt={user?.first_name}
                                        className="w-20 h-20 rounded-full object-cover shadow cursor-pointer"
                                        onClick={handleProfileOthers}
                                    />


                                    <h2 className="text-xl font-semibold text-gray-800" >
                                        {user?.first_name} {user?.last_name}
                                    </h2>
                                    <div className="my-2 ">
                                        <div className="flex justify-center items-center">
                                            <HeartUI
                                                isLike={isLike}
                                                postId={post.id}
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
                                    <PreviewWithCodeBlock
                                        content={addIdToHeadings(content)}
                                        onRendered={() => setIsContentReady(true)}
                                    />
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

                                <CommentBox initialComments={post.comments} postId={post.id} />

                            </div >
                            <PostList posts={postsTag} RenderItem={AddDetailsPostTags} titleContent={"Bài viết cùng tags"} />
                            <PostList posts={posts} RenderItem={AddDetailsPost} titleContent={"Bài blog của người viết"} />

                        </div >


                        <div className="hidden xl:col-span-1 xl:block">
                            <TableOfContents
                                tocList={tocList}
                                activeId={activeId}

                            />


                        </div>



                        {/* Sidebar dưới cho mobile (cố định ở dưới cùng màn hình) */}
                        <div className="fixed bottom-0 left-0 w-full xl:hidden bg-white border-t p-4 shadow z-50">
                            <div className="flex flex-raw  items-center justify-center gap-2">
                                <img
                                    src={user?.avatar.String || "/av.png"}
                                    alt={user?.first_name}
                                    className="w-10 h-10 rounded-full object-cover shadow"
                                    onClick={handleProfileOthers}
                                />

                                <div className="text-center">
                                    <h2 className="text-base font-semibold text-gray-800">
                                        {user?.first_name} {user?.last_name}
                                    </h2>
                                </div>
                                <div className="text-center">
                                    <div className="flex justify-center items-center">
                                        <HeartUI width={30} height={30} isLike={post.islike} postId={post.id} onLikeChange={(liked) => {
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



                    </div >
                </div>) : <div className="text-center text-red-500">Không tìm thấy bài viết.</div>}

        </>
    );
};

export default PostsDetail; 