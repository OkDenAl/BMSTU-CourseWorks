CREATE TABLE IF NOT EXISTS story_views_stat (
      story_id uuid,
      count counter,
      PRIMARY KEY (story_id)
)