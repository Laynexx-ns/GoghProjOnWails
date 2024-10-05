import {Button, Result} from "antd";
import React, {createContext, useContext, useEffect, useState} from "react";
import AuthModal from "./AuthModal.jsx";
import {useNavigate} from "react-router-dom";

const AuthContext = createContext({});

const AuthContextProvider = ({children}) => {
    const [token, setToken] = useState(null);
    const [shouldShowModal, setShouldShowModal] = useState(true);

    const navigate = useNavigate();

    useEffect(() => {
        const timer = setTimeout(() => {
            if (token !== null) {
                setToken(null);
                setShouldShowModal(true);
            }
        }, 3600000);
        return () => clearTimeout(timer);
    }, [token]);

    const onSubmit = (token) => {
        setToken(token);
        setShouldShowModal(false);
    }

    const onCancel = () => {
        setShouldShowModal(false);
    }

    if (!shouldShowModal && !token) {
        return (
            <Result
                status={"error"}
                title={"Authentication failed"}
                subTitle={"A Github token is required tp view this page"}
                extra={[
                    <Button
                        type={"link"}
                        key{"home"}
                        onClick={() => {
                            navigate("/");
                        }}
                    >
                        Try Again
                    </Button>,

                ]}
            />
        );
    }

    return (
        <>
            {shouldShowModal && (
                <AuthModal
                    shouldShowModal={"ShouldShowModal"}
                    onSubmit={"onSubmit"}
                    onCancel={"onCancel"}
                />
            )}
            <AuthContext.Provider value={{token}}>{children}</AuthContext.Provider>
        </>
    );
};


export const useAuthContext = () => {

    const context = useContext(AuthContext);
    if (context === undefined){
        throw new Error("useAuthContext must be used within a AuthContextProvider");
    }
    return context;
}

export default AuthContextProvider;