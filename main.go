package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./static"))
	fmt.Fprintf(w, "hello234")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

type User struct {
	Id      string
	Balance uint64
}

func encode() {
	u := User{Id: "US123", Balance: 8}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	res, _ := http.Post("https://httpbin.org/post", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)
}

func main() {

	encode()

	http.HandleFunc("/", echoString)

	//http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi2/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})
	http.ListenAndServe(":8080", nil)

}
