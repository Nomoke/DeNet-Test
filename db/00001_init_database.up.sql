CREATE TABLE IF NOT EXISTS users (
    id              SERIAL        PRIMARY KEY,
    nickname        VARCHAR(255)  NOT NULL UNIQUE,
    password        VARCHAR(50)   NOT NULL,
    referrer        VARCHAR(255),
    referals        INTEGER       NOT NULL DEFAULT 0,
    points          INTEGER       NOT NULL DEFAULT 0,
    tasks_completed INTEGER       NOT NULL DEFAULT 0,
    created_at      TIMESTAMP     WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP     WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP     WITH TIME ZONE
);

