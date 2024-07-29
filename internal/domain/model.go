package domain

type ActivityUser struct {
	ID       uint `gorm:"primaryKey"`
	Activity string
}
