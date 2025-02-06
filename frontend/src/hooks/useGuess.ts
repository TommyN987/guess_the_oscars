import { postGuess } from "../api/guess";
import { Guess, GuessDTO } from "../api/types";

function useGuess() {
    const submitGuess = async (g: Guess) => {
        const guess: GuessDTO = {
            nomination_id: g.nominationID,
            category_id: g.categoryID,
        };

        await postGuess(guess);
    };

    return { submitGuess };
}

export { useGuess };
