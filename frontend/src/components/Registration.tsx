import { useMemo, useCallback, useRef, useState } from "react";
import { useAuth } from "../hooks/useAuth";
import { User } from "../api/types";

enum Tab {
    Login,
    Register,
}

enum PasswordError {
    TooShort,
    NoUppercase,
    NoMatch,
    Invalid,
    None,
}

function Registration() {
    const {
        context: { setUser },
    } = useAuth();
    const [activeTab, setActiveTab] = useState<Tab>(Tab.Login);
    const [error, setError] = useState<PasswordError | string>(
        PasswordError.None,
    );

    const { register, login } = useAuth();

    const nameRef = useRef<HTMLInputElement>(null);
    const emailRef = useRef<HTMLInputElement>(null);
    const passwordRef = useRef<HTMLInputElement>(null);
    const repeatPasswordRef = useRef<HTMLInputElement>(null);

    const validatePassword = (): PasswordError => {
        if (passwordRef.current!.value.length < 8) {
            return PasswordError.TooShort;
        }

        if (!/[A-Z]/.test(passwordRef.current!.value)) {
            return PasswordError.NoUppercase;
        }

        if (passwordRef.current?.value !== repeatPasswordRef.current?.value) {
            return PasswordError.NoMatch;
        }

        return PasswordError.None;
    };

    const passwordErrorToString = useMemo((): string => {
        switch (error) {
            case PasswordError.TooShort:
                return "Password must be at least 8 characters long.";
            case PasswordError.NoUppercase:
                return "Password must contain at least 1 uppercase letter.";
            case PasswordError.NoMatch:
                return "Passwords don't match.";
            case PasswordError.Invalid:
                return "Invalid email or password";
            case PasswordError.None:
                return "";
        }
        return error;
    }, [error]);

    const handleSubmit = useCallback(async () => {
        if (activeTab === Tab.Login) {
            try {
                const user: User = await login(
                    emailRef.current!.value,
                    passwordRef.current!.value,
                );
                setUser(user);
            } catch (error) {
                if (error instanceof Error && parseInt(error.message) < 500) {
                    setError(PasswordError.Invalid);
                } else {
                    setError("Server error. Try again later.");
                }
            }
        } else {
            const error: PasswordError = validatePassword();

            if (error === PasswordError.None) {
                await register(
                    nameRef.current!.value,
                    emailRef.current!.value,
                    passwordRef.current!.value,
                );
                return;
            } else {
                setError(error);
                return;
            }
        }
    }, [activeTab, register, login, setUser]);

    return (
        <div className="flex flex-col w-[95%] sm:w-2/3 lg:w-1/2 bg-dark border-4 border-sunny p-8 z-10 text-sunny">
            <ul className="flex justify-evenly">
                <li
                    className={`${activeTab === Tab.Login && "border-b-2 border-sunny"} cursor-pointer`}
                    onClick={() => setActiveTab(Tab.Login)}
                >
                    Login
                </li>
                <li
                    className={`${activeTab === Tab.Register && "border-b-2 border-sunny"} cursor-pointer`}
                    onClick={() => setActiveTab(Tab.Register)}
                >
                    Register
                </li>
            </ul>
            <div className="flex flex-col gap-6">
                {activeTab === Tab.Register && (
                    <div className="flex flex-col gap-1">
                        <p>Name:</p>
                        <input
                            className="border-1 border-sunny rounded-xl p-2 outline-none"
                            type="text"
                            ref={nameRef}
                        />
                    </div>
                )}
                <div className="flex flex-col">
                    <p>Email:</p>
                    <input
                        className="border-1 border-sunny rounded-xl p-2 outline-none"
                        type="email"
                        ref={emailRef}
                    />
                </div>
                <div className="flex flex-col">
                    <p>Password:</p>
                    <input
                        className="border-1 border-sunny rounded-xl p-2 outline-none"
                        type="password"
                        ref={passwordRef}
                    />
                    {error !== PasswordError.None && (
                        <p className="text-red-500">{passwordErrorToString}</p>
                    )}
                </div>
                {activeTab === Tab.Register && (
                    <div className="flex flex-col">
                        <p>Repeat Password:</p>
                        <input
                            className="border-1 border-sunny rounded-xl p-2 outline-none"
                            type="password"
                            ref={repeatPasswordRef}
                        />
                    </div>
                )}
                <button
                    className="cursor-pointer border-2 border-sunny bg-sunny text-dark rounded-xl py-1.5 hover:bg-sunset hover:text-sunny"
                    onClick={handleSubmit}
                >
                    {activeTab === Tab.Login ? "Login" : "Register"}
                </button>
            </div>
        </div>
    );
}

export { Registration };
