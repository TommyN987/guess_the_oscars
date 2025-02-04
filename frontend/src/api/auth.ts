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

async function login(email: string, password: string) {
    const resp = await fetch(`${import.meta.env.VITE_API_BASE_URL}/login`, {
        method: "POST",
        body: JSON.stringify({ email, password }),
    });

    if (!resp.ok) {
        throw new Error("Invalid email or password.");
    }
}

async function logout() {
    fetch(`${import.meta.env.VITE_API_BASE_URL}/p/logout`, {
        method: "POST",
    });
}

export { getUser, register, login, logout };
