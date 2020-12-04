package auth

import (
	"fmt"
	"net/http"

	"github.com/cuonggoja/web_golang/models"

	"github.com/cuonggoja/web_golang/database"
	"github.com/julienschmidt/httprouter"
)

type Data struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetListUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := database.ConnUser()

	//gửi lệnh truy vấn xuống CSDL
	row, err := db.Query("select * from userlogin")

	if err != nil {
		fmt.Printf("Fail to query: %v \n", err)
	}
	for row.Next() {
		var user models.User

		//lấy dữ liệu từng row trong DB
		row.Scan(&user.Username, &user.Email, &user.Password)

		//sent to client
		fmt.Printf("username: %v, email: %v, password: %v \n", user.Username, user.Email, user.Password)
	}

}
