package models

type ParcelRes struct {
	SenderUsername string `json:"senderUsername"`
	SenderName     string `json:"senderName"`
	SenderPhone    string `json:"senderPhone"`
	SenderCity     uint64 `json:"senderCity"`
	RecipientName  string `json:"recipientName"`
	Battalion      uint64 `json:"battalion"`
	Company        string `json:"company"`
	Platoon        string `json:"platoon"`
	Department     string `json:"department"`
}

type Parcel struct {
	ID            uint64 `gorm:"primaryKey"`
	UserID        uint64
	CityID        uint64
	City          City `gorm:"foreignKey:CityID"`
	Recipient     string
	BattalionID   uint64
	Battalion     Battalion `gorm:"foreignKey:BattalionID"`
	SpecialUnitID uint64
	SpecialUnit   SpecialUnit `gorm:"foreignKey:SpecialUnitID"`
	Company       string
	Platoon       string
	Department    string
	Sender        string
	SenderNumber  uint64
}
