import { useNavigate, useSearchParams } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import { useEffect, useState } from "react";

function EmailValidation() {
    const [searchParams] = useSearchParams();
    const token = searchParams.get("token");
    const navigate = useNavigate();
    const {
        context: { setUser },
        validate,
    } = useAuth();
    const [status, setStatus] = useState<"loading" | "success" | "error">(
        "loading",
    );

    useEffect(() => {
        try {
            const validateEmail = async () => {
                if (token) {
                    const user = await validate(token);
                    setUser(user);
                    setStatus("success");
                    setTimeout(() => navigate("/"), 2000);
                }
            };
            validateEmail();
        } catch (error) {
            console.error(error);
            setStatus("error");
        }
    }, [token, setUser, validate, navigate]);

    return (
        <div className="flex flex-col items-center justify-center min-h-screen text-white bg-black">
            {status === "loading" && <p>Validating your email...</p>}
            {status === "success" && (
                <p className="text-gold font-bold">
                    Email validated successfully! Redirecting...
                </p>
            )}
            {status === "error" && (
                <p className="text-red-500">
                    Invalid or expired token. Please try again.
                </p>
            )}
        </div>
    );
}

export { EmailValidation };
