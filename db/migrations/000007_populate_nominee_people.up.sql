SELECT insert_nominee_people(
    'Best Director',
    '{
        "Anora": ["Sean Baker"],
        "A Complete Unknown": ["James Mangold"],
        "Emilia Pérez": ["Jacques Audiard"],
        "The Brutalist": ["Brady Corbet"],
        "The Substance": ["Coralie Fargeat"]
    }'::JSONB
);

SELECT insert_nominee_people(
    'Best Actress',
    '{
        "Wicked": ["Cynthia Erivo"],
        "Emilia Pérez": ["Karla Sofía Gascón"],
        "Anora": ["Mikey Madison"],
        "The Substance": ["Demi Moore"],
        "I''m Still Here": ["Fernanda Torres"]
    }'::JSONB
);

SELECT insert_nominee_people(
    'Best Actor',
    '{
        "The Brutalist": ["Adrien Brody"],
        "A Complete Unknown": ["Timothée Chalamet"],
        "Sing Sing": ["Colman Domingo"],
        "Conclave": ["Ralph Fiennes"],
        "The Apprentice": ["Sebastian Stan"]
    }'::JSONB
);

SELECT insert_nominee_people(
    'Best Supporting Actress',
    '{
        "A Complete Unknown": ["Monica Barbaro"],
        "Wicked": ["Ariana Grande"],
        "The Brutalist": ["Felicity Jones"],
        "Conclave": ["Isabella Rossellini"],
        "Emilia Pérez": ["Zoe Saldaña"]
    }'::JSONB
);

SELECT insert_nominee_people(
    'Best Supporting Actor',
    '{
        "Anora": ["Yura Borisov"],
        "A Real Pain": ["Kieran Culkin"],
        "A Complete Unknown": ["Edward Norton"],
        "The Brutalist": ["Guy Pearce"],
        "The Apprentice": ["Jeremy Strong"]
    }'::JSONB
);

SELECT insert_nominee_people(
    'Best Original Screenplay',
    '{
        "Anora": ["Sean Baker"],
        "The Brutalist": ["Brady Corbet", "Mona Fastvold"],
        "A Real Pain": ["Jesse Eisenberg"],
        "September 5": ["Tim Fehlbaum", "Moritz Binder"],
        "The Substance": ["Coralie Fargeat"]
    }'::JSONB
);

SELECT insert_nominee_people(
    'Best Adapted Screenplay',
    '{
        "A Complete Unknown": ["Jay Cocks", "James Mangold"],
        "Conclave": ["Peter Straughan"],
        "Emilia Pérez": ["Jacques Audiard", "Thomas Bidegain", "Nicolas Livecchi"],
        "Nickel Boys": ["RaMell Ross", "Joslyn Barnes"],
        "Sing Sing": ["Clint Bentley", "Greg Kwedar", "Clarence Maclin", "John ''Divine G'' Whitfield"]
    }'::JSONB
);

SELECT insert_nominee_people(
    'Best Original Score',
    '{
        "The Brutalist": ["Daniel Blumberg"],
        "Conclave": ["Volker Bertelmann"],
        "Emilia Pérez": ["Clément Ducol", "Camille"],
        "Wicked": ["John Powell", "Stephen Schwartz"],
        "The Wild Robot": ["Kris Bowers"]
    }'::JSONB
);

SELECT insert_nominee_people(
    'Best Original Song',
    '{
        "Elton John: Never Too Late": ["Elton John", "Brandi Carlile"],
        "Sing Sing": ["Adrian Quesada", "Abraham Alexander"],
        "The Six Triple Eight": ["Diane Warren"]
    }'::JSONB
);

-- El Mal (Emilia Pérez)
INSERT INTO nominee_people (nomination_id, person_id)
VALUES
((SELECT id FROM nominations WHERE category_id = (SELECT id FROM categories WHERE name = 'Best Original Song')
    AND movie_id = (SELECT id FROM movies WHERE title = 'Emilia Pérez') AND additional_info = 'El Mal'),
 (SELECT id FROM people WHERE name = 'Clément Ducol')),
((SELECT id FROM nominations WHERE category_id = (SELECT id FROM categories WHERE name = 'Best Original Song')
    AND movie_id = (SELECT id FROM movies WHERE title = 'Emilia Pérez') AND additional_info = 'El Mal'),
 (SELECT id FROM people WHERE name = 'Camille')),
((SELECT id FROM nominations WHERE category_id = (SELECT id FROM categories WHERE name = 'Best Original Song')
    AND movie_id = (SELECT id FROM movies WHERE title = 'Emilia Pérez') AND additional_info = 'El Mal'),
 (SELECT id FROM people WHERE name = 'Jacques Audiard'));

-- Mi Camino (Emilia Pérez)
INSERT INTO nominee_people (nomination_id, person_id)
VALUES
((SELECT id FROM nominations WHERE category_id = (SELECT id FROM categories WHERE name = 'Best Original Song')
    AND movie_id = (SELECT id FROM movies WHERE title = 'Emilia Pérez') AND additional_info = 'Mi Camino'),
 (SELECT id FROM people WHERE name = 'Clément Ducol')),
((SELECT id FROM nominations WHERE category_id = (SELECT id FROM categories WHERE name = 'Best Original Song')
    AND movie_id = (SELECT id FROM movies WHERE title = 'Emilia Pérez') AND additional_info = 'Mi Camino'),
 (SELECT id FROM people WHERE name = 'Camille'));
