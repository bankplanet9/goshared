package dto

import "time"

type Client struct {
	ID string `json:"id" validate:"required" gorm:"primaryKey"` // ID клиента

	// Scoring
	Limit           *float64   `json:"limit"`             // Сумма лимита
	LimitAssignedAt *time.Time `json:"limit_assigned_at"` // Дата и время определения лимита

	// Credit History
	HasGoodCreditHistory   bool       `json:"has_good_credit_history"`   // Индикатор качества кредитной истории (хорошая, плохая). Реализация определена микросервисом credit-history
	CreditHistoryCheckedAt *time.Time `json:"credit_history_checked_at"` // Дата и время осуществления проверки кредитной истории

	RemittanceLogs []RemittanceLog `json:"-"`

	CreatedUpdatedMixin
}

func (Client) TableName() string {
	return "clients_clients"
}

// IsGood determines whether the client is good according to these factors:
//
// 1. A client has limit
//
// 2. A client has a good credit history
func (c Client) IsGood() bool {
	return c.HasGoodCreditHistory && c.Limit != nil
}
