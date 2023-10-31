package endpoints


// These are the endpoints shortcuts to be shared and used across the microservices
// Only related to request-reply endpoints
// Template is <ServiceName><EndpointName>
type RequestReplyEndpoint = string


const (
	RemittanceFetchRemittanceLogs RequestReplyEndpoint = "remittance.fetch_remittance_logs"
	ScoringPerformCreditScoring RequestReplyEndpoint = "scoring.perform_credit_scoring"
)