package router

import (
	"github.com/go-chi/chi"
	eventing "github.com/haagor/RGB/domains/eventing/pkg"
)

func NewDefault() func(r chi.Router) {
	store := eventing.NewInMemoryEventStore()
	s := eventing.NewWithDefaultLock(store)

	return New(&s)
}

func New(s *eventing.Source) func(r chi.Router) {
	return func(r chi.Router) {
		r.Use()
		r.Post("/create", NewCreatePotHandler(s))
		r.Post("/addcolour", NewAddColourHandler(s))
		r.Get("/", NewGetPotHandler(s))
	}
}
