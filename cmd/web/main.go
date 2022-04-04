package main

import (
	"crypto/rand"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/pauloeduardods/e-commerce/pkg/config"
)

var AppConfig config.AppConfig

func main() {
	randBytes := make([]byte, 64)
	_, err := rand.Read(randBytes)
	if err != nil {
		panic(err)
	}
	AppConfig.HmacSecret = randBytes

	const port = ":3001"

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, map[string]string{
			"message": "GO e-commerce API REST",
		})
	})

	r.Route("/products", ProductsRoute)
	r.Route("/login", LoginRoute)

	http.ListenAndServe(port, r)
}
