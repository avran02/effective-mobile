CREATE TABLE time_logs (
    id SERIAL PRIMARY KEY,
    task_id INTEGER,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ,
    FOREIGN KEY (task_id) REFERENCES tasks(id)
);
