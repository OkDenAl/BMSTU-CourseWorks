CREATE TABLE IF NOT EXISTS story_stat
(
    story_id   varchar(100) PRIMARY KEY,
    author_id  varchar(100),
    story_json text,
    count      int,
    created_at timestamp
);