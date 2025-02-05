async function postGuess(nominationID: number) {
    const resp = await fetch(`${import.meta.env.VITE_API_BASE_URL}/p/guesses`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({ nomination_id: nominationID }),
    });

    if (!resp.ok) {
        throw new Error("Failed to post guess.");
    }
}

export { postGuess };
