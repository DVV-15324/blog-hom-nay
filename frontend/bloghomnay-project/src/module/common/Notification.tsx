
{/*https://mui.com/material-ui/react-menu/ */ }
import * as React from 'react';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';

import { useNavigate } from 'react-router-dom';
import axios, { AxiosError } from "axios";
import { Response } from "./model";
import { ApiGetNotification } from '../posts/services/api';
import { Notification } from '../posts/models/post';
import { CircularProgress } from '@mui/material';
import { enqueueSnackbar } from 'notistack';
import { useEffect } from 'react';
import { useHookAuth } from '../auth/hooks/authHooks';
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

interface notiMUIProps {
    children: React.ReactNode
}

export default function NotificationMUI({ children }: notiMUIProps) {
    const { profile } = useHookAuth()
    const [loading, setLoading] = React.useState<boolean>(true);
    const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
    const open = Boolean(anchorEl);
    const [noti, setNoties] = React.useState<Notification[]>([]);
    const handleToPosts = (id: string) => {
        navigate(`/post/${id}`);
    };
    const handleGetNotification = async () => {
        if (profile) {
            try {
                const res = await ApiGetNotification<Response<Notification[]>>();
                setNoties(res.data);
            } catch (error) {
                const err = ErrorHandle(error as AxiosError);
                enqueueSnackbar(err.message, { variant: "error" });
            } finally {
                setLoading(false);
            }
        }

    };

    useEffect(() => {
        handleGetNotification();
    }, []);
    const handleClick = (event: React.MouseEvent) => {
        setAnchorEl(event.currentTarget as HTMLElement);
        if (!profile) {
            navigate("/login");
        }
    };

    const handleClose = () => {
        setAnchorEl(null);
    };
    const navigate = useNavigate();
    return (
        <>
            <a

                className="flex items-center justify-center p-2 m-1 h-[48px] hover:text-amber-700 cursor-pointer rounded-sm text-xl text-green-900"
                onClick={handleClick}
            >
                {children}
            </a>
            <Menu
                disableScrollLock
                id="basic-menu"
                anchorEl={anchorEl}
                open={open}
                onClose={handleClose}
                slotProps={{
                    list: {
                        'aria-labelledby': 'basic-button',
                    },
                }}
            >
                {loading ? (
                    <MenuItem>
                        <DefaultLoading />
                    </MenuItem>
                ) : noti && noti.length > 0 ? (
                    noti.map((item, index) => (
                        <MenuItem
                            key={index}
                            onClick={() => {
                                handleToPosts(item.post_id);
                                handleClose();
                            }}
                            className="flex gap-2 items-start py-2 px-2 whitespace-normal"
                        >
                            <img
                                src={item.user.avatar.String || "/av.png"}
                                alt={item.user.first_name}
                                className="w-10 h-10 rounded-full object-cover"
                            />
                            <div className="flex flex-col">
                                <span className="text-sm font-semibold text-gray-900">
                                    {item.user.first_name + " " + item.user.last_name}
                                </span>
                                <span className="text-sm text-gray-700">
                                    🔔 {item.content}
                                </span>
                                <span className="text-xs text-gray-500 italic">
                                    {item.post?.title || "Không có tiêu đề"}
                                </span>
                            </div>
                        </MenuItem>
                    ))
                ) : (
                    <MenuItem disabled>Không có thông báo</MenuItem>
                )}
            </Menu>
        </>
    );


}
