ATTACH TABLE _ UUID 'ab79875a-3f12-44fa-b9c6-4ff7c2a1b685'
(
    `hostname` LowCardinality(String) COMMENT 'Hostname of the server executing the query.',
    `event_date` Date COMMENT 'Event date.',
    `event_time` DateTime COMMENT 'Event time.',
    `event_time_microseconds` DateTime64(6) COMMENT 'Event time with microseconds precision.',
    `query_id` String COMMENT 'Identifier of the query associated with the background task.',
    `database` LowCardinality(String) COMMENT 'Name of the database.',
    `table` LowCardinality(String) COMMENT 'Name of the table.',
    `table_uuid` UUID COMMENT 'UUID of the table the background task belongs to.',
    `log_name` LowCardinality(String) COMMENT 'Name of the background task.',
    `duration_ms` UInt64 COMMENT 'Duration of the task execution in milliseconds.',
    `error` UInt16 COMMENT 'The error code of the occurred exception.',
    `exception` String COMMENT 'Text message of the occurred error.'
)
ENGINE = MergeTree
PARTITION BY toYYYYMM(event_date)
ORDER BY (event_date, event_time)
SETTINGS index_granularity = 8192
COMMENT 'Contains history of background schedule pool task executions.\n\nIt is safe to truncate or drop this table at any time.'
