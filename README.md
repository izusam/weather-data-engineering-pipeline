# Weather Data Engineering Pipeline

This project is a complete end-to-end **data engineering pipeline** that ingests real-time weather data, processes it through Apache Kafka, stores it in PostgreSQL, and generates daily CSV summaries using Airflow. It also includes Metabase for interactive data visualization.

---

## Architecture Overview

**Tech Stack:**

- **Go**: Data ingestion and Kafka producer
- **Kafka**: Real-time data streaming
- **Python**: Kafka consumer + PostgreSQL loader
- **PostgreSQL**: Structured data storage
- **Apache Airflow**: ETL and batch processing
- **Metabase**: Data visualization and dashboarding
- **Docker Compose**: Container orchestration

---

## Pipeline Workflow

1. **Ingest Weather Data (Go)**  
   The Go service fetches weather data from [weatherapi.com](https://www.weatherapi.com/) and streams it to Kafka every minute.

2. **Stream Processing (Python)**  
   A Python Kafka consumer reads the weather data from Kafka and inserts it into a PostgreSQL database.

3. **Batch Processing (Airflow)**  
   A daily Airflow DAG exports the previous day's data into a CSV summary.

4. **Visualization (Metabase)**  
   Metabase provides charts and dashboards for exploring weather trends over time.

---

## Getting Started

### Prerequisites

- Docker + Docker Compose
- A [WeatherAPI](https://www.weatherapi.com/) API key

### Setup

1. Clone the repo:
   ```bash
   git clone https://github.com/izusam/weather-data-engineering-pipeline.git
   cd weather-data-engineering-pipeline
   ```
2. Replace YOUR_API_KEY in go-ingest/main.go with your actual API key.

3. Start the stack:
   ```
   docker-compose up --build
   ```
4. Access Services:

   - Airflow: http://localhost:8080

   - Metabase: http://localhost:3000

   - PostgreSQL: localhost:5432 (postgres/postgres)

```
.
├── go-ingest/              # Go service to produce weather data to Kafka
├── python-etl/             # Python consumer to store Kafka data into PostgreSQL
├── airflow/dags/           # Airflow DAG for CSV export
├── docker-compose.yml      # All-in-one service setup
└── README.md

```

## License

### MIT License

## Author

Built by Izuchukwu Samson – Data Engineer
