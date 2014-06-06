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

type ById []Fish

func (f ById) Len() int           { return len(f) }
func (f ById) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f ById) Less(i, j int) bool { return f[i].Id > f[j].Id }

var db = database.SetupDB()
