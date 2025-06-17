# Go Observability Demo

A comprehensive demonstration of observability practices in a Go microservice using Prometheus, Grafana, and AlertManager.

## Project Overview

This project demonstrates a complete observability stack for a Go microservice, including:
- Metrics collection with Prometheus
- Visualization with Grafana
- Alerting with AlertManager
- Horizontal Pod Autoscaling
- Kubernetes Ingress configuration
- PostgreSQL database integration

## Architecture

The project consists of the following components:

### Core Services
- **Product Service**: A Go microservice that manages product data
- **PostgreSQL**: Database for storing product information
- **Prometheus**: Metrics collection and storage
- **Grafana**: Metrics visualization and dashboards
- **AlertManager**: Alert handling and routing

### Infrastructure
- **Kubernetes**: Container orchestration
- **Nginx Ingress**: External access management
- **Horizontal Pod Autoscaler**: Automatic scaling based on metrics

## Prerequisites

- Docker Desktop
- Kubernetes cluster (local or remote)
- kubectl configured
- Go 1.x
- Make (optional)

## Project Structure

```
.
├── cmd/
│   └── server/         # Main application entry point
├── internal/
│   ├── handler/        # HTTP handlers
│   ├── metrics/        # Prometheus metrics
│   ├── model/          # Data models
│   └── repository/     # Database operations
├── k8s/
│   ├── alertmanager/   # AlertManager configuration
│   ├── grafana/        # Grafana configuration
│   ├── ingress/        # Ingress configuration
│   ├── prometheus/     # Prometheus configuration
│   └── *.yaml          # Kubernetes manifests
├── docker-compose.yml  # Local development setup
├── dockerfile         # Container configuration
└── go.mod             # Go dependencies
```

## Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/go-observability-demo.git
   cd go-observability-demo
   ```

2. **Build the application**
   ```bash
   go build -o productservice ./cmd/server
   ```

3. **Deploy to Kubernetes**
   ```bash
   kubectl apply -f k8s/
   ```

4. **Access the services**
   - Grafana: http://grafana.local
   - Prometheus: http://prometheus.local
   - Product Service: http://productservice.local

## Monitoring Stack

### Prometheus
- Collects metrics from the product service
- Stores time-series data
- Provides querying capabilities

### Grafana
- Visualizes metrics from Prometheus
- Custom dashboards for monitoring
- Alert visualization

### AlertManager
- Handles alerts from Prometheus
- Routes alerts to appropriate channels
- Manages alert grouping and silencing

## Kubernetes Configuration

### Ingress
- Routes external traffic to services
- Supports multiple domains
- SSL/TLS termination (if configured)

### Horizontal Pod Autoscaling
- Automatically scales the product service
- Based on CPU and memory metrics
- Configurable scaling policies

## Development

### Local Development
```bash
docker-compose up -d
```

### Building Docker Image
```bash
docker build -t go-observability-demo-productservice:latest .
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
