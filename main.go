package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "html/template"
    "fmt"
    "os"
    "github.com/gorilla/handlers"
    "strings"
)

type pageData struct {
    HeaderInfo headerInfo
    Data       interface{}
}

type headerInfo struct {
    Title      string
    Navigation []pathInfo
    OnHomePage bool
}

type pathInfo struct {
    Title   string
    Path    string
    Class   string
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

    // logging access logs
    accessLog, err := os.Create("access_log.txt")
    if err != nil {
        panic(err)
    }
    defer func() {
        if err := accessLog.Close(); err != nil {
            panic(err)
        }
    }()

    // start go routine to redirect every port 80 (http) to 443 (https)
    go func() {
       http.ListenAndServe(":80", http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
           hostDomain := strings.Split(req.Host, ":")[0]
           http.Redirect(wr, req, "https://" + hostDomain + ":443" + req.URL.Path, http.StatusTemporaryRedirect)
       }))
    }()

    s := http.Server{
        Addr:    ":443",
        Handler: handlers.LoggingHandler(accessLog, r),
    }

    fmt.Println("Running")
    err = s.ListenAndServeTLS("/etc/letsencrypt/live/eckon.rocks/fullchain.pem",
        "/etc/letsencrypt/live/eckon.rocks/privkey.pem")
    if err != nil {
        panic(err)
    }
}

// for ease handle all handler in this func (even from other go files)
func handleAllFunc(r *mux.Router) {
    // index path
    r.HandleFunc("/", indexHandler).Methods("GET")
    // paths: /phase (in phase.go)
    handlePhaseFunc(r)
    // paths: /authentication (in authentication.go)
    handleAuthenticationFunc(r)
    // paths: /calendar (in calendar.go)
    handleCalendarFunc(r)
    // paths: /release (in release.go)
    handleReleaseFunc(r)
    // redirect on 404
    r.NotFoundHandler = http.HandlerFunc(indexHandler)
}

func indexHandler(wr http.ResponseWriter, req *http.Request) {
    pageData := pageData{
        HeaderInfo: getHeaderInfo(req),
    }
    if err := tpl.ExecuteTemplate(wr, "index.gohtml", pageData); err != nil {
        http.Error(wr, err.Error(), http.StatusInternalServerError)
        return
    }
}

// build the headerInfo so we can use it everywhere and only have one place to edit it
func getHeaderInfo(req *http.Request) (h headerInfo) {
    // update the title if the user is already logged in
    status := pathInfo{"Login", "/authentication", "authentication"}
    if checkLoginStatus(req) {
        status.Title = "Logout"
    }

    // nav-bar information
    h = headerInfo{
        Title: "eckon.rocks",
        Navigation: []pathInfo{
            {
                "Phase",
                "/phase",
                "",
            },
            {
                "Calendar",
                "/calendar",
                "",
            },
            {
                "Release",
                "/release",
                "",
            },
            status,
        },
        OnHomePage: false,
    }

    for e := range h.Navigation {
        // if the path is in the headerInfo -> mark it to highlight it in the front end (append class list)
        if h.Navigation[e].Path == req.URL.String() {
            h.Navigation[e].Class += " highlighted-content"
            return
        }
    }
    // if no highlight -> it's on the home page
    h.OnHomePage = true

    return
}
