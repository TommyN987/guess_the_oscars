import { User } from "./types";

async function getUser(): Promise<User> {
    const resp = await fetch(`${import.meta.env.VITE_API_BASE_URL}/p/me`, {
        credentials: "include",
    });

    if (!resp.ok) {
        throw new Error("User not logged in.");
    }

    return resp.json();
}

async function register(name: string, email: string, password: string) {
    const resp = await fetch(`${import.meta.env.VITE_API_BASE_URL}/register`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify({
            name,
            email,
            password,
        }),
    });

    if (!resp.ok) {
        throw new Error("Failed to register.");
    }
}

async function validate(token: string): Promise<User> {
    try {
        const resp = await fetch(
            `${import.meta.env.VITE_API_BASE_URL}/validate?token=${token}`,
            {
                method: "GET",
                headers: { "Content-Type": "application/json" },
                credentials: "include",
            },
        );

        const user: User = await resp.json();
        return user;
    } catch (error) {
        throw new Error(`Validation error: ${error}`);
    }
}

async function login(email: string, password: string): Promise<User> {
    const resp = await fetch(`${import.meta.env.VITE_API_BASE_URL}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({ email, password }),
    });

    if (!resp.ok) {
        throw new Error("Invalid email or password.");
    } else {
        return resp.json();
    }
}

async function logout() {
    fetch(`${import.meta.env.VITE_API_BASE_URL}/p/logout`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
    });
}

export { getUser, register, validate, login, logout };
