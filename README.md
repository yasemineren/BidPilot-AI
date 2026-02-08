# BidPilot AI

**RTB için canlı karar destek ve otomatik guardrail sistemi**

**Slogan:** *Alarm değil, aksiyon üretir.*

BidPilot AI, RTB akışından gelen metrikleri izler, anomaliyi tespit eder, etkiyi hesaplar ve **aksiyon** önerir. Hedef: “model skor verdi” demek değil, **hangi segmentte hangi aksiyonun neden gerektiğini ve beklenen etkiyi** söylemek.

---

## Problem → Why it matters
RTB sistemlerinde küçük bir gecikme, düşük win-rate veya yanlış config **dakikalar içinde büyük gelir kaybı** yaratabilir. Sadece alarm üretmek yeterli değildir; **neden–etki–aksiyon** zinciri hızlıca kurulmalıdır.

BidPilot AI:
- **Anomaliyi** güvenilir şekilde yakalar
- **Etkiyi** TL/dk ve risk üzerinden açıklar
- **Aksiyon** önerir (throttle, weight, config check)
- **Sonucu öğrenir** (feedback loop)

---

## Architecture diagram (tek sayfalık)
```mermaid
flowchart LR
  subgraph Sim[RTB Metrics Generator]
    A1[bid_request]
    A2[bid_response]
    A3[win/click/revenue]
    A4[timeout/fraud_signal]
    A5[incident injector]
  end

  subgraph Ingest[Ingest Service (Go)]
    B1[Aggregation 1s/10s/1m]
    B2[Low-latency write]
  end

  subgraph CH[ClickHouse]
    C1[rtb_events_raw]
    C2[rtb_metrics_1m]
    C3[baseline_profiles]
  end

  subgraph Decision[Decision Engine]
    D1[Rule Guardrails]
    D2[ML Anomaly + Impact]
    D3[Action Selection]
  end

  subgraph Control[Control Plane (Postgres)]
    E1[config_rules]
    E2[actions]
    E3[action_outcomes]
  end

  subgraph UI[Ops Dashboard]
    F1[Live metrics]
    F2[Action queue]
    F3[Apply/Rollback]
  end

  subgraph Obs[Observability]
    G1[Prometheus]
    G2[Grafana]
  end

  Sim --> Ingest --> CH --> Decision --> Control
  Decision --> UI
  Ingest --> Obs
  Decision --> Obs
```

---

## Demo GIF
> **Anomali → aksiyon kartı → apply → recovery grafiği**

![BidPilot AI demo](docs/demo.gif)

---

## SLO / Latency budget
- **Decision API**: `p95 < 30ms`
- Ingest write: **< 50ms**
- UI refresh: **< 5s**

---

## Sistem Mimarisi (özet)
- **RTB Metrics Generator (Simülatör):** gerçek RTB event akışını ve kontrollü incident’ları üretir.
- **Ingest Service (Go):** 1s/10s/1m pencerelerde metrik agregasyonu + ClickHouse write.
- **Feature & Analytics (ClickHouse):** baseline + segment kırılımları + regressions sorguları.
- **Decision Engine (Python/Go):** rule guardrails + ML anomaly + impact + action selector.
- **Control Plane (Postgres):** config, actions, outcomes, approval workflow.
- **UI / Ops Dashboard:** canlı metrikler, öneri kuyruğu, apply/rollback.
- **Observability:** Prometheus + Grafana + structured logs.
- **Kubernetes:** HPA, Canary, ConfigMap/Secret.

---

## Veritabanı Tasarımı (özet)
### ClickHouse çekirdek tabloları
- `rtb_events_raw`
- `rtb_metrics_1m`
- `baseline_profiles`

### Postgres kontrol tabloları
- `config_rules`
- `actions`
- `action_outcomes`

Detaylı şema: [docs/architecture.md](docs/architecture.md)

---

## Model ve Karar Mantığı
**2 aşamalı hibrit yaklaşım:**
1) **Rule Guardrails**: extreme latency/timeout → throttle limit
2) **ML Layer**: anomaly score + incident type + impact + action selector

**Explainable çıktı şablonu:**
> *What*: Timeout spike detected in segment S
> *Why*: timeout_rate +312% vs baseline, p99 latency +180ms, win_rate down
> *Action*: throttle 18% for 10 min
> *Expected impact*: timeout_rate -40% ±10%, revenue -3% ±2%
> *Risk*: medium
> *Rollback condition*: revenue drop >5% for 3 min

Detaylar: [docs/decision_logic.md](docs/decision_logic.md)

