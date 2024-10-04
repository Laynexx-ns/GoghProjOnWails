import {Button, Result} from "antd";
import React, {createContext, useContext, useEffect, useState} from "react";
import AuthModal from "./Authmodal";
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
}