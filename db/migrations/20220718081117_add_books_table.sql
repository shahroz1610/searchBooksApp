-- migrate:up
CREATE TABLE books(
    id UUID PRIMARY KEY, 
    name text UNIQUE
);

-- migrate:down
DROP table books;
