package dto


type Client struct {
	ID string `json:"id" validate:"required" gorm:"primaryKey"` // ID клиента
	Limit *float64 `json:"limit"`

	RemittanceLogs []RemittanceLog `json:"-"`

	CreatedUpdatedMixin
}

func (Client) TableName() string {
	return "clients_clients"
}