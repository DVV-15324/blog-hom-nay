import { useEffect, useState } from "react";
import axios, { AxiosError } from "axios";
<<<<<<< HEAD
import { Response } from "../../common/model";
=======
import { Response } from "../../auth/model/auth";
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
import { enqueueSnackbar } from "notistack";
import CircularProgress from "@mui/material/CircularProgress";
import { PostResponse } from "../models/post";
import { ApiSearchPost } from "../services/api"; // Sửa lại API phù hợp
import ListPost from "./ItemsListPost";
import { useParams } from "react-router-dom";

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
export const SearchUI = () => {
    const { paramStr } = useParams();
    const [posts, setPosts] = useState<PostResponse[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const decodedString = decodeURIComponent(paramStr || "");
    const handleGetPostBySearch = async () => {
        try {
            if (!decodedString) return;
            const res = await ApiSearchPost<Response<PostResponse[]>>(decodedString);
            setPosts(res.data);
        } catch (error) {
            const err = ErrorHandle(error as AxiosError);
            enqueueSnackbar(err.message, { variant: "error" });
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        handleGetPostBySearch();
    }, []);

    return (
        <div className="xl:grid xl:grid-cols-5 h-full w-full gap-4 px-4">
            <div></div>

            <div className="col-span-3 bg-stone-50 rounded text-white items-center justify-center">
                {loading ? <DefaultLoading /> : <ListPost posts={posts} />}
            </div>

            <div></div>
        </div>
    );
};
