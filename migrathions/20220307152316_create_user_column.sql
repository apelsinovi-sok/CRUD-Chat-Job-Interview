-- +goose Up
-- +goose StatementBegin

CREATE TABLE chats (
    id INT GENERATED ALWAYS AS IDENTITY,
    chat_name VARCHAR(250) UNIQUE,
    chat_author VARCHAR(250) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE messages (
    id INT GENERATED ALWAYS AS IDENTITY,
    chat_name VARCHAR(250) NOT NULL,
    chat_id INT NOT NULL,
    message_author VARCHAR(250) NOT NULL,
    message_text VARCHAR(250),
    PRIMARY KEY(id),
    CONSTRAINT fk_chats FOREIGN KEY(chat_id) REFERENCES chats(id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE messages;
DROP TABLE chats;
-- +goose StatementEnd
