package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "io/ioutil"
    "fmt"
    "encoding/json"
    "os"
)

func handleCalendarFunc(r *mux.Router) {
    r.HandleFunc("/calendar", showCalendar).Methods("GET")
    r.HandleFunc("/calendar/ajax", ajaxCalendarAdd).Methods("POST")
    r.HandleFunc("/calendar/ajax", ajaxCalendarDelete).Methods("DELETE")
}

func showCalendar(wr http.ResponseWriter, req* http.Request) {
    pageData := pageData{
        HeaderInfo: getHeaderInfo(req),
    }
    var data []map[string]string
    pageData.Data, _ = readCalendarFile(data)

    if err := tpl.ExecuteTemplate(wr, "calendar.gohtml", pageData); err != nil {
        panic(err)
        http.Error(wr, err.Error(), http.StatusInternalServerError)
        return
    }
}

// get the old/new json data and append it
func ajaxCalendarAdd(wr http.ResponseWriter, req* http.Request) {
    if !checkLoginStatus(req) {
        return
    }
    var data []map[string]string
    msg := translateCalendarRequest(req)
    data, _ = readCalendarFile(data)
    data = append(data, msg)

    writeCalendarFile(data)
}

// get the to delete value, search and update the json accordingly
func ajaxCalendarDelete(wr http.ResponseWriter, req* http.Request) {
    if !checkLoginStatus(req) {
        return
    }
    var data []map[string]string
    msg := translateCalendarRequest(req)
    data, _ = readCalendarFile(data)

    // instead of deleting -> just make a new map without the deleted data
    var updatedData []map[string]string
    for _, d := range data {
        if !(d["date"] == msg["date"] && d["name"] == msg["name"]) {
            updatedData = append(updatedData, d)
        }
    }

    writeCalendarFile(updatedData)
}

// get the request data and translate it into a map
func translateCalendarRequest(req* http.Request) map[string]string{
    b, _ := ioutil.ReadAll(req.Body)
    defer req.Body.Close()
    var msg map[string]string
    // get the post values
    json.Unmarshal(b, &msg)

    return msg
}

// overwrite the json data
func writeCalendarFile(data []map[string]string) {
    a, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Error while writing File.", err.Error())
        return
    }
    ioutil.WriteFile("public/data/calendar/data.json", []byte(a), os.ModeAppend)
}

// read the json file and return a map, so we still have the keys:values format
func readCalendarFile(data []map[string]string) ([]map[string]string, error) {
    c, err := ioutil.ReadFile("public/data/calendar/data.json")
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
