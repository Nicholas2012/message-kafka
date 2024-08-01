-- +goose Up
-- +goose StatementBegin
CREATE TABLE messages (
                        id SERIAL PRIMARY KEY,
                        message VARCHAR NOT NULL,
                        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                        sent_at TIMESTAMPTZ
);
CREATE INDEX messages_sent_at_idx ON messages(sent_at) WHERE sent_at IS NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table messages;
-- +goose StatementEnd
