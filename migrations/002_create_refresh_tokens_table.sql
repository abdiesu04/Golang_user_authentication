CREATE TABLE refresh_tokens (
    user_id BIGINT PRIMARY KEY,
    token TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
