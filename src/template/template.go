package template

import (
	"html/template"
	"net/http"
)

type PageData struct {
	HeaderInfo HeaderInfo
	Data       interface{}
}

type HeaderInfo struct {
	Title      string
	Navigation []PathInfo
	OnHomePage bool
	User       string
}

type PathInfo struct {
	Title string
	Path  string
	Class string
}

var Template *template.Template

// Initialize : Initializes the tempalte for further use
func Initialize() {
	Template = template.Must(template.ParseGlob("src/template/view/*.gohtml"))
}

// GetHeaderInfo : build the headerInfo so we can use it everywhere and only have one place to edit it
func GetHeaderInfo(req *http.Request) (h HeaderInfo) {
	// update the title if the user is already logged in
	status := PathInfo{"Login", "/authentication", "authentication"}

	// nav-bar information
	h = HeaderInfo{
		Title: "eckon.dev",
		Navigation: []PathInfo{
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
