package main

import (
	"context"
	"fmt"
	"log"
	"crypto/tls"
	"os"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/gofiber/fiber/v2"
)

type RtbEvent struct {
	EventID   string  `json:"event_id"`
	Timestamp string  `json:"ts"`
	BidderID  string  `json:"bidder_id"`
	Geo       string  `json:"geo"`
	BidPrice  float64 `json:"bid_price"`
	Won       bool    `json:"won"`
}

func main() {
	// 1. Bağlantı Ayarları
	addrEnv := getenv("CLICKHOUSE_ADDR", "localhost:9000")
	addrs := strings.Split(addrEnv, ",")
	for i := range addrs {
		addrs[i] = strings.TrimSpace(addrs[i])
	}
	useTLS := getenv("CLICKHOUSE_SECURE", "") == "true"
	skipVerify := getenv("CLICKHOUSE_INSECURE_SKIP_VERIFY", "") == "true"

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: addrs,
		Auth: clickhouse.Auth{
			Database: getenv("CLICKHOUSE_DATABASE", "default"),
			Username: getenv("CLICKHOUSE_USER", "default"),
			Password: getenv("CLICKHOUSE_PASSWORD", ""),
		},
		TLS: func() *tls.Config {
			if !useTLS {
				return nil
			}
			return &tls.Config{InsecureSkipVerify: skipVerify}
		}(),
	})
	if err != nil {
		log.Fatal("❌ Bağlantı Kurulamadı (CLICKHOUSE_* env değişkenlerini kontrol et):", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal("❌ ClickHouse Cevap Vermiyor (şifre/host/port/tls kontrol et):", err)
	}
	fmt.Println("✅ ClickHouse Bağlantısı Başarılı! (Hazır)")

	// 2. Sunucu Başlat
	app := fiber.New()

	app.Post("/api/v1/event", func(c *fiber.Ctx) error {
		event := new(RtbEvent)
		if err := c.BodyParser(event); err != nil {
			return c.Status(400).SendString("Bozuk JSON")
		}

		// 3. ANINDA YAZ (Bekleme yok)
		ctx := context.Background()
		wonInt := uint8(0)
		if event.Won { wonInt = 1 }

		// Zaman formatını düzelt (ISO8601 -> DateTime)
		t, _ := time.Parse(time.RFC3339, event.Timestamp)

		query := `INSERT INTO rtb_events (event_id, ts, bidder_id, geo, bid_price, won) VALUES (?, ?, ?, ?, ?, ?)`
		err := conn.Exec(ctx, query, event.EventID, t, event.BidderID, event.Geo, event.BidPrice, wonInt)

		if err != nil {
			fmt.Println("❌ YAZMA HATASI:", err) // Hatayı ekrana bas
			return c.Status(500).SendString(err.Error())
		}

		fmt.Println("💾 YAZILDI:", event.EventID) // Başarıyı ekrana bas
		return c.SendStatus(200)
	})

	log.Fatal(app.Listen(":3000"))
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
