package main

import (
    "fmt"
    "os"
    "github.com/gorilla/mux"
    "net/http"
    "regexp"
    "io/ioutil"
    "encoding/json"
    "strconv"
)

/////////////////////////////////////
/* OLD SHIT CODE --- REFACTOR ASAP */
/////////////////////////////////////

type gameData struct {
    Player1 personData
    Player2 personData
    Player3 personData
    Player4 personData
    GameName string
}

type personData struct {
    Name string
    Points int
    Level int
    CanCompleteLevel bool
}

var gameDataTemplate = gameData{
    personData{
        "Niklas",
        0,
        1,
        true,
    },personData{
        "Alina",
        0,
        1,
        true,
    },personData{
        "Ludger",
        0,
        1,
        true,
    },personData{
        "Birgit",
        0,
        1,
        true,
    },
    readGameNumber(),
}

var game gameData
var err error

func init() {
    // Für Phase
    game.GameName = readGameNumber()
    game, err = readFile(game.GameName, game)
    if err != nil {
        // Falls keine Datei vorhanden  ist, so generiere eine neue mit dem Template
        fmt.Println("Created a new File with the default Template")
        writeFile(game.GameName, gameDataTemplate, os.ModeAppend)
    }
}

func handlePhaseFunc(r *mux.Router) {
    // old code -> phase and reset have the method check in the function itself
    r.HandleFunc("/phase", phaseHandler)
    r.HandleFunc("/phase/reset", resetHandler)
}

func phaseHandler(wr http.ResponseWriter, req *http.Request) {
    game, _ = readFile(game.GameName, game)
    // Wird ausgeführt, falls ein Post vorliegt
    // check if the user is logged in
    if checkLoginStatus(req) && req.Method == http.MethodPost {
        // Muss ausgeführt werden um die Post-Daten zu bekommen
        err := req.ParseForm()
        if err != nil {
            http.Error(wr, err.Error(), http.StatusInternalServerError)
            return
        }
        reNumbers := regexp.MustCompile("[0-9]+")
        reLetters := regexp.MustCompile("[a-z]+")
        // Durch iterieren durch die Post-Einträge
        for key, values := range req.PostForm {
            // Aufgabe und Position sind beide in key, also durch regex rausbekommen
            keyLetter := reLetters.FindAllString(key, -1)[0]
            keyNumber := reNumbers.FindAllString(key, -1)[0]
            // Update der jeweiligen Inhalte
            if keyLetter == "level" {
                game = addLevel(keyNumber, game)
            } else if keyLetter == "points" {
                game = addPoints(keyNumber, game, values[0])
            }
        }
        writeFile(game.GameName, game, os.ModeAppend)
    }

    game, _ = readFile(game.GameName, game)
    pd := pageData{
        HeaderInfo: getHeaderInfo(req),
        Data: game,
    }

    err := tpl.ExecuteTemplate(wr, "phase.gohtml", pd)
    if err != nil {
        http.Error(wr, err.Error(), http.StatusInternalServerError)
        return
    }
}

func resetHandler(wr http.ResponseWriter, req *http.Request)  {
    // check if the user is logged in
    if checkLoginStatus(req) {
        // Erhöhung der Spielnummer, und diese dem Template für das neue Spiel übergeben
        game.GameName = increaseGameNumber()
        gameDataTemplate.GameName = game.GameName
        writeFile(game.GameName, gameDataTemplate, os.ModeAppend)
    }
    http.Redirect(wr, req, "/phase", 301)
}


func readFile(filename string, data gameData) (gameData, error) {
    c, err := ioutil.ReadFile("public/data/phase/" + filename)
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

func writeFile(filename string, data gameData, mode os.FileMode) {
    // Encoden von gameData und in File schreiben
    a, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Error while writing File.",err.Error())
        return
    }
    ioutil.WriteFile("public/data/phase/" + filename, []byte(a), mode)
}

func readGameNumber() string {
    c, _ := ioutil.ReadFile("public/data/phase/gameNumber")
    return string(c)
}

func increaseGameNumber() string {
    numberString := readGameNumber()
    number, _ := strconv.Atoi(numberString)
    number++
    numberString = strconv.Itoa(number)
    ioutil.WriteFile("public/data/phase/gameNumber", []byte(numberString), os.ModeAppend)
    return numberString
}

func addLevel(value string, data gameData) gameData {
    // Erhöhung der Phase
    switch value {
    case "1":
        if data.Player1.Level < 10 && data.Player1.CanCompleteLevel{
            data.Player1.Level++
            data.Player1.CanCompleteLevel = false
        }
    case "2":
        if data.Player2.Level < 10 && data.Player2.CanCompleteLevel{
            data.Player2.Level++
            data.Player2.CanCompleteLevel = false
        }
    case "3":
        if data.Player3.Level < 10 && data.Player3.CanCompleteLevel{
            data.Player3.Level++
            data.Player3.CanCompleteLevel = false
        }
    case "4":
        if data.Player4.Level < 10 && data.Player4.CanCompleteLevel{
            data.Player4.Level++
            data.Player4.CanCompleteLevel = false
        }
    }
    return data
}

func addPoints(value string, data gameData, points string) gameData {
    // Converting String zu Int - Allgemein Bildung der Summe
    i, _ := strconv.Atoi(points)
    switch value {
    case "1": data.Player1.Points += i
    case "2": data.Player2.Points += i
    case "3": data.Player3.Points += i
    case "4": data.Player4.Points += i
    }
    // Reset der Phasenerhöhung, wenn Runde zuende
    data.Player1.CanCompleteLevel = true
    data.Player2.CanCompleteLevel = true
    data.Player3.CanCompleteLevel = true
    data.Player4.CanCompleteLevel = true
    return data
}
