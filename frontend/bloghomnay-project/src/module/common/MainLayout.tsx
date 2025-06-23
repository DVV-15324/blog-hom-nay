import React from "react"
import HeaderMain from "./Header"


interface AuthLayoutProps {
    children: React.ReactNode
}
export const MainLayout = ({ children }: AuthLayoutProps) => {
    return (
        <div className="flex bg-gray-50 ">
            <HeaderMain>

                {children}

            </HeaderMain >
        </div >
    )
}