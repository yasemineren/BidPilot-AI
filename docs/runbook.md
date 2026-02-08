# Runbook: Alarm geldiğinde ops ne yapar?

## 1) Olayı doğrula
- UI’da segment ve zaman penceresini kontrol et.
- Decision output: `What/Why/Action/Expected impact/Risk` kartını incele.

## 2) Aksiyon öncesi kontrol
- Impact estimate: expected revenue loss vs latency improvement.
- Risk seviyesi: high ise onay gerektirir.

## 3) Uygula
- UI’dan **Apply** veya Control API üzerinden `POST /v1/actions/apply`.
- Action audit log’a yazılır.

## 4) İzleme
- 10 dk boyunca metrik trendlerini takip et.
- Rollback şartı tetiklenirse `POST /v1/actions/rollback`.

## 5) Öğrenme
- Sonuçlar `action_outcomes` tablosuna yazılır.
- Haftalık retrain pipeline, yeni örneklerle güncellenir.
