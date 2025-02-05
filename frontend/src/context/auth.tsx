import React, { createContext, useEffect, useState } from "react";
import { User } from "../api/types";
import { getUser } from "../api/auth";

type AuthContextType = {
    user: User | null;
    setUser: (user: User | null) => void;
    loading: boolean; // Add loading state
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
    children,
}) => {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true); // Initialize as loading

    useEffect(() => {
        const checkAuth = async () => {
            try {
                const user = await getUser();
                setUser(user);
            } catch (error) {
                console.error(error);
                setUser(null);
            } finally {
                setLoading(false); // Done loading
            }
        };
        checkAuth();
    }, []);

    return (
        <AuthContext.Provider value={{ user, setUser, loading }}>
            {children}
        </AuthContext.Provider>
    );
};

export { AuthContext, AuthProvider };
