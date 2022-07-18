CREATE TABLE snippets
(
    id      INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title   VARCHAR(100) NOT NULL,
    content TEXT         NOT NULL,
    created DATETIME     NOT NULL,
    expires DATETIME     NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

-- Note we change the name and email fields from 255 to 100 because of MySQL ERROR 1071
-- https://stackoverflow.com/questions/1814532/mysql-error-1071-specified-key-was-too-long-max-key-length-is-767-bytes

CREATE TABLE users
(
    id              INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name            VARCHAR(100) NOT NULL,
    email           VARCHAR(100) NOT NULL,
    hashed_password CHAR(60)     NOT NULL,
    created         DATETIME     NOT NULL,
    active          BOOLEAN      NOT NULL DEFAULT TRUE
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

INSERT INTO users (name, email, hashed_password, created) VALUES ('Alice Jones', 'alice@example.com', '$2a$12$J3/y47mdRvuYH4TcRYjowOmOx2e7lillIbMc1vdFl0j8kIw2jQMOe', '2018-12-23 17:25:22');