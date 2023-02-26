package models

type Message struct {
	ID      int `gorm:"primaryKey"`
	Message string
}
