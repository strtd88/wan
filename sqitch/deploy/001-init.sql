-- sqitch/deploy/001-init.sql
CREATE TABLE projects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    description TEXT,
    dir TEXT,
    status TEXT DEFAULT 'new' CHECK (status IN ('new', 'active', 'paused', 'complete')),
    category TEXT DEFAULT 'devtools',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    business_value_short TEXT,
    business_value_long TEXT,
    business_value_achieved TEXT
);
