package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"ewallet/route"
	"ewallet/utill"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func newLoggingHandler(dst io.Writer) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(dst, h)
	}
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func setHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	})
}

func createLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("test")
		next.ServeHTTP(res, req)
	})
}

func main() {
	r := mux.NewRouter()
	r.Use(setHeader)
	r.Use(createLog)
	utill.DBConnection()
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}
	loggingHandler := newLoggingHandler(logFile)
	finalHandler := http.HandlerFunc(final)
	serveMux := http.NewServeMux()
	serveMux.Handle("/", loggingHandler(finalHandler))
	log.Println("Listening on :8080..")

	r = r.PathPrefix("/api/v1").Subrouter()
	err = http.ListenAndServe(":8080", route.Router(r))
	log.Fatal(err)
}
