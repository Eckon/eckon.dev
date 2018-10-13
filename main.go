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

// for ease handle all handler in this func (even from other go files)
func handleAllFunc(r *mux.Router) {
    // index path
    r.HandleFunc("/", indexHandler).Methods("GET")

    // paths: /phase
    handlePhaseFunc(r)

    // redirect on 404
    r.NotFoundHandler = http.HandlerFunc(indexHandler)
}

func indexHandler(wr http.ResponseWriter, req *http.Request) {
    pageData := pageData{
        HeaderInfo: getHeaderInfo(req),
    }
    err := tpl.ExecuteTemplate(wr, "index.gohtml", pageData)
    if err != nil {
        http.Error(wr, err.Error(), http.StatusInternalServerError)
        return
    }
}

// build the headerInfo so we can use it everywhere and only have one place to edit it
func getHeaderInfo(req *http.Request) (h headerInfo) {
    // nav-bar information
    h = headerInfo{
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

    for e := range h.Navigation {
        // if the path is in the headerInfo -> mark it to highlight it in the front end
        if h.Navigation[e].Path == req.URL.String() {
            h.Navigation[e].Current = true
            return
        }
    }

    return
}
