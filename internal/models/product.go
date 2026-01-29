package models

type Product struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Price int64  `gorm:"not null"`
}
