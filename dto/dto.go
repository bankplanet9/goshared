package dto

import "time"

type CreatedUpdatedMixin struct {
	CreatedAt time.Time `json:"created_at"` // Дата создания записи
	UpdatedAt time.Time `json:"updated_at"` // Дата обновления записи
}