-- +goose Up
    CREATE TABLE IF NOT EXISTS todo (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        title TEXT UNIQUE NOT NULL,
		status TEXT DEFAULT 'todo',
		priority TEXT NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

-- +goose Down
DROP TABLE todo;
