CREATE DATABASE brickholderdb;

DROP table series

CREATE TABLE series (
    series_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    parent_id INTEGER REFERENCES series (series_id),
    rebrickable_id INTEGER UNIQUE NOT NULL,
    rebrickable_parent_id INTEGER
);
