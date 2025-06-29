import { useEffect, useState } from "react";
import axios, { AxiosError } from "axios";
import { Response } from "../../common/model";
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

            <div className="col-span-3 bg-stone-50 rounded text-white items-center justify-center p-6">
                <div className="flex flex-col xl:flex-row gap-6 mb-6 items-center justify-start border-b border-gray-300 pb-6">
                    <img
                        src={profile?.avatar.String || "/av.png"}
                        alt={profile?.first_name}
                        className="w-24 h-24 xl:w-32 xl:h-32 rounded-full object-cover"
                    />
                    <div className="text-black space-y-2 text-center xl:text-left">
                        <div className="font-medium sm:text-sm md:text-md xl:text-xl">Họ và tên: {profile?.first_name} {profile?.last_name}</div>
                        <div className="font-medium sm:text-sm md:text-md xl:text-xl">Email: {profile?.email}</div>
                        <div className="font-medium sm:text-sm md:text-md xl:text-xl">Số điện thoại: {profile?.phone}</div>
                        <div className="font-medium sm:text-sm md:text-md xl:text-xl">Địa chỉ: {profile?.address.String || "---"}</div>
                    </div>
                </div>

                {loading ? <DefaultLoading /> : <OverViewListPost posts={posts} />}
            </div>


            <div></div>
        </div>
    );
};
