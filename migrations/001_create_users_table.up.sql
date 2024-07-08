CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    passport_number TEXT NOT NULL UNIQUE,
    passport_serie TEXT NOT NULL UNIQUE,
    name TEXT,
    surname TEXT,
    patronymic TEXT,
    address TEXT
);
