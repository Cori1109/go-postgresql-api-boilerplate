package sql

import "fmt"

func InsertUser(name, email, password string) string {
	return fmt.Sprintf("INSERT INTO users (name,email,password) VALUES (%s,%s,%s)", name, email, password)
}

func GetUserWithEmail(email string) string {
	return fmt.Sprintf("SELECT user_id,name,email,photo,password FROM users WHERE users.email = %s", email)
}

func GetUserWithId(id int) string {
	return fmt.Sprintf("SELECT user_id,name,email,photo,password,password_changed_at FROM users WHERE users.user_id = %s", id)
}

func UpdateUserPassResetData(id int, prt int64, pre int64) string {
	return fmt.Sprintf("UPDATE users SET users.password_reset_token = %s ,users.password_reset_expires = %s WHERE users.user_id = %s ", prt, pre, id)
}

func GetUserByResetToken(prt int64, now int64) string {
	return fmt.Sprintf("SELECT user_id,name,email,photo,password FROM users WHERE users.password_reset_token = %s AND users.password_reset_expires > %s", prt, now)
}

func ResetPassword(id int, pass int64, pca int64) string {
	return fmt.Sprintf("UPDATE users SET users.password = %s, users.password_reset_token = NULL, users.password_reset_expires = NULL, users.password_changed_at = %s WHERE users.user_id = %s", pass, pca, id)
}

func UpdateUserEmail(id int, email string) string {
	return fmt.Sprintf("UPDATE users SET email = %s WHERE users.user_id = %s", email, id)
}

func UpdateUserName(id int, name string) string {
	return fmt.Sprintf("UPDATE users SET name = %s WHERE users.user_id = %s", name, id)
}

func UpdateUserPhoto(id int, photo string) string {
	return fmt.Sprintf("UPDATE users SET photo = %s WHERE users.user_id = %s", photo, id)
}

func DeleteUser(id int) string {
	return fmt.Sprintf("DELETE FROM users WHERE users.user_id = %s", id)
}
