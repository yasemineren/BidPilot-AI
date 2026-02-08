# BidPilot AI

**RTB için canlı karar destek ve otomatik guardrail sistemi**

**Slogan:** Alarm değil, aksiyon üretir.

## Problem → Why it matters
RTB (Real-Time Bidding) dünyasında metrikler anlık dalgalanır: win-rate düşer, timeout artar, fraud-benzeri sinyaller patlar. Mevcut alarm sistemleri sadece uyarı üretir; operasyon ekipleri ise sorunun kaynağını, etkisini ve aksiyonu manuel çözmeye çalışır. BidPilot AI, bu süreci **anomali → etki → aksiyon → öğrenme** döngüsüne çevirir ve doğrudan uygulanabilir öneriler üretir.

## What it does
- RTB metriklerini zaman serisi olarak izler (QPS, win-rate, CPM, revenue, timeout, fraud-like sinyaller).
- **Anomali tespiti**: performans, gelir, latency, fraud pattern.
- **Traffic throttling & weight optimization**: segment bazlı öneri.
- **Config & rule discovery**: hangi kural/parametre düzeltilmeli?
- **Smart alerting**: “Ne oldu + neden + ne yapalım + beklenen etki”.
- **Feedback loop**: uygulanan aksiyonun sonucu öğrenmeye döner.

**Yıldız özellik:** “Model skor verdi” değil, “şu segmentte throttle %18 öneriyorum çünkü X/Y/Z ve beklenen revenue impact + risk” diyebilmesi.

## Architecture diagram (tek sayfa)
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

> Diagram görseli için: `docs/architecture.md` ve `docs/architecture.png` (placeholder).

## Demo GIF
- **Demo flow:** anomali geliyor → aksiyon kartı → apply → recovery grafiği
- **Yer tutucu:** `docs/demo.gif` (eklenmesi önerilir)

## SLO / Latency budget
- **Decision API**: p95 < 30ms

## Runbook (ops ne yapar?)
Özet runbook: `docs/runbook.md`

## Tradeoffs
- **Robust z-score + ML:** explainable ve hızlı.
- **Rule guardrails:** düşük riskli otomasyon.
- **LightGBM/XGBoost:** production seviyesinde hızlı ve yorumlanabilir.
- **Throttle optimizasyonu:** revenue kaybını minimize ederken timeout düşürür.

## Repo structure
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

## Quick pitch (LinkedIn/GitHub)
**Title:** Production-Grade RTB Decision Support: Anomaly → Impact → Action → Feedback Loop

**Lead:** Alarm üretmiyor, throttle/config aksiyonu öneriyor ve sonucu öğreniyor.

**Showcase:** dashboard + action card + before/after grafiği.

## Next steps
- Docs dosyalarını okuyun (`docs/`)
- Demo senaryolarını çalıştırın (`docs/incident_scenarios.md`)
- Decision API kontratını inceleyin (`docs/api_contract.md`)

---

> Bu repo MVP dokümantasyonu içerir; kod iskeleti ve demo akışları genişletmeye hazırdır.
