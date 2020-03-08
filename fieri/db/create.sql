CREATE TABLE task (
    id SERIAL PRIMARY KEY,
    gid varchar(32),
    todo text,
    created timestamp,
    UNIQUE(gid)
);
