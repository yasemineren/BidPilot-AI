# Sistem Mimarisi

BidPilot AI, RTB metriklerini gerçek zamanlı izler ve **actionable** kararlar üretir. Mimari, throughput ve düşük latency hedefiyle tasarlanmıştır.

## Bileşenler

### A) RTB Metrics Generator (Simülatör)
- Event üretir: `bid_request`, `bid_response`, `win`, `click`, `revenue`, `timeout`, `fraud_signal`.
- Segmentler: `geo`, `device`, `publisher`, `ad_format`, `hour_of_day`, `bidder_id`.
- Incident injector:
  - timeout spike
  - win-rate drop
  - revenue drop
  - fraud-like burst
  - config yanlışlığı (floor price değişimi)

### B) Ingest Service (Go)
- Event toplar, 1s/10s/1m pencerelerde **metrics aggregation** yapar.
- ClickHouse’a yazar.
- Latency bütçesi: hızlı ingest ve minimal buffer.

### C) Feature & Analytics Layer (ClickHouse)
- Zaman serisi metrikleri ve segment kırılımları.
- Sorgular: “son 5 dk vs baseline” ve “top regressions”.

### D) Decision Engine (FastAPI/Go)
- 2 aşamalı karar:
  - **Rule Guardrails:** extreme latency/timeout → throttle limit.
  - **ML Layer:** anomaly score + impact estimate + action selector.

### E) Control Plane (PostgreSQL)
- Config parametreleri, aksiyon geçmişi, onay akışı.

### F) UI / Ops Dashboard
- Canlı metrikler, anomali timeline, önerilen aksiyonlar.

### G) Observability
- Prometheus + Grafana + structured logs.

### H) Kubernetes Deployment
- Deployment + HPA + Canary rollout.

## Tek sayfa diagram
```mermaid
flowchart LR
  A[RTB Metrics Generator] --> B[Ingest Service (Go)]
  B --> C[ClickHouse: Feature & Analytics]
  C --> D[Decision Engine]
  D --> E[Control Plane (Postgres)]
  D --> F[UI / Ops Dashboard]
  D --> G[Observability: Prometheus/Grafana]
  E --> H[Approval Workflow]
```
