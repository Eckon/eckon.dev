package server

import (
	"net/http"
	"github.com/Eckon/eckon.dev/src/template"
	"github.com/gorilla/mux"
)

// CreateRouter : Creates the mux router, adds the different paths and handlers and returns the whole router
func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(wr http.ResponseWriter, req *http.Request) {
		pageData := template.PageData{
			HeaderInfo: template.GetHeaderInfo(req),
		}
	
		err := template.Template.ExecuteTemplate(wr, "index.gohtml", pageData)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		pageData := template.PageData{
			HeaderInfo: template.GetHeaderInfo(req),
		}
	
		err := template.Template.ExecuteTemplate(wr, "404.gohtml", pageData)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// fetch all the static data (public: js, css, etc.)
	router.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	return router
}
