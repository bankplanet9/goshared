package dto

import "time"

type Client struct {
	ID string `json:"id" validate:"required" gorm:"primaryKey"` // ID клиента

	// IsIdentified shows whether the client is identified or not
	IsIdentified bool `json:"is_identified"` // Идентифицирован ли клиент

	// Scoring
	Limit            *float64   `json:"limit"`              // Сумма лимита
	LimitAssignedAt  *time.Time `json:"limit_assigned_at"`  // Дата и время определения лимита
	LimitRequestedAt *time.Time `json:"limit_requested_at"` // Дата и время запроса на определение лимита

	// Credit History
	HasGoodCreditHistory     *bool      `json:"has_good_credit_history"`     // Индикатор качества кредитной истории (хорошая, плохая). Реализация определена микросервисом credit-history
	CreditHistoryRequestedAt *time.Time `json:"credit_history_requested_at"` // Дата и время запроса на проверку кредитной истории
	CreditHistoryCheckedAt   *time.Time `json:"credit_history_checked_at"`   // Дата и время осуществления проверки кредитной истории

	// Quota related
	// Quota owners are customers with bad credit history, but Planet9 assigns them a limit up to X amount,
	// defined in system configuration providing them an opportunity to improve their credit history
	HasQuotaAssigned *bool      `json:"has_quota_assigned"` // Получил квоту будучи плохим клиентом (да, нет)
	QuotaAssignedAt  *time.Time `json:"quota_assigned_at"`  // Дата и время назначения квоты

	RemittanceLogs []RemittanceLog `json:"-"`

	CreatedUpdatedMixin
}

func (Client) TableName() string {
	return "clients_clients"
}

// IsPreApproved determines whether we can provide a loan to the client or not.
func (c Client) IsPreApproved() bool {
	// If the client is not identified, they must at least have a limit to be pre-approved.
	if !c.IsIdentified && c.Limit != nil {
		return true
	}

	// If the client is identified and has a good credit history and limit assigned, they are pre-approved.
	if c.IsIdentified && c.HasGoodCreditHistory != nil && *c.HasGoodCreditHistory && c.Limit != nil {
		return true
	}

	// If the client is identified and has a bad credit history with a quota & limit assigned, they are pre-approved.
	if c.IsIdentified && c.HasGoodCreditHistory != nil && !*c.HasGoodCreditHistory && *c.HasQuotaAssigned && c.Limit != nil {
		return true
	}

	// Default case: not pre-approved.
	return false
}


// AssignLimit assigns a limit to client
func (c *Client) AssignLimit(limit float64) {
	now := time.Now().UTC()
	c.Limit = &limit
	c.LimitAssignedAt = &now
}


// AssignQuota assigns a limit to client and mark it as a quota
func (c *Client) AssignQuota(limit float64) {
	now := time.Now().UTC()
	quotaAssigned := true
	c.Limit = &limit
	c.HasQuotaAssigned = &quotaAssigned
	c.QuotaAssignedAt = &now
}
