BidPilot AI ⚡️

Production-grade RTB Decision Support: Anomaly → Impact → Action → Feedback Loop

BidPilot AI, RTB benzeri yüksek trafikli sistemlerde zaman serisi metriklerden anomali tespit eden, anomaliyi iş etkisine (revenue/latency) çeviren ve çıktıyı traffic throttling / weight optimization / config-check gibi aksiyonlara dönüştüren bir karar destek platformudur.

“Alarm üretmek” değil, operasyonel karar üretmek için tasarlandı.

Why this matters (Problem)

RTB platformlarında küçük bir gecikme artışı veya config hatası:

win-rate / revenue düşüşüne,

timeout patlamasına,

traffic quality bozulmasına
dakikalar içinde yol açabilir.

Klasik monitoring “grafik gösterir”; BidPilot AI ise:

ne oldu,

neden olabilir,

ne yapmalıyız,

beklenen etki + risk
çıktısını otomatik üretir.

What it does (Key Features)

✅ Anomaly Detection (Time-series & Robust stats + ML)

timeout_rate, win_rate, rpm, qps gibi metriklerde baseline’a göre sapma tespiti

incident type olasılıkları (latency spike, revenue drop, fraud-like burst, config mismatch)

✅ Impact Estimation

“Bu anomali devam ederse tahmini revenue kaybı/dakika”

“Bu throttle uygulanırsa beklenen recovery”

✅ Action Recommendations (Hybrid decisioning)

Rule-based guardrails + ML önerisi

segment bazlı: THROTTLE %x, WEIGHT adjust, CONFIG_CHECK

✅ Smart Alerting (Actionable outputs)
Her alarm şu formatta üretilir:

What: (ne oldu)

Why: (hangi sinyaller)

Action: (öneri)

Expected impact: (beklenen etki)

Risk & rollback condition: (risk + geri dönüş koşulu)

✅ Feedback Loop (Outcome → Learning → Action)
Uygulanan aksiyonların etkisi ölçülür ve modelin öğrenme datasına geri beslenir.

Architecture (High-level)

Go Ingestion Service: yüksek hacimli event ingest + aggregation

ClickHouse: time-series analytics, segment breakdown, baseline comparison

Python Decision Service (FastAPI): anomaly + impact + action selection

PostgreSQL Control Plane: config, rule versioning, action audit, approvals

Dashboard: anomalies + actions + apply/rollback

Observability: Prometheus metrics + Grafana dashboards

Deployment: Docker + Kubernetes (Helm)

Diagram
flowchart LR
  A[RTB Event Stream / Simulator] --> B[Go Ingestion Service]
  B --> C[ClickHouse: metrics & baselines]
  C --> D[Python Decision Service]
  D --> E[PostgreSQL Control Plane]
  D --> F[Alerts & Recommendations]
  F --> G[Ops Dashboard]
  G -->|Apply/Rollback| E
  E -->|Rules/Configs| D
  D --> H[Prometheus/Grafana]

Tech Stack

Go: ingestion & aggregation

Python (FastAPI): decision engine (anomaly/impact/action)

ClickHouse: large-scale time-series analytics

PostgreSQL: control plane + audit + config versioning

Docker / Kubernetes / Helm: containerization + deploy

Prometheus / Grafana: monitoring

Quickstart (Local)
1) Run everything
docker compose up --build

2) Generate traffic + incidents
python simulator/generator.py --scenario latency_spike

3) Open dashboard

Dashboard: http://localhost:8501

Decision API: http://localhost:8000/docs

Secrets: .env dosyanızı .env.example üzerinden oluşturun.

API (Core endpoints)
Decision

POST /v1/decide

input: segment_key + time_window_metrics

output: anomaly_score + incident_probs + action_proposal + explanation

Actions

POST /v1/actions/apply

POST /v1/actions/rollback

GET /v1/actions/history

Metrics

GET /metrics (Prometheus)

Incident Scenarios (Built-in)

latency_spike: timeout_rate ↑, p99 latency ↑

revenue_drop: rpm ↓, win_rate ↓

fraud_burst: fraud-like score burst

config_mismatch: baseline deviates after config change

Production-minded details (What makes this “not a notebook”)

Hybrid decisioning: rules + ML together

Action audit log: apply/rollback izlenebilir

Rollback conditions her öneride var

Observability: service latency, anomaly counts, action counts, error rates

Reproducible deploy: Docker + Helm

Targets / Benchmarks (fill as you measure)

Ingestion throughput: X events/sec

Decision latency: p95 < Y ms

ClickHouse “top regressions” query: < Z ms

Incident detection: TP/FP across 4 scenarios

Repo structure
bidpilot-ai/
  simulator/
  ingest-go/
  decision-service/
  control-plane/
  dashboards/
  infra/
  docs/

Roadmap

 Online bandit optimization for throttle % (Thompson Sampling)

 Canary rollout integration via control-plane

 Drift detection on segment distributions

 Automated retraining pipeline (weekly)

License

MIT
