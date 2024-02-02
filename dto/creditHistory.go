package dto

// CreditHistoryLog represents an ORM model for credit history checks stored in DB
type CreditHistoryLog struct {
	ID       uint   `json:"id" gorm:"primaryKey"` // ID записи в базе данных
	ClientID string `json:"client_id"`            // ID клиента

	ReportID          *string `json:"report_id" gorm:"unique"` // ID отчета кредитной истории в системе провайдера кредитной истории
	IsBadBorrower     *bool   `json:"is_bad_client"`           // Плохой заёмщик (binary)
	BadBorrowerReason *string `json:"bad_borrower_reason"`     // Критерии, по которым клиент был определен как плохой заёмщик

	// Validity
	IsValid       *bool   `json:"is_valid"`       // Флаг валидности клиента
	InvalidReason *string `json:"invalid_reason"` // Причина, по которой клиент не считается валидным

	CreatedUpdatedMixin
}

func (CreditHistoryLog) TableName() string {
	return "credithistory_credit_history_logs"
}

type CheckCreditHistoryRequest struct {
	ClientID string `json:"client_id" validate:"required"`
}
