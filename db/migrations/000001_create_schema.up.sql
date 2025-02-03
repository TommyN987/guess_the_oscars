CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    country VARCHAR(100)
);

CREATE TABLE people (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE nominations (
    id SERIAL PRIMARY KEY,
    category_id INT NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    movie_id INT NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    additional_info TEXT
);

CREATE TABLE nominee_people (
    id SERIAL PRIMARY KEY,
    nomination_id INT NOT NULL REFERENCES nominations(id) ON DELETE CASCADE,
    person_id INT NOT NULL REFERENCES people(id) ON DELETE CASCADE
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    email_confirmed BOOLEAN DEFAULT FALSE,
    password_hash VARCHAR(255) NOT NULL,
    confirmation_token VARCHAR(255),
    confirmation_expires TIMESTAMP
);

CREATE TABLE guesses (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    nomination_id INT NOT NULL REFERENCES nominations(id) ON DELETE CASCADE
);
