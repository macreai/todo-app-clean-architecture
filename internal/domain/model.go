package domain

type ActivityUser struct {
	ID       uint `gorm:"primaryKey"`
	Activity string
}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
