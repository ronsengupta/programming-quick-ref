package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

func hello(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"Body": "Hello World!"})
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/pastes/", hello},
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)

	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
