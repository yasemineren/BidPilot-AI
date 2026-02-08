# Demo Senaryoları

## 1) Latency/timeout incident
- **Trigger:** timeout spike
- **Beklenen:** throttle öneriyor, recovery grafiği yükseliyor.

## 2) Revenue drop
- **Trigger:** revenue düşüşü
- **Beklenen:** config check öneriyor (ör: floor price mismatch).

## 3) Fraud-like burst
- **Trigger:** impression/click pattern anomaly
- **Beklenen:** segment izolasyonu + alert severity high.

## 4) False positive
- **Trigger:** normal dalgalanma
- **Beklenen:** “no action, monitor” kararı.

## Otomasyon
- 4 senaryo, simülatör + incident injector ile otomatik üretilebilir.
