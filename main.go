package main

import (
	"fmt"
	"net/http"
	"oopLab1/core/user"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	user.NewUserService(*user.NewCustomerRepositoryPostgres())
	// http.HandleFunc("/", greet)
	// http.ListenAndServe(":8080", nil)
}
