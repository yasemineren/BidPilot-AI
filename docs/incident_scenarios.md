# Incident Scenarios

## 1) Latency/Timeout Incident
- **Belirti**: timeout_rate ↑, p99 latency ↑
- **Aksiyon**: throttle önerisi (%18 gibi)
- **Beklenen Etki**: timeout_rate -40% ±10%

## 2) Revenue Drop
- **Belirti**: win_rate stabil, revenue düşüyor
- **Aksiyon**: config check (floor price / pacing)

## 3) Fraud-like Burst
- **Belirti**: fraud_signal ↑, abnormal click/impression pattern
- **Aksiyon**: segment izolasyonu + high severity alert

## 4) False Positive
- **Belirti**: kısa süreli spike
- **Aksiyon**: “no action, monitor”
