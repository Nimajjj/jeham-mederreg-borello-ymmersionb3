CREATE DATABASE PetitsRochersGrosCailloux;
USE PetitsRochersGrosCailloux;

CREATE TABLE IF NOT EXISTS categories
(
    id    INT AUTO_INCREMENT
        PRIMARY KEY,
    title VARCHAR(255) NULL
);

CREATE TABLE IF NOT EXISTS pebbles
(
    id          INT AUTO_INCREMENT
        PRIMARY KEY,
    title       VARCHAR(255) NULL,
    description MEDIUMTEXT   NULL,
    price       FLOAT        NULL,
    creation    DATE         NULL,
    quantity    INT          NULL,
    breed       VARCHAR(255) NULL,
    weight      FLOAT        NULL
);

CREATE TABLE IF NOT EXISTS pebbles_categories
(
    id_pebble    INT NULL,
    id_categorie INT NULL,
    CONSTRAINT pebbles_categories_ibfk_1
        FOREIGN KEY (id_pebble) REFERENCES pebbles (id),
    CONSTRAINT pebbles_categories_ibfk_2
        FOREIGN KEY (id_categorie) REFERENCES categories (id)
);


CREATE TABLE IF NOT EXISTS photos
(
    id       INT AUTO_INCREMENT
        PRIMARY KEY,
    filepath VARCHAR(255) NULL
);

CREATE TABLE IF NOT EXISTS pebbles_photos
(
    id_pebble INT NULL,
    id_photo  INT NULL,
    CONSTRAINT pebbles_photos_ibfk_1
        FOREIGN KEY (id_pebble) REFERENCES pebbles (id),
    CONSTRAINT pebbles_photos_ibfk_2
        FOREIGN KEY (id_photo) REFERENCES photos (id)
);

CREATE TABLE IF NOT EXISTS users
(
    id           INT AUTO_INCREMENT
        PRIMARY KEY,
    email        VARCHAR(255) NULL,
    encryptedpwd VARCHAR(255) NULL
);

CREATE TABLE IF NOT EXISTS cart
(
    id         INT AUTO_INCREMENT
        PRIMARY KEY,
    fk_id_user INT NULL,
    CONSTRAINT cart_ibfk_1
        FOREIGN KEY (fk_id_user) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS pebbles_cart
(
    id_caillou INT NULL,
    id_basket  INT NULL,
    quantity   INT NULL,
    CONSTRAINT pebbles_cart_ibfk_1
        FOREIGN KEY (id_caillou) REFERENCES pebbles (id),
    CONSTRAINT pebbles_cart_ibfk_2
        FOREIGN KEY (id_basket) REFERENCES cart (id)
);

INSERT INTO users (email, encryptedpwd) VALUES ('test', 'test');
INSERT INTO cart (fk_id_user) VALUES (1);
