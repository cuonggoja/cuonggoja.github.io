package main

import (
	"fmt"
	"log"
	"net/http"

	auth "github.com/cuonggoja/web_golang/routes"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error getting env, %v", err)
	// }

	router := httprouter.New()

	router.GET("/", auth.Home)
	router.POST("/auth/login", auth.Login)
	router.POST("/auth/register", auth.Register)

	//sau khi đăng nhập mới thực hiện được chức năng
	//router.GET("/posts", middlewares.CheckJwt(auth.GetListUser))

	fmt.Println("listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
