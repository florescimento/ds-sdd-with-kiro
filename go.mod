module github.com/distributed-chat-api

go 1.21

require (
	github.com/golang-jwt/jwt/v5 v5.2.0
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/websocket v1.5.1
	github.com/segmentio/kafka-go v0.4.47
	go.mongodb.org/mongo-driver v1.13.1
	github.com/redis/go-redis/v9 v9.3.1
	github.com/minio/minio-go/v7 v7.0.66
	go.etcd.io/etcd/client/v3 v3.5.11
	github.com/prometheus/client_golang v1.18.0
	go.opentelemetry.io/otel v1.21.0
	go.opentelemetry.io/otel/trace v1.21.0
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0
	golang.org/x/crypto v0.17.0
	github.com/google/uuid v1.5.0
)
