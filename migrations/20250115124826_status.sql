-- +goose Up
-- +goose StatementBegin
CREATE TABLE status (
	id_status SERIAL PRIMARY KEY,
	nama_status VARCHAR(255) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE status;
-- +goose StatementEnd
