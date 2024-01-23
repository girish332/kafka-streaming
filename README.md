# Kafka Streaming

This project is a Kafka streaming application written in Go. It uses the confluent-kafka-go library to interact with Kafka.

## Prerequisites

- Docker
- Docker Compose
- Go

## Getting Started

1. Clone the repository:
    
    ```bash
   git clone https://github.com/girish332/kafka-streaming.git
   ```
   
2. Start the Kafka cluster:

    ```bash
   docker-compose up --build -d
   ```

This command will start all the services defined in your `docker-compose.yml` file in detached mode.

4. To stop the services, use:

    ```bash
   docker-compose down
   ```

## Project Structure

- `processor/main.go`: This file contains the code for the Kafka consumer that processes orders.
- `datateam/main.go`: This file contains the code for the Kafka consumer that reads orders for the data team.
- `main.go`: This file contains the code for the Kafka producer that places orders.