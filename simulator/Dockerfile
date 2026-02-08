# Python 3.9 kullan
FROM python:3.9-slim

WORKDIR /app

# Dosyaları kopyala
COPY . .

# Gerekli kütüphaneyi kur
RUN pip install requests

# Programı çalıştır (Hızlı modda)
CMD ["python", "generator.py"]