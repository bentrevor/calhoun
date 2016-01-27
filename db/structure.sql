CREATE TABLE users (
       id     serial primary key,
       name   varchar(40) NOT NULL CHECK (name <> '')
);

CREATE TABLE photos (
       id          serial primary key,
       user_id     integer references users(id)
);
