package models

type City struct {
	ID   uint64 `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
