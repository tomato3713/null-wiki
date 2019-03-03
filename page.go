package main

import (
	"io/ioutil"
	"path/filepath"
)

// Page struct describes how page data will be stored in memory.
type Page struct {
	Title  string
	Body   []byte
	Parsed []byte
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

	parsedBody := goMdParse(body)

	return &Page{Title: title, Body: body, Parsed: parsedBody}, nil
}
