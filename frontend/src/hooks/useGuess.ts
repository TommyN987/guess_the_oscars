import { postGuess } from "../api/guess";

function useGuess(nominationID: number) {
    const submitGuess = async () => {
        try {
            await postGuess(nominationID);
        } catch (error) {
            console.error(error);
        }
    };

    return { submitGuess };
}

export { useGuess };
