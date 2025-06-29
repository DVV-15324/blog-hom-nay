import { useEffect, useState } from "react";
import axios, { AxiosError } from "axios";
<<<<<<< HEAD
import { Response } from "../../common/model";
=======
import { Response } from "../../auth/model/auth";
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
import { enqueueSnackbar } from "notistack";
import CircularProgress from "@mui/material/CircularProgress";
import { Categories } from "../module/categories";
import { ApiGetAllCategories } from "../services/api";
import ListCategories from "./ListCategoriesUI";

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
export const HomeUI = () => {
    const [categories, setCategories] = useState<Categories[]>([]);
    const [loading, setLoading] = useState<boolean>(true);

    const handleGetAllCategories = async () => {
        try {
            const res = await ApiGetAllCategories<Response<Categories[]>>();
            setCategories(res.data);
        } catch (error) {
            const err = ErrorHandle(error as AxiosError);
            enqueueSnackbar(err.message, { variant: "error" });
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        handleGetAllCategories();
    }, []);

    return (
        <div className="w-full bg-stone-50 py-8">

            <div className="max-w-screen-xl mx-auto px-4">
                <div className="bg-white rounded-lg shadow p-6">
                    <h1 className="text-2xl font-semibold mb-2">
                        Hello Chào Mừng Các Bạn Đến Với BlogHomNay.
                    </h1>
                    <p className="text-gray-700">
                        Nơi Các Bạn Có Thể Viết Các Bài Blog Cá Nhân Và Nơi Học Hỏi <br />
                        Và Cùng Nhau Chia Sẻ Kiến Thức Về Công Nghệ Thông Tin
                    </p>
                </div>

                <div className="mt-6">
                    {loading ? <DefaultLoading /> : <ListCategories categories={categories} />}
                </div>
            </div>
        </div>
    );

};
