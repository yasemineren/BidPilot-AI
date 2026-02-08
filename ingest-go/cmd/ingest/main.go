package main

import (
	"context"
	"fmt"
	"log"
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
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
	})
	if err != nil {
		log.Fatal("❌ Bağlantı Kurulamadı:", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal("❌ ClickHouse Cevap Vermiyor:", err)
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