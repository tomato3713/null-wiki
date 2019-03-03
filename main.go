/*
+ wiki
+ wiki.go
+-+-data // Application Use data set
| |---tmpl
| |---img
|
+-+-pages
  |---text // User's wiki document text data
  |---img // User's wiki document image
*/

package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() // フラグを解釈します

	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("data/"))))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/", makeHandler(frontHandler))

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
