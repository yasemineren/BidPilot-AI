ATTACH TABLE _ UUID 'a7fafb61-5d25-488a-a37a-173a49373b58'
(
    `event_id` String,
    `ts` DateTime,
    `bidder_id` String,
    `geo` LowCardinality(String),
    `bid_price` Float32,
    `won` UInt8
)
ENGINE = MergeTree
ORDER BY (ts, bidder_id)
SETTINGS index_granularity = 8192
