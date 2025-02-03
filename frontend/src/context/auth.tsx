import React, { createContext, useEffect, useState } from "react";
import { User } from "../api/types";
import { getUser } from "../api/auth";

type AuthContextType = {
    user: User | null;
    setUser: (user: User) => void;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
    children,
}) => {
    const [user, setUser] = useState<User | null>(null);

    useEffect(() => {
        const checkAuth = async () => {
            try {
                const user = await getUser();
                setUser(user);
            } catch {
                setUser(null);
            }
        };
        checkAuth();
    }, []);

    return (
        <AuthContext.Provider value={{ user, setUser }}>
            {children}
        </AuthContext.Provider>
    );
};

export { AuthContext, AuthProvider };
