package mw

import (
	"fmt"
	"net/http"
	"time"
)

type Logger struct{}

func (*Logger) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("The logger middleware is executing!")
	t := time.Now()
	next.ServeHTTP(w, r)
	fmt.Printf("Execution time: %s \n", time.Now().Sub(t).String())
}
