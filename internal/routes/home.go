package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gungun974.com/viteintegration/resources/views/pages/home"
)

func IndexRouter() http.Handler {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		component := home.IndexPage()
		err := component.Render(r.Context(), w)
		if err != nil {
			handleHTTPError(err, w, r)
		}
	})

	FileRouter(router)

	return router
}
