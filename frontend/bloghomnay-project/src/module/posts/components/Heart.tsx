import { useState } from 'react';
import Heart from '@react-sandbox/heart';
import { ApiDisLike, ApiLike } from '../services/api';
import { Response } from "../../auth/model/auth";

interface HeartUIProps {
    width?: number;
    height?: number;
    isLike?: boolean;
    postId?: string;
    onLikeChange?: (newLikeState: boolean) => void;
}

export function HeartUI({ width = 64, height = 64, isLike = false, postId = "", onLikeChange }: HeartUIProps) {
    const [active, setActive] = useState(isLike);

    const handleLike = async () => {
        try {
            if (!active) {
                await ApiLike<Response<boolean>>(postId);
                setActive(true);
                onLikeChange?.(true);
            } else {
                await ApiDisLike<Response<boolean>>(postId);
                setActive(false);
                onLikeChange?.(false);
            }
        } catch (err) {
            console.error("Like/Dislike error:", err);
        }
    };

    return (
        <Heart
            width={width}
            height={height}
            active={active}
            onClick={handleLike} // dùng hàm đã xử lý API
        />
    );
}
