
import { useForm } from "react-hook-form"
import { yupResolver } from "@hookform/resolvers/yup";
import { useHookAuth } from "../hooks/authHooks";

import { LoginSchema } from "../model/schema";

export function LoginUI() {
    const { handleLogin } = useHookAuth();


    const methods = useForm({
        resolver: yupResolver(LoginSchema)

    })
    const { handleSubmit, register, formState: { errors } } = methods;
    return (
        <form className="flex flex-col bg-white p-12 rounded xl:w-xl md:w-base sm:w-sm lg:w-lg " onSubmit={handleSubmit(handleLogin)}>
            <h1 className="text-5xl text-center font-bold saira-font mb-6">
                Blog Hôm Nay
            </h1>
            <div className="mb-4">
                <label className="block mb-2" htmlFor="email">
                    Email
                </label>
                <input
                    className={`border rounded w-full px-3 py-2 ${errors.email ? "border-red-500" : "border-gray-300"
                        }`}
                    type="email"
                    id="email"
                    placeholder="Enter Email"
                    {...register("email")}
                />
                {errors.email && (
                    <p className="text-red-500 text-sm mt-1">{errors.email.message}</p>
                )}

            </div>

            <div className="mb-6">
                <label className="block mb-2" htmlFor="password">
                    Password
                </label>
                <input
                    className={`border rounded w-full px-3 py-2 ${errors.password ? "border-red-500" : "border-gray-300"
                        }`}
                    type="password"
                    id="password"
                    placeholder="Enter Password"
                    {...register("password")}
                />
                {errors.password && (
                    <p className="text-red-500 text-sm mt-1">{errors.password.message}</p>
                )}

            </div>

            <button type="submit" className="bg-red-900 text-white px-4 py-2 rounded w-full">
                Login
            </button>

            <h1 className="mt-4 text-center">
                Chưa có tài khoản?{" "}
                <a href="/register" className="text-blue-600 underline">
                    Đăng ký
                </a>
            </h1>
        </form>
    );
}
