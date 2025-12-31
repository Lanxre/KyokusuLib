-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_activities (
    id            SERIAL PRIMARY KEY ,
    user_id       INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    activity_type VARCHAR(50) NOT NULL,
    target_id     INTEGER,
    metadata      JSONB,
    timestamp     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_user_activities_user_id ON user_activities(user_id);
CREATE INDEX idx_user_activities_type ON user_activities(activity_type);
CREATE INDEX idx_user_activities_timestamp ON user_activities(timestamp DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_activities;
-- +goose StatementEnd
