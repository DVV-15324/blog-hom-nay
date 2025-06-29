import React, { useState } from "react";
import { Menu, MenuItem, Sidebar } from "react-pro-sidebar"
import MenuIcon from '@mui/icons-material/Menu';
import { useHookAuth } from "../auth/hooks/authHooks";
import { useNavigate } from "react-router-dom";
import MenuMUI from "./Menu";
import ISearchComponent from "./SearchComponent";
import NotificationMUI from "./Notification";
import NotificationsNoneIcon from '@mui/icons-material/NotificationsNone';
import CircularProgress from "@mui/material/CircularProgress";








interface HeaderMainProps {
    children: React.ReactNode;
}

const HeaderMain = ({ children }: HeaderMainProps) => {
    const { handleOut, profile, loading } = useHookAuth();
    const [toggled, setToggled] = useState(false);
    const navigate = useNavigate();


    if (loading) {
        return (
            <div className="flex justify-center items-center h-screen w-screen">
                <CircularProgress className="flex justify-center items-center"></CircularProgress >
            </div>
        );
    }

    return (
        <div className="flex rtl">
            <Sidebar
                backgroundColor="#ffffff"
                className="h-[90px]"
                rtl
                onBackdropClick={() => setToggled(false)}
                toggled={toggled}
                breakPoint="all"
            >
                <Menu>
                    <div className="flex items-center justify-center flex-col gap-3 my-10">
                        <img
                            src={profile?.avatar.String || "/av.png"}
                            alt="avatar"
                            className="w-20 h-20 rounded-xl object-contain"
                        />
                        <div className="font-bold text-xl">{`${profile?.first_name || ""} ${profile?.last_name || ""}`}</div>
                    </div>

                    <MenuItem
                        className="text-center hover:bg-gray-300 bg-gray-200 cursor-pointer rounded-sm text-xl text-green-900 underline"
                        onClick={() => { window.location.href = "/"; }}
                    >
                        Trang chủ
                    </MenuItem>
                    <MenuItem
                        className="text-center hover:bg-gray-300 bg-gray-200 cursor-pointer rounded-sm text-xl text-green-900 underline"
                        onClick={() => { navigate(`/user/${profile?.id!}`); }}
                    >
                        Trang hồ sơ
                    </MenuItem>

                    <MenuItem
                        className="text-center hover:bg-gray-300 bg-gray-200 cursor-pointer rounded-sm text-xl text-green-900 underline"
                        onClick={() => { navigate("my_post") }}
                    >
                        Quản lý bài viết
                    </MenuItem>
                    <MenuItem
                        className="text-center hover:bg-gray-300 bg-gray-200 cursor-pointer rounded-sm text-xl text-green-900 underline"
                        onClick={() => { navigate("profile") }}
                    >
                        Cá Nhân
                    </MenuItem>
                    <MenuItem
                        className="text-center hover:bg-gray-300 bg-gray-200 cursor-pointer rounded-sm text-xl text-green-900 underline"
                        onClick={() => { navigate("/create_post"); }}
                    >
                        Tạo Post
                    </MenuItem>
                    <MenuItem
                        className="text-center hover:bg-gray-300 bg-red-800 cursor-pointer rounded-sm text-xl text-white underline"
                        onClick={() => { handleOut() }}
                    >
                        Logout
                    </MenuItem>
                </Menu>
            </Sidebar>

            <main className="flex flex-col w-screen min-h-screen">
                <div className="flex fixed top-0 h-[90px] w-full items-center px-4 z-50 border-b border-gray-300 bg-white">
                    <div className="flex flex-1 items-center cursor-pointer">
                        <img src="/logo.png" alt="logo" className="w-12" onClick={() => { window.location.href = "/"; }} />
                        <h1 className="select-none hidden md:block text-xl text-black saira-font text-lg" onClick={() => { window.location.href = "/"; }}>Blog Hom Nay</h1>
                    </div>
                    <div className="flex flex-3 xl:flex-1 text-black justify-center w-full">
                        <div className="w-full">
                            <div className="relative">
                                <ISearchComponent />
                            </div>
                        </div>
                    </div>
                    <div className="flex flex-1 xl:flex-2 justify-end">
                        <a className="hidden xl:block p-3 m-2 hover:bg-gray-300 bg-gray-200 cursor-pointer rounded-sm text-xl text-green-900 underline" onClick={() => { window.location.href = "/"; }}>Trang chủ</a>
                        <a className="hidden xl:block p-3 m-2 hover:bg-gray-300 bg-gray-200 cursor-pointer rounded-sm text-xl text-green-900 underline" onClick={() => { navigate(`/user/${profile?.id!}`); }}>Trang hồ sơ</a>

                        <NotificationMUI>
                            <div className="hidden xl:block p-4 m-2 hover:bg-gray-300 bg-gray-200 cursor-pointer rounded-sm text-xl text-green-900 underline">Thông báo</div>
                        </NotificationMUI>

                        <a className="hidden xl:block p-3 m-2 hover:bg-gray-300 bg-gray-200 cursor-pointer rounded-sm text-xl text-green-900 underline" onClick={() => { navigate('/create_post'); }}>Tạo Post</a>

                        {profile ? (
                            <MenuMUI>
                                <img
                                    alt="avatar"
                                    className="hidden xl:block m-2 w-12 h-12 rounded-full cursor-pointer object-contain"
                                    src={profile?.avatar.String || "/av.png"}
                                />
                            </MenuMUI>
                        ) : (
                            <a
                                className="hidden xl:block p-3 m-2 hover:bg-gray-200 hover:text-black bg-red-900 cursor-pointer rounded-sm text-xl text-white"
                                onClick={() => { navigate('/login'); }}
                            >
                                Đăng nhập
                            </a>
                        )}

                        {profile ? (
                            <div className="xl:hidden flex items-center justify-end fixed right-3 z-50 gap-2">
                                <NotificationMUI>
                                    <NotificationsNoneIcon className="text-black cursor-pointer" />
                                </NotificationMUI>
                                <button className="sb-button" onClick={() => setToggled(!toggled)}>
                                    <MenuIcon className="text-black" />
                                </button>
                            </div>
                        ) : (
                            <div className="w-full xl:hidden">
                                <a
                                    className="block p-1 m-1 hover:bg-gray-200 hover:text-black bg-red-900 cursor-pointer rounded-sm text-xs text-white text-center"
                                    onClick={() => { navigate('/login'); }}
                                >
                                    Đăng nhập
                                </a>
                            </div>
                        )}
                    </div>
                </div>

                <div className="flex h-full w-full pt-[90px]">
                    {children}
                </div>
            </main>
        </div>
    );
};

export default HeaderMain;
