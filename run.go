package main

import (
	"PushKids/module/handle"
	"net/http"
)

func main() {

	http.HandleFunc("/", handle.Handler)
	http.HandleFunc("/home", handle.HomeHandler)
	/*
		http.HandleFunc("/home", homeHandler)
		http.HandleFunc("/admin", adminHandler)
		http.HandleFunc("/Sign", SignUp)
		http.HandleFunc("/addFish", FishForm)
	*/
	http.HandleFunc("/signin", handle.SignIn)
	http.HandleFunc("/logout", handle.LogOut)
	http.HandleFunc("/add", handle.AddUser)
	http.HandleFunc("/fish", handle.AddFish)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	http.ListenAndServe(":80", nil)
}
