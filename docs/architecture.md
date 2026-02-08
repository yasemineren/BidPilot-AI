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
