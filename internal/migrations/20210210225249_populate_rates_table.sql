-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO
    rates (name, taxi_service, min_price, minute_rate, meter_rate)
VALUES
    ('economy', 50, 150, 10, 0.02),
    ('business', 100, 250, 30, 0.05);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DELETE FROM rates WHERE name IN ('economy', 'business');
-- +goose StatementEnd
