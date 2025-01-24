CREATE OR REPLACE FUNCTION insert_nominee_people(
    category_name TEXT,
    nominations_data JSONB
) RETURNS VOID AS $$
DECLARE
    nomination_id INT;
    person_name TEXT;
    person_id INT;
    movie_title TEXT;
    people JSONB;
BEGIN
    -- Iterate through the array of nominations
    FOR movie_title, people IN
        SELECT key, value FROM jsonb_each(nominations_data)
    LOOP
        -- Find the nomination ID for the movie and category
        SELECT n.id INTO nomination_id
        FROM nominations n
        JOIN categories c ON n.category_id = c.id
        JOIN movies m ON n.movie_id = m.id
        WHERE c.name = category_name AND m.title = movie_title;

        IF nomination_id IS NULL THEN
            RAISE EXCEPTION 'Nomination not found for category "%" and movie "%"', category_name, movie_title;
        END IF;

        -- Insert each person associated with the movie
        FOR person_name IN SELECT * FROM jsonb_array_elements_text(people) LOOP
            -- Find the person ID
            SELECT id INTO person_id FROM people WHERE name = person_name;

            IF person_id IS NULL THEN
                RAISE EXCEPTION 'Person "%" does not exist', person_name;
            END IF;

            -- Insert into nominee_people
            INSERT INTO nominee_people (nomination_id, person_id)
            VALUES (nomination_id, person_id);
        END LOOP;
    END LOOP;
END;
$$ LANGUAGE plpgsql;

