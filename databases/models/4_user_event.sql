-- +migrate Up

-- Create Registration table
CREATE TABLE user_events (
    id SERIAL PRIMARY KEY,
    regis_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id INT,
    event_id INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (event_id) REFERENCES events(id)
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
BEFORE UPDATE ON user_events
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE user_events;