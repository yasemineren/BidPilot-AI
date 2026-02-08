# Decision Logic

## 1) Anomali Tespiti (Explainable)

### Katman 1: Robust Z-Score / MAD
- Metrikler: `timeout_rate`, `win_rate`, `rpm`, `qps`
- Baseline: aynı saat ve gün paternine göre
- Çıktı: `anomaly_score_robust`

### Katman 2: ML Anomaly
- Input: son N dakika features + trend/derivative + ratio
- Model: LightGBM / XGBoost classifier
- Label: simülatörden `incident_type`
- Çıktı: `incident_probabilities`, `confidence`

## 2) Impact Estimation
- “Incident devam ederse revenue kaybı/dk?”
- “Throttle uygularsak recovery etkisi?”

Çıktı:
- `expected_loss_per_min`
- `expected_recovery_if_throttle_X`

## 3) Action Selection

### Hard Rules
- Timeout_rate p99 üstü → throttle öner
- Fraud burst → throttle + segment flag

### ML Önerisi
- Multi-class recommendation: `THROTTLE`, `WEIGHT`, `CONFIG_CHECK`
- Basit arama: timeout düşsün, revenue loss minimize

### Bonus
- Thompson Sampling ile throttle seviyelerini test et

## 4) Explainability Şablonu
- What
- Why
- Action
- Expected impact
- Risk
- Rollback condition
