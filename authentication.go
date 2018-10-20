package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "github.com/gorilla/sessions"
)

// global variable so we can check it everywhere
var store = sessions.NewCookieStore([]byte("LOGIN_KEY"))

func handleAuthenticationFunc(r *mux.Router) {
    r.HandleFunc("/authentication", handleLoginPage).Methods("GET")
    r.HandleFunc("/authentication/login", handleLoginAttempt).Methods("POST")
}

func handleLoginPage(wr http.ResponseWriter, req *http.Request) {
    if checkLoginStatus(req) {
        // log out if we visit the login site again (while being logged in)
        session, _:= store.Get(req, "session")
        session.Values["status"] = "out"
        session.Save(req, wr)
        http.Redirect(wr, req, "/", http.StatusSeeOther)
    }

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
    session, _:= store.Get(req, "session")

    // just for only one user (later maybe more with database and general checking)
    if req.Form["name"][0] == "eckon" && req.Form["password"][0] == "123" {
        // safe the status of the one user
        session.Values["status"] = "in"
        session.Save(req, wr)
        http.Redirect(wr, req, "/", http.StatusSeeOther)
    } else {
        http.Redirect(wr, req, "/authentication", http.StatusSeeOther)
    }
}

func checkLoginStatus(req *http.Request) bool {
    session, _ := store.Get(req, "session")

    return session.Values["status"] == "in"
}
