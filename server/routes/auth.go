package auth

import (
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	jwt "github.com/cuonggoja/web_golang/auth"
	"github.com/cuonggoja/web_golang/database"
	"github.com/cuonggoja/web_golang/models"
	res "github.com/cuonggoja/web_golang/utils"
	"github.com/julienschmidt/httprouter"
)

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if govalidator.IsNull(username) || govalidator.IsNull(email) || govalidator.IsNull(password) {
		res.JSON(w, 400, "Data is not empty")
		return
	}

	if !govalidator.IsEmail(email) {
		res.JSON(w, 400, "Email is invalid")
		return
	}

	username = models.Santize(username)
	email = models.Santize(email)
	password = models.Santize(password)

	//kết nối database postgres
	db := database.ConnUser()
	defer db.Close()

	//kiểm tra xem tên và email đã tồn tại chưa
	SQLFindUsername := "select username from userlogin where username='" + username + "'"
	SQLFindemail := "select email from userlogin where email='" + email + "'"
	var usernameSQL string
	var emailSQL string

	rowU, errU := db.Query(SQLFindUsername)
	if errU != nil {
		fmt.Printf("Fail to query: %v \n", errU)
		return
	}
	for rowU.Next() {
		rowU.Scan(&usernameSQL)
	}

	rowP, errP := db.Query(SQLFindemail)
	if errP != nil {
		fmt.Printf("Fail to query: %v \n", errP)
		return
	}
	for rowP.Next() {
		rowP.Scan(&emailSQL)
	}

	if username == usernameSQL || email == emailSQL {
		res.JSON(w, 409, "User does exists")
		return
	}

	//băm password
	password, err := models.Hash(password)

	if err != nil {
		res.JSON(w, 500, "Register has failed (Hash password")
		return
	}

	//chèn dữ liệu mới vào postgres
	var newUser models.User
	newUser.Username = username
	newUser.Email = email
	newUser.Password = password

	fmt.Println(newUser.Username)
	fmt.Println(newUser.Email)
	fmt.Println(newUser.Password)

	SQLInsertUser := "INSERT INTO userlogin (username, email, password) VALUES ('" + newUser.Username + "', '" + newUser.Email + "','" + newUser.Password + "' );"
	_, errIN := db.Query(SQLInsertUser)

	if errIN != nil {
		res.JSON(w, 500, "Register has failed (InsertUser)")
		return
	}

}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	if govalidator.IsNull(username) || govalidator.IsNull(password) {
		res.JSON(w, 4, "Data can not empty")
		return
	}

	username = models.Santize(username)
	password = models.Santize(password)

	//kết nối Database
	db := database.ConnUser()

	// tìm username dưới database
	SQLFindUsername := "select username from userlogin where username='" + username + "'"
	var usernameSQL string

	rowU, errU := db.Query(SQLFindUsername)
	if errU != nil {
		fmt.Printf("Fail to query: %v \n", errU)
		return
	}
	for rowU.Next() {
		rowU.Scan(&usernameSQL)
	}
	if govalidator.IsNull(usernameSQL) {
		res.JSON(w, 401, "Username incorrect")
		return
	}

	// tìm password dưới database
	SQLFindPassword := "select password from userlogin where username='" + username + "'"
	var PasswordSQL string

	rowP, errP := db.Query(SQLFindPassword)
	if errP != nil {
		fmt.Printf("Fail to query: %v \n", errP)
		return
	}
	for rowP.Next() {
		rowP.Scan(&PasswordSQL)
	}

	//kiểm tra passwork của client gửi lên khi băm ra có giống trong Database không
	HashedPassword := fmt.Sprintf("%v", PasswordSQL)

	errCheck := models.CheckPasswordHash(HashedPassword, password)

	if errCheck != nil {
		res.JSON(w, 401, "Password incorrect")
		return
	}

	token, errCreate := jwt.Create(username)

	if errCreate != nil {
		res.JSON(w, 500, "Internal server Error")
		return
	}
	fmt.Println("Đăng nhập thành công")
	res.JSON(w, 200, token)
}
