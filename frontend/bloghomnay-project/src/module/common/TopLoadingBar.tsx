import { useRef, useEffect } from "react";
import LoadingBar from "react-top-loading-bar";
import { useLocation } from "react-router-dom";

export const TopLoadingBar = () => {
    const ref = useRef<any>(null);
    const location = useLocation();

    useEffect(() => {
        ref.current?.continuousStart(); // Bắt đầu chạy thanh loading
        // Giả lập thời gian load, bạn có thể gọi done() khi dữ liệu load xong thật sự
        const timer = setTimeout(() => {
            ref.current?.complete(); // Hoàn thành thanh loading
        }, 500);

        return () => {
            clearTimeout(timer);
            ref.current?.complete();
        };
    }, [location.pathname]);

    return <LoadingBar
        color="#ff4500"
        ref={ref}
        className="shadow-[0_0_10px_#ff4500]"
    />

};
