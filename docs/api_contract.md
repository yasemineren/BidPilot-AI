# API Tasarımı

## Decision API
**POST /v1/decide**

**Input**
```json
{
  "segment_key": "geo|device|publisher|format|bidder",
  "last_metrics_window": {
    "qps": 1200,
    "win_rate": 0.24,
    "rpm": 3.1,
    "timeout_rate": 0.08,
    "fraud_score_avg": 0.12,
    "window_minutes": 5
  }
}
```

**Output**
```json
{
  "anomaly": true,
  "incident_type": "latency",
  "action_proposal": {
    "type": "THROTTLE",
    "value": 0.18,
    "duration_minutes": 10
  },
  "explanation": {
    "what": "Timeout spike detected in segment S",
    "why": "timeout_rate +312% vs baseline, p99 latency +180ms",
    "expected_impact": "timeout_rate -40% ±10%, revenue -3% ±2%",
    "risk": "medium",
    "rollback_condition": "revenue drop >5% for 3 min"
  }
}
```

## Control API

**POST /v1/actions/apply**
```json
{
  "segment_key": "geo|device|publisher|format|bidder",
  "action_type": "THROTTLE",
  "value": 0.18,
  "mode": "APPROVED"
}
```

**POST /v1/actions/rollback**
```json
{
  "action_id": "uuid",
  "reason": "revenue drop >5% for 3 min"
}
```

**GET /v1/actions/history**
- Aksiyon geçmişi, status ve outcome ile birlikte döner.

## Metrics
**GET /metrics** (Prometheus)
