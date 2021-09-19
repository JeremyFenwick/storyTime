package storyTime

import (
	"encoding/json"
	"net/http"
	"os"
	"text/template"
)

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(defaultHanderTemplate))
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

func JsonStory(file *os.File) (Story, error) {
	r := json.NewDecoder(file)
	var narrative Story
	if err := r.Decode(&narrative); err != nil {
		return nil, err
	}
	return narrative, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"story"`
	Options   []Option `json:"options"`
}

type Option struct {
	Text        string `json:"text"`
	NextChapter string `json:"arc"`
}

var defaultHanderTemplate = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Story Time!</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Paragraph}}
        <p>{{.}}</p>
        {{end}}
        <ul>
            {{range .Options}}
            <li><a href=/{{.NextChapter}}>{{.Text}}</a></li>
            {{end}}
        </ul>
    </body>
</html>
`
