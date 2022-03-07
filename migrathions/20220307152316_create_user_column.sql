-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id UUID NOT NULL,
    name VARCHAR(250) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE chats (
    id UUID NOT NULL,
    chat_name VARCHAR(250) UNIQUE,
    author_id UUID NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_author FOREIGN KEY(author_id) REFERENCES users(id) ON DELETE CASCADE
);


CREATE TABLE messages (
    id INT GENERATED ALWAYS AS IDENTITY,
    chat_id UUID NOT NULL,
    author_id UUID NOT NULL,
    message_text VARCHAR(250),
    PRIMARY KEY(id),
    CONSTRAINT fk_chats FOREIGN KEY(chat_id) REFERENCES chats(id) ON DELETE CASCADE,
    CONSTRAINT fk_author FOREIGN KEY(author_id) REFERENCES users(id) ON DELETE CASCADE
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE messages;
DROP TABLE chats;
DROP TABLE users;
-- +goose StatementEnd
