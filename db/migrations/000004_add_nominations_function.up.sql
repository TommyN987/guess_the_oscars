CREATE OR REPLACE FUNCTION insert_nominations(category_name TEXT, movies TEXT[]) RETURNS VOID AS $$
DECLARE
    category_id INT;
    movie_title TEXT;
    movie_id INT;
BEGIN
    -- Find the category ID
    SELECT id INTO category_id FROM categories WHERE name = category_name;
    IF category_id IS NULL THEN
        RAISE EXCEPTION 'Category "%" does not exist', category_name;
    END IF;

    -- Loop through each movie title and insert into nominations
    FOREACH movie_title IN ARRAY movies LOOP
        -- Find the movie ID
        SELECT id INTO movie_id FROM movies WHERE title = movie_title;
        IF movie_id IS NULL THEN
            RAISE EXCEPTION 'Movie "%" does not exist', movie_title;
        END IF;

        -- Insert into nominations
        INSERT INTO nominations (category_id, movie_id) VALUES (category_id, movie_id);
    END LOOP;
END;
$$ LANGUAGE plpgsql;
