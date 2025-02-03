import { useCallback, useRef, useState } from "react";
import { useAuth } from "../hooks/useAuth";

enum Tab {
    Login,
    Register,
}

function Registration() {
    const [activeTab, setActiveTab] = useState<Tab>(Tab.Login);

    const { register, login } = useAuth();

    const nameRef = useRef<HTMLInputElement>(null);
    const emailRef = useRef<HTMLInputElement>(null);
    const passwordRef = useRef<HTMLInputElement>(null);
    const repeatPasswordRef = useRef<HTMLInputElement>(null);

    const handleSubmit = useCallback(() => {
        if (activeTab === Tab.Login) {
            console.log({
                email: emailRef.current?.value,
                password: passwordRef.current?.value,
            });
        } else {
            console.log({
                name: nameRef.current?.value,
                pass: repeatPasswordRef.current?.value,
            });
        }
    }, [activeTab]);

    return (
        <div className="flex flex-col h-1/2 w-1/2 bg-zinc-800 border-4 border-white p-8 z-10">
            <ul className="flex justify-evenly">
                <li
                    className={`${activeTab === Tab.Login && "border-b-2 border-white"} cursor-pointer`}
                    onClick={() => setActiveTab(Tab.Login)}
                >
                    Login
                </li>
                <li
                    className={`${activeTab === Tab.Register && "border-b-2 border-white"} cursor-pointer`}
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
                            className="border-1 border-white rounded-xl p-2 outline-none"
                            type="text"
                            ref={nameRef}
                        />
                    </div>
                )}
                <div className="flex flex-col">
                    <p>Email:</p>
                    <input
                        className="border-1 border-white rounded-xl p-2 outline-none"
                        type="email"
                        ref={emailRef}
                    />
                </div>
                <div className="flex flex-col">
                    <p>Password:</p>
                    <input
                        className="border-1 border-white rounded-xl p-2 outline-none"
                        type="password"
                        ref={passwordRef}
                    />
                </div>
                {activeTab === Tab.Register && (
                    <div className="flex flex-col">
                        <p>Repeat Password:</p>
                        <input
                            className="border-1 border-white rounded-xl p-2 outline-none"
                            type="password"
                            ref={repeatPasswordRef}
                        />
                    </div>
                )}
                <button onClick={handleSubmit}>Submit</button>
            </div>
        </div>
    );
}

export { Registration };
