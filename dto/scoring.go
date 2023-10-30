package dto

// ForecastLimitAssigned is a body payload for message coming from SCORING.forecast_limit_assigned topic
type ForecastLimitAssigned struct {
	ClientID *string  `validate:"required" json:"client_id"`
	LimitAmount    *float64 `validate:"required" json:"limit_amount"`
}