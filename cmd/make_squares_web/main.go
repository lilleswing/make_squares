package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


func SquareIt(a int) string {
	return strconv.Itoa(a * a)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	var nr NumberRequest

	err := json.NewDecoder(r.Body).Decode(&nr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(nr.Value)
	strResponse := SquareIt(nr.Value)
	w.Write([]byte(strResponse))
}
func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
