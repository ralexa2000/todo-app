-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.tasks (
    username varchar(255) NOT NULL CONSTRAINT username_right CHECK(username ~ '^[A-Za-z0-9_-\.]*$'),
    task_id int NOT NULL,
    task text NOT NULL,
    due_date date,
    PRIMARY KEY(username, task_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.tasks;
-- +goose StatementEnd
