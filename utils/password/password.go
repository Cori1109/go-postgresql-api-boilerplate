package password

import (
	"crypto/rand"
	"crypto/sha256"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(currentPassword, newPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(newPassword))
	return err == nil
}

func changedPasswordAfter(jwtTimestamp, passwordChangedAt float64) bool {
	return jwtTimestamp < passwordChangedAt
}

func GenRandomBytes(size int) ([]byte, error) {
	blk := make([]byte, size)
	_, err := rand.Read(blk)
	return blk, err
}

type createPasswordResetTokenResult struct {
	resetToken            []byte
	password_reset_token  []byte
	password_reset_expire int64
}

func CreatePasswordResetToken() createPasswordResetTokenResult {

	resetToken, err := GenRandomBytes(32)
	if err != nil {
		panic(err)
	}
	h := sha256.New()
	h.Write(resetToken)

	password_reset_token := h.Sum(nil)

	password_reset_expire := time.Now().Unix() + 10*60

	return createPasswordResetTokenResult{
		resetToken,
		password_reset_token,
		password_reset_expire,
	}
}
