import { useContext } from "react";
import { AuthContext } from "../context/auth";
import { register, login, logout, validate } from "../api/auth";

function useAuth() {
    const context = useContext(AuthContext);

    console.log(context);

    if (!context) {
        throw new Error("useAuth must be used within an AuthProvider.");
    }

    return { context, register, validate, login, logout };
}

export { useAuth };
