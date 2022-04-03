package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	// fmt.Println(models.GetAllProducts())
	// fmt.Println(models.InsertProducts(models.Product{Name: "testeeeeee", Quantity: 1, Price: 1.0}))
	// fmt.Println(models.GetAllProducts())
	// fmt.Println(models.GetProduct(2))

	const port = ":3001"

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, map[string]string{"message": "GO e-commerce API REST"})
	})

	r.Route("/products", ProductsRoute)

	http.ListenAndServe(port, r)
}
