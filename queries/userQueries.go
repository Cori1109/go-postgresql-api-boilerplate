package queries

import "fmt"

func InsertUser(name, email, password string) string {
	return fmt.Sprintf("INSERT INTO users_user (name,email,password) VALUES ('%s','%s','%s') RETURNING id,name,email", name, email, password)
}

func GetUserWithEmail(email string) string {
	return fmt.Sprintf("SELECT id,name,email,password FROM users_user WHERE users_user.email = '%s'", email)
}

func GetUserWithId(id int) string {
	return fmt.Sprintf("SELECT id,name,email,photo,password,password_changed_at FROM users_user WHERE users_user.id = %d", id)
}

func UpdateUserPassResetData(id int, prt string, pre int64) string {
	return fmt.Sprintf("UPDATE users_user SET password_reset_token = '%s' ,password_reset_expires = %d WHERE id = %d", prt, pre, id)
}

func GetUserByResetToken(prt int64, now int64) string {
	return fmt.Sprintf("SELECT id,name,email,photo,password FROM users_user WHERE users_user.password_reset_token = '%s' AND users_user.password_reset_expires > %d", prt, now)
}

func ResetPassword(id int, pass int64, pca int64) string {
	return fmt.Sprintf("UPDATE users_user SET users_user.password = '%s', users_user.password_reset_token = NULL, users_user.password_reset_expires = NULL, users_user.password_changed_at = '%s' WHERE users_user.id = %d", pass, pca, id)
}

func UpdateUserEmail(id int, email string) string {
	return fmt.Sprintf("UPDATE users_user SET email = '%s' WHERE users_user.id = %d", email, id)
}

func UpdateUserName(id int, name string) string {
	return fmt.Sprintf("UPDATE users_user SET name = '%s' WHERE users_user.id = %d", name, id)
}

func UpdateUserPhoto(id int, photo string) string {
	return fmt.Sprintf("UPDATE users_user SET photo = '%s' WHERE users_user.id = %d", photo, id)
}

func DeleteUser(id int) string {
	return fmt.Sprintf("DELETE FROM users_user WHERE users_user.id = %d", id)
}
