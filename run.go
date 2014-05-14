package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save(path string) error {
	filename := path + "/" + p.Title + ".pk"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".pk"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func shitAppend(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/video", videoHandler)
	http.HandleFunc("/admin", adminHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.ListenAndServe(":8080", nil)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	file, err := loadPage("article/push")
	shitAppend(err)

	temp, err := template.ParseFiles("template/push.html")
	shitAppend(err)

	fmt.Println("[*]Handling Request from : " + req.RemoteAddr)
	temp.Execute(rw, file)
}

func videoHandler(rw http.ResponseWriter, req *http.Request) {
	http.RedirectHandler("/admin", 303)
}

func adminHandler(rw http.ResponseWriter, req *http.Request) {
	temp, _ := template.ParseFiles("template/admin.html")
	temp.Execute(rw, nil)
	if req.Method == "POST" {
		page := &Page{Title: req.FormValue("title"), Body: []byte(req.FormValue("content"))}
		page.save("article")
	}
}
