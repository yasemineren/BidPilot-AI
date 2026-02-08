# Runbook

## Alarm geldiğinde
1. Segment etkisi ve risk skoru kontrol edilir
2. Önerilen aksiyonun beklenen etkisi incelenir
3. Rollback şartı netleştirilir
4. **Apply** veya **Manual** mod seçilir

## Rollback
- Revenue drop >5% (3 dakika üst üste)
- Timeout iyileşmedi

## ClickHouse bağlantı sorunları (şifre/erişim)
ClickHouse çalışıyor gibi görünse bile aşağıdaki hata ile karşılaşırsan:
```
If you have installed ClickHouse and forgot password you can reset it in the configuration file.
The password for default user is typically located at /etc/clickhouse-server/users.d/default-password.xml
```

Kontrol listesi:
1. **ClickHouse Cloud** kullanıyorsan servis ayarlarından şifreyi doğrula ve bağlantı string’ini güncelle.
2. **Self-hosted** kurulumda `default` kullanıcısının şifresini `/etc/clickhouse-server/users.d/default-password.xml`
   dosyasından kontrol et (dosyayı silmek şifreyi sıfırlar).
3. Uygulama tarafında **env/config** güncellemesi yap:
   - `CLICKHOUSE_ADDR` (ör: `localhost:9000` veya `host:9440`)
   - `CLICKHOUSE_USER`
   - `CLICKHOUSE_PASSWORD`
   - `CLICKHOUSE_DATABASE` (opsiyonel, varsayılan `default`)
   - `CLICKHOUSE_SECURE=true` (TLS gerekiyorsa)
   - `CLICKHOUSE_INSECURE_SKIP_VERIFY=true` (sadece lokal test)

Amaç: yanlış şifre yüzünden ClickHouse bağlantısının düşmesini hızlıca ayırt etmek.

## Kapanış
- `action_outcomes` tablosuna sonuç yazılır
- Feedback loop için dataset güncellenir
