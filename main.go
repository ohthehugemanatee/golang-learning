package main

import (
	"net/http"

	"github.com/ohthehugemanatee/golang-learning/injection"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	injection.Greet(w, "my man")
}

func __main() {
	http.ListenAndServe(":9000", http.HandlerFunc(MyGreeterHandler))
}
