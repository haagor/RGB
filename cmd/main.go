package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/haagor/RGB/domains/colour/pkg/router"
	"github.com/rs/cors"
)

func Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/pot", router.NewDefault())

	co := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})
	handler := co.Handler(r)
	err := http.ListenAndServe(":2222", handler)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return nil
}

func main() {
	err := Run()
	if err != nil {
		panic(err.Error())
	}
}
