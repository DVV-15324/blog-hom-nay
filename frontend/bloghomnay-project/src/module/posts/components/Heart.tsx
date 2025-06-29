import { useState } from 'react';
import Heart from '@react-sandbox/heart';
import { ApiDisLike, ApiLike } from '../services/api';
<<<<<<< HEAD
import { Response } from "../../common/model";
import { useNavigate } from 'react-router-dom';
import { useHookAuth } from '../../auth/hooks/authHooks';
=======
import { Response } from "../../auth/model/auth";
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20

interface HeartUIProps {
    width?: number;
    height?: number;
    isLike?: boolean;
    postId?: string;
    onLikeChange?: (newLikeState: boolean) => void;
}

export function HeartUI({ width = 64, height = 64, isLike = false, postId = "", onLikeChange }: HeartUIProps) {
    const [active, setActive] = useState(isLike);
<<<<<<< HEAD
    const { profile } = useHookAuth();
    const navigate = useNavigate()
    const handleLike = async () => {
        try {
            if (profile != null) {
                if (!active) {
                    await ApiLike<Response<boolean>>(postId);
                    setActive(true);
                    onLikeChange?.(true);
                } else {
                    await ApiDisLike<Response<boolean>>(postId);
                    setActive(false);
                    onLikeChange?.(false);
                }
            } else {
                navigate("/login")
            }

=======

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
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
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
