package web

import (
	"fmt"
	"log"
	"main.go/logic"
	"net/http"
	"os"
	"strconv"
)

func serveHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("./web/index.html")
	if err != nil {
		fmt.Println("GG")
		return
	}
	fmt.Fprintf(w, "%s", file)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	fs := logic.FibonacciService{}

	number, err := strconv.Atoi(r.FormValue("numberValue"))

	if err != nil || number < 0 {
		fmt.Fprintf(w, "Введённое число некорректно!!!")
		return
	}

	fmt.Fprintf(w, "Введённое число: %d\n", number)

	if fs.IsFibonacci(number) {
		prev, next := fs.GetAdjacentFibonacci(number)
		fmt.Fprintf(w, "Предыдущее число Фибоначчи: %d\n", prev)
		fmt.Fprintf(w, "Следующее число Фибоначчи: %d\n", next)
	} else {
		closest := fs.GetNearestFibonacci(number)
		fmt.Fprintf(w, "Ближайшее число Фибоначчи: %d\n", closest)
	}
}

func StartServer() {
	http.HandleFunc("/", serveHandler)
	http.HandleFunc("/result", formHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
