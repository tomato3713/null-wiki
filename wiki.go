/*
+ wiki
+ wiki.go
+-+-data // Application Use data set
| |---template
| |---img
|
+-+-pages
  |---text // User's wiki document text data
  |---img // User's wiki document image
*/

package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

// Page struct describes how page data will be stored in memory.
type Page struct {
	Title string
	Body  []byte
}

// Save method will save the Page's Body to a text file.
// For simplicity, use file name as title
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filepath.Join(usrDir, filename), p.Body, 0600)
}

// loadPage method return the Page struct pointer.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filepath.Join(usrDir, filename))

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func getBinDir() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(exe), nil
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func init() {
	binDir, err := getBinDir()
	if err != nil {
		os.Exit(1)
	}
	usrDir = filepath.Join(binDir, `pages`)
	dataDir := filepath.Join(binDir, `data`)

	for _, tmpl := range []string{"edit", "view"} {
		file := tmpl + ".html"
		t := template.Must(template.ParseFiles(filepath.Join(dataDir, `tmpl`, file)))
		templates[tmpl] = t
	}
}

var templates = make(map[string]*template.Template)
var usrDir, dataDir string

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates[tmpl].Execute(w, p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
