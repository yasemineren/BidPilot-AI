# 🚀 BidPilot: Real-Time Bidding (RTB) Analytics Pipeline

![Go](https://img.shields.io/badge/Go-1.21-00ADD8?style=flat&logo=go)
![ClickHouse](https://img.shields.io/badge/ClickHouse-OLAP-FFCC00?style=flat&logo=clickhouse)
![Docker](https://img.shields.io/badge/Docker-Container-2496ED?style=flat&logo=docker)
![Grafana](https://img.shields.io/badge/Grafana-Dashboard-F46800?style=flat&logo=grafana)

**BidPilot** is a high-performance data engineering project that simulates, ingests, stores, and visualizes Real-Time Bidding (RTB) ad auction data. It demonstrates a modern **Microservices** architecture using **Golang** for high-throughput ingestion and **ClickHouse** for real-time big data analytics.

---

## 🏗 Architecture

The system consists of 4 main containerized components orchestrated via Docker Compose:

```mermaid
graph LR
    A[🐍 Python Simulator] -- POST /api/bid (JSON) --> B[🐹 Go Ingest Service]
    B -- Batch/Stream Insert --> C[🏛️ ClickHouse DB]
    D[📊 Grafana UI] -- SQL Queries --> C
