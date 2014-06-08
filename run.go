package main

import (
	"PushKids/module/handle"
	"flag"
	"net/http"
)

func main() {
	port := flag.String("port", "80", "-port [your port]")
	flag.Parse()
	// GET Handler
	http.HandleFunc("/", handle.Handler)
	http.HandleFunc("/profil", handle.ProfilHandler)
	http.HandleFunc("/home", handle.HomeHandler)
	http.HandleFunc("/find", handle.FindHandler)
	// POST Handler
	http.HandleFunc("/signin", handle.SignIn)
	http.HandleFunc("/logout", handle.LogOut)
	http.HandleFunc("/add", handle.AddUser)
	http.HandleFunc("/fish", handle.AddFish)
	http.HandleFunc("/delete", handle.DeleteFish)
	// Ressources
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))
	// Start to serve
	http.ListenAndServe(":"+*port, nil)
}
