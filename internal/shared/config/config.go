package config

import (
	"os"
	"strconv"
	"time"
)

// Config holds all application configuration
type Config struct {
	Service    ServiceConfig
	Kafka      KafkaConfig
	MongoDB    MongoDBConfig
	Redis      RedisConfig
	MinIO      MinIOConfig
	Etcd       EtcdConfig
	Observability ObservabilityConfig
}

// ServiceConfig holds service-specific configuration
type ServiceConfig struct {
	Name string
	Port int
	Env  string
}

// KafkaConfig holds Kafka configuration
type KafkaConfig struct {
	Brokers []string
	GroupID string
}

// MongoDBConfig holds MongoDB configuration
type MongoDBConfig struct {
	URI      string
	Database string
	Timeout  time.Duration
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// MinIOConfig holds MinIO configuration
type MinIOConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
	Bucket    string
}

// EtcdConfig holds etcd configuration
type EtcdConfig struct {
	Endpoints []string
	Timeout   time.Duration
}

// ObservabilityConfig holds observability configuration
type ObservabilityConfig struct {
	PrometheusPort int
	JaegerEndpoint string
	LogLevel       string
}

// LoadConfig loads configuration from environment variables
func LoadConfig(serviceName string) *Config {
	return &Config{
		Service: ServiceConfig{
			Name: serviceName,
			Port: getEnvAsInt("SERVICE_PORT", 8080),
			Env:  getEnv("ENVIRONMENT", "development"),
		},
		Kafka: KafkaConfig{
			Brokers: []string{getEnv("KAFKA_BROKERS", "localhost:9092")},
			GroupID: getEnv("KAFKA_GROUP_ID", serviceName+"-group"),
		},
		MongoDB: MongoDBConfig{
			URI:      getEnv("MONGODB_URI", "mongodb://localhost:27017"),
			Database: getEnv("MONGODB_DATABASE", "chat_platform"),
			Timeout:  time.Duration(getEnvAsInt("MONGODB_TIMEOUT_SEC", 10)) * time.Second,
		},
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		MinIO: MinIOConfig{
			Endpoint:  getEnv("MINIO_ENDPOINT", "localhost:9000"),
			AccessKey: getEnv("MINIO_ACCESS_KEY", "minioadmin"),
			SecretKey: getEnv("MINIO_SECRET_KEY", "minioadmin"),
			UseSSL:    getEnvAsBool("MINIO_USE_SSL", false),
			Bucket:    getEnv("MINIO_BUCKET", "chat-files"),
		},
		Etcd: EtcdConfig{
			Endpoints: []string{getEnv("ETCD_ENDPOINTS", "localhost:2379")},
			Timeout:   time.Duration(getEnvAsInt("ETCD_TIMEOUT_SEC", 5)) * time.Second,
		},
		Observability: ObservabilityConfig{
			PrometheusPort: getEnvAsInt("PROMETHEUS_PORT", 9090),
			JaegerEndpoint: getEnv("JAEGER_ENDPOINT", "http://localhost:14268/api/traces"),
			LogLevel:       getEnv("LOG_LEVEL", "info"),
		},
	}
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return boolVal
		}
	}
	return defaultValue
}
