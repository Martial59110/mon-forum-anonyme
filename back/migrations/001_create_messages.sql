CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    pseudonym VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    avatar VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_messages_created_at ON messages(created_at DESC);

