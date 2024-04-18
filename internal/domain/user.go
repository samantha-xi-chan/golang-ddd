package domain

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"index;unique"`
	Email string
}
