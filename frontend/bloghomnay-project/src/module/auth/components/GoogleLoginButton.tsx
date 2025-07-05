import { useContext, useEffect } from "react";
import { AuthContext } from "../context/authContext";


declare global {
    interface Window {
        google: any;
    }
}

const GoogleLoginButton = () => {
    const { handleCredentialResponse } = useContext(AuthContext);

    useEffect(() => {
        if (window.google) {
            window.google.accounts.id.initialize({
                client_id: "101056890896-k2a2c4u0ga85ie1kk1gul45qcmedvqnf.apps.googleusercontent.com",
                callback: handleCredentialResponse,
            });

            window.google.accounts.id.renderButton(
                document.getElementById("google-signin"),
                { theme: "outline", size: "large" }
            );
        }
    }, [handleCredentialResponse]);

    return <div id="google-signin"></div>;
};

export default GoogleLoginButton;
