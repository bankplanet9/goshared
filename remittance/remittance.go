package remittance

import "time"


// Модель запроса на создание записи международных переводов
type CreateRemittanceLogParams struct {
	ClientID            string `json:"client_id" validate:"required"`
	SentViaSystem       *string `json:"sent_via_system"`
	ReceivedAt          time.Time `json:"received_at" validate:"required"`
	OriginCountry       string  `json:"origin_country" validate:"required"`
	DestinationCountry  string  `json:"destination_country" validate:"required"`
	OriginAmount        float64 `json:"origin_amount" validate:"required"`
	OriginCurrency      string  `json:"origin_currency" validate:"required"`
	DestinationAmount   float64 `json:"destination_amount" validate:"required"`
	DestinationCurrency string  `json:"destination_currency" validate:"required"`
}


// Запись международного перевода
type RemittanceLog struct {
	ID                  uint `gorm:"primaryKey" json:"id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	CreateRemittanceLogParams
}