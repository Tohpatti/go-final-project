-- +migrate Up

-- Create EventCategory table
CREATE TABLE event_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;

-- Create trigger to update updated_at timestamp
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON event_categories
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE event_categories;