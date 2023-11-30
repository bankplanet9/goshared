package endpoints

import "fmt"

// These are the endpoints shortcuts to be shared and used across the microservices
// Only related to pub-sub endpoints
// Template is <ServiceName><EndpointName>
type PubSubEndpoint = string

const (
	RemittanceCreateRemittanceLog       PubSubEndpoint = "REMITTANCE.create_remittance_log"
	RemittanceCreateRemittanceLogsBatch PubSubEndpoint = "REMITTANCE.create_remittance_logs_batch"
	RemittanceClientUpdated             PubSubEndpoint = "REMITTANCE.client_updated"
	ScoringForecastLimitAssigned        PubSubEndpoint = "SCORING.forecast_limit_assigned"
	CreditHistoryCheck                  PubSubEndpoint = "CREDIT_HISTORY.check"
	CreditHistoryChecked                PubSubEndpoint = "CREDIT_HISTORY.checked"
	LoanApplicationCompleted             PubSubEndpoint = "LOANAPP.completed"
)

func ScoringForecast(clientID string) PubSubEndpoint {
	return PubSubEndpoint(fmt.Sprintf("FORECAST.%v", clientID))
}
