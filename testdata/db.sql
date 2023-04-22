CREATE TABLE IF NOT EXISTS users
(
    id         BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name  VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL,
    phone      VARCHAR(255)
);

INSERT INTO users (first_name, last_name, email, phone)
VALUES ('Alice', 'Foo', 'alice@example.com', '000-000-0000'),
       ('Bob', 'Foo', 'bob@example.com', NULL);

CREATE TABLE IF NOT EXISTS messages
(
    id           BIGSERIAL PRIMARY KEY,
    sender_id    VARCHAR(64) NOT NULL,
    recipient_id VARCHAR(64) NOT NULL,
    title        TEXT        NOT NULL,
    body         TEXT        NOT NULL
);
