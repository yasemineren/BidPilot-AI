# API Contract

## Decision API
**POST** `/v1/decide`

Request:
```json
{
  "segment_key": "geo|device|publisher|format|bidder",
  "window": "5m",
  "metrics": {
    "qps": 1200,
    "win_rate": 0.18,
    "timeout_rate": 0.07,
    "rpm": 2.4
  }
}
```

Response:
```json
{
  "anomaly_score": 0.86,
  "incident_type": "timeout_spike",
  "action_proposal": {
    "type": "THROTTLE",
    "value": 18,
    "duration_minutes": 10
  },
  "explanation": {
    "what": "Timeout spike detected",
    "why": "timeout_rate +312% vs baseline",
    "expected_impact": "timeout_rate -40% ±10%",
    "risk": "medium",
    "rollback_condition": "revenue drop >5% for 3 min"
  }
}
```

## Control API
- **POST** `/v1/actions/apply`
- **POST** `/v1/actions/rollback`
- **GET** `/v1/actions/history`

## Metrics
- **GET** `/metrics`
