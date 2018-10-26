package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "github.com/gorilla/sessions"
    "io/ioutil"
    "fmt"
    "encoding/json"
)

// global variable so we can check it everywhere
var store = sessions.NewCookieStore([]byte("LOGIN_KEY"))
var sessionName = "ECKONID"

func handleAuthenticationFunc(r *mux.Router) {
    r.HandleFunc("/authentication", handleLoginPage).Methods("GET")
    r.HandleFunc("/authentication/login", handleLoginAttempt).Methods("POST")
}

func handleLoginPage(wr http.ResponseWriter, req *http.Request) {
    if checkLoginStatus(req) {
        // log out if we visit the login site again (while being logged in)
        session, _:= store.Get(req, sessionName)
        session.Values["token"] = ""
        session.Values["user_name"] = ""
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
    session, _ := store.Get(req, sessionName)

    var users []map[string]string
    users, _ = readAuthenticationFile(users)

    success := false
    for _, user := range users {
        if req.Form["name"][0] == user["user_name"] && req.Form["password"][0] == user["password_hash"] {
            // safe the status of the one user
            session.Values["token"] = user["token"]
            session.Values["user_name"] = user["user_name"]
            session.Save(req, wr)
            success = true
        }
    }

    // do it with this "ugly" way, because with 2 redirect the compiler will throw warnings (it wants to use 2 wr !)
    if success {
        http.Redirect(wr, req, "/", http.StatusSeeOther)
    } else {
        // no valid user -> redirect
        http.Redirect(wr, req, "/authentication", http.StatusSeeOther)
    }
}

func checkLoginStatus(req *http.Request) bool {
    session, _ := store.Get(req, sessionName)

    var users []map[string]string
    users, _ = readAuthenticationFile(users)

    // check if the username and the token are valid in the cookie
    for _, user := range users {
        if user["user_name"] == session.Values["user_name"] && user["token"] == session.Values["token"] {
            return true
        }
    }

    return false
}

func getCurrentUsername(req *http.Request) string {
    session, _ := store.Get(req, sessionName)

    // check if the conversion can be used - if not return an empty string (preventing error)
    if name, ok := session.Values["user_name"].(string); ok {
        return name
    }

    return ""
}

// read the json file and return a map, so we still have the keys:values format
func readAuthenticationFile(data []map[string]string) ([]map[string]string, error) {
    c, err := ioutil.ReadFile("public/data/authentication/data.json")
    if err != nil {
        fmt.Println("Error while reading File.", err.Error())
        return data, err
    }

    err = json.Unmarshal(c, &data)
    if err != nil {
        fmt.Println("Error while reading File.", err.Error())
    }

    return data, nil
}
