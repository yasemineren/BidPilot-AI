# Runbook

## Alarm geldiğinde
1. Segment etkisi ve risk skoru kontrol edilir
2. Önerilen aksiyonun beklenen etkisi incelenir
3. Rollback şartı netleştirilir
4. **Apply** veya **Manual** mod seçilir

## Rollback
- Revenue drop >5% (3 dakika üst üste)
- Timeout iyileşmedi

## Kapanış
- `action_outcomes` tablosuna sonuç yazılır
- Feedback loop için dataset güncellenir
