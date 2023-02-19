package models

import "time"

type StatusRes struct {
	Message uint64 `json:"message"`
	City    uint64 `json:"city"`
}

type Status struct {
	ID        uint64 `gorm:"primaryKey"`
	ParcelID  uint64
	Parcel    Parcel `gorm:"foreignKey:ParcelID"`
	MessageID uint64
	Message   Message `gorm:"foreignKey:MessageID"`
	CityID    uint64
	City      City      `gorm:"foreignKey:CityID"`
	Timestamp time.Time `gorm:"type:timestamp;default:current_timestamp"`
}
