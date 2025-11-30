package health

import (
	"encoding/json"
	"net/http"
	"time"
)

// Status represents the health status of a component
type Status string

const (
	StatusHealthy   Status = "healthy"
	StatusUnhealthy Status = "unhealthy"
	StatusDegraded  Status = "degraded"
)

// HealthCheck represents the health status of the service
type HealthCheck struct {
	Status    Status            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Service   string            `json:"service"`
	Version   string            `json:"version"`
	Checks    map[string]Status `json:"checks,omitempty"`
}

// Handler provides HTTP handlers for health checks
type Handler struct {
	serviceName string
	version     string
	checks      map[string]func() Status
}

// NewHandler creates a new health check handler
func NewHandler(serviceName, version string) *Handler {
	return &Handler{
		serviceName: serviceName,
		version:     version,
		checks:      make(map[string]func() Status),
	}
}

// AddCheck adds a health check function
func (h *Handler) AddCheck(name string, checkFn func() Status) {
	h.checks[name] = checkFn
}

// LivenessHandler handles liveness probe requests
func (h *Handler) LivenessHandler(w http.ResponseWriter, r *http.Request) {
	health := HealthCheck{
		Status:    StatusHealthy,
		Timestamp: time.Now().UTC(),
		Service:   h.serviceName,
		Version:   h.version,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(health)
}

// ReadinessHandler handles readiness probe requests
func (h *Handler) ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	checks := make(map[string]Status)
	overallStatus := StatusHealthy

	for name, checkFn := range h.checks {
		status := checkFn()
		checks[name] = status
		if status == StatusUnhealthy {
			overallStatus = StatusUnhealthy
		} else if status == StatusDegraded && overallStatus != StatusUnhealthy {
			overallStatus = StatusDegraded
		}
	}

	health := HealthCheck{
		Status:    overallStatus,
		Timestamp: time.Now().UTC(),
		Service:   h.serviceName,
		Version:   h.version,
		Checks:    checks,
	}

	w.Header().Set("Content-Type", "application/json")
	
	statusCode := http.StatusOK
	if overallStatus == StatusUnhealthy {
		statusCode = http.StatusServiceUnavailable
	}
	
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(health)
}
