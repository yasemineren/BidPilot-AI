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
  - Amaç: `timeout_rate` düşüşü
  - Kısıt: `revenue` kaybı minimum

### Bonus
- Thompson Sampling ile throttle seviyelerini test et

## 4) Explainability Şablonu
- What
- Why
- Action
- Expected impact
- Risk
- Rollback condition

## 5) Feedback Loop
- Aksiyon uygulandıktan sonra **10 dk window** ölçülür
- `success_score` hesaplanır (revenue + latency + fraud)
- Sonuç `action_outcomes` tablosuna yazılır
- Dataset’e yeni örnek eklenir, haftalık retrain pipeline beslenir
