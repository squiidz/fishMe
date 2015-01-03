package main

import (
	"flag"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/go-zoo/claw"
	"github.com/go-zoo/claw/mw"
	"github.com/squiidz/fishMe/module/handle"
)

func main() {
	port := flag.String("port", "80", "-port [your port]")
	flag.Parse()

	mux := bone.New()
	clw := claw.New(mw.Logger)
	// GET Handler
	mux.Get("/", clw.Use(handle.Handler))
	mux.Get("/profil", clw.Use(handle.ProfilHandler))
	mux.Get("/home", clw.Use(handle.HomeHandler))
	mux.Get("/find", clw.Use(handle.FindHandler))
	// POST Handler
	mux.Post("/signin", clw.Use(handle.SignIn))
	mux.Post("/logout", clw.Use(handle.LogOut))
	mux.Post("/add", clw.Use(handle.AddUser))
	mux.Post("/fish", clw.Use(handle.AddFish))
	mux.Post("/delete", clw.Use(handle.DeleteFish))
	// Ressources
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	// Start to serve
	http.ListenAndServe(":"+*port, mux)
}
