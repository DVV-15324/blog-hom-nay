import { useEffect, useState } from "react";
import axios, { AxiosError } from "axios";
import { Response } from "../../auth/model/auth";
import { enqueueSnackbar } from "notistack";
import CircularProgress from "@mui/material/CircularProgress";
import { PostResponse } from "../models/post";
import { ApiGetPostByUser } from "../services/api"; // Sửa lại API phù hợp
import MyListPost from "./MyListPost";


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
export const DefaultLoading = () => {
    return (
        <div className="flex justify-center items-center h-40">
            <CircularProgress />
        </div>
    );
};

// Component chính
export const MyPostUI = () => {
    const [posts, setPosts] = useState<PostResponse[]>([]);
    const [loading, setLoading] = useState<boolean>(true);

    const handleGetPostByUser = async () => {
        try {

            const res = await ApiGetPostByUser<Response<PostResponse[]>>();
            setPosts(res.data);
        } catch (error) {
            const err = ErrorHandle(error as AxiosError);
            enqueueSnackbar(err.message, { variant: "error" });
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        handleGetPostByUser();
    }, []);

    return (
        <div className="xl:grid xl:grid-cols-5 h-full w-full gap-4 px-4">
            <div></div>

            <div className="col-span-3 bg-stone-50 rounded text-white items-center justify-center">
                {loading ? <DefaultLoading /> : <MyListPost posts={posts} />}
            </div>

            <div></div>
        </div>
    );
};
