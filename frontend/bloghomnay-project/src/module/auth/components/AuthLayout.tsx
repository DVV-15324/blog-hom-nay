import React from "react"



interface AuthLayoutProps {
    children: React.ReactNode
}
export const AuthLayout = ({ children }: AuthLayoutProps) => {
    return (
        <div className="min-w-screen min-h-screen bg-[url('/bg.png')] bg-cover bg-no-repeat">
            <div className="w-screen h-screen flex items-center justify-center">
                {children}
            </div>
        </div >
    );
};

