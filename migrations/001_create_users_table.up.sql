CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    passport_number TEXT NOT NULL UNIQUE,
    name TEXT,
    surname TEXT,
    patronymic TEXT,
    address TEXT
);
