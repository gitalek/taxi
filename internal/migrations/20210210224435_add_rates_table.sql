-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE rates (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    taxi_service FLOAT NOT NULL,
    min_price FLOAT NOT NULL,
    minute_rate FLOAT NOT NULL,
    meter_rate FLOAT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE rates;
-- +goose StatementEnd
