CREATE TABLE IF NOT EXISTS users
(
    id            character varying PRIMARY KEY,
    email         character varying NOT NULL,
    password_hash character varying NOT NULL,
    updated_at    timestamp WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at    timestamp WITH TIME ZONE DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_active_users ON users (email) WHERE deleted_at IS NULL;

CREATE TABLE IF NOT EXISTS refresh_sessions
(
    id            SERIAL PRIMARY KEY,
    user_id       character varying,
    device_id     character varying NOT NULL,
    refresh_token character varying NOT NULL,
    last_visit_at timestamp WITH TIME ZONE NOT NULL DEFAULT now(),
    expires_at    timestamp WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_active_sessions ON refresh_sessions (user_id, refresh_token);

CREATE TABLE IF NOT EXISTS user_devices
(
    id              character varying PRIMARY KEY,
    user_id         character varying NOT NULL,
    user_agent      character varying NOT NULL,
    ip              character varying NOT NULL,
    detached        boolean NOT NULL,
    latest_login_at timestamp WITH TIME ZONE NOT NULL DEFAULT now(),
    detached_at     timestamp WITH TIME ZONE DEFAULT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_user_device ON user_devices (user_id, user_agent, ip);

CREATE TABLE IF NOT EXISTS lists
(
    id          character varying PRIMARY KEY,
    title       character varying NOT NULL,
    user_id     character varying NOT NULL,
    updated_at  timestamp WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at  timestamp WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_list_user_id ON lists (user_id);

CREATE TABLE IF NOT EXISTS tasks
(
    id          character varying PRIMARY KEY,
    title       character varying NOT NULL,
    description character varying,
    start_date  timestamp WITH TIME ZONE,
    deadline    timestamp WITH TIME ZONE,
    start_time  time WITH TIME ZONE,
    end_time    time WITH TIME ZONE,
    status_id   int NOT NULL,
    list_id     character varying NOT NULL,
    heading_id  character varying,
    user_id     character varying NOT NULL,
    updated_at  timestamp WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at  timestamp WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_task_user_id ON tasks(user_id);


CREATE TABLE IF NOT EXISTS statuses
(
    id          int PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    status_name character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS headings
(
    id         character varying PRIMARY KEY,
    title      character varying NOT NULL,
    list_id    character varying NOT NULL,
    user_id    character varying NOT NULL,
    is_default boolean NOT NULL DEFAULT false,
    updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at timestamp WITH TIME ZONE DEFAULT NULL
);


CREATE TABLE IF NOT EXISTS tags
(
    id         character varying PRIMARY KEY,
    title      character varying NOT NULL,
    user_id    character varying NOT NULL,
    updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at timestamp WITH TIME ZONE DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS tasks_tags
(
    task_id character varying NOT NULL,
    tag_id  character varying NOT NULL,
    CONSTRAINT tasks_tags_pkey PRIMARY KEY (task_id, tag_id)
);

CREATE INDEX IF NOT EXISTS idx_tag_user_id ON tags(user_id);

CREATE TABLE IF NOT EXISTS reminders
(
    id         character varying PRIMARY KEY,
    content    character varying NOT NULL,
    read       boolean NOT NULL,
    task_id    character varying NOT NULL,
    user_id    character varying NOT NULL,
    updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at timestamp WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_remind_task_id ON reminders(task_id);

CREATE TABLE IF NOT EXISTS reminder_settings
(
    id       int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    interval character varying NOT NULL
);

ALTER TABLE refresh_sessions ADD FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE user_devices ADD FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE lists ADD FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE tasks ADD FOREIGN KEY (status_id) REFERENCES statuses (id);
ALTER TABLE tasks ADD FOREIGN KEY (list_id) REFERENCES lists (id);
ALTER TABLE tasks ADD FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE tasks ADD FOREIGN KEY (heading_id) REFERENCES headings(id);
ALTER TABLE headings ADD FOREIGN KEY (list_id) REFERENCES lists(id);
ALTER TABLE headings ADD FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE tags ADD FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE reminders ADD FOREIGN KEY (task_id) REFERENCES tasks(id);
ALTER TABLE reminders ADD FOREIGN KEY (user_id) REFERENCES users(id);
