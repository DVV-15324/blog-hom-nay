import { Outlet, Route, Routes } from "react-router-dom"
import { AuthLayout } from "../auth/components/AuthLayout"
import { RegisterUI } from "../auth/components/RegisterUI"
import { LoginUI } from "../auth/components/LoginUI"
import { MainLayout } from "./MainLayout"

import { HomeUI } from "../categories/components/HomeUI"
import { PrivateComponent } from "./PrivateComponent"
import { PostUI } from "../posts/components/ItemsPostUI"
import { PublicOnlyRoute } from "../auth/components/PublicOnlyRoute "
import CreatePostUI from "../posts/components/CreatePost"

import { MyPostUI } from "../posts/components/MyPostUI"
import { SearchUI } from "../posts/components/SearchPostUI"
import { NotFound } from "./ErrorUI"
import UpdatePost from "../posts/components/UpdatePost"
import { OverViewPostUI } from "../posts/components/OverViewPostUI"

import { ProfileUI } from "../user/components/ProfileUI"


import { PostDetailWrapper } from "./PostDetailWrapper"


export const MainRoutes = () => {
    return (
        <Routes>
            {/* Auth routes - Public only */}
            <Route
                path="/register"
                element={
                    <PublicOnlyRoute>
                        <AuthLayout>
                            <RegisterUI />
                        </AuthLayout>
                    </PublicOnlyRoute>
                }
            />
            <Route
                path="/login"
                element={
                    <PublicOnlyRoute>
                        <AuthLayout>
                            <LoginUI />
                        </AuthLayout>
                    </PublicOnlyRoute>
                }
            />

            {/* Public Routes - dùng MainLayout nhưng không cần đăng nhập */}
            <Route
                element={
                    <MainLayout>
                        <Outlet />
                    </MainLayout>
                }
            >
                <Route path="/" element={<HomeUI />} />
                <Route path="post/categories/:id" element={<PostUI />} />
                <Route path="search" element={<SearchUI />} />

                <Route path="/post/:id" element={<PostDetailWrapper key={location.pathname} />} />
            </Route>

            {/* Private Routes - cần đăng nhập */}
            <Route
                element={
                    <PrivateComponent>
                        <MainLayout>
                            <Outlet />
                        </MainLayout>
                    </PrivateComponent>
                }
            >
                <Route path="user/:id" element={<OverViewPostUI />} />
                <Route path="create_post" element={<CreatePostUI />} />
                <Route path="profile" element={<ProfileUI />} />
                <Route path="my_post" element={<MyPostUI />} />
                <Route path="my_post/update_post/:id" element={<UpdatePost />} />
            </Route>

            {/* 404 fallback */}
            <Route path="*" element={<NotFound />} />
        </Routes>
    );
};
