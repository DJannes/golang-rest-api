package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gitlab.com/janneseffendi/rest-api/internal/controller"
)

// https://go-chi.io/#/pages/middleware
// Mount => Menggabungkan beberapa Mux menjadi satu
// newRouter => membuat mux baru
// router.Route => create new mux with handler and Mount it to main mux.
func main() {
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	r.Route("/api/v1", func(r chi.Router) {
		// controller.AddPublicRouter(r)
		controller.AddAuthRouter(r)
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println(err.Error())
	}
}
