package main

import (
	"PushKids/module/handle"
	"net/http"
)

func main() {
	// GET Handler
	http.HandleFunc("/", handle.Handler)
	http.HandleFunc("/profil", handle.ProfilHandler)
	http.HandleFunc("/home", handle.HomeHandler)
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
	// Start to serve
	http.ListenAndServe(":80", nil)
}
