package main

import (
	"encoding/json"
	"fmt"
	"github.com/lilleswing/make_squares/pkg/make_squares"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	var nr make_squares.NumberRequest

	err := json.NewDecoder(r.Body).Decode(&nr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(nr.Value)
	strResponse := make_squares.SquareIt(nr.Value)
	w.Write([]byte(strResponse))
}
func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
