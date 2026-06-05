
CREATE TABLE refresh_tokens(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    token varchar(255) NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    revoked_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
)