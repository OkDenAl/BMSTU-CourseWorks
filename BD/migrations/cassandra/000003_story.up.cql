CREATE TABLE IF NOT EXISTS story (
                                          story_id uuid,
                                          author_id uuid,
                                          story_json text,
                                          created_at timestamp,
                                          PRIMARY KEY (story_id))
    WITH COMPRESSION = { 'class' : 'ZstdCompressor' }
     AND COMPACTION = { 'class' : 'LeveledCompactionStrategy' };