---

## Feedback Loop (Outcome → Learning → Action)
- Aksiyon uygulandıktan **10 dk sonrası** metrikler ölçülür
- `success_score` üretilir (revenue, latency, fraud hedefleri birlikte)
- Yeni örnek **dataset’e eklenir**, haftalık retrain pipeline beslenir

Detaylı süreç: [docs/decision_logic.md](docs/decision_logic.md)

---

## Demo Senaryoları (en az 4)
1. **Latency/timeout incident** → throttle önerisi + recovery
2. **Revenue drop** → config check önerisi (floor price mismatch)
3. **Fraud-like burst** → segment izolasyonu + high severity
4. **False positive** → “no action, monitor”

Detaylar: [docs/incident_scenarios.md](docs/incident_scenarios.md)

---

## Runbook (Ops ne yapar?)
Alarm geldiğinde:
1. Segment etkisi ve beklenen risk skoru kontrol edilir
2. Action card üzerindeki öneri ve rollback şartı değerlendirilir
3. **Apply** veya **Manual override** seçilir

Detaylı runbook: [docs/runbook.md](docs/runbook.md)

---

## Tradeoffs
- **Robust z-score + ML**: düşük yanlış-pozitif + explainability
- **Throttle yerine config**: revenue kaybını minimize etme
- **Semi-automated mod**: operatör onayıyla risk düşürme

---

## Production Kalitesi: Eksiği Olmasın Checklist
### 1) API Tasarımı
- **Decision API**: `POST /v1/decide`
- **Control API**: `POST /v1/actions/apply`, `POST /v1/actions/rollback`, `GET /v1/actions/history`
- **Metrics**: `GET /metrics`

### 2) Güvenlik & Operasyon
- Secrets: DB creds, API keys → **K8s Secret**
- **RBAC**: Apply/Rollback endpoint sadece operator token ile
- **Audit log**: action uygulayan kim, ne zaman?

### 3) Testler
- Unit: feature builder, anomaly scoring, action policy
- Integration: ClickHouse write/read, Decision API schema
- Load: k6 ile `/v1/decide` latency
- Regression: 5 incident senaryosu, beklenen aksiyon doğru mu?

### 4) CI/CD
- GitHub Actions: lint + tests
- Docker build + Trivy security scan
- Helm chart validate

### 5) Deploy
- `helm/` chart ile tek komut kurulum
- HPA: CPU + custom metric (QPS)
- Canary rollout (%10 → %100)

---

## Repo Yapısı (target)
```
bidpilot-ai/
  README.md
  docs/
    architecture.md
    decision_logic.md
    incident_scenarios.md
    runbook.md
    api_contract.md
  simulator/
    generator.py
    incident_injector.py
  ingest-go/
    cmd/ingest/main.go
    internal/...
  decision-service/
    app/
    models/
    feature_pipeline/
    policy/
    main.py
  control-plane/
    migrations/
    api/
  dashboards/
    streamlit_app.py
    grafana/
  infra/
    docker/
    k8s/
    helm/
  tests/
  scripts/
```

---

## API Surface (özet)
- **Decision API**: `POST /v1/decide`
- **Control API**: `POST /v1/actions/apply`, `POST /v1/actions/rollback`, `GET /v1/actions/history`
- **Metrics**: `GET /metrics`

Detaylı contract: [docs/api_contract.md](docs/api_contract.md)

---

## “Beni işe aldıracak” ekstra dokunuşlar
- **Actionable Alert Composer** (şablon + guardrail)
- **Root-Cause Candidate Ranking** (top contributing segments)
- **Cost-Aware Decision** (opportunity cost vs latency)

---

## Kabul kriterleri (Definition of Done)
- ClickHouse’da metrik tabloları ve baseline var
- 4 incident senaryosu otomatik üretilebiliyor
- Decision engine en az 2 katmanlı
- Her öneri explainable + rollback condition içeriyor
- Apply/Rollback audit log yazıyor
- Prometheus + Grafana mevcut
- K8s deploy tek komutla kalkıyor (helm)
- CI testleri geçiyor
- README + runbook + architecture diagram + demo GIF var

---

## LinkedIn / GitHub sunumu
**Başlık:** *Production-Grade RTB Decision Support: Anomaly → Impact → Action → Feedback Loop*

**İlk cümle:** *“Alarm üretmiyor, throttle/config aksiyonu öneriyor ve sonucu öğreniyor.”*

**Önerilen görsel:** dashboard + action card + before/after grafiği
