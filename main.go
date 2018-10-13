package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "html/template"
)

type pageData struct {
    HeaderInfo headerInfo
    Data interface{}
}

type headerInfo struct {
    Title string
    Navigation []pathInfo
}

type pathInfo struct {
    Title string
    Path string
    Current bool
}

// nav-bar information
var defaultHeaderInfo = headerInfo{
    Title: "eckon.rocks",
    Navigation: []pathInfo{
        {
            "Phase",
            "/phase",
            false,
        },
        {
            "Phase",
            "/phasde",
            false,
        },
    },
}

var tpl *template.Template

func init() {
    // import templates
    tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
    r := mux.NewRouter()

    // setup the routes from all places
    handleAllFunc(r)

    // fetch all the static data (public: js, css, etc.)
    r.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

    s := http.Server{
        Addr: ":80",
        Handler: r,
    }
    s.ListenAndServe()
}

func handleAllFunc(r *mux.Router) {
    // index path
    r.HandleFunc("/", indexHandler).Methods("GET")

    // paths: /phase
    handlePhaseFunc(r)

    // redirect on 404
    r.NotFoundHandler = http.HandlerFunc(indexHandler)
}

func indexHandler(wr http.ResponseWriter, req *http.Request) {
    defaultHeaderInfo = updateCurrentPage(defaultHeaderInfo, req)
    pageData := pageData{
        HeaderInfo: defaultHeaderInfo,
    }
    err := tpl.ExecuteTemplate(wr, "index.gohtml", pageData)
    if err != nil {
        http.Error(wr, err.Error(), http.StatusInternalServerError)
        return
    }
}

func updateCurrentPage(header headerInfo, req *http.Request) (h headerInfo) {
    h = header
    for e := range h.Navigation {
        // if the path is in the headerInfo -> mark it to highlight it in the front end
        h.Navigation[e].Current = h.Navigation[e].Path == req.URL.String()
    }

    return
}
