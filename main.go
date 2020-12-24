package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)

}

// func RootEndPoint(response http.ResponseWriter, request *http.Request) {
// 	response.Write([]byte("Hello World"))
// }

// func main() {
// 	router := mux.NewRouter()
// 	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
// 	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
// 	origins := handlers.AllowedOrigins([]string{"*"})
// 	router.HandleFunc("/", RootEndPoint).Methods("GET")
// 	http.ListenAndServe(":9091", handlers.CORS(headers, methods, origins)(router))
// }
