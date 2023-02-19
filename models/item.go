package models

import "time"

type ItemUpdate struct {
	Name     string `json:"name"`
	Unit     string `json:"unit"`
	Quantity int    `json:"quantity"`
}

type Item struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	ParcelID  uint64    `json:"parcelId"`
	Num       uint64    `json:"num"`
	Name      string    `json:"name"`
	Unit      string    `json:"unit"`
	Quantity  int       `json:"quantity"`
	Timestamp time.Time `gorm:"type:timestamp;default:current_timestamp" json:"timestamp"`
}
