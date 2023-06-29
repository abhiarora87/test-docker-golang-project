package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var serverID string

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	serverID = strconv.Itoa(rand.Intn(100) + 100)
	logger := log.New(os.Stdout, "", log.Lmicroseconds)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load env file")
	}
	port := os.Getenv("PORT")

	logger.Printf("server %s starting on port %s ...", serverID, port)

	http.HandleFunc("/health", Middle(logger, Health))
	http.HandleFunc("/hello", Middle(logger, Hello))
	http.HandleFunc("/keyword", Middle(logger, Keyword))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		logger.Panicln(err)
	}
}

func Middle(l *log.Logger, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
		f(w, r)
	}
}
