import time
import json
import random
import uuid
import requests  # Bu kütüphaneyi kullanacağız
from datetime import datetime

# RTB Sabitleri
GEO_LIST = ["US", "TR", "UK", "DE", "FR"]
BIDDERS = [f"bidder_{i:02d}" for i in range(1, 6)]

def generate_event():
    event_id = str(uuid.uuid4())
    ts = datetime.utcnow().isoformat()
    
    # Basit bir event oluştur
    event = {
        "event_id": event_id,
        "ts": ts,
        "bidder_id": random.choice(BIDDERS),
        "geo": random.choice(GEO_LIST),
        "bid_price": round(random.uniform(0.1, 5.0), 4),
        "won": random.choice([True, False])
    }
    return event

if __name__ == "__main__":
    print("🚀 Traffic Generator Started... Sending to Go Server!")
    target_url = "http://127.0.0.1:3000/api/v1/event"

    try:
        while True:
            # Event oluştur
            data = generate_event()
            
            try:
                # Go servisine POST isteği at
                response = requests.post(target_url, json=data)
                
                # Sonucu ekrana yaz (Sadece hata varsa veya başarılıysa)
                if response.status_code == 200:
                    print(f"✅ Sent: {data['event_id']} | Go: 200 OK")
                else:
                    print(f"❌ Error: {response.status_code}")
            except Exception as e:
                print(f"⚠️ Connection Error: {e}")
                print("   (Go servisi çalışıyor mu?)")

            time.sleep(0.01)
    except KeyboardInterrupt:
        print("\n🛑 Stopped.")