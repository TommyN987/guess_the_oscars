SELECT insert_nominations(
    'Best Picture',
    ARRAY[
        'Anora',
        'The Brutalist',
        'A Complete Unknown',
        'Conclave',
        'Dune: Part Two',
        'Emilia Pérez',
        'I''m Still Here',
        'Nickel Boys',
        'The Substance',
        'Wicked'
    ]
);

SELECT insert_nominations(
    'Best Director',
    ARRAY[
        'Anora',
        'The Brutalist',
        'A Complete Unknown',
        'Emilia Pérez',
        'The Substance'
    ]
);

SELECT insert_nominations(
    'Best Actress',
    ARRAY[
        'Wicked',
        'Emilia Pérez',
        'Anora',
        'The Substance',
        'I''m Still Here'
    ]
);

SELECT insert_nominations(
    'Best Actor',
    ARRAY[
        'The Brutalist',
        'A Complete Unknown',
        'Sing Sing',
        'Conclave',
        'The Apprentice'
    ]
);

SELECT insert_nominations(
    'Best Supporting Actress',
    ARRAY[
        'The Brutalist',
        'A Complete Unknown',
        'Wicked',
        'Conclave',
        'Emilia Pérez'
    ]
);

SELECT insert_nominations(
    'Best Supporting Actor',
    ARRAY[
        'The Brutalist',
        'A Complete Unknown',
        'Anora',
        'A Real Pain',
        'The Apprentice'
    ]
);

SELECT insert_nominations(
    'Best Original Screenplay',
    ARRAY[
        'Anora',
        'The Brutalist',
        'A Real Pain',
        'September 5',
        'The Substance'
    ]
);

SELECT insert_nominations(
    'Best Adapted Screenplay',
    ARRAY[
        'Conclave',
        'Nickel Boys',
        'A Complete Unknown',
        'Emilia Pérez',
        'Sing Sing'
    ]
);

SELECT insert_nominations(
    'Best International Feature',
    ARRAY[
        'I''m Still Here',
        'The Girl with the Needle',
        'The Seed of the Sacred Fig',
        'Emilia Pérez',
        'Flow'
    ]
);

SELECT insert_nominations(
    'Best Documentary Feature',
    ARRAY[
        'Black Box Diaries',
        'No Other Land',
        'Porcelain War',
        'Soundtrack to a Coup d''Etat',
        'Sugarcane'
    ]
);

SELECT insert_nominations(
    'Best Animated Feature',
    ARRAY[
        'Inside Out 2',
        'Memoir of a Snail',
        'Wallace & Gromit: Vengeance Most Fowl',
        'The Wild Robot',
        'Flow'
    ]
);

SELECT insert_nominations(
    'Best Cinematography',
    ARRAY[
        'The Brutalist',
        'Dune: Part Two',
        'Emilia Pérez',
        'Maria',
        'Nosferatu'
    ]
);

SELECT insert_nominations(
    'Best Costume Design',
    ARRAY[
        'A Complete Unknown',
        'Conclave',
        'Gladiator II',
        'Wicked',
        'Nosferatu'
    ]
);

SELECT insert_nominations(
    'Best Film Editing',
    ARRAY[
        'Anora',
        'The Brutalist',
        'Emilia Pérez',
        'Conclave',
        'Wicked'
    ]
);

SELECT insert_nominations(
    'Best Production Design',
    ARRAY[
        'The Brutalist',
        'Dune: Part Two',
        'Conclave',
        'Wicked',
        'Nosferatu'
    ]
);

SELECT insert_nominations(
    'Best Original Score',
    ARRAY[
        'The Brutalist',
        'Conclave',
        'Emilia Pérez',
        'The Wild Robot',
        'Wicked'
    ]
);

SELECT insert_nominations(
    'Best Sound',
    ARRAY[
        'A Complete Unknown',
        'Dune: Part Two',
        'Emilia Pérez',
        'The Wild Robot',
        'Wicked'
    ]
);

SELECT insert_nominations(
    'Best Visual Effects',
    ARRAY[
        'Alien: Romulus',
        'Better Man',
        'Dune: Part Two',
        'Kingdom of the Planet of the Apes',
        'Wicked'
    ]
);

SELECT insert_nominations(
    'Best Makeup and Hairstyling',
    ARRAY[
        'A Different Man',
        'Emilia Pérez',
        'The Substance',
        'Wicked',
        'Nosferatu'
    ]
);

SELECT insert_nominations(
    'Best Documentary Short',
    ARRAY[
        'Death by Numbers',
        'I Am Ready, Warden',
        'Incident',
        'Instruments of a Beating Heart',
        'The Only Girl in the Orchestra'
    ]
);

SELECT insert_nominations(
    'Best Animated Short',
    ARRAY[
        'Beautiful Men',
        'In the Shadow of the Cypress',
        'Magic Candies',
        'Wander to Wonder',
        'Yuck!'
    ]
);

SELECT insert_nominations(
    'Best Live Action Short',
    ARRAY[
        'A Lien',
        'Anuja',
        'I''m Not a Robot',
        'The Last Ranger',
        'The Man Who Could Not Remain Silent'
    ]
);

INSERT INTO nominations (category_id, movie_id, additional_info)
VALUES
((SELECT id FROM categories WHERE name = 'Best Original Song'), (SELECT id FROM movies WHERE title = 'Emilia Pérez'), 'El Mal');

INSERT INTO nominations (category_id, movie_id, additional_info)
VALUES
((SELECT id FROM categories WHERE name = 'Best Original Song'), (SELECT id FROM movies WHERE title = 'Emilia Pérez'), 'Mi Camino');

INSERT INTO nominations (category_id, movie_id, additional_info)
VALUES
((SELECT id FROM categories WHERE name = 'Best Original Song'), (SELECT id FROM movies WHERE title = 'Elton John: Never Too Late'), 'Never Too Late');

INSERT INTO nominations (category_id, movie_id, additional_info)
VALUES
((SELECT id FROM categories WHERE name = 'Best Original Song'), (SELECT id FROM movies WHERE title = 'Sing Sing'), 'Like a Bird');

INSERT INTO nominations (category_id, movie_id, additional_info)
VALUES
((SELECT id FROM categories WHERE name = 'Best Original Song'), (SELECT id FROM movies WHERE title = 'The Six Triple Eight'), 'The Journey');
