package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "github.com/gorilla/sessions"
)

// global variable so we can check it everywhere
var store = sessions.NewCookieStore([]byte("LOGIN_KEY"))

func handleLoginFunc(r *mux.Router) {
    r.HandleFunc("/login", handleLoginPage).Methods("GET")
    r.HandleFunc("/login", handleLoginAttempt).Methods("POST")
}

func handleLoginPage(wr http.ResponseWriter, req *http.Request) {
    pageData := pageData{
        HeaderInfo: getHeaderInfo(req),
    }
    if err := tpl.ExecuteTemplate(wr, "login.gohtml", pageData); err != nil {
        http.Error(wr, err.Error(), http.StatusInternalServerError)
        return
    }
}

func handleLoginAttempt(wr http.ResponseWriter, req *http.Request) {
    req.ParseForm()

    // just for only one user (later maybe more with database and general checking)
    if req.Form["name"][0] == "eckon" && req.Form["password"][0] == "123" {

        session, _:= store.Get(req, "session")
        // safe the status of the one user
        session.Values["status"] = "in"
        session.Save(req, wr)
        http.Redirect(wr, req, "/", http.StatusSeeOther)
    } else {
        http.Redirect(wr, req, "/login", http.StatusSeeOther)
    }
}

func checkLoginStatus(req *http.Request) bool {
    session, _ := store.Get(req, "session")

    return session.Values["status"] == "in"
}
