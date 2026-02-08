# Architecture

## Sistem Bileşenleri

### A) RTB Metrics Generator (Simülatör)
- Event tipleri: `bid_request`, `bid_response`, `win`, `click`, `revenue`, `timeout`, `fraud_signal`
- Segmentler: `geo`, `device`, `publisher`, `ad_format`, `hour_of_day`, `bidder_id`
- **Incident Injector**: timeout spike, win-rate drop, revenue drop, fraud-like burst, config mismatch

### B) Ingest Service (Go)
- 1s/10s/1m pencerelerde metrik agregasyonu
- ClickHouse write, düşük latency bütçesi

### C) Feature & Analytics (ClickHouse)
- `rtb_events_raw`, `rtb_metrics_1m`, `baseline_profiles`
- Sorgular: “son 5 dk vs baseline”, “top regressions”, “pareto segments”

### D) Decision Engine (Python FastAPI / Go)
- **Rule Guardrails**: extreme timeout/latency → throttle
- **ML Layer**: anomaly score + incident type + impact + action selector

### E) Control Plane (Postgres)
- `config_rules`, `actions`, `action_outcomes`
- **Approval workflow**: semi-automated mod

### F) UI / Ops Dashboard
- Live metrics
- Anomali timeline
- Önerilen aksiyonlar kuyruğu
- One-click apply/rollback

### G) Observability
- Prometheus metrics + Grafana
- Structured JSON logs

### H) Kubernetes Deployment
- Deployment + HPA
- Canary rollout
- ConfigMap/Secret yönetimi

## Veri Akışı
1. Simülatör event üretir
2. Ingest servisinde aggregation + ClickHouse write
3. Decision engine ClickHouse’dan metrikleri okur
4. Guardrail + ML katmanları aksiyon önerir
5. Control plane aksiyonu uygular ve loglar
6. UI ve Grafana üzerinden izlenir

## Veritabanı Tasarımı (ClickHouse + Postgres)

### ClickHouse (Time-series & Analytics)
**`rtb_events_raw`**
- `ts`, `request_id`, `bidder_id`, `geo`, `device`, `publisher`, `ad_format`
- `bid_price`, `won` (bool), `revenue`, `timeout` (bool), `fraud_signal` (float)

**`rtb_metrics_1m`**
- `ts_minute`
- `segment_key` (ör: `geo|device|publisher|format|bidder`)
- `qps`, `bids`, `wins`, `win_rate`
- `revenue`, `rpm`, `cpm`
- `timeouts`, `timeout_rate`
- `fraud_score_avg`, `fraud_spike_count`

**`baseline_profiles`**
- `segment_key`, `hour_of_day`, `day_of_week`
- `expected_mean`, `expected_std` (win_rate, rpm, timeout_rate vb.)
- Rolling quantiles: `p50`, `p90`, `p99`

### Postgres (Control + Audit)
**`config_rules`**
- `rule_id`, `rule_type`, `segment_filter`, `params` (jsonb)
- `enabled`, `version`

**`actions`**
- `action_id`, `ts`, `segment_key`
- `action_type` (THROTTLE, WEIGHT, CONFIG_SUGGEST)
- `proposed_value`, `applied_value`
- `mode` (AUTO, APPROVED, MANUAL)
- `status` (PROPOSED, APPLIED, ROLLED_BACK)
- `expected_impact`, `risk_level`

**`action_outcomes`**
- `action_id`, `window_after_minutes`
- `delta_revenue`, `delta_timeout_rate`, `delta_win_rate`
- `success_score`
