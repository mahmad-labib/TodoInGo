-- +goose Up
    CREATE TABLE IF NOT EXISTS todo (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        title TEXT UNIQUE NOT NULL,
		status TEXT DEFAULT 'inProgress',
		priority TEXT NOT NULL,
		created_at DATE DEFAULT CURRENT_DATE
    );
-- +goose Down
DROP TABLE todo;
