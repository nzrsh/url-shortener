package services

type Link struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	SourseLink string `gorm:"not null"`
	ShortLink  string `gorm:"not null"`
}
