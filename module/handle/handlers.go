package handle

import (
	"PushKids/module/database"
	"time"
)

type User struct {
	Username string
	Email    string
	Password string
	Date     time.Time
}

type Fish struct {
	Id       int
	Type     string
	Username string
	Weight   float64
	Length   float64
	Location string
	Date     string
	Lure     string
	Info     string
	Picture  string
}

var db = database.SetupDB()
