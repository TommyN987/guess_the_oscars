CREATE OR REPLACE FUNCTION set_winner(
    category_name TEXT,
    movie_title TEXT
) RETURNS VOID AS $$
DECLARE
    winner_nomination_id INT;
BEGIN
    -- Find the nomination ID for the winner
    SELECT n.id INTO winner_nomination_id
    FROM nominations n
    JOIN categories c ON n.category_id = c.id
    JOIN movies m ON n.movie_id = m.id
    WHERE c.name = category_name AND m.title = movie_title;

    IF winner_nomination_id IS NULL THEN
        RAISE EXCEPTION 'Nomination not found for category "%" and movie "%"', category_name, movie_title;
    END IF;

    -- Update the categories table to set the winner
    UPDATE categories
    SET winner_id = winner_nomination_id
    WHERE name = category_name;

    -- Update the guesses table
    UPDATE guesses
    SET is_correct = (nomination_id = winner_nomination_id)
    WHERE category_id = (SELECT id FROM categories WHERE name = category_name);
END;
$$ LANGUAGE plpgsql;
