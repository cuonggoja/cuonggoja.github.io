package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Email    string
	Password string
}
type JSONUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

//hàm Santize dùng để loại bỏ các dấu cách thừa, encode các ký tự đặc biệt của dữ liệu (tránh phần nào các lỗ hổng injection) trước khi lưu vào db .
func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
