package adventure

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

// Story exported type
type Story map[string]Storyarc

//var tpl *template.Template

var defaulthandlertemplate = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Choose your own ADVENTURE</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Story}}
        <p> {{.}}</p>
        {{end}}
        <ul>
            {{range .Options}}
			<li> <a href="/{{.Arc}}">{{.Text}}</a></li>
			{{end}}
        </ul>
    </body>
</html>`

//func init() {
//	tpl = template.Must(template.New("").Parse(defaulthandlertemplate))
//
//}

type handler struct {
	s Story
}

//Newhandler function to handle web requests
func Newhandler(s Story) http.Handler {
	return handler{s}

}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(defaulthandlertemplate))
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	// "/intro" -> "intro"
	path = path[1:]

	if text, ok := h.s[path]; ok {
		err := tpl.Execute(w, text)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found...", http.StatusNotFound)
}

// Jsonparser to parse the json data to a struct type
func Jsonparser(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	err := d.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil

}

//Option USED TO STORE OPTION
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

//Storyarc USED TO STORE STORY
type Storyarc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}
