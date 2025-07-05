import { useEffect, useState } from "react";
import axios, { AxiosError } from "axios";
import { Response } from "../../common/model";
import { enqueueSnackbar } from "notistack";
import CircularProgress from "@mui/material/CircularProgress";
import { PostResponse } from "../models/post";
import { ApiGetPostByUser } from "../services/api"; // Sửa lại API phù hợp
import OverViewListPost from "./OverViewListPost";

import { useParams } from "react-router-dom";
import { ProfileType } from "../../auth/model/auth";
import { ApiProfileP } from "../../auth/services/api";
import { GetLastString } from "../../common/profile.slug.ts";


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
        <div className="flex justify-center items-center h-40 min-w-screen">
            <CircularProgress />
        </div>
    );
};

// Component chính
export const OverViewPostUI = () => {
    const { id } = useParams();
    const [profileOver, setProfileOver] = useState<ProfileType | null>(null)
    const [posts, setPosts] = useState<PostResponse[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const lastString = id ? GetLastString(id) : null;
    const handleProfile = async (id: string) => {
        try {
            if (!id) return
            const profile = await ApiProfileP<Response<ProfileType>>(id)
            setProfileOver(profile.data)

            setLoading(false)
        } catch (error) {
            setProfileOver(null)
            setLoading(false)

        }
    }

    const handleGetPostByUser = async () => {
        try {
            if (!lastString) return;
            const res = await ApiGetPostByUser<Response<PostResponse[]>>(lastString);
            setPosts(res.data);
            await handleProfile(lastString)
        } catch (error) {
            const err = ErrorHandle(error as AxiosError);
            enqueueSnackbar(err.message, { variant: "error" });
        } finally {
            setLoading(false);
        }
    };
    useEffect(() => {
        handleGetPostByUser();
    }, [id]);


    return loading ? (
        <DefaultLoading />
    ) : (
        <div className="xl:grid xl:grid-cols-5 h-full w-full gap-4 xl:px-4 py-10">
            <div></div>

            <div className="col-span-3 bg-stone-50 rounded text-white items-center justify-center p-6">
                <div className="container mx-auto xl:px-4 flex justify-center items-center">
                    <div className="flex flex-col xl:flex-row gap-6 mb-6 items-center justify-start border-b border-gray-300 pb-6">
                        <img
                            src={profileOver?.avatar.String || "/av.png"}
                            alt={profileOver?.first_name}
                            className="w-24 h-24 xl:w-32 xl:h-32 rounded-full object-cover"
                        />
                        <div className="text-black space-y-2 text-center xl:text-left">
                            <div className="font-medium sm:text-sm md:text-md xl:text-xl">
                                Họ và tên: {profileOver?.first_name} {profileOver?.last_name}
                            </div>
                            <div className="font-medium sm:text-sm md:text-md xl:text-xl">
                                Email: {profileOver?.email}
                            </div>
                            <div className="font-medium sm:text-sm md:text-md xl:text-xl">
                                Số điện thoại: {profileOver?.phone}
                            </div>
                            <div className="font-medium sm:text-sm md:text-md xl:text-xl">
                                Địa chỉ: {profileOver?.address.String || "---"}
                            </div>
                        </div>
                    </div>
                </div>

                <OverViewListPost posts={posts} />
            </div>

            <div></div>
        </div>
    );

};
