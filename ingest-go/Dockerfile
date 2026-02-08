# Hafif bir Linux versiyonu kullan
FROM golang:alpine

# Çalışma klasörü oluştur
WORKDIR /app

# Dosyaları kopyala
COPY . .

# Kütüphaneleri indir
RUN go mod download

# Programı derle (build)
RUN go build -o main cmd/ingest/main.go

# Çalıştır
CMD ["./main"]