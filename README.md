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
Simulator (Python): Generates high-frequency synthetic bid requests (JSON) with random parameters (price, geo, device, etc.).

Ingestion Layer (Golang + Fiber): A lightweight, high-performance REST API that receives the stream, validates data, and writes to the database.

Storage Layer (ClickHouse): A column-oriented OLAP database optimized for lightning-fast analytical queries on large datasets.

Visualization (Grafana): connects to ClickHouse to visualize KPIs like Average Bid Price, Win Rate, and Traffic by Country in real-time.

🛠 Tech Stack
Language: Go (Golang) 1.21

Web Framework: Fiber (Express-inspired, zero allocation)

Database: ClickHouse (Time-Series / OLAP)

Dashboard: Grafana

Scripting: Python 3.9 (Data Generation)

Infrastructure: Docker & Docker Compose

🚀 Getting Started
Prerequisites
Docker & Docker Compose installed on your machine.

Installation
Clone the repository:

Bash

git clone [https://github.com/YOUR_USERNAME/BidPilot-AI.git](https://github.com/YOUR_USERNAME/BidPilot-AI.git)
cd BidPilot-AI
Build and Run the Stack: This command will build the Go and Python images and start all services (ClickHouse, Grafana, Ingest, Simulator).

Bash

docker-compose up --build -d
Verify Running Services:

Bash

docker ps
You should see bidpilot-ingest, bidpilot-clickhouse, bidpilot-grafana, and bidpilot-simulator running.

📊 Dashboard Configuration (Grafana)
Once the containers are up, you can access the dashboard:

Open your browser and go to: http://localhost:3001

Login: admin / admin

Add Data Source:

Go to Connections > Data Sources > Add data source.

Search for ClickHouse.

Address/URL: bidpilot-clickhouse

Port: 8123

Database: default

Click Save & Test.

📈 Sample Queries for Panels
You can use these SQL queries to create widgets in Grafana:

1. Average Bid Price (Stat Panel):

SQL

SELECT avg(bid_price) FROM rtb_events WHERE ts >= now() - INTERVAL 1 MINUTE
2. Live Traffic Feed (Table Panel):

SQL

SELECT * FROM rtb_events ORDER BY ts DESC LIMIT 10
3. Request Count by Country (Pie Chart):

SQL

SELECT country, count(*) as count 
FROM rtb_events 
GROUP BY country 
ORDER BY count DESC
📂 Project Structure
Bash

BidPilot-AI/
├── cmd/
│   └── ingest/
│       └── main.go        # Go Entrypoint (Ingestion Service)
├── simulator/
│   ├── generator.py       # Python Traffic Generator
│   └── Dockerfile         # Python Environment
├── ingest-go/
│   ├── go.mod             # Go Dependencies
│   └── Dockerfile         # Go Build Instructions
├── docker-compose.yml     # Orchestration Config
└── README.md              # Project Documentation
🔮 Future Improvements
[ ] Implement Kafka buffer between Ingest and ClickHouse for better scalability.

[ ] Add Redis for deduplication of bid requests.

[ ] Deploy to Kubernetes (K8s).

📝 License
This project is open source and available under the MIT License.

Developed with ❤️ by YASEMİN EREN




