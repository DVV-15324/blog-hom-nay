import React from "react";
import { Menu, MenuItem, Sidebar } from "react-pro-sidebar"
import MenuIcon from '@mui/icons-material/Menu';
import { useHookAuth } from "../auth/hooks/authHooks";
import { useNavigate } from "react-router-dom";
import MenuMUI from "./Menu";
import ISearchComponent from "./SearchComponent";

interface HeaderMainProps {
    children: React.ReactNode
}
// HeaderAuth.tsx
const HeaderMain = ({ children }: HeaderMainProps) => {

    const [toggled, setToggled] = React.useState(false);
    const { profile } = useHookAuth();
    const navigate = useNavigate();
    return (
        <div className="flex rtl">
            <Sidebar className="h-[90px]" rtl onBackdropClick={() => setToggled(false)} toggled={toggled} breakPoint="all">
                <Menu>
                    <MenuItem>
                        <img
                            src={useHookAuth().profile?.avatar.String}
                            alt="avatar"
                            className="m-2 w-12 h-12 rounded-xl object-contain"
                        />
                    </MenuItem>

                    <MenuItem> Trang chủ</MenuItem>
                    <MenuItem> Blog của bạn</MenuItem>
                    <MenuItem>Thông báo</MenuItem>

                </Menu>
            </Sidebar>
            <main className="flex flex-col w-screen min-h-screen">

                <div className="flex fixed top-0 h-[90px] w-full items-center px-4 z-50 border-b border-gray-300 bg-white">
                    <div className="flex flex-1 items-center cursor-pointer">

                        <img src="/logo.png" alt="logo" className="w-12 " onClick={() => {
                            window.location.href = "/";
                        }} />
                        <h1 className="select-none hidden md:block text-xl text-black saira-font text-lg" onClick={() => {
                            window.location.href = "/";
                        }}>Blog Hom Nay</h1>

                    </div>
                    {/*https://www.material-tailwind.com/docs/html/input-search*/}
                    <div className="flex flex-2 xl:flex-1 text-black justify-center w-full">
                        <div className="w-full">
                            <div className="relative">
                                <ISearchComponent />
                            </div>
                        </div>

                    </div>
                    <div className="flex flex-1 xl:flex-2 justify-end">


                        <a className="hidden xl:block p-3 m-2 hover:bg-gray-300 bg-gray-200 cursor-pointer rounded-sm  text-xl  text-green-900 underline" onClick={() => {
                            window.location.href = "/";
                        }}>Trang chủ</a>
                        <a className="hidden xl:block p-3 m-2 hover:bg-gray-300 bg-gray-200 cursor-pointer rounded-sm  text-xl  text-green-900 underline" onClick={() => {
                            navigate(`/user/${profile?.id!}`);
                        }}>Blog của bạn</a>
                        <a className="hidden xl:block p-3 m-2 hover:bg-gray-300 bg-gray-200 cursor-pointer rounded-sm  text-xl  text-green-900 underline">Thông báo</a>
                        <a
                            className="hidden xl:block p-3 m-2 hover:bg-gray-300 bg-gray-200 cursor-pointer rounded-sm text-xl text-green-900 underline"
                            onClick={() => {
                                navigate('/create_post');
                            }}
                        >
                            Tạo Post
                        </a>
                        <div>
                            <MenuMUI>
                                <img
                                    alt="avatar"
                                    className="hidden xl:block m-2 w-12 h-12 rounded-full cursor-pointer object-contain"
                                    src={profile?.avatar.String || "/av.png"}

                                />
                            </MenuMUI>

                        </div>

                        <div className="xl:hidden flex flex-1 fixed top-1 right-1 z-50">
                            <button className="sb-button flex justify-end" onClick={() => setToggled(!toggled)}>
                                <MenuIcon />
                            </button>
                        </div>
                    </div>
                </div>

                <div className="flex h-full w-full pt-[90px]">
                    {children}
                </div>
            </main >
        </div >


    );
};

export default HeaderMain