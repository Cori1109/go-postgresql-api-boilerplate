package password

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
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

func ChangedPasswordAfter(jwtIat, passwordChangedAt float64) bool {
	return jwtIat < passwordChangedAt
}

type createPasswordResetTokenResult struct {
	Rt  string
	Prt string
	Pre int64
}

func GenRandomStringFromBytes(size int) (string, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	str := base64.URLEncoding.EncodeToString(b)
	return str, err
}

func CryptString(str string) string {
	b := []byte(str)
	h := sha256.New()
	h.Write(b)
	c := h.Sum(nil)
	crypted := hex.EncodeToString(c) // String representation
	return crypted
}

func CreatePasswordResetToken() createPasswordResetTokenResult {
	resetToken, err := GenRandomStringFromBytes(32)
	if err != nil {
		panic(err)
	}

	password_reset_token := CryptString(resetToken)

	password_reset_expire := time.Now().Unix() + 10*60

	return createPasswordResetTokenResult{
		resetToken,
		password_reset_token,
		password_reset_expire,
	}
}
