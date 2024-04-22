package web

import (
	"fmt"
	"log"
	"main.go/logic"
	"net/http"
	"os"
	"strconv"
)

func render(w http.ResponseWriter, r *http.Request, filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("GG")
		return
	}
	fmt.Fprintf(w, "%s", file)
}

func serveHandler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "./web/index.html")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	render(w, r, "./web/result.html")

	fs := logic.FibonacciService{}

	number, err := strconv.Atoi(r.FormValue("numberValue"))

	if err != nil || number < 0 {

		fmt.Fprintf(w, "<h2>Введённое число некорректно!!!</h2>")

	} else {

		fmt.Fprintf(w, "<a>Введённое число: %d\n</a>", number)

		if fs.IsFibonacci(number) {
			prev, next := fs.GetAdjacentFibonacci(number)
			fmt.Fprintf(w, "<a>Предыдущее число Фибоначчи: %d</a>", prev)
			fmt.Fprintf(w, "<a>Следующее число Фибоначчи: %d</a>", next)
		} else {
			closest := fs.GetNearestFibonacci(number)
			fmt.Fprintf(w, "<a>Ближайшее число Фибоначчи: %d</a>", closest)
		}

	}

	fmt.Fprintf(w, "<a href=\"/\">Вернуться</a>")
}

func StartServer() {
	http.HandleFunc("/", serveHandler)
	http.HandleFunc("/result", formHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
