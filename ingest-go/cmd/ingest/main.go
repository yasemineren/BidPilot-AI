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
	// Bağlantı ayarları (Localhost)
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"bidpilot-clickhouse:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "123",
		},
	})
	if err != nil {
		log.Fatal("❌ Bağlantı hatası:", err)
	}

	app := fiber.New()

	app.Post("/api/v1/event", func(c *fiber.Ctx) error {
		event := new(RtbEvent)
		if err := c.BodyParser(event); err != nil {
			return c.Status(400).SendString("Bozuk JSON")
		}
		
		ctx := context.Background()
		wonInt := uint8(0)
		if event.Won { wonInt = 1 }

		// --- KRİTİK NOKTA: ŞU ANKİ ZAMAN ---
		t := time.Now() 
		// -----------------------------------

		// İSPİYONCU: Terminale tarihi yazdırıyoruz
		fmt.Println("🕒 VERİTABANINA GİDEN TARİH:", t.Format("2006-01-02 15:04:05"))

		query := "INSERT INTO rtb_events (event_id, ts, bidder_id, geo, bid_price, won) VALUES (?, ?, ?, ?, ?, ?)"
		err := conn.Exec(ctx, query, event.EventID, t, event.BidderID, event.Geo, event.BidPrice, wonInt)

		if err != nil {
			fmt.Println("❌ YAZMA HATASI:", err)
		}
		return c.SendStatus(200)
	})

	log.Fatal(app.Listen(":3000"))
}