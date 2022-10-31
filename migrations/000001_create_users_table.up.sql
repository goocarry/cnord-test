CREATE TABLE IF NOT EXISTS users(
   user_id serial PRIMARY KEY,
   firstname VARCHAR (50) NOT NULL,
   lastname VARCHAR (50) NOT NULL,
    UNIQUE (firstname, lastname)
);
