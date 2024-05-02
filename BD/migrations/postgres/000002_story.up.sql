CREATE TABLE IF NOT EXISTS story (
    story_id   varchar(100) PRIMARY KEY,
    author_id  varchar(100),
    story_json text,
    created_at timestamp
);