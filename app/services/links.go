package services

import (
	"errors"

	"gorm.io/gorm"
)

func CreateShortURL(url string) *Link {
	shortUrl := generateShortURL(url)

	link := &Link{
		SourseLink: url,
		ShortLink:  shortUrl,
	}
	SqlDB.Create(link)
	return link
}

func GetLongUrl(shortUrl string) (string, error) {
	link := &Link{}

	err := SqlDB.Where("short_link = ?", shortUrl).First(&link).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "no", nil // Возвращаем nil для пользователя и nil для ошибки
		}
		return "nil", err // Возвращаем ошибку, если она другая
	}
	// Сохраняем изменения
	return link.SourseLink, nil
}
