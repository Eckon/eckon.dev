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
    // read the json body
    b, _ := ioutil.ReadAll(req.Body)
    defer req.Body.Close()
    var msg map[string]string
    // get the post values
    json.Unmarshal(b, &msg)

    // get the old json data and append it with the new post
    var data []map[string]string
    data, _ = readCalendarFile(data)
    data = append(data, msg)

    // overwrite the file
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
