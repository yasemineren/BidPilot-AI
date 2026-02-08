# Model ve Karar Mantığı (Hibrit Tasarım)

## 1) Anomali tespiti

### Katman 1: Robust z-score / MAD
- Metric’ler: `timeout_rate`, `win_rate`, `rpm`, `qps`.
- Baseline: aynı saat ve aynı gün paternine göre.
- Çıkış: `anomaly_score_robust`.

### Katman 2: ML Anomaly
- Input: son N dakika features + derivatives (trend) + ratio’lar.
- Model: LightGBM / XGBoost classifier.
- Label’lar: simülatörden gelen `incident_type`.
- Çıkış:
  - `incident_probabilities` (latency, revenue_drop, fraud_spike, config_mismatch)
  - `confidence`

## 2) Impact estimation
- “Bu incident bu segmentte devam ederse revenue kaybı/dakika?”
- “Throttle uygularsak beklenen recovery?”
- Çıkış:
  - `expected_loss_per_min`
  - `expected_recovery_if_throttle_X`

## 3) Action selection

### Hard rules
- `timeout_rate` p99 üstü → throttle öner (risk high).
- `fraud_score` burst → throttle + segment flag.

### ML önerisi
- Multi-class recommendation:
  - `THROTTLE %x`
  - `WEIGHT -y`
  - `CONFIG_CHECK` (floor price / bidder timeout / pacing)

### Optimization
- “Throttle %” seçimi için basit arama:
  - hedef: timeout_rate düşsün
  - constraint: revenue loss minimize

### Bonus (nice-to-have)
- Thompson Sampling bandit ile farklı throttle seviyelerini test edip en iyiyi öğrenme.

## 4) Explainability şablonu
```
What: Timeout spike detected in segment S
Why: timeout_rate +312% vs baseline, p99 latency +180ms, win_rate down
Action: throttle 18% for 10 min
Expected impact: timeout_rate -40% ±10%, revenue -3% ±2%
Risk: medium
Rollback condition: revenue drop >5% for 3 min
```
