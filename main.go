package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"

	"github.com/dee-ex/gocarest/router"
)

func greeting() {
	fmt.Println()
	fmt.Println("  /$$$$$$             /$$$$$$   /$$$$$$  /$$$$$$$  /$$$$$$$$  /$$$$$$  /$$$$$$$$")
	fmt.Println(" /$$__  $$           /$$__  $$ /$$__  $$| $$__  $$| $$_____/ /$$__  $$|__  $$__/")
	fmt.Println("| $$  \\__/  /$$$$$$ | $$  \\__/| $$  \\ $$| $$  \\ $$| $$      | $$  \\__/   | $$   ")
	fmt.Println("| $$ /$$$$ /$$__  $$| $$      | $$$$$$$$| $$$$$$$/| $$$$$   |  $$$$$$    | $$   ")
	fmt.Println("| $$|_  $$| $$  \\ $$| $$      | $$__  $$| $$__  $$| $$__/    \\____  $$   | $$   ")
	fmt.Println("| $$  \\ $$| $$  | $$| $$    $$| $$  | $$| $$  \\ $$| $$       /$$  \\ $$   | $$   ")
	fmt.Println("|  $$$$$$/|  $$$$$$/|  $$$$$$/| $$  | $$| $$  | $$| $$$$$$$$|  $$$$$$/   | $$   ")
	fmt.Println(" \\______/  \\______/  \\______/ |__/  |__/|__/  |__/|________/ \\______/    |__/   ")
	fmt.Println()
}

func serveHTTP(port string) {
	r := router.NewRouter()
	c := router.NewCors()

	enR := handlers.LoggingHandler(os.Stdout, c.Handler(r))

	greeting()
	fmt.Println("(ɔ◔‿◔)ɔ ♥ Serving at: http://localhost:" + port)

	log.Fatal(http.ListenAndServe(":"+port, enR))
}

func main() {
	var port string
	flag.StringVar(&port, "p", "6969", "Server port")
	flag.Parse()

	serveHTTP(port)
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
