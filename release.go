package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "encoding/json"
    "fmt"
)

type JsonData struct {
    Releases interface{}
}

func handleReleaseFunc(r *mux.Router) {
    r.HandleFunc("/release", releaseHandler).Methods("GET")
}

func releaseHandler(wr http.ResponseWriter, req *http.Request) {
    pageData := pageData{
        HeaderInfo: getHeaderInfo(req),
    }

    jsonData := new(JsonData)
    // id from normal search https://www.discogs.com/de/artist/40027-Disturbed
    url := "https://api.discogs.com/artists/40027/releases"
    // id from main_release
    // https://api.discogs.com/releases/12678059
    // http://jsonviewer.stack.hu/
    requestApi(url, jsonData)
    fmt.Println(jsonData.Releases)

    if err := tpl.ExecuteTemplate(wr, "release.gohtml", pageData); err != nil {
        http.Error(wr, err.Error(), http.StatusInternalServerError)
        return
    }
}

func requestApi(url string, target *JsonData) {
    res, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    json.NewDecoder(res.Body).Decode(target)
}
