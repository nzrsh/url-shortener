package services

import (
	"crypto/sha256"
	"encoding/base64"

	_ "github.com/gofiber/fiber/v3/log"
)

func generateShortURL(originalURL string) string {
	hash := sha256.Sum256([]byte(originalURL))
	encoded := base64.URLEncoding.EncodeToString(hash[:])
	// Обрезаем до первых N символов
	return encoded[:6] // Например, 8 символов
}
