package main

import "net/http"

func main() {
	api := &api{addr: ":8080"}

	//Initialize the ServeMux
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getusersHandler)
	mux.HandleFunc("POST /users", api.createusersHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
