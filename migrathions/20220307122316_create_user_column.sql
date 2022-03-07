-- +goose Up
-- +goose StatementBegin


CREATE TABLE chats (
    id INT GENERATED ALWAYS AS IDENTITY,
    chat_name VARCHAR(250),
    chat_author VARCHAR(250),
    PRIMARY KEY(id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE chats;

-- +goose StatementEnd
