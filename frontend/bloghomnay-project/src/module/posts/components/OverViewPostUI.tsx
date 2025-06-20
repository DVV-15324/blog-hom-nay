import { useEffect, useState } from "react";
import axios, { AxiosError } from "axios";
import { Response } from "../../auth/model/auth";
import { enqueueSnackbar } from "notistack";
import CircularProgress from "@mui/material/CircularProgress";
import { PostResponse } from "../models/post";
import { ApiGetPostByUser } from "../services/api"; // Sửa lại API phù hợp
import OverViewListPost from "./OverViewListPost";
import { useHookAuth } from "../../auth/hooks/authHooks";


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
export const OverViewPostUI = () => {
    const { profile } = useHookAuth();
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
        <div className="xl:grid xl:grid-cols-5 h-full w-full gap-4 px-4 py-10">
            <div></div>

            <div className="col-span-3 bg-stone-50 rounded text-white items-center justify-center">
                <div className="flex gap-4 mb-4 items-center justify-start border-b border-gray-300">
                    <img
                        src={profile?.avatar.String}
                        alt={profile?.first_name}
                        className="w-50 h-50 rounded-full object-contain mb-5 mr-2 "
                    />
                    <div>
                        <div className="font-medium text-black text-xl">Họ và tên: {profile?.first_name} {profile?.last_name}</div>
                        <div className="font-medium text-black text-xl">Email: {profile?.email}</div>
                        <div className="font-medium text-black text-xl">Số điện thoại: {profile?.phone}</div>
                        <div className="font-medium text-black text-xl">Địa chỉ: {profile?.address.String || "---"}</div>
                    </div>
                </div>
                {loading ? <DefaultLoading /> : <OverViewListPost posts={posts} />}
            </div>

            <div></div>
        </div>
    );
};
