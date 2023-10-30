package dto

// ForecastLimitAssigned is a body payload for message coming from SCORING.forecast_limit_assigned topic
type ForecastLimitAssigned struct {
	ClientID *string  `validate:"required"`
	Limit    *float64 `validate:"required"`
}