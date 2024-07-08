CREATE TABLE time_logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_id INTEGER,
    start_time DATETIME NOT NULL,
    end_time DATETIME,
    FOREIGN KEY (task_id) REFERENCES tasks(id)
);
