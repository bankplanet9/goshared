package endpoints

import "fmt"

// These are the endpoints shortcuts to be shared and used across the microservices
// Only related to pub-sub endpoints
// Template is <ServiceName><EndpointName>
type PubSubEndpoint = string

const (
	RemittanceCreateRemittanceLog PubSubEndpoint = "REMITTANCE.create_remittance_log"
	ScoringForecastLimitAssigned  PubSubEndpoint = "SCORING.forecast_limit_assigned"
)

func ScoringForecast(clientID string) PubSubEndpoint {
	return PubSubEndpoint(fmt.Sprintf("FORECAST.%v", clientID))
}
