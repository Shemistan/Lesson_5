package converters

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/Shemistan/Lesson_5/internal/models"
)

func ApiModelToServiceModel(req models.AddRequest) *models.User {
	hash := sha256.New()
	hash.Write([]byte(req.AuthParams.Password))
	hashB := hash.Sum(nil)
	hashPassword := hex.EncodeToString(hashB)
	res := &models.User{
		Login:        req.AuthParams.Login,
		PasswordHash: hashPassword,
		Name:         req.Date.Name,
		Age:          int64(req.Date.Age),
	}

	return res
}
