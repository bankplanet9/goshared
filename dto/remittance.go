package dto

import "time"


// RemittanceLog
type RemittanceLog struct {
	ID                  uint `gorm:"primaryKey" json:"id"`  // ID записи
	ClientID            string `json:"client_id" validate:"required"`  // ID клиента в системе
	SentViaSystem       *string `json:"sent_via_system"`  // Система, с помощью которой отправлен/получен перевод
	ReceivedAt          time.Time `json:"received_at" validate:"required"`  // Дата получения перевода
	OriginCountry       string  `json:"origin_country" validate:"required"`  // Страна отправителя
	DestinationCountry  string  `json:"destination_country" validate:"required"`  // Страна получателя
	OriginAmount        float64 `json:"origin_amount" validate:"required"`  // Сумма отправителя
	OriginCurrency      string  `json:"origin_currency" validate:"required"` // Валюта отправителя
	DestinationAmount   float64 `json:"destination_amount" validate:"required"`  // Сумма получателя
	DestinationCurrency string  `json:"destination_currency" validate:"required"` // Валюта получателя

	CreatedUpdatedMixin
}


func (RemittanceLog) TableName() string {
	return "clients__remittanceLogs"
}


// CreateRemittanceLogParams
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


// FetchRemittanceLogsFilters is a struct containing all the filters to be applied to the query
type FetchRemittanceLogsFilters struct {
	ClientID string `json:"client_id" validate:"required"`
}