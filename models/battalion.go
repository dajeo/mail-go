package models

type Battalion struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
