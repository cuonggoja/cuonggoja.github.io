package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cuonggoja/web_golang/middlewares"
	auth "github.com/cuonggoja/web_golang/routes"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error getting env, %v", err)
	// }

	router := httprouter.New()

	router.POST("/auth/login", auth.Login)
	router.POST("/auth/register", auth.Register)

	router.GET("/posts", middlewares.CheckJwt(auth.GetListUser))

	fmt.Println("listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
