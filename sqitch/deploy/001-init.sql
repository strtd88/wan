-- sqitch/deploy/001-init.sql
CREATE TABLE projects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    description TEXT,
    dir TEXT,
    status TEXT DEFAULT 'active' CHECK(status IN ('active', 'paused', 'complete')),
    category TEXT DEFAULT 'devtools',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
