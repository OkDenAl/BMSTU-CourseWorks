CREATE KEYSPACE IF NOT EXISTS story_stat
WITH replication = {
    'class': 'NetworkTopologyStrategy',
    'replication_factor': 3
}
AND durable_writes = true;