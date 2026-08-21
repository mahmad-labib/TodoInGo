-- +goose Up
SELECT 'up SQL query';
    CREATE TABLE IF NOT EXISTS todo (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		created_at DATE DEFAULT CURRENT_DATE,
        title TEXT UNIQUE NOT NULL,
		status TEXT DEFAULT 'inProgress',
		priority TEXT NOT NULL CHECK(LOWER(priority) IN ('high','normal','low'))
    );
-- +goose Down
DROP TABLE todo;
